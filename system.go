package main

import (
	"fmt"
	"reflect"
)

var nextSysId uint64 = 1

func newSysId() uint64 {
	currId := nextSysId
	nextSysId++
	if nextSysId < currId {
		panic("Gecs: System Id's have overflowed.")
	}
	return currId
}

func RegisterSystem(ecs *ecs, sys system, sysId uint64) {
	ecs.validCache = false
	ecs.batches = nil
	newHeader := systemHeader{
		writesMask:     0,
		readsMask:      0,
		system:         sys,
		inDegree:       0,
		ioReadyToInfer: false,
		id:             sysId,
	}
	ecs.sysHeaders = append(ecs.sysHeaders, newHeader)
	ecs.sysToHeader[sysId] = &newHeader
}

type systemHeader struct {
	id             uint64
	ioReadyToInfer bool //Tells scheduler that it will have io set next run
	writesMask     uint64
	readsMask      uint64
	system
	inDegree int
}

type system interface {
	run(ecs *ecs)
}

type ConditionalSys struct {
	op        func(ecs *ecs)
	condition func() bool
}

func (cs ConditionalSys) run(ecs *ecs) {
	cs.op(ecs)
}

type TimedSys struct {
	op func(ecs *ecs)
	dt float64
}

func (ds TimedSys) run(ecs *ecs) {
	ds.op(ecs)
}

type SimpleSys struct {
	op func(ecs *ecs)
}

func (ss SimpleSys) run(ecs *ecs) {
	ss.op(ecs)
}

func SystemWrite[T any](ecs *ecs, e entity, comp T, sysId uint64) {
	sysHeader := ecs.sysToHeader[sysId]
	tType := reflect.TypeOf((*T)(nil)).Elem()
	sysHeader.writesMask |= ecs.componentIds[tType]
	compType := reflect.TypeOf(comp)

	if _, ok := ecs.componentIds[compType]; !ok {
		panic(fmt.Sprintf("Gecs: Component:%v is not registered with ecs", compType))
	}

	if tType != compType {
		panic(fmt.Sprintf("Gecs: Generic Type:%v is not the same type as argument comp type:%v", tType, compType))
	}

	compSig := ecs.componentIds[compType]
	a, ok := ecs.entToArche[e]
	if !ok {
		panic("Gecs: Could not find matching archetype")
	}

	store, ok := a.compStores[compSig].(*compStore[T])
	if !ok {
		panic("Gecs: Could not type assert")
	}

	store.data[a.entityIdx[e]] = comp
}

func SystemRead[T any](ecs *ecs, e entity, sysId uint64) T {
	sysHeader := ecs.sysToHeader[sysId]
	tType := reflect.TypeOf((*T)(nil)).Elem()
	sysHeader.readsMask |= ecs.componentIds[tType]
	_, ok := ecs.componentIds[tType]
	if !ok {
		panic(fmt.Sprintf("Gecs: Parameter Type:%v is not registred in ecs", tType))
	}
	a := ecs.entToArche[e]
	if a.mutableComposition {
		panic("Gecs: Archetype composition must be finalized by calling ecs's method EmbedArchetype() before use.")
	}
	untypeStore := a.compStores[ecs.componentIds[tType]]
	typedStore, _ := untypeStore.(*compStore[T])
	return typedStore.data[a.entityIdx[e]]
}
