package main

import (
	"fmt"
)

type compStorer interface {
	add(data any, e entity, a *archetype)
	growTo(idx int)
	delete(e entity, a *archetype)
	get(idx int) compStorer
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

func (cs *compStore[T]) delete(e entity, a *archetype) {
	removalIdx := a.entityIdx[e]
	lastIdx := len(cs.data) - 1
	cs.data[removalIdx] = cs.data[lastIdx]
	cs.data = cs.data[:lastIdx]
}

func (cs *compStore[T]) get(idx int) compStorer {
	return &compStore[T]{
		data: []T{cs.data[idx]},
	}
}
