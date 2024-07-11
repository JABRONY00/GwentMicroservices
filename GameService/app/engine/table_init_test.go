package engine_test

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"slices"
	"sort"
	"testing"
)

var TestStack1 = []uint{2, 6, 7, 8, 9, 11, 12, 13, 14, 15, 18, 20, 21, 24, 25, 26, 29, 31, 34}
var TestStack2 = []uint{1, 6, 7, 8, 9, 10, 15, 16, 17, 22, 24, 27}

func TestInitPlayerField(t *testing.T) {
	pf := engine.PlayerField{}

	pf.InitPlayerField(&models.Client{Name: "testname"})

	if &pf.AssaultField.CardField != &pf.Fields[engine.Field.Assault].CardField {
		t.Error("Fields map has no access to the CardField failed ")
	}
	if &pf.DistantField.CardField != &pf.Fields[engine.Field.Distant].CardField {
		t.Error("Fields map has no access to the CardField failed ")
	}
	if &pf.SiegeField.CardField != &pf.Fields[engine.Field.Siege].CardField {
		t.Error("Fields map has no access to the CardField failed ")
	}
}

func TestNewTable(t *testing.T) {

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	if &table.PlayerA.AssaultField != table.Players[table.PlayerA.Name].Fields[engine.Field.Assault] {
		t.Errorf("Access to the field failed ")
	}
	if &table.PlayerA.DistantField != table.Players[table.PlayerA.Name].Fields[engine.Field.Distant] {
		t.Errorf("Access to the field failed ")
	}
	if &table.PlayerA.SiegeField != table.Players[table.PlayerA.Name].Fields[engine.Field.Siege] {
		t.Errorf("Access to the field failed ")
	}

	if &table.PlayerB.AssaultField != table.Players[table.PlayerB.Name].Fields[engine.Field.Assault] {
		t.Errorf("Access to the field failed ")
	}
	if &table.PlayerB.DistantField != table.Players[table.PlayerB.Name].Fields[engine.Field.Distant] {
		t.Errorf("Access to the field failed ")
	}
	if &table.PlayerB.SiegeField != table.Players[table.PlayerB.Name].Fields[engine.Field.Siege] {
		t.Errorf("Access to the field failed ")
	}
}

////////////////////////////////////////////////////////////

func TestGetStartCards(t *testing.T) {

	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)
	koef := uint(1)

	cards := table.GetStartCards(koef, models.PlayerPreset{Race: "nrth", Stack: TestStack1})
	if len(cards) < 1 {
		t.Error("No cards returned")
	}

	for i := range cards {
		if !slices.Contains(TestStack1, cards[i].ID-(koef*100)) ||
			len(cards) != len(TestStack1) {
			t.Error("Cards scaned incorrectly")
		}
	}
}

func TestStartStack(t *testing.T) {
	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	cards := table.GetStartCards(1, models.PlayerPreset{Race: "nrth", Stack: TestStack1})
	if len(cards) < 11 {
		t.Error("Failed to get start cards")
		return
	}

	table.ActiveCards = append(table.ActiveCards, cards...)
	sort.SliceStable(table.ActiveCards,
		func(i, j int) bool {
			return table.ActiveCards[i].ID < table.ActiveCards[j].ID
		},
	)

	table.StartStack()

	if table.PlayerA.LeaderCard == nil {
		t.Error("Failed to recognize leader card")
	}

	if table.PlayerB.LeaderCard != nil ||
		len(table.PlayerB.Stack) != 0 {
		t.Error("Failed to recognize the owner")
	}

	switch {
	case table.PlayerA.Stack == nil:
		{
			t.Error("Failed to fill stack")
		}
	case len(table.PlayerA.Stack) != len(table.ActiveCards)-1:
		{
			t.Error("Stack filled incorrectly")
		}
	}
}

func TestStartHand(t *testing.T) {
	table := engine.NewTable(
		&models.Client{Name: "testname1"},
		&models.Client{Name: "testname2"},
	)

	cards := table.GetStartCards(1, models.PlayerPreset{Race: "nrth", Stack: TestStack1})
	table.ActiveCards = append(table.ActiveCards, cards...)
	cards = table.GetStartCards(2, models.PlayerPreset{Race: "nrth", Stack: TestStack2})
	table.ActiveCards = append(table.ActiveCards, cards...)
	sort.SliceStable(table.ActiveCards,
		func(i, j int) bool {
			return table.ActiveCards[i].ID < table.ActiveCards[j].ID
		},
	)

	table.StartStack()

	lenStackA := len(table.PlayerA.Stack)
	lenStackB := len(table.PlayerB.Stack)

	table.PlayerA.StartHand()
	table.PlayerB.StartHand()

	if len(table.PlayerA.Hand) != 10 || len(table.PlayerB.Hand) != 10 {
		t.Error("Failed to get hand")
	}

	if lenStackA-len(table.PlayerA.Stack) != 10 ||
		lenStackB-len(table.PlayerB.Stack) != 10 {
		t.Error("Something goes wrong")
	}
}

func TestInitGame(t *testing.T) {

	for i := 0; i < 10; i++ {
		table := engine.NewTable(
			&models.Client{Name: "testname1"},
			&models.Client{Name: "testname2"},
		)

		presets := make(map[string]models.PlayerPreset)
		presets["testname1"] = models.PlayerPreset{Race: "nrth", Stack: TestStack1}
		presets["testname2"] = models.PlayerPreset{Race: "nilf", Stack: TestStack2}

		table.InitGame(presets)

		if table.Pm.ActPlr != "testname2" {
			t.Error("Failed to define first player correctly")
		}

		if len(table.Pm.IDs) != 10 {
			t.Error("Incorrect table permission state")
		}
	}
}
