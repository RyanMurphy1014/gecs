package main

type archetype struct {
	signature uint64
	compStore map[uint64]compStorer
	entityIdx map[entity]uint //Index into the various component slices
	nextIdx   uint            //Next availale index
	openIdxs  []uint
}

func (a *archetype) linkEntity(e entity) {
	if len(a.openIdxs) > 0 {
		a.entityIdx[e] = a.openIdxs[len(a.openIdxs)-1] //Pop from slice
		a.openIdxs = a.openIdxs[:len(a.openIdxs)-1]
	} else {
		a.entityIdx[e] = a.nextIdx
		a.nextIdx++
	}
}
