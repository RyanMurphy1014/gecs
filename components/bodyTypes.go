package components

import (
	"github.com/RyanMurphy1014/LoreWeaver/entity"
)

var Humanoid_BodyType CompLabel = "BodyType_Humanoid"

type BodyType struct {
	CompLabel
	heads int
	eyes  int
	arms  int
	legs  int
}

func WithBodyType(prefab BodyType) entity.CompBuilder {
	return CompOptionGenerator(prefab, prefab.CompLabel)
}

var BodyType_Humanoid = BodyType{
	heads:     1,
	eyes:      2,
	arms:      2,
	legs:      2,
	CompLabel: Humanoid_BodyType,
}
