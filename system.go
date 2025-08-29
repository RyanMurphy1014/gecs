package main

func RegisterSystem(ecs *ecs, sys system, writesMask uint64, readsMask uint64) {
	ecs.sysHeaders = append(ecs.sysHeaders, systemHeader{
		writesMask: writesMask,
		readsMask:  readsMask,
		system:     sys,
		inDegree:   0,
	})

}

type systemHeader struct {
	writesMask uint64
	readsMask  uint64
	system
	inDegree int
}

type system interface {
	run(ecs *ecs)
}

type conditionalSys struct {
	op        func(ecs *ecs)
	condition bool
}

func (cs conditionalSys) run(ecs *ecs) {
	cs.op(ecs)
}

type timedSys struct {
	op func(ecs *ecs)
	dt float64
}

func (ds timedSys) run(ecs *ecs) {
	ds.op(ecs)
}

type simpleSys struct {
	op func(ecs *ecs)
}

func (ss simpleSys) run(ecs *ecs) {
	ss.op(ecs)
}
