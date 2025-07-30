package main

type CompId uint64

type CompLabel string

type CompData interface {
	Id() CompId
	SetId(CompId)
}

type Component struct {
	CompLabel
	CompData
}

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
