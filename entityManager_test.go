package main

import "testing"

func TestNewEcsSingleComp(t *testing.T) {
	attr := WithAttributes(true)
	ecs, e := NewECS(attr)
	returnedComp := *ecs.Component(e, Attributes_Comp)
	if returnedComp != attr {
		t.Errorf("Got %v, want %v", returnedComp, attr)
	}
}

func TestMultiEntSingleComp(t *testing.T) {
	attr1 := WithAttributes(true)
	attr2 := WithAttributes(true)
	ecs, e1 := NewECS(attr1)
	e2 := ecs.AddEntity(attr2)
	if *ecs.Component(e1, Attributes_Comp) != attr1 {
		t.Errorf("Got %v, want %v", *ecs.Component(e1, Attributes_Comp), attr1)
	}
	if *ecs.Component(e2, Attributes_Comp) != attr2 {
		t.Errorf("Got %v, want %v", *ecs.Component(e2, Attributes_Comp), attr2)
	}
}
