package main

import "testing"

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
	ecs, e := NewECS(attr, pers)
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
