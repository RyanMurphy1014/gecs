package main

import (
	"testing"
)

func TestQuerying(t *testing.T) {

	ecs := NewEcs()
	vectorId := RegisterComp[vector](ecs)
	RegisterComp[location](ecs)

	vectorArche := NewArchetype()
	With[vector](ecs, vectorArche)
	ecs.LinkArchetype(vectorArche)

	locationArche := NewArchetype()
	With[location](ecs, locationArche)
	ecs.LinkArchetype(locationArche)

	e := ecs.NewEntity(vectorArche)

	vector1 := vector{}
	UpdateEntity(ecs, e, vector1)

	t.Run("Entity with single component", func(t *testing.T) {
		queriedComp := Query[vector](ecs, e)
		if queriedComp != vector1 {
			t.Fatalf("Queried Comp:\n%v does not match control comp:\n%v", queriedComp, vector1)
		}
	})

	vecAndLocArche := NewArchetype()
	With[vector](ecs, vecAndLocArche)
	With[location](ecs, vecAndLocArche)
	ecs.LinkArchetype(vecAndLocArche)

	e2 := ecs.NewEntity(vecAndLocArche)

	location1 := location{}
	UpdateEntity(ecs, e2, vector1)
	UpdateEntity(ecs, e2, location1)

	t.Run("Entity with multiple components", func(t *testing.T) {
		queriedVec := Query[vector](ecs, e2)
		queriedLoc := Query[location](ecs, e2)
		if queriedVec != vector1 {
			t.Fatalf("Comps do not match. Got:%v - Want:%v", queriedVec, vector1)
		}
		if queriedLoc != location1 {
			t.Fatalf("Comps do not match. Got:%v - Want:%v", queriedLoc, location1)
		}
	})

	e3 := ecs.NewEntity(vecAndLocArche)
	vector2 := vector{2, 2, 2}
	location2 := location{5, 5}
	UpdateEntity(ecs, e3, vector2)
	UpdateEntity(ecs, e3, location2)

	t.Run("Archetype with multiple entities", func(t *testing.T) {
		queriedVec := Query[vector](ecs, e3)
		queriedLoc := Query[location](ecs, e3)
		if queriedVec != vector2 {
			t.Fatalf("Comps do not match. Got:%v - Want:%v", queriedVec, vector2)
		}
		if queriedLoc != location2 {
			t.Fatalf("Comps do not match. Got:%v - Want:%v", queriedLoc, location2)
		}
	})

	t.Run("Removed Component", func(t *testing.T) {
		newlyCreatedArche := DeleteComponent(ecs, e3, vectorId)
		querriedLocation := Query[location](ecs, e3)
		if querriedLocation != location2 {
			t.Fatalf("Comps do not match. Got:%v - Want:%v", querriedLocation, location2)
		}
		t.Log(newlyCreatedArche)
	})

	e4 := ecs.NewEntity(vectorArche)
	t.Run("Entity with default value", func(t *testing.T) {
		if Query[vector](ecs, e4) != (vector{}) {
			t.Fatalf("Corresponding compStore is not initialized")
		}
	})

	e5 := ecs.NewEntity(vectorArche)
	nonDefaultVec := vector{
		x: 5,
		y: 15,
		z: 20,
	}
	UpdateEntity(ecs, e5, nonDefaultVec)
	t.Run("Entity with non default value", func(t *testing.T) {
		if Query[vector](ecs, e5) != nonDefaultVec {
			t.Fatalf("Comps do not match")
		}
	})
}

type vector struct {
	x, y, z float64
}

type location struct {
	x, y float64
}
