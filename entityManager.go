package main

type ecs struct {
	Archetypes map[uint64]*Archetype
	EntToArche map[Entity]*Archetype //What archetype an entity belongs to
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
			Components: map[CompLabel][]Component{},
			EntityIdx:  map[Entity]int{},
			nextIdx:    0,
		}
		a.Entities = append(a.Entities, e)
		a.AssignEntity(e)
		for _, comp := range comps {
			compSlice := a.Components[comp.CompLabel]
			compSlice = append(compSlice, comp)
		}

		ecs.Archetypes[compSig] = &a
	} else { //									Insert into exsisting Archetype
		a := ecs.Archetypes[matchedArche]
		a.Entities = append(a.Entities, e)
		a.AssignEntity(e)
		for _, comp := range comps {
			compSlice := a.Components[comp.CompLabel]
			compSlice = append(compSlice, comp)
		}
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
			Components: map[CompLabel][]Component{},
			EntityIdx:  map[Entity]int{},
			nextIdx:    0,
		}
		a.Entities = append(a.Entities, e)
		a.AssignEntity(e)
		for _, comp := range comps {
			a.Components[comp.CompLabel] = append(a.Components[comp.CompLabel], comp)
		}

		ecs.EntToArche[e] = &a
		ecs.Archetypes[compSig] = &a
	} else { //									Insert into exsisting Archetype
		a := ecs.Archetypes[matchedArche]
		a.Entities = append(a.Entities, e)
		a.AssignEntity(e)
		for _, comp := range comps {
			a.Components[comp.CompLabel] = append(a.Components[comp.CompLabel], comp)
		}
		ecs.EntToArche[e] = a
	}

	return &ecs, e
}
