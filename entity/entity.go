package entity

var curId uint64 = 0

func newId() uint64 {
	newId := curId
	curId++
	if newId > curId {
		panic("Entity ID's have overflown")
	}
	return newId
}

type Entity uint64

func NewEntity(c ...Component) Entity {
	e := Entity(newId())
	//Generate a bitfield signature to find/create an archetype by looping over passed comps
	return e
}
