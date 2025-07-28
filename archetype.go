package main

//archetypes == list of entities with matching compnents
//Matched by a bitfield signature^

// Do not access nextIdx, call AssignEntity(e Entity)
type archetype struct {
	signature  uint64
	components map[CompLabel]*[]Component
	entityIdx  map[Entity]int //Index into the various component slices
	nextIdx    int            //Next availale index
	openIdxs   []int
}

func (a *archetype) assignEntity(e Entity) {
	if len(a.openIdxs) > 0 {
		a.entityIdx[e] = a.openIdxs[len(a.openIdxs)-1] //Pop from slice
		a.openIdxs = a.openIdxs[:len(a.openIdxs)-1]
	} else {
		a.entityIdx[e] = a.nextIdx
		a.nextIdx++
	}
}

type CompLabel string

type CompData interface {
	Id() CompId
	SetId(CompId)
}

type Component struct {
	CompLabel
	CompData
}
