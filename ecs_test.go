package main

import (
	"testing"
)

func TestQuerying(t *testing.T) {

	ecs := NewEcs()
	register[Attributes](ecs)

	attributesArche := NewArchetype()
	With[Attributes](ecs, attributesArche)
	ecs.EmbedArchetype(attributesArche)

	e := ecs.NewEntity(attributesArche)

	attrComp := WithAttributes(true)
	MutateEntity(ecs, e, attrComp)

	t.Run("Querying an entity with a single component", func(t *testing.T) {
		queriedComp := Query[Attributes](ecs, e)
		if queriedComp != attrComp {
			t.Fatalf("Queried Comp:\n%v does not match control comp:\n%v", queriedComp, attrComp)
		}
	})

	attrAndPersonalityArche := NewArchetype()
	With[Attributes](ecs, attrAndPersonalityArche)
	With[Personality](ecs, attrAndPersonalityArche)
	ecs.EmbedArchetype(attrAndPersonalityArche)

	e2 := ecs.NewEntity(attrAndPersonalityArche)

	personalityComp := WithPersonality(true)
	MutateEntity(ecs, e2, attrComp)
	MutateEntity(ecs, e2, personalityComp)

	t.Run("Querying an entity with multiple components", func(t *testing.T) {
		queriedAttr := Query[Attributes](ecs, e2)
		queriedPersonality := Query[Personality](ecs, e2)
		if queriedAttr != attrComp {
			t.Fatalf("Attributes does not match")
		}
		if queriedPersonality != personalityComp {
			t.Fatalf("Personality does not match")
		}
	})

	e3 := ecs.NewEntity(attrAndPersonalityArche)
	attrComp2 := WithAttributes(true)
	personalityComp2 := WithPersonality(true)
	MutateEntity(ecs, e3, attrComp2)
	MutateEntity(ecs, e3, personalityComp2)

	t.Run("Querying an archetype with multiple entities", func(t *testing.T) {
		queriedAttr := Query[Attributes](ecs, e3)
		queriedPersonality := Query[Personality](ecs, e3)
		if queriedAttr != attrComp2 {
			t.Fatalf("Attributes does not match")
		}
		if queriedPersonality != personalityComp2 {
			t.Fatalf("Personlaity does not match")
		}
	})

	e4 := ecs.NewEntity(attributesArche)
	t.Run("Querying an entity with default value", func(t *testing.T) {
		if Query[Attributes](ecs, e4) != WithAttributes(false) {
			t.Fatalf("Corresponding compStore is not initialized")
		}
	})

}
