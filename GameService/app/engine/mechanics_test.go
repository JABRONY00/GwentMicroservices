package engine_test

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"testing"
)

func TestPutRandCardFromStackToHand(t *testing.T) {
	cards := getTestCards(1)

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.PlayerA.PutRandCardFromStackToHand(1)
	if table.PlayerA.Hand != nil {
		t.Error("pick card from empty stack")
	}

	table.PlayerA.PutRandCardFromStackToHand(2)
	if table.PlayerA.Hand != nil {
		t.Error("pick cards from empty stack")
	}

	table.PlayerA.Stack = append(table.PlayerA.Stack, cards["defAssault"])
	table.PlayerA.PutRandCardFromStackToHand(2)
	if len(table.PlayerA.Stack) != 0 {
		t.Error("failed to delete card from stack")
	}
	if table.PlayerA.Hand[0] != cards["defAssault"] {
		t.Error("failed to put card to hand")
	}

	table.PlayerA.Hand = nil

	table.PlayerA.Stack = append(table.PlayerA.Stack, []*engine.Card{
		cards["defAssault"],
		cards["defDistant"],
		cards["rareSiege"],
	}...)

	table.PlayerA.PutRandCardFromStackToHand(2)
	if len(table.PlayerA.Stack) != 1 {
		t.Error("failed to delete cards from stack")
	}
	switch {
	case table.PlayerA.Hand == nil:
		fallthrough
	case len(table.PlayerA.Hand) != 2:
		{
			t.Error("failed to put cards to hand")
		}
	}
}

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

	table.PlayerA.PlayerFieldScoreCounter(true, false, true)

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
		t.Errorf("General score is wrong %d", table.PlayerA.Score)
	}
}

func TestTableScoreCounter(t *testing.T) {
	cardsA := getTestCards(1)
	cardsB := getTestCards(2)

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"
	table.Players[table.Pm.ActPlr].Race = engine.Race.Nilf
	table.Players[table.Pm.PasPlr].Race = engine.Race.Nrth

	table.TableScoreCounter()
	if table.MaxCardScore != 0 {
		t.Error("empty table max score should be 0")
	}
	if table.Winner != table.Players[table.Pm.ActPlr].Name {
		t.Error("Nilfgaard race must be winner in withdraw")
	}

	///////////////////////////////////////////

	table.Players[table.Pm.ActPlr].AssaultField.CardField = append(table.Players[table.Pm.ActPlr].AssaultField.CardField, []*engine.Card{
		cardsA["rareAssault"],
		cardsA["squadA"],
		cardsA["squadA2"],
	}...)
	table.Players[table.Pm.ActPlr].DistantField.CardField = append(table.Players[table.Pm.ActPlr].DistantField.CardField, []*engine.Card{
		cardsA["squadB"],
		cardsA["squadB2"],
		cardsA["squadB3"],
	}...)
	table.Players[table.Pm.ActPlr].SiegeField.CardField = append(table.Players[table.Pm.ActPlr].SiegeField.CardField, []*engine.Card{
		cardsA["boostSiege"],
		cardsA["boostSiege2"],
		cardsA["rareSiege"],
	}...)

	table.Players[table.Pm.PasPlr].AssaultField.CardField = append(table.Players[table.Pm.PasPlr].AssaultField.CardField, []*engine.Card{
		cardsB["rareAssault"],
		cardsB["squadA"],
		cardsB["squadA2"],
	}...)
	table.Players[table.Pm.PasPlr].DistantField.CardField = append(table.Players[table.Pm.PasPlr].DistantField.CardField, []*engine.Card{
		cardsB["squadB"],
		cardsB["squadB2"],
		cardsB["squadB3"],
	}...)
	table.Players[table.Pm.PasPlr].SiegeField.CardField = append(table.Players[table.Pm.PasPlr].SiegeField.CardField, []*engine.Card{
		cardsB["boostSiege"],
		cardsB["boostSiege2"],
		cardsB["rareSiege"],
	}...)

	table.TableScoreCounter()
	if table.MaxCardScore != 6 {
		t.Errorf("table max score should be 6, now %d", table.MaxCardScore)
	}
	if table.Winner != table.Players[table.Pm.ActPlr].Name {
		t.Error("Nilfgaard race must be winner in withdraw")
	}

	//////////////////////////////

	table.Players[table.Pm.ActPlr].SiegeField.CardField = append(table.Players[table.Pm.PasPlr].SiegeField.CardField, cardsA["boostRareSiege"])

	table.TableScoreCounter()
	if table.MaxCardScore != 6 {
		t.Errorf("table max score should be 6, now %d", table.MaxCardScore)
	}
	if table.Winner != table.Players[table.Pm.ActPlr].Name {
		t.Error("Winner has less score than looser")
	}

	//////////////////////////////////

	table.Players[table.Pm.PasPlr].DistantField.HornField = append(table.Players[table.Pm.PasPlr].DistantField.HornField, cardsB["horn"])

	table.TableScoreCounter()
	if table.MaxCardScore != 12 {
		t.Errorf("table max score should be 12, now %d", table.MaxCardScore)
	}
	if table.Winner != table.Players[table.Pm.PasPlr].Name {
		t.Error("Winner has less score than looser")
	}
}

/*func TestExecution(t *testing.T) {

	cardsA := getTestCards(1)
	cardsB := getTestCards(2)

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	table.Pm.ActPlr = "testname1"
	table.Pm.PasPlr = "testname2"

	table.

}*/
