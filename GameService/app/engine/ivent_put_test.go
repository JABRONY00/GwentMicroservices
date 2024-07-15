package engine_test

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"slices"
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
			Role:     "",
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
			Role:     "",
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
		"agile": {
			ID:       50 + (100 * idKoef),
			Name:     "Agile",
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
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
		"agile2": {
			ID:       51 + (100 * idKoef),
			Name:     "Agile",
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
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
		"agile3": {
			ID:       52 + (100 * idKoef),
			Name:     "Agile",
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
				engine.Field.Distant: true,
				engine.Field.Siege:   false,
			},
		},
	}

	return testCards
}

func TestPutWeatherCard(t *testing.T) {
	cardsA := getTestCards(1)
	cardsB := getTestCards(2)

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"

	if table.WeatherFlags.Frost || table.WeatherFlags.Fog || table.WeatherFlags.Rain {
		t.Error("initial flags condition is wrong")
	}
	if table.Players[table.Pm.ActPlr].Grave != nil || table.Players[table.Pm.PasPlr].Grave != nil {
		t.Error("initial flags condition is wrong")
	}

	//////////////////

	err := table.PutWeatherCard(cardsA["frost"])
	if err != nil {
		t.Errorf("%v", err)
	}
	if !table.WeatherFlags.Frost || table.WeatherFlags.Fog || table.WeatherFlags.Rain {
		t.Error("failed to switch flag correctly")
	}
	if table.Players[table.Pm.ActPlr].WeatherField == nil {
		t.Error("failed to put card on weather field")
	}
	if table.Players[table.Pm.PasPlr].WeatherField != nil {
		t.Error("failed to put card on weather correctly field")
	}
	if table.Players[table.Pm.ActPlr].Grave != nil || table.Players[table.Pm.PasPlr].Grave != nil {
		t.Error("somehow card appeared to be on grave field")
	}
	if table.Pm.Instr != engine.Instr.PmSwitch {
		t.Error("failed to set instruction")
	}

	////////////////////////

	err = table.PutWeatherCard(cardsA["rain"])
	if err != nil {
		t.Errorf("%v", err)
	}
	if !table.WeatherFlags.Frost || table.WeatherFlags.Fog || !table.WeatherFlags.Rain {
		t.Error("failed to switch flag correctly")
	}
	if len(table.Players[table.Pm.ActPlr].WeatherField) != 2 {
		t.Error("failed to put card on weather field")
	}
	if table.Players[table.Pm.PasPlr].WeatherField != nil {
		t.Error("card apeared to be on the other side")
	}
	if table.Players[table.Pm.ActPlr].Grave != nil || table.Players[table.Pm.PasPlr].Grave != nil {
		t.Error("somehow card appeared to be on grave field")
	}

	//////////////////////////

	err = table.PutWeatherCard(cardsA["frost2"])
	if err != nil {
		t.Errorf("%v", err)
	}
	if !table.WeatherFlags.Frost || table.WeatherFlags.Fog || !table.WeatherFlags.Rain {
		t.Error("failed to switch flag correctly")
	}
	if len(table.Players[table.Pm.ActPlr].WeatherField) != 2 {
		t.Error("duplicating card on weather field")
	}
	if table.Players[table.Pm.PasPlr].WeatherField != nil {
		t.Error("card apeared to be on the other side")
	}
	if table.Players[table.Pm.ActPlr].Grave == nil || table.Players[table.Pm.PasPlr].Grave != nil {
		t.Error("failed to put duplicating card on grave field correctly")
	}

	//////////////////////////////

	table.PermissionSwitch()

	err = table.PutWeatherCard(cardsB["fog"])
	if err != nil {
		t.Errorf("%v", err)
	}
	if !table.WeatherFlags.Frost || !table.WeatherFlags.Fog || !table.WeatherFlags.Rain {
		t.Error("failed to switch flag correctly")
	}
	if table.Players[table.Pm.ActPlr].WeatherField == nil {
		t.Error("failed to put card on weather field")
	}
	if len(table.Players[table.Pm.PasPlr].WeatherField) != 2 {
		t.Error("failed to put card on weather correctly field")
	}
	if table.Players[table.Pm.ActPlr].Grave != nil {
		t.Error("somehow card appeared to be on grave field")
	}

	/////////////////////

	err = table.PutWeatherCard(cardsB["rain"])
	if err != nil {
		t.Errorf("%v", err)
	}
	if !table.WeatherFlags.Frost || !table.WeatherFlags.Fog || !table.WeatherFlags.Rain {
		t.Error("failed to switch flag correctly")
	}
	if len(table.Players[table.Pm.ActPlr].WeatherField) != 2 {
		t.Error("failed to change duplicating enemy card")
	}
	if len(table.Players[table.Pm.PasPlr].WeatherField) != 1 {
		t.Error("failed to delete duplicating card from enemy weather field")
	}
	if table.Players[table.Pm.ActPlr].Grave != nil {
		t.Error("somehow duplicating card appeared on own grave field")
	}
	if len(table.Players[table.Pm.PasPlr].Grave) != 2 {
		t.Error("failed to put duplicating card on enemy grave field")
	}

	/////////////////////

	table.PermissionSwitch()

	err = table.PutWeatherCard(cardsA["sun"])
	if err != nil {
		t.Errorf("%v", err)
	}
	if table.WeatherFlags.Frost || table.WeatherFlags.Fog || table.WeatherFlags.Rain {
		t.Error("failed to switch flag correctly")
	}
	if table.Players[table.Pm.ActPlr].WeatherField != nil || table.Players[table.Pm.PasPlr].WeatherField != nil {
		t.Error("failed clear weather fields")
	}
	if len(table.Players[table.Pm.ActPlr].Grave) != 4 {
		t.Error("failed to put own cards to grave field correctly")
	}
	if len(table.Players[table.Pm.PasPlr].Grave) != 2 {
		t.Error("failed to put enemy cards to grave field correctly")
	}

	///////////////////////////////

	err = table.PutWeatherCard(cardsA["sun2"])
	if err != nil {
		t.Errorf("%v", err)
	}
	if table.WeatherFlags.Frost || table.WeatherFlags.Fog || table.WeatherFlags.Rain {
		t.Error("failed to switch flag correctly")
	}
	if table.Players[table.Pm.ActPlr].WeatherField != nil || table.Players[table.Pm.PasPlr].WeatherField != nil {
		t.Error("failed clear weather fields")
	}
	if len(table.Players[table.Pm.ActPlr].Grave) != 5 {
		t.Error("failed to put own sun card to grave field correctly")
	}
	if len(table.Players[table.Pm.PasPlr].Grave) != 2 {
		t.Error("failed to put enemy sun card to grave field correctly")
	}

}

func TestPutDefaultCard(t *testing.T) {
	cards := getTestCards(1)

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"

	table.PutDefaultCard(cards["defAssault"], engine.Field.Assault)

	if table.Players[table.Pm.ActPlr].AssaultField.CardField[0] != cards["defAssault"] {
		t.Error("failed to put card on field")
	}

	table.PutDefaultCard(cards["defAssault"], engine.Field.Assault)

	if len(table.Players[table.Pm.ActPlr].AssaultField.CardField) != 2 {
		t.Error("failed to put second card on field")
	}
	if table.Pm.Instr != engine.Instr.PmSwitch {
		t.Error("failed to set instruction")
	}
}

func TestPutDecoyCard(t *testing.T) {
	cards := getTestCards(1)

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"

	table.Players[table.Pm.ActPlr].AssaultField.CardField = append(table.Players[table.Pm.ActPlr].AssaultField.CardField, []*engine.Card{
		cards["defAssault"],
		cards["defAssault2"],
		cards["rareAssault"],
	}...)

	if len(table.Players[table.Pm.ActPlr].AssaultField.CardField) != 3 || table.Players[table.Pm.ActPlr].Hand != nil {
		t.Error("initial table condition is wrong")
	}

	///////////////////////////////////////////////////////

	err := table.PutDecoyCard(cards["decoy"], engine.Field.Assault, cards["defAssault"].ID)
	if err != nil {
		t.Errorf("non existing error:%v", err)
	}
	if len(table.Players[table.Pm.ActPlr].AssaultField.CardField) != 3 {
		t.Error("card field contains wrong quantity of cards")
	}
	if slices.Contains(table.Players[table.Pm.ActPlr].AssaultField.CardField, cards["defAssault"]) ||
		!slices.Contains(table.Players[table.Pm.ActPlr].AssaultField.CardField, cards["decoy"]) {
		t.Error("card field modified incorrectly")
	}
	switch {
	case table.Players[table.Pm.ActPlr].Hand == nil:
		{
			t.Error("no cards was put to hand")
		}
	case table.Players[table.Pm.ActPlr].Hand[0].ID != cards["defAssault"].ID:
		{
			t.Error("wrong card exchanged")
		}
	}
	if table.Pm.Instr != engine.Instr.PmSwitch {
		t.Error("failed to set instruction")
	}

	//////////////////////////////////////////////////

	err = table.PutDecoyCard(cards["decoy2"], engine.Field.Assault, cards["rareAssault"].ID)
	if err.Error() != engine.Instr.ForbMove {
		t.Error("failed to form error while exchanging rare card")
	}
	if len(table.Players[table.Pm.ActPlr].AssaultField.CardField) != 3 ||
		slices.Contains(table.Players[table.Pm.ActPlr].AssaultField.CardField, cards["decoy2"]) ||
		!slices.Contains(table.Players[table.Pm.ActPlr].AssaultField.CardField, cards["rareAssault"]) {
		t.Error("card field was changed besides an error")
	}
	if slices.Contains(table.Players[table.Pm.ActPlr].Hand, cards["rareAssault"]) {
		t.Error("card exchanged besides an error")
	}

	////////////////////////////////////

	err = table.PutDecoyCard(cards["decoy2"], engine.Field.Assault, cards["decoy"].ID)
	if err.Error() != engine.Instr.ForbMove {
		t.Error("failed to form error while exchanging decoy card")
	}
	if len(table.Players[table.Pm.ActPlr].AssaultField.CardField) != 3 ||
		slices.Contains(table.Players[table.Pm.ActPlr].AssaultField.CardField, cards["decoy2"]) ||
		!slices.Contains(table.Players[table.Pm.ActPlr].AssaultField.CardField, cards["decoy"]) {
		t.Error("card field was changed besides an error")
	}
	if slices.Contains(table.Players[table.Pm.ActPlr].Hand, cards["decoy"]) {
		t.Error("card exchanged besides an error")
	}

	///////////////////////////////////////////////////////////////////

	err = table.PutDecoyCard(cards["decoy2"], engine.Field.Assault, cards["defSiege"].ID)
	if err.Error() != engine.Instr.ForbMove {
		t.Error("failed to form error while exchanging non existing card")
	}
	if len(table.Players[table.Pm.ActPlr].AssaultField.CardField) != 3 ||
		slices.Contains(table.Players[table.Pm.ActPlr].AssaultField.CardField, cards["decoy2"]) {
		t.Error("card field was changed besides an error")
	}
	if slices.Contains(table.Players[table.Pm.ActPlr].Hand, cards["defaultSiege"]) {
		t.Error("card exchanged besides an error")
	}

	/////////////////////////////////////////////

	err = table.PutDecoyCard(cards["decoy2"], engine.Field.Assault, 0)
	if err != nil {
		t.Error("non existing error")
	}
	if len(table.Players[table.Pm.ActPlr].AssaultField.CardField) != 4 ||
		!slices.Contains(table.Players[table.Pm.ActPlr].AssaultField.CardField, cards["decoy2"]) {
		t.Error("card was not put on the field")
	}
	if len(table.Players[table.Pm.ActPlr].Hand) != 1 {
		t.Error("wrong cards quantity in hand")
	}
}

func TestPutSpyCard(t *testing.T) {
	cards := getTestCards(1)

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"

	table.PutSpyCard(cards["spy"], engine.Field.Assault)
	if table.Players[table.Pm.ActPlr].Hand != nil {
		t.Error("non existing cards appeared in own hand")
	}
	if table.Players[table.Pm.PasPlr].Hand != nil {
		t.Error("non existing cards appeared in enemy hand")
	}
	if !slices.Contains(table.Players[table.Pm.PasPlr].AssaultField.CardField, cards["spy"]) {
		t.Error("failed to put spy card on enemy card field")
	}

	table.Players[table.Pm.PasPlr].AssaultField.CardField = nil

	//////////////////////////

	table.Players[table.Pm.ActPlr].Stack = append(table.Players[table.Pm.ActPlr].Stack, cards["defAssault"])

	table.PutSpyCard(cards["spy"], engine.Field.Assault)
	if !slices.Contains(table.Players[table.Pm.ActPlr].Hand, cards["defAssault"]) {
		t.Error("failed to put correct card to hand")
	}
	if table.Players[table.Pm.PasPlr].Hand != nil {
		t.Error("card appeared in enemy hand")
	}
	if !slices.Contains(table.Players[table.Pm.PasPlr].AssaultField.CardField, cards["spy"]) {
		t.Error("failed to put spy card on enemy card field")
	}

	table.Players[table.Pm.PasPlr].AssaultField.CardField = nil
	table.Players[table.Pm.ActPlr].Hand = nil
	table.Players[table.Pm.ActPlr].Stack = nil

	////////////////////////

	table.Players[table.Pm.ActPlr].Stack = append(table.Players[table.Pm.ActPlr].Stack, []*engine.Card{
		cards["defAssault"],
		cards["defAssault2"],
		cards["rareAssault"],
	}...)

	table.PutSpyCard(cards["spy"], engine.Field.Assault)
	if len(table.Players[table.Pm.ActPlr].Hand) != 2 {
		t.Error("failed to put cards to hand")
	}
	if slices.Contains(table.Players[table.Pm.ActPlr].Hand, cards["defAssault"]) &&
		slices.Contains(table.Players[table.Pm.ActPlr].Stack, cards["defAssault"]) {
		t.Error("failed to modify player field correctly")
	}
	if slices.Contains(table.Players[table.Pm.ActPlr].Hand, cards["defAssault2"]) &&
		slices.Contains(table.Players[table.Pm.ActPlr].Stack, cards["defAssault2"]) {
		t.Error("failed to modify player field correctly")
	}
	if slices.Contains(table.Players[table.Pm.ActPlr].Hand, cards["rareSiege"]) &&
		slices.Contains(table.Players[table.Pm.ActPlr].Stack, cards["rareSiege"]) {
		t.Error("failed to modify player field correctly")
	}
	if table.Players[table.Pm.PasPlr].Hand != nil {
		t.Error("cards appeared in enemy hand")
	}
	if !slices.Contains(table.Players[table.Pm.PasPlr].AssaultField.CardField, cards["spy"]) {
		t.Error("failed to put spy card on enemy card field")
	}
}

/*func TestPutExecutionCard(t *testing.T) {

	cardsA := getTestCards(1)
	cardsB := getTestCards(2)

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"
}*/
