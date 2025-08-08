package main

var curId uint64 = 0

func newId() uint64 {
	newId := curId
	curId++
	if newId > curId {
		panic("Entity ID's have overflown")
	}
	return newId
}

type entity uint64

func newEntity() entity {
	e := entity(newId())
	return e
}
