package components

import (
	"math/rand"

	"github.com/RyanMurphy1014/LoreWeaver/entity"
)

type CompLabel string
type CompId uint64

func RandomStat() int {
	return rand.Intn(100) + 1
}

// Must have struct created and label set before creating CompOptionGenerator()
// Injects the Component into the Entity
func CompOptionGenerator(comp any, label CompLabel) entity.CompBuilder {
	return func(e *entity.Entity) {
		e.Components[label] = comp
	}
}
