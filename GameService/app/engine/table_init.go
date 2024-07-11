package engine

import (
	"GwentMicroservices/GameService/app/api/models"
	log "GwentMicroservices/GameService/app/helpers/log"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"sort"
	"time"

	"github.com/google/uuid"
)

func NewTable(client1 *models.Client, client2 *models.Client) *Table {
	t := new(Table)
	t.TableID = uuid.New().String()
	t.PlayerA.Name = client1.Name
	t.PlayerA.Conn = client1.Conn
	t.PlayerB.Name = client2.Name
	t.PlayerB.Conn = client2.Conn

	t.PlayerA.InitPlayerField(client1)
	t.PlayerB.InitPlayerField(client2)

	t.Players = make(map[string]*PlayerField)
	t.Players = map[string]*PlayerField{
		t.PlayerA.Name: &t.PlayerA,
		t.PlayerB.Name: &t.PlayerB,
	}

	return t
}

func (pf *PlayerField) InitPlayerField(client *models.Client) {
	pf.Name = client.Name
	pf.Conn = client.Conn

	pf.AssaultField.ActiveBonuses.Squads = make(map[string]uint)
	pf.DistantField.ActiveBonuses.Squads = make(map[string]uint)
	pf.SiegeField.ActiveBonuses.Squads = make(map[string]uint)

	pf.Fields = make(map[string]*GameField)
	pf.Fields = map[string]*GameField{
		Field.Assault: &pf.AssaultField,
		Field.Distant: &pf.DistantField,
		Field.Siege:   &pf.SiegeField,
	}

}

////////////////////////////////////////////////////////////////////////

func (t *Table) InitGame(presets map[string]models.PlayerPreset) {

	t.PlayerA.WinTokens = 2
	t.PlayerB.WinTokens = 2

	t.PlayerA.Race = presets[t.PlayerA.Name].Race
	t.PlayerB.Race = presets[t.PlayerB.Name].Race

	cards := t.GetStartCards(1, presets[t.PlayerA.Name])
	if len(cards) < 1 {
		log.ServerLog(log.Error, "StartStack", "There is no cards to work with")
	}
	t.ActiveCards = append(t.ActiveCards, cards...)
	cards = t.GetStartCards(2, presets[t.PlayerB.Name])
	if len(cards) < 1 {
		log.ServerLog(log.Error, "StartStack", "There is no cards to work with")
	}
	t.ActiveCards = append(t.ActiveCards, cards...)
	sort.SliceStable(t.ActiveCards,
		func(i, j int) bool {
			return t.ActiveCards[i].ID < t.ActiveCards[j].ID
		},
	)

	t.StartStack()

	t.PlayerA.StartHand()
	t.PlayerB.StartHand()

	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	switch r.Intn(2) {
	case 0:
		{
			t.Pm.ActPlr = t.PlayerA.Name
			t.Pm.PasPlr = t.PlayerB.Name
		}
	case 1:
		{
			t.Pm.ActPlr = t.PlayerB.Name
			t.Pm.PasPlr = t.PlayerA.Name
		}
	}

	if t.Players[t.Pm.PasPlr].Race == Race.Nilf &&
		t.Players[t.Pm.ActPlr].Race != Race.Nilf {
		t.PermissionSwitch()
	}

	t.Pm.Instr = Instr.Move
	t.Pm.IDs = t.Players[t.Pm.ActPlr].GetIDsHand()
}

func (t *Table) GetStartCards(koef uint, preset models.PlayerPreset) []Card {
	var (
		all      []Card
		newstack []Card
	)

	storage, err := os.OpenFile(fmt.Sprintf("./cards/%s.json", preset.Race), os.O_RDONLY, 0666)
	if err != nil {
		log.ServerLog(log.Error, "GetStartCards", err.Error())
	}
	defer storage.Close()

	decoder := json.NewDecoder(storage)
	err = decoder.Decode(&all)
	if err != nil {
		log.ServerLog(log.Error, "GetStartCards", err.Error())
	}

	for _, card := range all {
		if slices.Contains(preset.Stack, card.ID) {
			card.ID += koef * 100
			newstack = append(newstack, card)
		}
	}

	return newstack
}

func (t *Table) StartStack() {
	for i := range t.ActiveCards {
		switch {
		case t.ActiveCards[i].Role == "leader":
			{
				if t.ActiveCards[i].ID < 200 {
					t.PlayerA.LeaderCard = &t.ActiveCards[i]
				} else {
					t.PlayerB.LeaderCard = &t.ActiveCards[i]
				}
			}
		default:
			{
				if t.ActiveCards[i].ID < 200 {
					t.PlayerA.Stack = append(t.PlayerA.Stack, &t.ActiveCards[i])
				} else {
					t.PlayerB.Stack = append(t.PlayerB.Stack, &t.ActiveCards[i])
				}
			}
		}
	}
}

func (pf *PlayerField) StartHand() {
	switch {
	default:
		{
			pf.PutRandCardFromStackToHand(10)
		}
	}

	sort.SliceStable(
		pf.Hand,
		func(i, j int) bool {
			return pf.Hand[i].Score < pf.Hand[j].Score
		},
	)
}
