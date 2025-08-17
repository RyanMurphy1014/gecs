package main

import "reflect"

type archetype struct {
	signature  uint64
	compStores map[uint64]compStorer
	entityIdx  map[entity]int //Index into the various component slices
	entities   []entity       //Stores ordered list of entities. Used for removal
	mutable    bool
}

func NewArchetype() *archetype {
	a := archetype{
		signature:  0,
		compStores: map[uint64]compStorer{},
		entityIdx:  map[entity]int{},
		mutable:    true,
		entities:   []entity{},
	}
	return &a
}

// Sets up the archetype compStore
func With[T any](ecs *ecs, a *archetype) {
	if !a.mutable {
		panic("Gecs: An archetype's composition cannot be changed after being embed into an ECS.")
	}
	compSig := register[T](ecs)
	a.compStores[compSig] = &compStore[T]{
		data: []T{},
	}
	a.signature |= ecs.componentIds[reflect.TypeOf((*T)(nil)).Elem()]
}

// Inserts entity into the next available idx
func (a *archetype) insertEntity(e entity) {
	idx := len(a.entities) + 1
	a.entityIdx[e] = idx
	for _, store := range a.compStores {
		store.growTo(idx)
	}
}
