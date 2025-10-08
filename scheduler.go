package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrSystemsNeedReBatching = errors.New("Gecs: On next system run, newly infered systems can be batched.")
var ErrCircularDependency = errors.New("Gecs: Circular system dependency found. Systems prevented from running due to detected race condition.")

type dependencyCache struct {
	depGraph   map[uint64][]systemHeader
	validCache bool
	batches    [][]systemHeader
}

func (ecs *ecs) ExecuteSystems() error {
	needsRebatch := false
	if !ecs.validCache {
		cacheError := cacheDepGraph(ecs)
		if !errors.Is(cacheError, ErrSystemsNeedReBatching) {
			return cacheError
		} else {
			needsRebatch = true
		}
	}

	var wg sync.WaitGroup
	for _, batch := range ecs.batches {
		for _, sys := range batch {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sys.run(ecs)
			}()
		}
		wg.Wait()
	}

	if needsRebatch {
		cacheDepGraph(ecs)
	}

	return nil
}

func cacheDepGraph(ecs *ecs) error {
	ecs.batches = make([][]systemHeader, 64)
	sysHeaders, err := applicableSystems(ecs.sysHeaders)
	if err != nil && len(ecs.sysHeaders) == 0 {
		return fmt.Errorf("Gecs: No systems are registered with this ECS ~ %w", err)
	}
	if err != nil {
		return err
	}
	//Generate dependency graph
	depGraph := ecs.depGraph
	if len(sysHeaders) > 1 {
		for _, sysA := range sysHeaders {
			for _, sysB := range sysHeaders {
				if conflict(sysA, sysB) {
					deps := depGraph[sysA.id]
					deps = append(deps, sysB)
					depGraph[sysA.id] = deps
					sysB.inDegree++
				}
			}
		}
	}
	//Mark batchable systems
	acyclicDependency := false
	batchingComplete := false
	batchCount := 0
	newSystemAdded := false
	for !batchingComplete {
		batch := make([]systemHeader, 0, len(sysHeaders))
		for i := len(sysHeaders) - 1; i >= 0; i-- {
			header := sysHeaders[i]
			if header.ioReadyToInfer == false {
				ecs.validCache = false
				newSystemAdded = true
				batch = append(batch, header)
				lastIdx := len(sysHeaders) - 1
				//Remove marked headers from list
				sysHeaders[i], sysHeaders[lastIdx] = sysHeaders[lastIdx], sysHeaders[i]
				sysHeaders = sysHeaders[:lastIdx]
				break // NOTE: Breaks so an uninfered systems goes in their own batch
			}
			if header.inDegree == 0 && len(sysHeaders) > 0 {
				batch = append(batch, header)
				lastIdx := len(sysHeaders) - 1
				//Remove marked headers from list
				sysHeaders[i], sysHeaders[lastIdx] = sysHeaders[lastIdx], sysHeaders[i]
				sysHeaders = sysHeaders[:lastIdx]
			} else {
				i--
				//Only look to the next index if order of sysHeaders is unchanged
			}
		}

		if len(batch) != 0 {
			acyclicDependency = true
		}
		if len(sysHeaders) == 0 {
			batchingComplete = true
		}

		//Write batches
		currBatch := make([]systemHeader, 0, 64)
		for _, header := range batch {
			header.ioReadyToInfer = true
			currBatch = append(currBatch, header)
			ecs.batches[batchCount] = currBatch

			for _, dep := range ecs.depGraph[header.id] {
				dep.inDegree--
			}
		}
	}

	if newSystemAdded {
		return ErrSystemsNeedReBatching
	}

	if !acyclicDependency {
		return ErrCircularDependency
	}

	ecs.validCache = true
	return nil
}

func applicableSystems(sysHeaders []systemHeader) ([]systemHeader, error) {
	if len(sysHeaders) == 0 {
		return nil, errors.New("Gecs: No systems to check if applicable")
	}
	applicable := make([]systemHeader, 0, 64)
	for i, header := range sysHeaders {
		sys := header.system
		switch v := sys.(type) {
		case SimpleSys:
			applicable = append(applicable, sysHeaders[i])
		case ConditionalSys:
			if v.condition() == true {
				applicable = append(applicable, sysHeaders[i])
			}
		case TimedSys:
			if v.dt > 0 { // TODO: PLACEHOLDER
				applicable = append(applicable, sysHeaders[i])
			}
		}

	}
	return applicable, nil
}

func conflict(sysA systemHeader, sysB systemHeader) bool {
	if sysA.ioReadyToInfer == false || sysB.ioReadyToInfer == false {
		return true
	}
	if (sysA.writesMask & sysB.writesMask) != 0 {
		return true
	}
	if (sysA.writesMask & sysB.readsMask) != 0 {
		return true
	}
	if (sysA.readsMask & sysB.writesMask) != 0 {
		return true
	}
	return false
}
