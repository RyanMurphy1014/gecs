package main

var Attributes_Comp CompLabel = "Attributes"

type Attributes struct {
	CompId
	Strength     int
	Dexterity    int
	Constitution int
	Wisdom       int
	Intelligence int
	Charisma     int
}

func (comp Attributes) Id() CompId {
	return comp.CompId
}

func WithAttributes(random bool) Component {
	compId, err := nextCompId(string(Attributes_Comp))
	if err != nil {
		panic("Component ID has already been registered - WithAttributes()")
	}

	if random == true {
		return Component{
			CompLabel: Attributes_Comp,
			CompData: Attributes{
				CompId:       compId,
				Strength:     RandomStat(),
				Dexterity:    RandomStat(),
				Constitution: RandomStat(),
				Wisdom:       RandomStat(),
				Intelligence: RandomStat(),
				Charisma:     RandomStat(),
			},
		}
	}
	return Component{
		CompLabel: Attributes_Comp,
		CompData: Attributes{
			CompId:       compId,
			Strength:     0,
			Dexterity:    0,
			Constitution: 0,
			Wisdom:       0,
			Intelligence: 0,
			Charisma:     0,
		},
	}
}
