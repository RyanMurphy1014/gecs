package main

import "sync"

func (ecs *ecs) ExecuteSystems() {
	//Generate dependency graph
	sysHeaders := applicableSystems(ecs.sysHeaders)
	depGraph := make(map[systemHeader][]systemHeader)
	for sysA := range 64 {
		for sysB := range 64 {
			if conflict(sysHeaders[sysA], sysHeaders[sysB]) {
				deps := depGraph[sysHeaders[sysA]]
				deps = append(deps, sysHeaders[sysB])
				depGraph[sysHeaders[sysA]] = deps
				sysHeaders[sysB].inDegree++
			}
		}
	}

	//Mark batchable systems
	acyclicDependency := false
	batchingComplete := false
	for !batchingComplete {
		markedToRun := make([]systemHeader, 0, len(sysHeaders))
		for i := len(sysHeaders) - 1; i >= 0; i-- {
			header := sysHeaders[i]
			if header.inDegree == 0 && len(sysHeaders) > 0 {
				markedToRun = append(markedToRun, header)
				lastIdx := len(sysHeaders) - 1
				//Remove marked heders from list
				sysHeaders[i], sysHeaders[lastIdx] = sysHeaders[lastIdx], sysHeaders[i]
				sysHeaders = sysHeaders[:lastIdx]
			} else {
				i--
				//Only look to the next index if order of sysHeaders is unchanged
			}
		}

		if len(markedToRun) != 0 {
			acyclicDependency = true
		}
		if len(sysHeaders) == 0 {
			batchingComplete = true
		}

		//Execute batch of systems - will be skipped if batch is empty
		var wg sync.WaitGroup
		for _, header := range markedToRun {
			wg.Add(1)
			go func() {
				defer wg.Done()
				header.run(ecs)
			}()
			for _, dep := range depGraph[header] {
				dep.inDegree--
			}
		}
		wg.Wait()

	}

	if !acyclicDependency {
		panic("Gecs: Circular system dependency found.")
	}
}

func applicableSystems(sysHeaders []systemHeader) []systemHeader {
	applicable := make([]systemHeader, 0, 64)
	for i, header := range sysHeaders {
		sys := header.system
		switch v := sys.(type) {
		case simpleSys:
			applicable = append(applicable, sysHeaders[i])
		case conditionalSys:
			if v.condition {
				applicable = append(applicable, sysHeaders[i])
			}
		case timedSys:
			if v.dt > 0 { //PLACEHOLDER
				applicable = append(applicable, sysHeaders[i])
			}
		}

	}
	return applicable
}

func conflict(sysA systemHeader, sysB systemHeader) bool {
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
