package engine

import (
	"GwentMicroservices/GameService/app/api/models"
	log "GwentMicroservices/GameService/app/helpers/log"
	"encoding/json"
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
	t.PlayerA.WinTokens = 2
	t.PlayerB.WinTokens = 2

	t.PlayerA.Fields = make(map[string]*GameField)
	t.PlayerA.Fields = map[string]*GameField{
		Field.Assault: &t.PlayerA.AssaultField,
		Field.Distant: &t.PlayerA.DistantField,
		Field.Siege:   &t.PlayerA.SiegeField,
	}
	t.PlayerA.Fields[Field.Assault].ActiveBonuses.Squads = make(map[string]uint)
	t.PlayerA.Fields[Field.Distant].ActiveBonuses.Squads = make(map[string]uint)
	t.PlayerA.Fields[Field.Siege].ActiveBonuses.Squads = make(map[string]uint)

	t.PlayerB.Fields = make(map[string]*GameField)
	t.PlayerB.Fields = map[string]*GameField{
		Field.Assault: &t.PlayerB.AssaultField,
		Field.Distant: &t.PlayerB.DistantField,
		Field.Siege:   &t.PlayerB.SiegeField,
	}
	t.PlayerB.Fields[Field.Assault].ActiveBonuses.Squads = make(map[string]uint)
	t.PlayerB.Fields[Field.Distant].ActiveBonuses.Squads = make(map[string]uint)
	t.PlayerB.Fields[Field.Siege].ActiveBonuses.Squads = make(map[string]uint)

	t.Players = make(map[string]*PlayerField)
	t.Players = map[string]*PlayerField{
		t.PlayerA.Name: &t.PlayerA,
		t.PlayerB.Name: &t.PlayerB,
	}
	return t
}

func (t *Table) InitTable(presets *map[string]models.PlayerPreset) {

	/*sort.SliceStable(t.ActiveCards,
	func(i, j int) bool { return t.ActiveCards[i].ID < t.ActiveCards[j].ID })*/
	t.StartStack(presets)
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
	t.Pm.Instr = Instr.ForbMove
	t.Pm.IDs = t.Players[t.Pm.ActPlr].GetIDsHand()
	t.Players[t.Pm.ActPlr].Conn.Mut.Lock()
	t.Players[t.Pm.ActPlr].Conn.WriteJSON("Game is running")
	t.Players[t.Pm.ActPlr].Conn.WriteJSON(models.ResponseData{Instr: Instr.Move, Data: t.Pm.IDs})
	t.Players[t.Pm.ActPlr].Conn.Mut.Unlock()
	t.Players[t.Pm.PasPlr].Conn.Mut.Lock()
	t.Players[t.Pm.PasPlr].Conn.WriteJSON("Game is running")
	t.Players[t.Pm.PasPlr].Conn.WriteJSON(models.ResponseData{Instr: Instr.Wait})
	t.Players[t.Pm.PasPlr].Conn.Mut.Unlock()
}

func (t *Table) StartStack(presets *map[string]models.PlayerPreset) {
	for i := range t.ActiveCards {
		t.ActiveCards[i].Score = t.ActiveCards[i].Cost
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

	sort.SliceStable(pf.Hand,
		func(i, j int) bool { return pf.Hand[i].Score < pf.Hand[j].Score })
}

func (pf *PlayerField) GetStartCards(koef uint, preset []uint, source string) []Card {
	var (
		all      []Card
		newstack []Card
	)
	storage, err := os.OpenFile(source, os.O_RDONLY, 0666)
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
		if slices.Contains(preset, card.ID) {
			card.ID += koef
			newstack = append(newstack, card)
		}
	}
	return newstack
}
