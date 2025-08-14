package main

import (
	"fmt"
	"reflect"
)

type ecs struct {
	archetypes   map[uint64]*archetype
	entToArche   map[entity]*archetype //What archetype an entity belongs to
	componentIds map[reflect.Type]uint64
	nextCompID   uint64
}

func NewEcs() *ecs {
	return &ecs{
		archetypes:   map[uint64]*archetype{},
		entToArche:   map[entity]*archetype{},
		componentIds: map[reflect.Type]uint64{},
		nextCompID:   0,
	}
}

func register[T any](ecs *ecs) uint64 {
	key := reflect.TypeOf((*T)(nil)).Elem()
	if compId, ok := ecs.componentIds[key]; ok {
		return compId
	}

	var compId uint64 = 1 << ecs.nextCompID
	ecs.componentIds[key] = compId
	ecs.nextCompID++
	return compId
}

func (ecs *ecs) NewEntity(a *archetype) entity {
	e := newEntity()
	a.insertEntity(e)
	ecs.entToArche[e] = a
	return e
}

func Query[T any](ecs *ecs, e entity) T {
	tType := reflect.TypeOf((*T)(nil)).Elem()
	_, ok := ecs.componentIds[tType]
	if !ok {
		panic(fmt.Sprintf("Parameter Type:%v is not registred in ecs", tType))
	}
	a := ecs.entToArche[e]
	untypeStore := a.compStore[ecs.componentIds[tType]]
	typedStore, _ := untypeStore.(*compStore[T])
	return typedStore.data[a.entityIdx[e]]
}

func (ecs *ecs) EmbedArchetype(a *archetype) {
	ecs.archetypes[a.signature] = a
	a.mutable = false
}

func MutateEntity[T any](ecs *ecs, e entity, comp T) {
	TType := reflect.TypeOf((*T)(nil)).Elem()
	compType := reflect.TypeOf(comp)
	ecs.isCompRegistered(comp)
	if TType != compType {
		panic(fmt.Sprintf("Generic Type:%v is not the same type as argument comp type:%v", TType, compType))
	}
	compSig := ecs.componentIds[compType]
	a, ok := ecs.entToArche[e]
	if !ok {
		panic("Could not find matching archetype")
	}
	store, ok := a.compStore[compSig].(*compStore[T])
	if !ok {
		panic("Could not type assert")
	}
	store.data[a.entityIdx[e]] = comp
}

func (ecs *ecs) isCompRegistered(comp any) {
	compType := reflect.TypeOf(comp)
	if _, ok := ecs.componentIds[compType]; !ok {
		panic(fmt.Sprintf("Component:%v is not registered with ecs", compType))
	}
}

func (ecs *ecs) AddEntToArche(a archetype, compSet ...any) entity {
	e := newEntity()
	for _, comp := range compSet {
		compType := reflect.TypeOf(comp)
		compSig, ok := ecs.componentIds[compType]
		//If comp is registered
		if !ok {
			panic(fmt.Sprintf("Component:%v is registered", compType))
		}

		//If component does not belong to archetype
		if (compSig & a.signature) != compSig {
			panic(fmt.Sprintf("Component:%v is not a part of archetype with signature:%v", compType, a.signature))
		}

		ecs.entToArche[e] = &a
		a.insertEntity(e)
		compStore := a.compStore[compSig]
		compStore.add(comp, e, &a)
		a.compStore[compSig] = compStore
	}
	return e
}
