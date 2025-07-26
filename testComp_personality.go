package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Personality struct {
	CompId
	Aggression int // 0=Pacifist, 100=Aggressive
	Curiosity  int // 0=Traditionalist, 100=Curious
	Drive      int // 0=Lazy, 100=Industrious
	Empathy    int // 0=Calloused, 100=Empath
	Honesty    int // 0=Liar, 100=Truthful
	Loyalty    int // 0=Spineless, 100=Loyal
	Optimism   int // 0=Pessimistic, 100=Optimistic
	Reason     int // 0=Impulsive, 100=Rational
	Socialness int // 0=Solitary, 100=Gregarious
}

func (p Personality) String() string {
	var sb strings.Builder

	val := reflect.ValueOf(p)
	typ := val.Type()
	sb.WriteString("{")
	for i := range val.NumField() {
		sb.WriteString(fmt.Sprintf("\n\t\t%v : %+v", typ.Field(i).Name, val.Field(i)))
	}
	sb.WriteString(" }")
	return sb.String()
}

func (comp Personality) Id() CompId {
	return comp.CompId
}

type GenMethod string

var Random GenMethod = "Random"
var Empty GenMethod = "Empty"

var Personality_Comp CompLabel = "Personality"

func WithPersonality(genMethod GenMethod) Component {
	CompId := nextCompId(string(Personality_Comp))
	switch genMethod {
	case "Empty":
		return Component{
			CompLabel: Personality_Comp,
			CompData: Personality{
				CompId:     CompId,
				Aggression: 0,
				Curiosity:  0,
				Drive:      0,
				Empathy:    0,
				Honesty:    0,
				Loyalty:    0,
				Optimism:   0,
				Reason:     0,
				Socialness: 0,
			},
		}
	case "Random":
		return Component{
			CompLabel: Personality_Comp,
			CompData: Personality{
				Aggression: RandomStat(),
				Curiosity:  RandomStat(),
				Drive:      RandomStat(),
				Empathy:    RandomStat(),
				Honesty:    RandomStat(),
				Loyalty:    RandomStat(),
				Optimism:   RandomStat(),
				Reason:     RandomStat(),
				Socialness: RandomStat(),
				CompId:     CompId,
			},
		}
	default:
		return Component{
			CompLabel: Personality_Comp,
			CompData: Personality{
				CompId:     CompId,
				Aggression: 0,
				Curiosity:  0,
				Drive:      0,
				Empathy:    0,
				Honesty:    0,
				Loyalty:    0,
				Optimism:   0,
				Reason:     0,
				Socialness: 0,
			},
		}
	}
}
