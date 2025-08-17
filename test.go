package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
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

	t.Run("Entity with single component", func(t *testing.T) {
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

	t.Run("Entity with multiple components", func(t *testing.T) {
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

	t.Run("Archetype with multiple entities", func(t *testing.T) {
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
	t.Run("Entity with default value", func(t *testing.T) {
		if Query[Attributes](ecs, e4) != WithAttributes(false) {
			t.Fatalf("Corresponding compStore is not initialized")
		}
	})

}

type Attributes struct {
	Strength     int
	Dexterity    int
	Constitution int
	Wisdom       int
	Intelligence int
	Charisma     int
}

func WithAttributes(random bool) Attributes {
	if random == true {
		return Attributes{
			Strength:     RandomStat(),
			Dexterity:    RandomStat(),
			Constitution: RandomStat(),
			Wisdom:       RandomStat(),
			Intelligence: RandomStat(),
			Charisma:     RandomStat(),
		}
	}
	return Attributes{
		Strength:     0,
		Dexterity:    0,
		Constitution: 0,
		Wisdom:       0,
		Intelligence: 0,
		Charisma:     0,
	}
}

type Personality struct {
	Aggression int // 0=Pacifist, 100=Aggressive
	Curiosity  int // 0=Traditionalist, 100=Curious
	Drive      int // 0=Lazy, 100=Industrious
	Empathy    int // 0=Calloused, 100=Empath
	Honesty    int // 0=Liar, 100=Truthful
	Loyalty    int // 0=Spineless, 100=Loyal
	Optimism   int // 0=Pessimistic, 100=Optimistic
	Reason     int // 0=Impulsive, 100=Rational
	Socialness int // 0=Solitary, 100=Gregarious
}

func (p Personality) String() string {
	var sb strings.Builder

	val := reflect.ValueOf(p)
	typ := val.Type()
	sb.WriteString("{")
	for i := range val.NumField() {
		sb.WriteString(fmt.Sprintf("\n\t\t%v : %+v", typ.Field(i).Name, val.Field(i)))
	}
	sb.WriteString(" }")
	return sb.String()
}

func WithPersonality(random bool) Personality {
	if !random {
		return Personality{
			Aggression: 0,
			Curiosity:  0,
			Drive:      0,
			Empathy:    0,
			Honesty:    0,
			Loyalty:    0,
			Optimism:   0,
			Reason:     0,
			Socialness: 0,
		}
	} else {
		return Personality{
			Aggression: RandomStat(),
			Curiosity:  RandomStat(),
			Drive:      RandomStat(),
			Empathy:    RandomStat(),
			Honesty:    RandomStat(),
			Loyalty:    RandomStat(),
			Optimism:   RandomStat(),
			Reason:     RandomStat(),
			Socialness: RandomStat(),
		}
	}
}

func RandomStat() int {
	return rand.Intn(100) + 1
}
