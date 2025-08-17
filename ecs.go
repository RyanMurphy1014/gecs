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

func register[T any](ecs *ecs) uint64 {
	if ecs.nextCompID > 64 {
		panic("Gecs: Limit has been reached on number of components that can be registered")
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

func UpdateEntity[T any](ecs *ecs, e entity, comp T) {
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

func DeleteComponent(ecs *ecs, e entity, compSig uint64) *archetype {
	currArche := ecs.entToArche[e]
	copiedCompStores := make(map[uint64]compStorer)
	targetArcheSig := currArche.signature ^ compSig
	var idx uint64 = 0
	for i := range 64 {
		idx = 1 << i
		targetCompSig := targetArcheSig & idx
		if targetCompSig > 0 {
			copiedCompStores[targetCompSig] = currArche.compStores[targetCompSig].get(currArche.entityIdx[e])
		}
	}
	currArche.unlinkEntity(e)

	targetArcheSig = currArche.signature ^ compSig
	var outputArche *archetype = nil
	if exsistingArche, ok := ecs.archetypes[targetArcheSig]; !ok {
		newArche := &archetype{
			signature:          targetArcheSig,
			compStores:         copiedCompStores,
			entityIdx:          map[entity]int{},
			entities:           []entity{},
			mutableComposition: true,
		}
		newArche.entityIdx[e] = 0
		if newArche.signature == 0 {
			panic("Gecs: Attempted to remove from a single component entity. Empty entities are not allowed.")
		}

		ecs.archetypes[newArche.signature] = newArche
		ecs.LinkArchetype(newArche)
		outputArche = newArche
	} else {
		outputArche = exsistingArche
	}
	ecs.entToArche[e] = outputArche
	outputArche.insertEntity(e)
	outputArche.entities = append(outputArche.entities, e)

	return outputArche
}

func Query[T any](ecs *ecs, e entity) T {
	tType := reflect.TypeOf((*T)(nil)).Elem()
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

func NewEcs() *ecs {
	return &ecs{
		archetypes:   map[uint64]*archetype{},
		entToArche:   map[entity]*archetype{},
		componentIds: map[reflect.Type]uint64{},
		nextCompID:   0,
	}
}

func (ecs *ecs) NewEntity(a *archetype) entity {
	e := newEntity()
	a.entities = append(a.entities, e)
	a.insertEntity(e)
	ecs.entToArche[e] = a
	return e
}

// Adds archetype to ecs's archetypes map and make archetype composition read-only
func (ecs *ecs) LinkArchetype(a *archetype) {
	ecs.archetypes[a.signature] = a
	a.mutableComposition = false
}

// func RemoveComponent[T any](ecs *ecs, e entity) *archetype {
// 	a := ecs.entToArche[e]
// 	entityComps := make([]any, 0)
// 	compTypes := make([]reflect.Type, 0)
// 	for _, store := range a.compStores {
// 		entityComps = append(entityComps, store.get(a.entityIdx[e]))
// 		compTypes = append(compTypes, store.storeType())
// 		store.remove(e, a)
// 	}
// 	newArche := NewArchetype()
//
// }
