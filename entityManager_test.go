package main

import "testing"

func TestNewECS(t *testing.T) {
	attr := WithAttributes(true)
	ecs, e := NewECS(attr)
	if ecs.EntToArche[e].Components[Attributes_Comp][0] != attr {
		t.Errorf("Got %v, want %v", ecs.EntToArche[e].Components[Attributes_Comp][0], attr)
	}
}
