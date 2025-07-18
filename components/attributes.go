package components

var Attributes_Comp CompLabel = "Attributes"

type Attributes struct {
	CompLabel
	Strength     int
	Dexterity    int
	Constitution int
	Wisdom       int
	Intelligence int
	Charisma     int
}

func WithAttributes() Attributes {
	return Attributes{
		CompLabel:    Attributes_Comp,
		Strength:     RandomStat(),
		Dexterity:    RandomStat(),
		Constitution: RandomStat(),
		Wisdom:       RandomStat(),
		Intelligence: RandomStat(),
		Charisma:     RandomStat(),
	}
}
