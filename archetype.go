package main

import "reflect"

type archetype struct {
	signature          uint64
	compStores         map[uint64]compStorer
	entityIdx          map[entity]int //Index into compStores. Index is the same in all stores
	entities           []entity       //Stores ordered list of entities. Used for removal
	mutableComposition bool
}

// Initializes archetype's component store for passed in type
func With[T any](ecs *ecs, a *archetype) {
	if !a.mutableComposition {
		panic("Gecs: An archetype's composition cannot be changed after being embed into an ECS.")
	}
	if _, ok := ecs.componentIds[reflect.TypeOf((*T)(nil)).Elem()]; !ok {
		panic("Gecs: Components must first be registered to an ECS with the Register function.")
	}
	compSig := RegisterComp[T](ecs)
	a.compStores[compSig] = &compStore[T]{
		data: []T{},
	}
	a.signature |= ecs.componentIds[reflect.TypeOf((*T)(nil)).Elem()]
}

func NewArchetype() *archetype {
	a := archetype{
		signature:          0,
		compStores:         map[uint64]compStorer{},
		entityIdx:          map[entity]int{},
		mutableComposition: true,
		entities:           []entity{},
	}
	return &a
}

func (a *archetype) insertEntity(e entity) {
	idx := len(a.entities)
	a.entityIdx[e] = idx
	for _, store := range a.compStores {
		store.growTo(idx)
	}
}

func (a *archetype) unlinkEntity(e entity) {
	for _, store := range a.compStores {
		store.delete(e, a)
	}
	delete(a.entityIdx, e)
}
