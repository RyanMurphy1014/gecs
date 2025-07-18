package components

import (
	"github.com/RyanMurphy1014/LoreWeaver/entity"
)

type Creature struct {
	CompLabel
}

func WithCreature(prefab Creature) entity.CompBuilder {
	return CompOptionGenerator(prefab, prefab.CompLabel)
}

var Human_Creature CompLabel = "Human"
var Creature_Human = Creature{CompLabel: Human_Creature}
