package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Personality struct {
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

func WithPersonality(random bool) Personality {
	if !random {
		return Personality{
			Aggression: 0,
			Curiosity:  0,
			Drive:      0,
			Empathy:    0,
			Honesty:    0,
			Loyalty:    0,
			Optimism:   0,
			Reason:     0,
			Socialness: 0,
		}
	} else {
		return Personality{
			Aggression: RandomStat(),
			Curiosity:  RandomStat(),
			Drive:      RandomStat(),
			Empathy:    RandomStat(),
			Honesty:    RandomStat(),
			Loyalty:    RandomStat(),
			Optimism:   RandomStat(),
			Reason:     RandomStat(),
			Socialness: RandomStat(),
		}
	}
}
