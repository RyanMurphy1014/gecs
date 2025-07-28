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
func (comp Attributes) SetId(id CompId) {
	comp.CompId = id
}

func WithAttributes(random bool) Component {
	if random == true {
		return Component{
			CompLabel: Attributes_Comp,
			CompData: Attributes{
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
			Strength:     0,
			Dexterity:    0,
			Constitution: 0,
			Wisdom:       0,
			Intelligence: 0,
			Charisma:     0,
		},
	}
}
