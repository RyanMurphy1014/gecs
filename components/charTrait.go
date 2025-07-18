package components

import (
	"github.com/RyanMurphy1014/LoreWeaver/entity"
)

type Trait struct {
	CompLabel
}

func WithCharacterTrait(t Trait) entity.CompBuilder {
	return CompOptionGenerator(t, t.CompLabel)
}

var Trait_AbsentMinded CompLabel = "AbsentMinded"
var AbsentMinded = Trait{CompLabel: Trait_AbsentMinded}

var Trait_Adventurous CompLabel = "Adventurous"
var Adventurous = Trait{CompLabel: Trait_Adventurous}

var Trait_Ambitious CompLabel = "Ambitious"
var Ambitious = Trait{CompLabel: Trait_Ambitious}

var Trait_Angler CompLabel = "Angler"
var Angler = Trait{CompLabel: Trait_Angler}

var Trait_AnimalLover CompLabel = "AnimalLover"
var AnimalLover = Trait{CompLabel: Trait_AnimalLover}

var Trait_Artistic CompLabel = "Artistic"
var Artistic = Trait{CompLabel: Trait_Artistic}

var Trait_Athletic CompLabel = "Athletic"
var Athletic = Trait{CompLabel: Trait_Athletic}

var Trait_Bookworm CompLabel = "Bookworm"
var Bookworm = Trait{CompLabel: Trait_Bookworm}

var Trait_Salesperson CompLabel = "Salesperson"
var Salesperson = Trait{CompLabel: Trait_Salesperson}

var Trait_Brave CompLabel = "Brave"
var Brave = Trait{CompLabel: Trait_Brave}

var Trait_Brooding CompLabel = "Brooding"
var Brooding = Trait{CompLabel: Trait_Brooding}

var Trait_CatPerson CompLabel = "CatPerson"
var CatPerson = Trait{CompLabel: Trait_CatPerson}

var Trait_Charismatic CompLabel = "Charismatic"
var Charismatic = Trait{CompLabel: Trait_Charismatic}

var Trait_Childish CompLabel = "Childish"
var Childish = Trait{CompLabel: Trait_Childish}

var Trait_Clumsy CompLabel = "Clumsy"
var Clumsy = Trait{CompLabel: Trait_Clumsy}

var Trait_CommitmentIssues CompLabel = "CommitmentIssues"
var CommitmentIssues = Trait{CompLabel: Trait_CommitmentIssues}

var Trait_CouchPotato CompLabel = "CouchPotato"
var CouchPotato = Trait{CompLabel: Trait_CouchPotato}

var Trait_Coward CompLabel = "Coward"
var Coward = Trait{CompLabel: Trait_Coward}

var Trait_Daredevil CompLabel = "Daredevil"
var Daredevil = Trait{CompLabel: Trait_Daredevil}

var Trait_Disciplined CompLabel = "Disciplined"
var Disciplined = Trait{CompLabel: Trait_Disciplined}

var Trait_DislikesChildren CompLabel = "DislikesChildren"
var DislikesChildren = Trait{CompLabel: Trait_DislikesChildren}

var Trait_Diva CompLabel = "Diva"
var Diva = Trait{CompLabel: Trait_Diva}

var Trait_DogPerson CompLabel = "DogPerson"
var DogPerson = Trait{CompLabel: Trait_DogPerson}

var Trait_Dramatic CompLabel = "Dramatic"
var Dramatic = Trait{CompLabel: Trait_Dramatic}

var Trait_EasilyImpressed CompLabel = "EasilyImpressed"
var EasilyImpressed = Trait{CompLabel: Trait_EasilyImpressed}

var Trait_Eccentric CompLabel = "Eccentric"
var Eccentric = Trait{CompLabel: Trait_Eccentric}

var Trait_EcoFriendly CompLabel = "EcoFriendly"
var EcoFriendly = Trait{CompLabel: Trait_EcoFriendly}

var Trait_Equestrian CompLabel = "Equestrian"
var Equestrian = Trait{CompLabel: Trait_Equestrian}

var Trait_Evil CompLabel = "Evil"
var Evil = Trait{CompLabel: Trait_Evil}

var Trait_Excitable CompLabel = "Excitable"
var Excitable = Trait{CompLabel: Trait_Excitable}

var Trait_FamilyOriented CompLabel = "FamilyOriented"
var FamilyOriented = Trait{CompLabel: Trait_FamilyOriented}

var Trait_Flirty CompLabel = "Flirty"
var Flirty = Trait{CompLabel: Trait_Flirty}

var Trait_Friendly CompLabel = "Friendly"
var Friendly = Trait{CompLabel: Trait_Friendly}

var Trait_Frugal CompLabel = "Frugal"
var Frugal = Trait{CompLabel: Trait_Frugal}

var Trait_Gatherer CompLabel = "Gatherer"
var Gatherer = Trait{CompLabel: Trait_Gatherer}

var Trait_Genius CompLabel = "Genius"
var Genius = Trait{CompLabel: Trait_Genius}

var Trait_Good CompLabel = "Good"
var Good = Trait{CompLabel: Trait_Good}

var Trait_Humorous CompLabel = "Humorous"
var Humorous = Trait{CompLabel: Trait_Humorous}

var Trait_GreenThumb CompLabel = "GreenThumb"
var GreenThumb = Trait{CompLabel: Trait_GreenThumb}

var Trait_Grumpy CompLabel = "Grumpy"
var Grumpy = Trait{CompLabel: Trait_Grumpy}

var Trait_Handy CompLabel = "Handy"
var Handy = Trait{CompLabel: Trait_Handy}

var Trait_Homebody CompLabel = "Homebody"
var Homebody = Trait{CompLabel: Trait_Homebody}

var Trait_HopelessRomantic CompLabel = "HopelessRomantic"
var HopelessRomantic = Trait{CompLabel: Trait_HopelessRomantic}

var Trait_HotHeaded CompLabel = "HotHeaded"
var HotHeaded = Trait{CompLabel: Trait_HotHeaded}

var Trait_Hydrophobic CompLabel = "Hydrophobic"
var Hydrophobic = Trait{CompLabel: Trait_Hydrophobic}

var Trait_Inappropriate CompLabel = "Inappropriate"
var Inappropriate = Trait{CompLabel: Trait_Inappropriate}

var Trait_Insane CompLabel = "Insane"
var Insane = Trait{CompLabel: Trait_Insane}

var Trait_Irresistible CompLabel = "Irresistible"
var Irresistible = Trait{CompLabel: Trait_Irresistible}

var Trait_Kleptomaniac CompLabel = "Kleptomaniac"
var Kleptomaniac = Trait{CompLabel: Trait_Kleptomaniac}

var Trait_Loner CompLabel = "Loner"
var Loner = Trait{CompLabel: Trait_Loner}

var Trait_Loser CompLabel = "Loser"
var Loser = Trait{CompLabel: Trait_Loser}

var Trait_Outdoorsy CompLabel = "Outdoorsy"
var Outdoorsy = Trait{CompLabel: Trait_Outdoorsy}

var Trait_Lucky CompLabel = "Lucky"
var Lucky = Trait{CompLabel: Trait_Lucky}

var Trait_Mean CompLabel = "Mean"
var Mean = Trait{CompLabel: Trait_Mean}

var Trait_Mooch CompLabel = "Mooch"
var Mooch = Trait{CompLabel: Trait_Mooch}

var Trait_Performer CompLabel = "Performer"
var Performer = Trait{CompLabel: Trait_Performer}

var Trait_NaturalCook CompLabel = "NaturalCook"
var NaturalCook = Trait{CompLabel: Trait_NaturalCook}

var Trait_Neat CompLabel = "Neat"
var Neat = Trait{CompLabel: Trait_Neat}

var Trait_Neurotic CompLabel = "Neurotic"
var Neurotic = Trait{CompLabel: Trait_Neurotic}

var Trait_NightOwl CompLabel = "NightOwl"
var NightOwl = Trait{CompLabel: Trait_NightOwl}

var Trait_NoHumor CompLabel = "NoHumor"
var NoHumor = Trait{CompLabel: Trait_NoHumor}

var Trait_Nurturing CompLabel = "Nurturing"
var Nurturing = Trait{CompLabel: Trait_Nurturing}

var Trait_Emotional CompLabel = "Emotional"
var Emotional = Trait{CompLabel: Trait_Emotional}

var Trait_PartyAnimal CompLabel = "PartyAnimal"
var PartyAnimal = Trait{CompLabel: Trait_PartyAnimal}

var Trait_Perceptive CompLabel = "Perceptive"
var Perceptive = Trait{CompLabel: Trait_Perceptive}

var Trait_Perfectionist CompLabel = "Perfectionist"
var Perfectionist = Trait{CompLabel: Trait_Perfectionist}

var Trait_Proper CompLabel = "Proper"
var Proper = Trait{CompLabel: Trait_Proper}

var Trait_Rebellious CompLabel = "Rebellious"
var Rebellious = Trait{CompLabel: Trait_Rebellious}

var Trait_Sailor CompLabel = "Sailor"
var Sailor = Trait{CompLabel: Trait_Sailor}

var Trait_Artist CompLabel = "Artist"
var Artist = Trait{CompLabel: Trait_Artist}

var Trait_Schmoozer CompLabel = "Schmoozer"
var Schmoozer = Trait{CompLabel: Trait_Schmoozer}

var Trait_Shy CompLabel = "Shy"
var Shy = Trait{CompLabel: Trait_Shy}

var Trait_Slob CompLabel = "Slob"
var Slob = Trait{CompLabel: Trait_Slob}

var Trait_Snob CompLabel = "Snob"
var Snob = Trait{CompLabel: Trait_Snob}

var Trait_SocialButterfly CompLabel = "SocialButterfly"
var SocialButterfly = Trait{CompLabel: Trait_SocialButterfly}

var Trait_Awkward CompLabel = "Awkward"
var Awkward = Trait{CompLabel: Trait_Awkward}

var Trait_SupernaturalFan CompLabel = "SupernaturalFan"
var SupernaturalFan = Trait{CompLabel: Trait_SupernaturalFan}

var Trait_SupernaturalSkeptic CompLabel = "SupernaturalSkeptic"
var SupernaturalSkeptic = Trait{CompLabel: Trait_SupernaturalSkeptic}

var Trait_Unflirty CompLabel = "Unflirty"
var Unflirty = Trait{CompLabel: Trait_Unflirty}

var Trait_Unlucky CompLabel = "Unlucky"
var Unlucky = Trait{CompLabel: Trait_Unlucky}

var Trait_Unstable CompLabel = "Unstable"
var Unstable = Trait{CompLabel: Trait_Unstable}

var Trait_Vegetarian CompLabel = "Vegetarian"
var Vegetarian = Trait{CompLabel: Trait_Vegetarian}

var Trait_Virtuoso CompLabel = "Virtuoso"
var Virtuoso = Trait{CompLabel: Trait_Virtuoso}

var Trait_Workaholic CompLabel = "Workaholic"
var Workaholic = Trait{CompLabel: Trait_Workaholic}
