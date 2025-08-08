package main

import (
	"errors"
	"fmt"
	"reflect"
)

type ecs struct {
	archetypes    map[uint64]*archetype
	entToArche    map[entity]*archetype //What archetype an entity belongs to
	archeTypeSigs []uint64
	componentIds  map[reflect.Type]uint64
	nextCompID    uint64
}

func QueryEntity[T any](e entity, ecs *ecs) (*T, error) {
	a, ok := ecs.entToArche[e]
	if !ok {
		return nil, errors.New("This entity does not belong to the ecs")
	}
	idx := a.entityIdx[e]
	compId := Register[T](ecs)
	store, ok := a.compStore[compId]
	if !ok {
		return nil, errors.New("Entity does not have the coresponding component")
	}
	typedStore := store.(*compStore[T])
	return &typedStore.data[idx], nil
}

func (ecs *ecs) Remove(e entity) error {
	a, ok := ecs.entToArche[e]
	if !ok {
		return fmt.Errorf("Entity:%v - does not exsist in this ECS", e)
	}
	a.openIdxs = append(a.openIdxs, a.entityIdx[e])
	a.entityIdx[entity(a.entityIdx[e])] = 0
	ecs.entToArche[e] = nil
	return nil
}

func (ecs *ecs) AddEntity(compSet ...any) entity {
	e := newEntity()

	//Get Comp signatures
	var compSetSig uint64 = 0
	for _, comp := range compSet {
		compType := reflect.TypeOf(comp)
		compId, registered := ecs.componentIds[compType]
		if !registered {
			panic("Attempting to add unregistered component of type: " + compType.Name())
		}
		compSetSig |= compId
	}

	//Find matching arche
	arche, ok := ecs.archetypes[compSetSig]
	if !ok { //No match. Must create archetype
		a := &archetype{
			signature: compSetSig,
			compStore: map[uint64]compStorer{},
			entityIdx: map[entity]uint{},
			nextIdx:   0,
			openIdxs:  []uint{},
		}

		for _, comp := range compSet {
			compType := reflect.TypeOf(comp)
			a.compStore[ecs.componentIds[compType]] = &compStore[any]{
				data: []any{},
			}
		}

		arche = a
		ecs.archeTypeSigs = append(ecs.archeTypeSigs, compSetSig)
		ecs.archetypes[compSetSig] = a
	} else {
		arche = ecs.archetypes[arche.signature]
	}

	arche.linkEntity(e)
	ecs.entToArche[e] = arche

	//Store
	for _, comp := range compSet {
		arche.compStore[compSetSig].add(comp, e, arche)
	}
	return e
}

func NewECS() *ecs {
	return &ecs{
		archetypes:   map[uint64]*archetype{},
		entToArche:   map[entity]*archetype{},
		componentIds: map[reflect.Type]uint64{},
		nextCompID:   0,
	}
}

func (ecs *ecs) entityIdx(e entity) uint {
	return ecs.entToArche[e].entityIdx[e]
}
