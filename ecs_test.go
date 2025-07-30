package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewEcsSingleComp(t *testing.T) {
	attr := WithAttributes(true)
	ecs, e := NewECS(attr)
	returnedComp, _ := ecs.Component(e, Attributes_Comp)
	if *returnedComp != attr {
		t.Errorf("Got %v, want %v", *returnedComp, attr)
	}
}

func TestMultiEntSingleComp(t *testing.T) {
	attr1 := WithAttributes(true)
	attr2 := WithAttributes(true)
	ecs, e1 := NewECS(attr1)
	e2 := ecs.AddEntity(attr2)
	returnComp1, _ := ecs.Component(e1, Attributes_Comp)
	if *returnComp1 != attr1 {
		t.Errorf("Got %v, want %v", *returnComp1, attr1)
	}
	returnComp2, _ := ecs.Component(e2, Attributes_Comp)
	if *returnComp2 != attr2 {
		t.Errorf("Got %v, want %v", *returnComp2, attr2)
	}
}

func TestSingleEntMultiComp(t *testing.T) {
	attr := WithAttributes(true)
	pers := WithPersonality(Random)
	ecs, e := NewECS(pers, attr)
	returnComp1, _ := ecs.Component(e, Attributes_Comp)
	if *returnComp1 != attr {
		t.Errorf("Got %v, want %v", *returnComp1, attr)
	}
	returnComp2, _ := ecs.Component(e, Personality_Comp)
	if *returnComp2 != pers {
		t.Errorf("Got %v, want %v", *returnComp2, pers)
	}
}

func TestSingleArcheMultiComp(t *testing.T) {
	attr1 := WithAttributes(true)
	pers1 := WithPersonality(Random)
	attr2 := WithAttributes(true)
	pers2 := WithPersonality(Random)
	ecs, e1 := NewECS(attr1, pers1)
	e2 := ecs.AddEntity(attr2, pers2)

	returnAttr1, _ := ecs.Component(e1, Attributes_Comp)
	if *returnAttr1 != attr1 {
		t.Errorf("Got %v, want %v", *returnAttr1, attr1)
	}
	returnPers1, _ := ecs.Component(e1, Personality_Comp)
	if *returnPers1 != pers1 {
		t.Errorf("Got %v, want %v", *returnPers1, pers1)
	}

	returnAttr2, _ := ecs.Component(e2, Attributes_Comp)
	if *returnAttr2 != attr2 {
		t.Errorf("Got %v, want %v", *returnAttr2, attr2)
	}
	returnPers2, _ := ecs.Component(e2, Personality_Comp)
	if *returnPers2 != pers2 {
		t.Errorf("Got %v, want %v", *returnPers2, pers2)
	}
}

func TestMultiArcheSingleComp(t *testing.T) {
	attr := WithAttributes(true)
	pers := WithPersonality(Random)
	ecs, e1 := NewECS(attr)
	e2 := ecs.AddEntity(pers)

	returnAttr, _ := ecs.Component(e1, Attributes_Comp)
	if *returnAttr != attr {
		t.Errorf("Got %v, want %v", *returnAttr, attr)
	}
	returnPers, _ := ecs.Component(e2, Personality_Comp)
	if *returnPers != pers {
		t.Errorf("Got %v, want %v", *returnPers, pers)
	}
}

func TestMultiArcheMultiEnt(t *testing.T) {
	attr1 := WithAttributes(true)
	attr2 := WithAttributes(true)
	pers1 := WithPersonality(Random)
	pers2 := WithPersonality(Random)

	ecs, eAttr1 := NewECS(attr1)
	eAttr2 := ecs.AddEntity(attr2)

	ePers1 := ecs.AddEntity(pers1)
	ePers2 := ecs.AddEntity(pers2)

	returnAttr1, _ := ecs.Component(eAttr1, Attributes_Comp)
	if *returnAttr1 != attr1 {
		t.Errorf("Got %v, want %v", *returnAttr1, attr1)
	}
	returnAttr2, _ := ecs.Component(eAttr2, Attributes_Comp)
	if *returnAttr2 != attr2 {
		t.Errorf("Got %v, want %v", *returnAttr2, attr2)
	}

	returnPers1, _ := ecs.Component(ePers1, Personality_Comp)
	if *returnPers1 != pers1 {
		t.Errorf("Got %v, want %v", *returnPers1, pers1)
	}

	returnPers2, _ := ecs.Component(ePers2, Personality_Comp)
	if *returnPers2 != pers2 {
		t.Errorf("Got %v, want %v", *returnPers2, pers2)
	}
}

func TestInvalidCompLookup(t *testing.T) {
	attr := WithAttributes(true)
	ecs, e := NewECS(attr)
	_, err := ecs.Component(e, Personality_Comp)
	if err != nil {
		t.Log(err)
	} else {
		t.Fatal("Invalid component lookup did not error")
	}
}

func TestEntityRemoval(t *testing.T) {
	attr := WithAttributes(true)
	ecs, e1 := NewECS(attr)
	e1Idx := ecs.EntToArche[e1].entityIdx[e1]

	ecs.Remove(e1)

	t.Run("Entity Removal", func(t *testing.T) {
		comp, err := ecs.Component(e1, Attributes_Comp)
		if err == nil {
			t.Log(comp)
			t.Fatal("Invalid Component lookup did not error")
		}
	})
	e2 := ecs.AddEntity(WithAttributes(true))
	e2Idx := ecs.EntToArche[e2].entityIdx[e2]
	t.Run("Entity Insertion", func(t *testing.T) {
		if e1Idx != e2Idx {
			var sb strings.Builder
			for ent, idx := range ecs.EntToArche[e2].entityIdx {
				sb.WriteString(fmt.Sprintf("ent:%v - Index:%v\n", ent, idx))
			}
			t.Log(sb.String())
			t.Fatalf("First archetype IDX:%v  is not resued for Second Entity. Second IDX:%v", e1Idx, e2Idx)
		}
	})
	t.Run("Mutltiple Empty Slots", func(t *testing.T) {
		e3 := ecs.AddEntity(WithAttributes(true))
		e3Idx := ecs.entityIdx(e3)
		e4 := ecs.AddEntity(WithAttributes(true))
		e4Idx := ecs.entityIdx(e4)
		ecs.Remove(e3)
		ecs.Remove(e4)
		e5 := ecs.AddEntity(WithAttributes(true))
		e5Idx := ecs.entityIdx(e5)
		e6 := ecs.AddEntity(WithAttributes(true))
		e6Idx := ecs.entityIdx(e6)

		if e5Idx != e4Idx {
			t.Fatalf("e5Idx:%v != e4Idx:%v", e5Idx, e4Idx)
		}
		if e6Idx != e3Idx {
			t.Fatalf("e6Idx:%v != e3Idx:%v", e6Idx, e3Idx)
		}
	})
}
