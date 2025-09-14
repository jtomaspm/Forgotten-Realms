package game

type Faction string

const (
	Caldari      Faction = "caldari"      //A highly advanced magical-tech faction (inspired by Protoss/High Elves)
	Varnak       Faction = "varnak"       //Brutal, war-driven creatures with strong melee units (inspired by Orcs/Zerg)
	Dawnhold     Faction = "dawnhold"     //Strategic human-like settlers focused on balance and economics (inspired by Travian Romans)
	TheForgotten Faction = "theforgotten" //Ancient native defenders of the realm with powerful NPC-controlled villages and elite troops
)

type Direction string

const (
	NorthEast Direction = "NorthEast"
	NorthWest Direction = "NorthWest"
	SouthEast Direction = "SouthEast"
	SouthWest Direction = "SouthWest"
)

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Resources struct {
	Wood float32 `json:"wood"`
	Clay float32 `json:"clay"`
	Iron float32 `json:"iron"`
}

type ResourceProduction struct {
	Base        Resources `json:"base"`
	Buildings   Resources `json:"buildings"`
	Multipliers Resources `json:"multipliers"`
}

type Population struct {
	Population int `json:"population"`
}

type Cost struct {
	Resources
	Population int `json:"population"`
}
