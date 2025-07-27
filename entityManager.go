package main

import (
	"fmt"
)

type ecs struct {
	Archetypes map[uint64]*archetype
	EntToArche map[Entity]*archetype //What archetype an entity belongs to
}

func (ecs *ecs) Component(e Entity, compLabel CompLabel) (*Component, error) {
	arche := ecs.EntToArche[e]
	if arche == nil {
		return nil, fmt.Errorf("No matching archetype for entity:%v", e)
	}
	compList, ok := arche.components[compLabel]
	if !ok {
		return nil, fmt.Errorf("Matching entity does not have a matching component of compLabel:%v", compLabel)
	}
	return &(*compList)[arche.entityIdx[e]], nil
}

func (ecs *ecs) Remove(e Entity) error {
	a, ok := ecs.EntToArche[e]
	if !ok {
		return fmt.Errorf("Entity:%v - does not exsist in this ECS", e)
	}
	a.openIdxs = append(a.openIdxs, a.entityIdx[e])
	a.entityIdx[Entity(a.entityIdx[e])] = 0
	ecs.EntToArche[e] = nil
	return nil
}

func (ecs *ecs) AddEntity(comps ...Component) Entity {
	e := NewEntity()

	//Get Comp signatures
	var compSig uint64 = 0
	for _, comp := range comps {
		compSig |= uint64(comp.Id()) //Bitwise OR bitfield signature of components
	}
	//Append/Create archetype
	var matchedArche uint64
	for _, arche := range ecs.Archetypes {
		if arche.signature == compSig {
			matchedArche = arche.signature
		}
	}

	if matchedArche == 0 { //					No matching Archetype
		a := archetype{
			signature:  compSig,
			components: map[CompLabel]*[]Component{},
			entityIdx:  map[Entity]int{},
			nextIdx:    0,
			openIdxs:   []int{},
		}
		a.AssignEntity(e)
		for _, comp := range comps {
			if a.components[comp.CompLabel] == nil {
				a.components[comp.CompLabel] = &[]Component{}
			}
			compSlice := a.components[comp.CompLabel]
			*compSlice = append(*compSlice, comp)
		}
		ecs.EntToArche[e] = &a
		ecs.Archetypes[compSig] = &a
	} else { //									Insert into exsisting Archetype
		a := ecs.Archetypes[matchedArche]
		a.AssignEntity(e)
		for _, comp := range comps {
			compSlice := a.components[comp.CompLabel]
			*compSlice = append(*compSlice, comp)
		}
		ecs.EntToArche[e] = a
	}
	return e
}

func NewECS(comps ...Component) (*ecs, Entity) {
	ecs := ecs{
		Archetypes: map[uint64]*archetype{},
		EntToArche: map[Entity]*archetype{},
	}
	e := NewEntity()

	//Get Comp signatures
	var compSig uint64 = 0
	for _, comp := range comps {
		compSig |= uint64(comp.Id()) //Bitwise OR bitfield signature of components
	}
	//Append/Create archetype
	var matchedArche uint64
	for _, arche := range ecs.Archetypes {
		if arche.signature == compSig {
			matchedArche = arche.signature
		}
	}

	if matchedArche == 0 { //					No matching Archetype
		a := archetype{
			signature:  compSig,
			components: map[CompLabel]*[]Component{},
			entityIdx:  map[Entity]int{},
			nextIdx:    0,
			openIdxs:   []int{},
		}
		a.AssignEntity(e)
		for _, comp := range comps {
			if a.components[comp.CompLabel] == nil {
				a.components[comp.CompLabel] = &[]Component{}
			}
			*a.components[comp.CompLabel] = append(*a.components[comp.CompLabel], comp)
		}

		ecs.EntToArche[e] = &a
		ecs.Archetypes[compSig] = &a
	} else { //									Insert into exsisting Archetype
		a := ecs.Archetypes[matchedArche]
		a.AssignEntity(e)
		for _, comp := range comps {
			*a.components[comp.CompLabel] = append(*a.components[comp.CompLabel], comp)
		}
		ecs.EntToArche[e] = a
	}

	return &ecs, e
}
