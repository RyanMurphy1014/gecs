package main

//archetypes == list of entities with matching compnents
//Matched by a bitfield signature^

// Do not access nextIdx, call AssignEntity(e Entity)
type Archetype struct {
	Signature  uint64
	Entities   []Entity
	Components map[CompLabel]*[]Component
	EntityIdx  map[Entity]int //Index into the various component slices
	nextIdx    int            //Next availale index
	// TODO: Swap/Remove/Defragment indexcies
}

func (a *Archetype) AssignEntity(e Entity) {
	a.EntityIdx[e] = a.nextIdx
	a.nextIdx++
}

type CompLabel string

type CompData interface {
	Id() CompId
}

type Component struct {
	CompLabel
	CompData
}
