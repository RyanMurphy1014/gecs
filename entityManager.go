package main

import (
	"fmt"
)

type ecs struct {
	Archetypes map[uint64]*Archetype
	EntToArche map[Entity]*Archetype //What archetype an entity belongs to
}

func (ecs *ecs) Component(e Entity, compLabel CompLabel) (*Component, error) {
	arche := ecs.EntToArche[e]
	if arche == nil {
		return nil, fmt.Errorf("No matching archetype for entity:%v", e)
	}
	compList, ok := arche.Components[compLabel]
	if !ok {
		return nil, fmt.Errorf("Matching entity does not have a matching component of compLabel:%v", compLabel)
	}
	return &(*compList)[arche.EntityIdx[e]], nil
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
		if arche.Signature == compSig {
			matchedArche = arche.Signature
		}
	}

	if matchedArche == 0 { //					No matching Archetype
		a := Archetype{
			Signature:  compSig,
			Entities:   []Entity{},
			Components: map[CompLabel]*[]Component{},
			EntityIdx:  map[Entity]int{},
			nextIdx:    0,
		}
		a.Entities = append(a.Entities, e)
		a.AssignEntity(e)
		for _, comp := range comps {
			if a.Components[comp.CompLabel] == nil {
				a.Components[comp.CompLabel] = &[]Component{}
			}
			compSlice := a.Components[comp.CompLabel]
			*compSlice = append(*compSlice, comp)
		}
		ecs.EntToArche[e] = &a
		ecs.Archetypes[compSig] = &a
	} else { //									Insert into exsisting Archetype
		a := ecs.Archetypes[matchedArche]
		a.Entities = append(a.Entities, e)
		a.AssignEntity(e)
		for _, comp := range comps {
			compSlice := a.Components[comp.CompLabel]
			*compSlice = append(*compSlice, comp)
		}
		ecs.EntToArche[e] = a
	}
	return e
}

func NewECS(comps ...Component) (*ecs, Entity) {
	ecs := ecs{
		Archetypes: map[uint64]*Archetype{},
		EntToArche: map[Entity]*Archetype{},
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
		if arche.Signature == compSig {
			matchedArche = arche.Signature
		}
	}

	if matchedArche == 0 { //					No matching Archetype
		a := Archetype{
			Signature:  compSig,
			Entities:   []Entity{},
			Components: map[CompLabel]*[]Component{},
			EntityIdx:  map[Entity]int{},
			nextIdx:    0,
		}
		a.Entities = append(a.Entities, e)
		a.AssignEntity(e)
		for _, comp := range comps {
			if a.Components[comp.CompLabel] == nil {
				a.Components[comp.CompLabel] = &[]Component{}
			}
			*a.Components[comp.CompLabel] = append(*a.Components[comp.CompLabel], comp)
		}

		ecs.EntToArche[e] = &a
		ecs.Archetypes[compSig] = &a
	} else { //									Insert into exsisting Archetype
		a := ecs.Archetypes[matchedArche]
		a.Entities = append(a.Entities, e)
		a.AssignEntity(e)
		for _, comp := range comps {
			*a.Components[comp.CompLabel] = append(*a.Components[comp.CompLabel], comp)
		}
		ecs.EntToArche[e] = a
	}

	return &ecs, e
}
