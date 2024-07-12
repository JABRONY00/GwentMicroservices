package engine_test

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"testing"
)

func getTestCards() map[string]*engine.Card {
	testCards := map[string]*engine.Card{
		"default1": {
			ID:       1,
			Name:     "def1",
			Rareness: false,
			Role:     "",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
	}

	return testCards
}

func TestPutWeatherCard(t *testing.T) {
	cards := []*engine.Card{
		{
			ID:   1,
			Name: engine.Weather.Frost,
			Role: engine.Role.Weather,
		},
		{
			ID:   2,
			Name: engine.Weather.Frost,
			Role: engine.Role.Weather,
		},
		{
			ID:   3,
			Name: engine.Weather.Fog,
			Role: engine.Role.Weather,
		},
		{
			ID:   4,
			Name: engine.Weather.Fog,
			Role: engine.Role.Weather,
		},
		{
			ID:   5,
			Name: engine.Weather.Rain,
			Role: engine.Role.Weather,
		},
		{
			ID:   6,
			Name: engine.Weather.Rain,
			Role: engine.Role.Weather,
		},
		{
			ID:   7,
			Name: engine.Weather.Sun,
			Role: engine.Role.Weather,
		},
		{
			ID:   8,
			Name: engine.Weather.Sun,
			Role: engine.Role.Weather,
		},
	}

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"

	table.PutWeatherCard(cards[0])
	table.PutWeatherCard(cards[2])
	table.PutWeatherCard(cards[4])

	if len(table.Players[table.Pm.ActPlr].WeatherField) != 3 {
		t.Error("failed to put cards on weather field")
	}

	if !table.WeatherFlags.Frost ||
		!table.WeatherFlags.Fog ||
		!table.WeatherFlags.Rain {
		t.Error("failed to switch the flags")
	}

	table.PermissionSwitch()

	table.PutWeatherCard(cards[1])
	table.PutWeatherCard(cards[3])
	table.PutWeatherCard(cards[5])

	if len(table.Players[table.Pm.ActPlr].WeatherField) != 3 {
		t.Error("failed to put cards on weather field")
	}

	if len(table.Players[table.Pm.PasPlr].Grave) != 3 {
		t.Error("failed to put cards to passive players grave field")
	}

	if !table.WeatherFlags.Frost ||
		!table.WeatherFlags.Fog ||
		!table.WeatherFlags.Rain {
		t.Error("failed to switch the flags")
	}

	table.PermissionSwitch()

	table.PutWeatherCard(cards[6])

	if len(table.Players[table.Pm.ActPlr].WeatherField) != 0 &&
		len(table.Players[table.Pm.PasPlr].WeatherField) != 0 {
		t.Error("failed to clean weather field")
	}

	if len(table.Players[table.Pm.ActPlr].Grave) != 4 {
		t.Error("failed to put sun card to active players grave field")
	}

	if len(table.Players[table.Pm.PasPlr].Grave) != 3 {
		t.Error("failed to put cards to passive players grave field")
	}

	if table.WeatherFlags.Frost ||
		table.WeatherFlags.Fog ||
		table.WeatherFlags.Rain {
		t.Error("failed to switch the flags")
	}

	table.PermissionSwitch()

	table.PutWeatherCard(cards[7])

	if len(table.Players[table.Pm.ActPlr].Grave) != 4 {
		t.Error("failed to put sun card to active players grave field")
	}

	if len(table.Players[table.Pm.PasPlr].Grave) != 4 {
		t.Error("something went wrong")
	}

	if table.WeatherFlags.Frost ||
		table.WeatherFlags.Fog ||
		table.WeatherFlags.Rain {
		t.Error("failed to switch the flags")
	}
}

func TestPutDefaultCard(t *testing.T) {
	cards := []*engine.Card{
		{
			ID:   1,
			Name: "1",
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		{
			ID:   2,
			Name: "2",
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
		{
			ID:   3,
			Name: "3",
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   true,
			},
		},
		{
			ID:   4,
			Name: "4",
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
	}

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"

	table.PutDefaultCard(cards[0], engine.Field.Assault)

	if table.Players[table.Pm.ActPlr].AssaultField.CardField[0] != cards[0] {
		t.Error("failed to put card on field")
	}

	table.PutDefaultCard(cards[3], engine.Field.Assault)

	if len(table.Players[table.Pm.ActPlr].AssaultField.CardField) != 2 {
		t.Error("failed to put second card on field")
	}
}

func TestPutDecoyCard(t *testing.T) {
	cards := []*engine.Card{
		{
			ID:   1,
			Name: "1",
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		{
			ID:       2,
			Name:     "Rare",
			Rareness: true,
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		{
			ID:   3,
			Name: "Decoy1",
			Role: engine.Role.Decoy,
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		{
			ID:   4,
			Name: "Decoy2",
			Role: engine.Role.Decoy,
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
	}

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"

	table.Players[table.Pm.ActPlr].AssaultField.CardField = append(table.Players[table.Pm.ActPlr].AssaultField.CardField, cards[:2]...)

	err := table.PutDecoyCard(cards[2], engine.Field.Assault, 2)
	if err != nil {
		t.Error("Rare card was exchanged")
	}
	t.Logf("")
}
