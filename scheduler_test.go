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
	SetEntity(schedEcs, velocityEnt, testVelocity)
	return velocityEnt
}

func TestExecuteSystems(t *testing.T) {
	ent := setupWorld()
	increaseSysId := newSysId()
	increase5 := System{
		op: func(ecs *ecs) {
			vel := SystemRead[velocity](ecs, ent, increaseSysId)
			vel.x += 5
			vel.y += 5
			vel.z += 5
			SystemWrite(ecs, ent, vel, increaseSysId)
		},
	}
	RegisterSystem(schedEcs, increase5, increaseSysId)

	t.Log(schedEcs.TickSystems())

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

func TestConflictingSystems(t *testing.T) {
	ecs2 := NewEcs()
	vel := velocity{
		x: 0,
		y: 0,
		z: 0,
	}

	type rotation struct {
		angle float32
	}

	rot := rotation{
		angle: 180,
	}

	RegisterComp[velocity](ecs2)
	RegisterComp[rotation](ecs2)
	velAndRotArche := NewArchetype().AddComponents(ecs2, With[velocity](), With[rotation]())
	ecs2.AddArchetype(velAndRotArche)
	ent := ecs2.NewEntity(velAndRotArche)
	SetEntity(ecs2, ent, rot)
	SetEntity(ecs2, ent, vel)

	rotateCounterClockSysId := newSysId()
	rotateCounterClockSys := System{
		op: func(ecs *ecs) {
			val := SystemRead[rotation](ecs2, ent, rotateCounterClockSysId)
			val.angle += 5
			if val.angle > 360 {
				val.angle = 0
			}
			SystemWrite(ecs2, ent, val, rotateCounterClockSysId)
		},
	}

	RegisterSystem(ecs2, rotateCounterClockSys, rotateCounterClockSysId)

	rotateClockwiseSysId := newSysId()
	rotateClockwiseSys := System{
		op: func(ecs *ecs) {
			val := SystemRead[rotation](ecs2, ent, rotateClockwiseSysId)
			val.angle--
			if val.angle < 0 {
				val.angle = 0
			}
			SystemWrite(ecs2, ent, val, rotateClockwiseSysId)
		},
	}
	RegisterSystem(ecs2, rotateClockwiseSys, rotateClockwiseSysId)

	incVelocitySysId := newSysId()
	incVelocitySys := System{
		op: func(ecs *ecs) {
			vel := SystemRead[velocity](ecs2, ent, incVelocitySysId)
			vel.x += 6
			vel.y += 5
			vel.z += 5
			SystemWrite(ecs2, ent, vel, incVelocitySysId)
		},
	}
	RegisterSystem(ecs2, incVelocitySys, incVelocitySysId)

	// TODO: Investigate why the systems are not both running. Probably a caching issue.
	//Rotate Counter-Clockwise Sys ID = 2
	//Rotate Clockwise Sys ID = 3
	t.Log(ecs2.TickSystems())
	t.Log(Query[rotation](ecs2, ent))
	t.Log(ecs2.batches)
	t.Log(Query[velocity](ecs2, ent))
	t.Log(ecs2.TickSystems())
}
