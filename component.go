package main

import "reflect"

type compStorer interface {
	add(data any, e entity, a *archetype)
}

type compStore[T any] struct {
	data []T
}

func (cs *compStore[T]) add(data any, e entity, a *archetype) {
	cs.data[a.entityIdx[e]] = data.(T)
}

func Register[T any](ecs *ecs) uint64 {
	key := reflect.TypeOf((*T)(nil)).Elem()
	if compId, ok := ecs.componentIds[key]; ok {
		return compId
	}

	var compId uint64 = 1 << ecs.nextCompID
	ecs.componentIds[key] = compId
	ecs.nextCompID++
	return compId
}
