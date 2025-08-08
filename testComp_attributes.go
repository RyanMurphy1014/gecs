package main

type Attributes struct {
	Strength     int
	Dexterity    int
	Constitution int
	Wisdom       int
	Intelligence int
	Charisma     int
}

func WithAttributes(random bool) Attributes {
	if random == true {
		return Attributes{
			Strength:     RandomStat(),
			Dexterity:    RandomStat(),
			Constitution: RandomStat(),
			Wisdom:       RandomStat(),
			Intelligence: RandomStat(),
			Charisma:     RandomStat(),
		}
	}
	return Attributes{
		Strength:     0,
		Dexterity:    0,
		Constitution: 0,
		Wisdom:       0,
		Intelligence: 0,
		Charisma:     0,
	}
}
