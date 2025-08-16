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
	if ecs.nextCompID > 64 {
		panic("Limit has been reached on number of components that can be registered")
	}
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
		panic(fmt.Sprintf("Gecs: Parameter Type:%v is not registred in ecs", tType))
	}
	a := ecs.entToArche[e]
	if a.mutable {
		panic("Gecs: Archetype composition must be finalized by calling ecs's method EmbedArchetype() before use.")
	}
	untypeStore := a.compStores[ecs.componentIds[tType]]
	typedStore, _ := untypeStore.(*compStore[T])
	return typedStore.data[a.entityIdx[e]]
}

func (ecs *ecs) EmbedArchetype(a *archetype) {
	ecs.archetypes[a.signature] = a
	a.mutable = false
}

func MutateEntity[T any](ecs *ecs, e entity, comp T) {
	tType := reflect.TypeOf((*T)(nil)).Elem()
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
