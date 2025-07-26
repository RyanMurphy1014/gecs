package main

type CompId uint64

var currId CompId = 1
var registerdComps = make(map[string]CompId)

func nextCompId(label string) CompId {
	if reggedId, ok := registerdComps[label]; ok {
		return reggedId
	}

	nextId := currId
	currId++
	if nextId > currId {
		panic("Component ID's have overflown")
	}
	registerdComps[label] = nextId
	return nextId
}
