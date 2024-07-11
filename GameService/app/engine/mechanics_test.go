package engine_test

/*import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"sort"
	"testing"
)

var Cards []engine.Card = []engine.Card{
	{
		Name:        "BlueStripesCommando1",
		ID:          8,
		Rareness:    false,
		Cost:        4,
		Score:       4,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: true,
			engine.Field.Distant: false,
			engine.Field.Siege:   false,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "BlueStripes",
			Horn:        false,
			Boost:       false,
		},
	},
	{
		Name:        "BlueStripesCommando2",
		ID:          9,
		Rareness:    false,
		Cost:        4,
		Score:       4,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: true,
			engine.Field.Distant: false,
			engine.Field.Siege:   false,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "BlueStripes",
			Horn:        false,
			Boost:       false,
		},
	},
	{
		Name:        "Catapult1",
		ID:          11,
		Rareness:    false,
		Cost:        8,
		Score:       8,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: false,
			engine.Field.Distant: false,
			engine.Field.Siege:   true,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "Catapult",
			Horn:        false,
			Boost:       false,
		},
	},
	{
		Name:        "Catapult2",
		ID:          12,
		Rareness:    false,
		Cost:        8,
		Score:       8,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: false,
			engine.Field.Distant: false,
			engine.Field.Siege:   true,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "Catapult",
			Horn:        false,
			Boost:       false,
		},
	},
	{
		Name:        "EsteradThyssen",
		ID:          18,
		Rareness:    true,
		Cost:        10,
		Score:       10,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: true,
			engine.Field.Distant: false,
			engine.Field.Siege:   false,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "",
			Horn:        false,
			Boost:       false,
		},
	},
	{
		Name:        "KaedweniSiegeExpert1",
		ID:          20,
		Rareness:    false,
		Cost:        1,
		Score:       1,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: false,
			engine.Field.Distant: false,
			engine.Field.Siege:   true,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "",
			Horn:        false,
			Boost:       true,
		},
	},
	{
		Name:        "KaedweniSiegeExpert2",
		ID:          21,
		Rareness:    false,
		Cost:        1,
		Score:       1,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: false,
			engine.Field.Distant: false,
			engine.Field.Siege:   true,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "",
			Horn:        false,
			Boost:       true,
		},
	},
	{
		Name:        "PoorFuckingInfantry1",
		ID:          25,
		Rareness:    false,
		Cost:        1,
		Score:       1,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: true,
			engine.Field.Distant: false,
			engine.Field.Siege:   false,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "Infantry",
			Horn:        false,
			Boost:       false,
		},
	},
	{
		Name:        "PoorFuckingInfantry2",
		ID:          26,
		Rareness:    false,
		Cost:        1,
		Score:       1,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: true,
			engine.Field.Distant: false,
			engine.Field.Siege:   false,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "Infantry",
			Horn:        false,
			Boost:       false,
		},
	},
	{
		Name:        "SiegfriedOfDenesle",
		ID:          34,
		Rareness:    false,
		Cost:        5,
		Score:       5,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: true,
			engine.Field.Distant: false,
			engine.Field.Siege:   false,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "",
			Horn:        false,
			Boost:       false,
		},
	},
	{
		Name:        "Trebuchet1",
		ID:          38,
		Rareness:    false,
		Cost:        6,
		Score:       6,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: false,
			engine.Field.Distant: false,
			engine.Field.Siege:   true,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "",
			Horn:        false,
			Boost:       false,
		},
	},
	{
		Name:        "Trebuchet2",
		ID:          39,
		Rareness:    false,
		Cost:        6,
		Score:       6,
		Role:        "",
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: false,
			engine.Field.Distant: false,
			engine.Field.Siege:   true,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "",
			Horn:        false,
			Boost:       false,
		},
	},
}

var T = engine.NewTable(&models.Client{Name: "player1"}, &models.Client{Name: "player2"})

func TestGameFieldBonusCounter(t *testing.T) {
	var gameField engine.GameField

	for index := range Cards {
		if Cards[index].TargetField[engine.Field.Assault] {
			gameField.CardField = append(gameField.CardField, &Cards[index])
		}
	}

	if len(gameField.CardField) > 1 {
		sort.SliceStable(gameField.CardField,
			func(i, j int) bool { return gameField.CardField[i].Cost < gameField.CardField[j].Cost })
	}

	gameField.GameFieldBonusCounter(false)

	t.Logf("Results: %+v", gameField)

}

func TestGameFieldScoreCounter(t *testing.T) {
	var gameField engine.GameField

	for index := range Cards {
		if Cards[index].TargetField[engine.Field.Siege] {
			gameField.CardField = append(gameField.CardField, &Cards[index])
		}
	}

	if len(gameField.CardField) > 1 {
		sort.SliceStable(gameField.CardField,
			func(i, j int) bool { return gameField.CardField[i].Cost < gameField.CardField[j].Cost })
	}

	gameField.HornField = append(gameField.HornField, &engine.Card{
		Name:        "Horn",
		ID:          50,
		Rareness:    false,
		Cost:        0,
		Score:       0,
		Role:        engine.Role.Horn,
		Description: "-",
		TargetField: map[string]bool{
			engine.Field.Assault: true,
			engine.Field.Distant: true,
			engine.Field.Siege:   true,
		},
		CardBonus: engine.CardBonus{
			Passive:     false,
			LeaderBonus: "",
			Squad:       "",
			Horn:        true,
			Boost:       false,
		},
	},
	)

	gameField.GameFieldBonusCounter(true)
	gameField.GameFieldScoreCounter()

	for i := range Cards {
		if Cards[i].TargetField[engine.Field.Siege] {
			t.Logf("Card: %s, Score: %d", Cards[i].Name, Cards[i].Score)
		}
	}

	t.Logf("Results: %+v", gameField)

}*/
