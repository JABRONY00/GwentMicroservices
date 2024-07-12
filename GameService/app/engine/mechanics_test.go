package engine_test

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"sync"
	"testing"
)

func TestPlayerFieldScoreCounter(t *testing.T) {
	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	cards := table.GetStartCards(1, models.PlayerPreset{Race: "nrth", Stack: TestStack1})

	for i := range cards {
		switch {
		case cards[i].TargetField[engine.Field.Assault]:
			{
				table.PlayerA.AssaultField.CardField = append(table.PlayerA.AssaultField.CardField, &cards[i])
			}
		case cards[i].TargetField[engine.Field.Distant]:
			{
				table.PlayerA.DistantField.CardField = append(table.PlayerA.DistantField.CardField, &cards[i])
			}
		case cards[i].TargetField[engine.Field.Siege]:
			{
				table.PlayerA.SiegeField.CardField = append(table.PlayerA.SiegeField.CardField, &cards[i])
			}
		}
	}

	table.PlayerA.DistantField.HornField = append(table.PlayerA.DistantField.HornField, &engine.Card{ID: 1})
	table.PlayerA.SiegeField.HornField = append(table.PlayerA.SiegeField.HornField, &engine.Card{ID: 2})

	twg := &sync.WaitGroup{}
	twg.Add(1)
	table.PlayerA.PlayerFieldScoreCounter(twg, true, false, true)
	twg.Wait()

	if table.PlayerA.AssaultField.Score != 20 {
		t.Error("Assault score is wrong")
	}

	if table.PlayerA.DistantField.Score != 108 {
		t.Error("Distant score is wrong")
	}

	if table.PlayerA.SiegeField.Score != 36 {
		t.Error("Siege score is wrong")
	}

	if table.PlayerA.MaxCardScore != 30 {
		t.Error("Max card score is wrong")
	}

	if table.PlayerA.Score != 164 {
		t.Error("General score is wrong")
	}
}
