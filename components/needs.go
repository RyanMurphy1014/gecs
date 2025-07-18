package components

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/RyanMurphy1014/LoreWeaver/entity"
)

type Needs struct {
	CompLabel
	basicNeeds
	safetyNeeds
	socialNeeds
	esteemNeeds
	selfActNeeds
}

type basicNeeds struct {
	Hunger float64
	Thirst float64
	Sleep  float64
}

func (n *basicNeeds) decayBasicNeeds() {
	n.Hunger -= 0.00001929 //Stavation in 2 months
	n.Thirst -= 0.000386   //Thirst in 3 days death at 0
	n.Sleep -= 0.000386    //3 days till 0. Significant debuffs
}

type safetyNeeds struct {
	PhysicalSecurity float64
	ResourceSecurity float64
	Health           float64
	Shelter          float64
}

func (n *safetyNeeds) decaySafetyNeeds() {
	n.PhysicalSecurity -= 0.00011
	n.ResourceSecurity -= 0.000047
	n.Health -= 0.0000257
	n.Shelter -= 0.00011
}

type socialNeeds struct {
	Social     float64
	Affection  float64
	Acceptance float64
	Family     float64
}

func (n *socialNeeds) decaySocialNeeds() {
	n.Social -= 0.000066
	n.Affection -= 0.0000257
	n.Acceptance -= 0.000066
	n.Family -= 0.0000154
}

type esteemNeeds struct {
	SelfEsteem float64
	Respect    float64
	Mastery    float64
}

func (n *esteemNeeds) decayEsteemNeeds() {
	n.SelfEsteem -= 0.0000257
	n.Respect -= 0.0000257
	n.Mastery -= 0.0000154
}

type selfActNeeds struct {
	PersonalGrowth float64
	Creativity     float64
	Purpose        float64
	Exploration    float64
}

func (n *selfActNeeds) decaySelfActNeeds() {
	n.PersonalGrowth -= 0.00000857
	n.Creativity -= 0.00000857
	n.Purpose -= 0.00000428
	n.Exploration -= 0.00000857
}

func (n *Needs) Decay() {
	n.decayBasicNeeds()
	n.decaySafetyNeeds()
	n.decaySocialNeeds()
	n.decayEsteemNeeds()
	n.decaySelfActNeeds()
}

func (n Needs) String() string {
	var sb strings.Builder

	val := reflect.ValueOf(n)
	typ := val.Type()
	sb.WriteString("{")
	for i := range val.NumField() {
		sb.WriteString(fmt.Sprintf("\n\t\t%v : %+v", typ.Field(i).Name, val.Field(i)))
	}
	sb.WriteString(" }")
	return sb.String()
}

var Needs_Comp CompLabel = "Needs"

// Needs to also have a personality to drive the needs
func WithNeeds() entity.CompBuilder {
	n := Needs{
		basicNeeds: basicNeeds{
			Hunger: 100,
			Thirst: 100,
			Sleep:  100,
		},
		safetyNeeds: safetyNeeds{
			PhysicalSecurity: 100,
			ResourceSecurity: 100,
			Health:           100,
			Shelter:          100,
		},
		socialNeeds: socialNeeds{
			Social:     100,
			Affection:  100,
			Acceptance: 100,
			Family:     100,
		},
		esteemNeeds: esteemNeeds{
			SelfEsteem: 100,
			Respect:    100,
			Mastery:    100,
		},
		selfActNeeds: selfActNeeds{
			PersonalGrowth: 100,
			Creativity:     100,
			Purpose:        100,
			Exploration:    100,
		},
		CompLabel: Needs_Comp,
	}
	return CompOptionGenerator(n, n.CompLabel)
}
