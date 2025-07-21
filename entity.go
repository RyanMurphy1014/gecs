package main

var curId uint64 = 0

func newEntId() uint64 {
	newId := curId
	curId++
	if newId > curId {
		panic("Entity ID's have overflown")
	}
	return newId
}

type Entity uint64

func NewEntity() Entity {
	e := Entity(newEntId())
	return e
}
