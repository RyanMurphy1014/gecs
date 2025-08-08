package main

import (
	"testing"
)

func TestECS(t *testing.T) {

}

func TestNewEcsSingleComp(t *testing.T) {
	ecs := NewECS()
	Register[Attributes](ecs)
	attr := WithAttributes(true)
	e1 := ecs.AddEntity(attr)
	if comp, _ := QueryEntity[Attributes](e1, ecs); *comp != attr {
		t.Fatal()
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
