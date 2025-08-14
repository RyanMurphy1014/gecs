package main

import "reflect"

type archetype struct {
	signature uint64
	compStore map[uint64]compStorer
	entityIdx map[entity]uint //Index into the various component slices
	nextIdx   uint            //Next availale index
	openIdxs  []uint
	mutable   bool
}

func NewArchetype() *archetype {
	a := archetype{
		signature: 0,
		compStore: map[uint64]compStorer{},
		entityIdx: map[entity]uint{},
		nextIdx:   0,
		openIdxs:  []uint{},
		mutable:   true,
	}
	return &a
}

// Sets up the archetype compStore
func With[T any](ecs *ecs, a *archetype) {
	if !a.mutable {
		panic("Gecs: An archetype's composition cannot be changed after being embed into an ECS.")
	}
	compSig := register[T](ecs)
	a.compStore[compSig] = &compStore[T]{
		data: []T{},
	}
	a.signature |= ecs.componentIds[reflect.TypeOf((*T)(nil)).Elem()]
}

// Inserts entity into the next available idx
func (a *archetype) insertEntity(e entity) {
	var idx uint
	if len(a.openIdxs) > 0 {
		idx = a.openIdxs[len(a.openIdxs)-1] //Pop from slice
		a.openIdxs = a.openIdxs[:len(a.openIdxs)-1]
	} else {
		idx = a.nextIdx
		a.nextIdx++
	}
	a.entityIdx[e] = idx
	for _, store := range a.compStore {
		store.grow(idx)
	}
}
