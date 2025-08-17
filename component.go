package main

import (
	"fmt"
)

type compStorer interface {
	add(data any, e entity, a *archetype)
	growTo(idx int)
	delete(e entity, a *archetype)
	get(idx int) any
}

type compStore[T any] struct {
	data []T
}

func UpdateComp[T any](cs compStore[T], data T, e entity, a *archetype) {
	idx := a.entityIdx[e]
	cStore := cs.data
	fmt.Println(cStore)
	cs.data[idx] = data
}

func (cs *compStore[T]) add(data any, e entity, a *archetype) {
	cs.data = append(cs.data, data.(T))
	a.entityIdx[e] = len(cs.data) - 1
}

func (cs *compStore[T]) growTo(idx int) {
	if int(idx) >= len(cs.data) {
		newSlice := make([]T, idx+1)
		copy(newSlice, cs.data)
		cs.data = newSlice
	}
}

// MUST be done after system has ran. Could lead to swapped entity not being ran
func (cs *compStore[T]) delete(e entity, a *archetype) {
	removalIdx := a.entityIdx[e]
	preOpLen := len(cs.data)

	delete(a.entityIdx, e)
	a.entityIdx[a.entities[len(a.entities)-1]] = removalIdx

	cs.data[removalIdx] = cs.data[preOpLen-1]
	cs.data = cs.data[:preOpLen-1]

	a.entities[removalIdx] = a.entities[preOpLen-1]
	a.entities = a.entities[:preOpLen-1]
}

func (cs *compStore[T]) get(idx int) any {
	return cs.data[idx]
}
