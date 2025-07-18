package manager

import (
	"github.com/RyanMurphy1014/gecs/entity"
)

//archetypes == list of entities with matching compnents
//Matched by a bitfield signature^

var archetypeSigs []uint64

type Archetype struct {
	Entities   []entity.Entity
	Components map[string][]any
	EntityIdx  map[entity.Entity]int
}
