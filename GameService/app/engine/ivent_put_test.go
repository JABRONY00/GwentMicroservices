package engine_test

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"testing"
)

func getTestCards(idKoef uint) map[string]*engine.Card {
	testCards := map[string]*engine.Card{
		"defAssault": {
			ID:       1 + (100 * idKoef),
			Name:     "defAssault",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"defAssault2": {
			ID:       2 + (100 * idKoef),
			Name:     "defAssault2",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"defDistant": {
			ID:       3 + (100 * idKoef),
			Name:     "defDistant",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
		"defDistant2": {
			ID:       4 + (100 * idKoef),
			Name:     "defDistant2",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
		"defSiege": {
			ID:       5 + (100 * idKoef),
			Name:     "defSiege",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   true,
			},
		},
		"defSiege2": {
			ID:       6 + (100 * idKoef),
			Name:     "defSiege2",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   true,
			},
		},
		"rareAssault": {
			ID:       7 + (100 * idKoef),
			Name:     "rareAssault",
			Cost:     3,
			Score:    3,
			Rareness: true,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"rareDistant": {
			ID:       8 + (100 * idKoef),
			Name:     "rareDistant",
			Cost:     3,
			Score:    3,
			Rareness: true,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
		"rareSiege": {
			ID:       9 + (100 * idKoef),
			Name:     "rareSiege",
			Cost:     3,
			Score:    3,
			Rareness: true,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   true,
			},
		},
		"sun": {
			ID:       10 + (100 * idKoef),
			Name:     "Sun",
			Rareness: false,
			Role:     "weather",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"sun2": {
			ID:       11 + (100 * idKoef),
			Name:     "Sun",
			Rareness: false,
			Role:     "weather",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"frost": {
			ID:       12 + (100 * idKoef),
			Name:     "Frost",
			Rareness: false,
			Role:     "weather",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"frost2": {
			ID:       13 + (100 * idKoef),
			Name:     "Frost",
			Rareness: false,
			Role:     "weather",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"fog": {
			ID:       14 + (100 * idKoef),
			Name:     "Fog",
			Rareness: false,
			Role:     "weather",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"fog2": {
			ID:       15 + (100 * idKoef),
			Name:     "Fog",
			Rareness: false,
			Role:     "weather",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"rain": {
			ID:       16 + (100 * idKoef),
			Name:     "Rain",
			Rareness: false,
			Role:     "weather",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"rain2": {
			ID:       17 + (100 * idKoef),
			Name:     "Rain",
			Rareness: false,
			Role:     "weather",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"horn": {
			ID:       18 + (100 * idKoef),
			Name:     "Horn",
			Rareness: false,
			Role:     "horn",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"horn1": {
			ID:       19 + (100 * idKoef),
			Name:     "Frost",
			Rareness: false,
			Role:     "weather",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"horn2": {
			ID:       20 + (100 * idKoef),
			Name:     "Horn",
			Rareness: false,
			Role:     "horn",
			Bonuses:  map[string]string{},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"boostAssault": {
			ID:       21 + (100 * idKoef),
			Name:     "BoostAssault",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "boost",
			Bonuses: map[string]string{
				"boost":      "y",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"boostAssault2": {
			ID:       22 + (100 * idKoef),
			Name:     "BoostAssault",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "boost",
			Bonuses: map[string]string{
				"boost":      "y",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"boostDistant": {
			ID:       23 + (100 * idKoef),
			Name:     "BoostDistant",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "boost",
			Bonuses: map[string]string{
				"boost":      "y",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
		"boostDistant2": {
			ID:       24 + (100 * idKoef),
			Name:     "BoostDistant",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "boost",
			Bonuses: map[string]string{
				"boost":      "y",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
		"boostSiege": {
			ID:       25 + (100 * idKoef),
			Name:     "Boost",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "boost",
			Bonuses: map[string]string{
				"boost":      "y",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   true,
			},
		},
		"boostSiege2": {
			ID:       26 + (100 * idKoef),
			Name:     "Boost",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "boost",
			Bonuses: map[string]string{
				"boost":      "y",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   true,
			},
		},
		"boostRareAssault": {
			ID:       27 + (100 * idKoef),
			Name:     "BoostRare",
			Cost:     3,
			Score:    3,
			Rareness: true,
			Role:     "boost",
			Bonuses: map[string]string{
				"boost":      "y",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"boostRareDistant": {
			ID:       28 + (100 * idKoef),
			Name:     "BoostRare",
			Cost:     3,
			Score:    3,
			Rareness: true,
			Role:     "boost",
			Bonuses: map[string]string{
				"boost":      "y",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
		"boostRareSiege": {
			ID:       29 + (100 * idKoef),
			Name:     "BoostRare",
			Cost:     3,
			Score:    3,
			Rareness: true,
			Role:     "boost",
			Bonuses: map[string]string{
				"boost":      "y",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   true,
			},
		},
		"decoy": {
			ID:       30 + (100 * idKoef),
			Name:     "Decoy",
			Cost:     0,
			Score:    0,
			Rareness: true,
			Role:     "decoy",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"decoy2": {
			ID:       31 + (100 * idKoef),
			Name:     "Decoy",
			Cost:     0,
			Score:    0,
			Rareness: true,
			Role:     "decoy",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"decoy3": {
			ID:       32 + (100 * idKoef),
			Name:     "Decoy",
			Cost:     0,
			Score:    0,
			Rareness: true,
			Role:     "decoy",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"spy": {
			ID:       33 + (100 * idKoef),
			Name:     "Spy",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "spy",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"spy2": {
			ID:       34 + (100 * idKoef),
			Name:     "Spy",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "spy",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"squadA": {
			ID:       35 + (100 * idKoef),
			Name:     "SquadA",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "A",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"squadA2": {
			ID:       36 + (100 * idKoef),
			Name:     "SquadA",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "A",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"squadA3": {
			ID:       37 + (100 * idKoef),
			Name:     "SquadA",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "A",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"squadB": {
			ID:       38 + (100 * idKoef),
			Name:     "SquadB",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "B",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"squadB2": {
			ID:       39 + (100 * idKoef),
			Name:     "SquadB",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "B",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"squadB3": {
			ID:       40 + (100 * idKoef),
			Name:     "SquadB",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "B",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"horner": {
			ID:       41 + (100 * idKoef),
			Name:     "Horner",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "y",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"horner2": {
			ID:       42 + (100 * idKoef),
			Name:     "Horner",
			Cost:     2,
			Score:    2,
			Rareness: false,
			Role:     "",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "y",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"execution": {
			ID:       43 + (100 * idKoef),
			Name:     "Execution",
			Cost:     0,
			Score:    0,
			Rareness: true,
			Role:     "execution",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"execution2": {
			ID:       44 + (100 * idKoef),
			Name:     "Execution",
			Cost:     0,
			Score:    0,
			Rareness: true,
			Role:     "execution",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: false,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"executorAssault": {
			ID:       45 + (100 * idKoef),
			Name:     "ExecutorAssault",
			Cost:     3,
			Score:    3,
			Rareness: false,
			Role:     "executor",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"executorAssault2": {
			ID:       46 + (100 * idKoef),
			Name:     "ExecutorAssault",
			Cost:     3,
			Score:    3,
			Rareness: false,
			Role:     "executor",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: false,
				engine.Field.Siege:   false,
			},
		},
		"healer": {
			ID:       47 + (100 * idKoef),
			Name:     "Healer",
			Cost:     0,
			Score:    0,
			Rareness: false,
			Role:     "healer",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"healer2": {
			ID:       48 + (100 * idKoef),
			Name:     "Healer",
			Cost:     0,
			Score:    0,
			Rareness: false,
			Role:     "healer",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
			},
		},
		"healer3": {
			ID:       49 + (100 * idKoef),
			Name:     "Healer",
			Cost:     0,
			Score:    0,
			Rareness: false,
			Role:     "healer",
			Bonuses: map[string]string{
				"boost":      "",
				"horn":       "",
				"leader-act": "",
				"leader-pas": "",
				"squad":      "",
			},
			TargetField: map[string]bool{
				engine.Field.Assault: true,
				engine.Field.Distant: true,
				engine.Field.Siege:   true,
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
