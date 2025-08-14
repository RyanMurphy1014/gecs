package main

import "fmt"

type compStorer interface {
	add(data any, e entity, a *archetype)
	grow(idx uint)
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
	idx := a.entityIdx[e]
	cStore := cs.data
	fmt.Println(cStore)
	cs.data[idx] = data.(T)
}

func (cs *compStore[T]) grow(idx uint) {
	if int(idx) >= len(cs.data) {
		newSlice := make([]T, idx+1)
		copy(newSlice, cs.data)
		cs.data = newSlice
	}
}
