package main

//archetypes == list of entities with matching compnents
//Matched by a bitfield signature^

// Do not access nextIdx, call AssignEntity(e Entity)
type archetype struct {
	signature  uint64
	components map[CompLabel]*[]Component
	entityIdx  map[Entity]int //Index into the various component slices
	nextIdx    int            //Next availale index
	// TODO: Swap/Remove/Defragment indexcies
}

func (a *archetype) AssignEntity(e Entity) {
	a.entityIdx[e] = a.nextIdx
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
