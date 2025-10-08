package main

import (
	"testing"
)

var schedEcs *ecs = NewEcs()

type velocity struct {
	x, y, z float64
}

func setupWorld() entity {
	testVelocity := velocity{
		x: 5,
		y: 10,
		z: 15,
	}
	RegisterComp[velocity](schedEcs)
	velocityArche := NewArchetype().AddComponents(schedEcs, With[velocity]())
	schedEcs.AddArchetype(velocityArche)
	velocityEnt := schedEcs.NewEntity(velocityArche)
	UpdateEntity(schedEcs, velocityEnt, testVelocity)
	return velocityEnt
}

func TestExecuteSystems(t *testing.T) {
	ent := setupWorld()
	increaseSysId := newSysId()
	increase5 := SimpleSys{
		op: func(ecs *ecs) {
			vel := SystemRead[velocity](ecs, ent, increaseSysId)
			vel.x += 5
			vel.y += 5
			vel.z += 5
			SystemWrite(ecs, ent, vel, increaseSysId)
		},
	}
	RegisterSystem(schedEcs, increase5, increaseSysId)

	t.Log(schedEcs.ExecuteSystems())

	gotVel := Query[velocity](schedEcs, ent)
	wantVel := velocity{
		x: 10,
		y: 15,
		z: 20,
	}
	if gotVel != wantVel {
		t.Fatalf("Got:%+v \n Want:%+v", gotVel, wantVel)
	}
}
