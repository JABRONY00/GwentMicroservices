package engine

import (
	"GwentMicroservices/GameService/app/api/models"
)

//COMMUNICATION OBJECTS

type RequestData struct {
	Instr       string `json:"instruction"`
	CardID      uint   `json:"card-id"`
	TargetField string `json:"target-field"`
	TargetID    uint   `json:"target-id"`
}

//GAME OBJECTS

type Table struct {
	TableID string
	Pm      struct {
		ActPlr string
		PasPlr string
		Instr  string
		IDs    []uint
	}
	PlayerA      PlayerField
	PlayerB      PlayerField
	Players      map[string]*PlayerField
	WeatherFlags struct {
		Frost bool
		Fog   bool
		Rain  bool
	}
	MaxCardScore uint
	Winner       string
	ActiveCards  []Card
}

type PlayerField struct {
	Name         string
	Conn         *models.Connection
	Score        uint
	WinTokens    uint
	MaxCardScore uint
	Race         string
	LeaderFlag   bool
	PassFlag     bool
	LeaderCard   *Card
	Stack        []*Card
	Hand         []*Card
	Fields       map[string]*GameField
	WeatherField []*Card
	AssaultField GameField
	DistantField GameField
	SiegeField   GameField
	Grave        []*Card
}

type GameField struct {
	Score         uint    // Common Score of Field
	MaxCardScore  uint    // Max Score on Field(Rare cards Excluded).  This parameter needs for execution
	CardField     []*Card // Cards placed here
	HornField     []*Card // Place for keeping a Doubler card
	ActiveBonuses struct {
		Weather uint
		Squads  map[string]uint
		Horn    uint
		Boost   uint
	} // Wich bonuses affect cards on field
}

type Card struct {
	Name        string          `json:"name"`
	ID          uint            `json:"id"`
	Rareness    bool            `json:"rareness"`
	Cost        uint            `json:"cost"`
	Score       uint            `json:"score"`
	Role        string          `json:"role"`
	Description string          `json:"description"`
	TargetField map[string]bool `json:"targetfield"`
	CardBonus
}

type CardBonus struct {
	Passive     bool   `json:"passive"`
	LeaderBonus string `json:"leaderbonus"`
	Squad       string `json:"squad"`
	Horn        bool   `json:"horn"`
	Boost       bool   `json:"boost"`
}

type TableShort struct {
	Winner       bool `json:"winner"`
	ActivePlayer bool `json:"active-player"`
	WinTokens    struct {
		Player uint `json:"player"`
		Enemy  uint `json:"enemy"`
	} `json:"win-tokens"`
	WeatherField struct {
		Frost bool   `json:"frost"`
		Fog   bool   `json:"fog"`
		Rain  bool   `json:"rain"`
		Cards []uint `json:"cards"`
	} `json:"weather-field"`
	Assaultfield struct {
		Player     []CardShort `json:"player-cards"`
		HornPlayer bool        `json:"player-horn"`
		Enemy      []CardShort `json:"enemy-cards"`
		HornEnemy  bool        `json:"enemy-horn"`
	} `json:"assault-field"`
	Distantfield struct {
		Player     []CardShort `json:"player-cards"`
		HornPlayer bool        `json:"player-horn"`
		Enemy      []CardShort `json:"enemy-cards"`
		HornEnemy  bool        `json:"enemy-horn"`
	} `json:"distant-field"`
	SiegeField struct {
		Player     []CardShort `json:"player-cards"`
		HornPlayer bool        `json:"player-horn"`
		Enemy      []CardShort `json:"enemy-cards"`
		HornEnemy  bool        `json:"enemy-horn"`
	} `json:"siege-field"`
	Gravefield struct {
		Player []uint `json:"player"`
		Enemy  []uint `json:"enemy"`
	} `json:"grave-field"`
	Stack struct {
		Player []uint `json:"player"`
		Enemy  uint   `json:"enemy"`
	} `json:"stack"`
	Hand struct {
		Player []uint `json:"player"`
		Enemy  uint   `json:"enemy"`
	} `json:"hand"`
	Leaderfield struct {
		Player bool `json:"player"`
		Enemy  bool `json:"enemy"`
	} `json:"leader-field"`
}

type CardShort struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
	Score  uint   `json:"score"`
}
