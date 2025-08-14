package main

import (
	"testing"
)

func TestNewEcsSingleComp(t *testing.T) {
	attrComp := WithAttributes(true)
	ecs := NewEcs()
	register[Attributes](ecs)
	a := NewArchetype()
	With[Attributes](ecs, a)
	NewEcs().EmbedArchetype(a)
	e := ecs.NewEntity(a)
	MutateEntity(ecs, e, attrComp)
	queriedComp := Query[Attributes](ecs, e)
	if queriedComp != attrComp {
		t.Fatalf("Queried Comp:\n%v does not match control comp:\n%v", queriedComp, attrComp)
	}

}

func TestMultiEntSingleComp(t *testing.T) {
	t.Skip()
}

func TestSingleEntMultiComp(t *testing.T) {
	t.Skip()
}

func TestSingleArcheMultiComp(t *testing.T) {
	t.Skip()
}

func TestMultiArcheSingleComp(t *testing.T) {
	t.Skip()
}

func TestMultiArcheMultiEnt(t *testing.T) {
	t.Skip()
}

func TestInvalidCompLookup(t *testing.T) {
	t.Skip()
}

func TestEntityRemoval(t *testing.T) {
	t.Skip()
}
