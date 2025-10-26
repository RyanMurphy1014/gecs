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

func validateCache(ecs *ecs) error {
	if !ecs.validCache {
		err := cacheDepGraph(ecs)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ecs *ecs) TickSystems() error {
	err := validateCache(ecs)

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

	return err
}

func generateDepGraph(ecs *ecs, sysHeaders []systemHeader) {
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
}

func batch(ecs *ecs, sysHeaders []systemHeader) (err error, batch []systemHeader) {
	acyclicDependency := false
	batchingComplete := false
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

		writeBatches(ecs, batch)
	}
	if newSystemAdded {
		return ErrSystemsNeedReBatching, nil
	}

	if !acyclicDependency {
		return ErrCircularDependency, nil
	}
	return nil, batch
}

func writeBatches(ecs *ecs, batch []systemHeader) {
	currBatch := make([]systemHeader, 0, 64)
	for _, header := range batch {
		header.ioReadyToInfer = true
		currBatch = append(currBatch, header)
		ecs.batches = append(ecs.batches, currBatch)

		for _, dep := range ecs.depGraph[header.id] {
			dep.inDegree--
		}
	}
}

func cacheDepGraph(ecs *ecs) error {
	ecs.batches = make([][]systemHeader, 0, 64)
	sysHeaders, err := applicableSystems(ecs.sysHeaders)
	if err != nil && len(ecs.sysHeaders) == 0 {
		return fmt.Errorf("Gecs: No systems are registered with this ECS ~ %w", err)
	}
	if err != nil {
		return err
	}
	generateDepGraph(ecs, sysHeaders)
	batch(ecs, sysHeaders)

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
		case System:
			applicable = append(applicable, sysHeaders[i])
		case ConditionalSystem:
			if v.condition() == true {
				applicable = append(applicable, sysHeaders[i])
			}
		}

	}
	return applicable, nil
}

func conflict(sysA systemHeader, sysB systemHeader) bool {
	switch {
	case sysA.id == sysB.id:
		return false

	case sysA.ioReadyToInfer == false || sysB.ioReadyToInfer == false:
		return true

	case (sysA.writesMask & sysB.writesMask) != 0:
		return true

	case (sysA.writesMask & sysB.readsMask) != 0:
		return true

	case (sysA.readsMask & sysB.writesMask) != 0:
		return true

	default:
		return false
	}
}
