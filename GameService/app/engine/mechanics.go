package engine

import (
	"GwentMicroservices/GameService/app/api/models"
	"errors"
	"math"
	"math/rand"
	"slices"
	"sort"
	"sync"
	"time"
)

func (t *Table) EndGame() {
	t.PlayerA.SendJson(models.ResponseData{Instr: "end-game"})
	t.PlayerB.SendJson(models.ResponseData{Instr: "end-game"})

	if t.Winner != "" {

		t.Players[t.Winner].SendJson("You Win")

	}

	if t.PlayerA.Name != t.Winner {
		t.PlayerA.SendJson("You Lose")
	}

	if t.PlayerB.Name != t.Winner {
		t.PlayerB.SendJson("You Lose")
	}
}

func (t *Table) EndRound() {
	if t.Winner != t.PlayerA.Name {
		t.PlayerA.WinTokens--
	}
	if t.Winner != t.PlayerB.Name {
		t.PlayerB.WinTokens--
	}
	t.WeatherFlags.Frost = false
	t.WeatherFlags.Fog = false
	t.WeatherFlags.Rain = false
	/*t.PlayerA.EndRound()
	t.PlayerB.EndRound()*/
	switch {
	case t.PlayerA.WinTokens == 0:
		fallthrough
	case t.PlayerB.WinTokens == 0:
		{
			t.EndGame()
		}
	}
}

func (t *Table) Pass() {
	if t.Players[t.Pm.PasPlr].PassFlag {
		t.EndRound()
	}
	t.PermissionSwitch()
	t.Players[t.Pm.PasPlr].PassFlag = true
}

/////////////////////////////////////
/////////////////////////////////////

func (t *Table) CardByID(id uint) *Card {
	if id == 0 {
		return nil
	}
	leftside := 0
	rightside := len(t.ActiveCards) - 1
	for leftside <= rightside {
		middle := (leftside + rightside) / 2
		if t.ActiveCards[middle].ID == id {
			return (&t.ActiveCards[middle])
		} else {
			if t.ActiveCards[middle].ID < id {
				leftside = middle + 1
			} else {
				rightside = middle - 1
			}
		}
	}
	return nil
}

func (t *Table) PermissionSwitch() {
	if !t.Players[t.Pm.PasPlr].PassFlag {
		buf := t.Pm.ActPlr
		t.Pm.ActPlr = t.Pm.PasPlr
		t.Pm.PasPlr = buf
	}
	t.Pm.IDs = t.Players[t.Pm.ActPlr].GetIDsHand()
	t.Pm.Instr = Instr.Move
}

////////////////////////////
////////////////////////////

func (t *Table) TableScoreCounter() {

	t.MaxCardScore = 0

	twg := &sync.WaitGroup{}
	twg.Add(2)

	go t.Players[t.Pm.ActPlr].PlayerFieldScoreCounter(twg, t.WeatherFlags.Frost, t.WeatherFlags.Fog, t.WeatherFlags.Rain)
	go t.Players[t.Pm.PasPlr].PlayerFieldScoreCounter(twg, t.WeatherFlags.Frost, t.WeatherFlags.Fog, t.WeatherFlags.Rain)
	twg.Wait()

	if t.Players[t.Pm.ActPlr].MaxCardScore > t.Players[t.Pm.PasPlr].MaxCardScore {
		t.MaxCardScore = t.Players[t.Pm.ActPlr].MaxCardScore
	} else {
		t.MaxCardScore = t.Players[t.Pm.PasPlr].MaxCardScore
	}

	switch {
	case int(t.Players[t.Pm.ActPlr].Score-t.Players[t.Pm.PasPlr].Score) > 0:
		{
			t.Winner = t.Pm.ActPlr
		}
	case int(t.Players[t.Pm.ActPlr].Score-t.Players[t.Pm.PasPlr].Score) < 0:
		{
			t.Winner = t.Pm.PasPlr
		}
	case t.PlayerA.Score-t.PlayerB.Score == 0:
		{
			t.Winner = "None"
			if t.Players[t.Pm.ActPlr].Race == Race.Nilf && t.Players[t.Pm.ActPlr].Race != t.Players[t.Pm.PasPlr].Race {
				t.Winner = t.Pm.ActPlr
			}
			if t.Players[t.Pm.PasPlr].Race == Race.Nilf && t.Players[t.Pm.PasPlr].Race != t.Players[t.Pm.ActPlr].Race {
				t.Winner = t.Pm.PasPlr
			}
		}
	}
}

func (pf *PlayerField) PlayerFieldScoreCounter(twg *sync.WaitGroup, frost bool, fog bool, rain bool) {
	defer twg.Done()

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go pf.Fields[Field.Assault].GameFieldBonusCounter(wg, frost)
	go pf.Fields[Field.Distant].GameFieldBonusCounter(wg, fog)
	go pf.Fields[Field.Siege].GameFieldBonusCounter(wg, rain)
	wg.Wait()

	wg.Add(3)
	pf.Score = pf.Fields[Field.Assault].GameFieldScoreCounter(wg)
	pf.Score += pf.Fields[Field.Distant].GameFieldScoreCounter(wg)
	pf.Score += pf.Fields[Field.Siege].GameFieldScoreCounter(wg)
	wg.Wait()

	pf.MaxCardScore = pf.Fields[Field.Assault].MaxCardScore
	if pf.MaxCardScore < pf.Fields[Field.Distant].MaxCardScore {
		pf.MaxCardScore = pf.Fields[Field.Distant].MaxCardScore
	}
	if pf.MaxCardScore < pf.Fields[Field.Siege].MaxCardScore {
		pf.MaxCardScore = pf.Fields[Field.Siege].MaxCardScore
	}
}

func (gf *GameField) GameFieldBonusCounter(wg *sync.WaitGroup, weather bool) {
	defer wg.Done()

	if weather {
		gf.ActiveBonuses.Weather = 1
	} else {
		gf.ActiveBonuses.Weather = 0
	}
	gf.ActiveBonuses.Squads = make(map[string]uint)
	gf.ActiveBonuses.Horn = 0
	gf.ActiveBonuses.Boost = 0

	if gf.HornField != nil {
		gf.ActiveBonuses.Horn = 1
	}

	for i := range gf.CardField {
		if gf.CardField[i].CardBonus.Squad != "" {
			koef, ok := gf.ActiveBonuses.Squads[gf.CardField[i].Squad]
			if ok {
				gf.ActiveBonuses.Squads[gf.CardField[i].Squad] = koef + 1
			} else {
				gf.ActiveBonuses.Squads[gf.CardField[i].Squad] = 1
			}

		}
		if gf.CardField[i].CardBonus.Horn {
			gf.ActiveBonuses.Horn++
		}
		if gf.CardField[i].CardBonus.Boost {
			gf.ActiveBonuses.Boost++
		}
	}
}

func (gf *GameField) GameFieldScoreCounter(wg *sync.WaitGroup) uint {
	defer wg.Done()

	gf.MaxCardScore = 0
	gf.Score = 0
	squadKoef := 1
	hornKoef := gf.ActiveBonuses.Horn
	hornfix := uint(0)
	weatherfix := uint(0)
	boostfix := uint(0)

	switch {
	case len(gf.CardField) < 1:
		{
			return 0
		}
	case len(gf.CardField) > 0:
		{
			for i := range gf.CardField {
				if gf.CardField[i].Rareness {
					gf.CardField[i].Score = gf.CardField[i].Cost
					gf.Score += gf.CardField[i].Score
					continue
				}
				if gf.CardField[i].CardBonus.Squad != "" {
					squadKoef = int(gf.ActiveBonuses.Squads[gf.CardField[i].CardBonus.Squad])
				}
				if gf.CardField[i].Cost == 0 && gf.ActiveBonuses.Weather == 1 {
					weatherfix = 1
				}
				if gf.CardField[i].CardBonus.Horn && gf.HornField == nil && gf.ActiveBonuses.Horn < 2 {
					hornfix = 1
				}
				if hornKoef > 1 {
					hornKoef = 1
				}
				if gf.CardField[i].CardBonus.Boost {
					boostfix = 1
				} else {
					boostfix = 0
				}

				gf.CardField[i].Score = ((gf.CardField[i].Cost-gf.ActiveBonuses.Weather*(gf.CardField[i].Cost-1)-weatherfix)*uint(squadKoef) + (gf.ActiveBonuses.Boost - boostfix)) * uint(math.Pow(2, float64(hornKoef-hornfix)))
				gf.Score += gf.CardField[i].Score
				if gf.CardField[i].Score > gf.MaxCardScore {
					gf.MaxCardScore = gf.CardField[i].Score
				}

				squadKoef = 1
				hornfix = 0
				weatherfix = 0
				hornKoef = gf.ActiveBonuses.Horn
			}
		}
	}
	return gf.Score
}

//////////////////////////
//////////////////////////

func (pf *PlayerField) PutRandCardFromStackToHand(number uint) {
	for number != 0 {
		if len(pf.Stack) < 1 {
			return
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		index := r.Intn(len(pf.Stack))
		pf.Hand = append(pf.Hand, pf.Stack[index])
		pf.Stack[index] = nil
		pf.Stack = slices.Delete(pf.Stack, index, index+1)
		pf.Stack = slices.Clip(pf.Stack)
		number--
	}
	if len(pf.Hand) > 1 {
		sort.SliceStable(pf.Hand,
			func(i, j int) bool { return pf.Hand[i].Cost < pf.Hand[j].Cost })
	}
}

func (pf *PlayerField) PutCardToGrave(card *Card) {
	pf.Grave = append(pf.Grave, card)
	if len(pf.Grave) > 1 {
		sort.SliceStable(pf.Grave,
			func(i, j int) bool { return pf.Grave[i].Cost < pf.Grave[j].Cost })
	}
}

func (pf *PlayerField) DeleteCardFromGrave(targetID uint) (*Card, error) {
	for i, card := range pf.Grave {
		if card.ID == targetID {
			pf.Grave = slices.Delete(pf.Grave, i, i+1)
			pf.Grave = slices.Clip(pf.Grave)
			return card, nil
		}
	}
	return nil, errors.New(Instr.ForbMove)
}

func (pf *PlayerField) PutCardToHand(card *Card) {
	pf.Hand = append(pf.Hand, card)
	if len(pf.Hand) > 1 {
		sort.SliceStable(pf.Hand,
			func(i, j int) bool { return pf.Hand[i].Cost < pf.Hand[j].Cost })
	}
}

func (pf *PlayerField) DeleteCardFromHand(targetID uint) (*Card, error) {
	for i, card := range pf.Hand {
		if card.ID == targetID {
			pf.Hand = slices.Delete(pf.Hand, i, i+1)
			pf.Hand = slices.Clip(pf.Hand)
			return card, nil
		}
	}
	return nil, errors.New(Instr.ForbMove)
}

func (pf *PlayerField) PutCardOnWeatherField(card *Card) {
	pf.WeatherField = append(pf.WeatherField, card)
}

func (pf *PlayerField) DeleteCardFromWeatherField(targetID uint) (*Card, error) {
	for i, card := range pf.WeatherField {
		if card.ID == targetID {
			pf.WeatherField = slices.Delete(pf.WeatherField, i, i+1)
			pf.WeatherField = slices.Clip(pf.WeatherField)
			return card, nil
		}
	}
	return nil, errors.New(Instr.ForbMove)
}

func (pf *PlayerField) PutCardToStack(card *Card) {
	pf.Stack = append(pf.Stack, card)
}

func (pf *PlayerField) DeleteCardFromStack(targetID uint) (*Card, error) {
	for i, card := range pf.Stack {
		if card.ID == targetID {
			pf.Stack = slices.Delete(pf.Stack, i, i+1)
			pf.Stack = slices.Clip(pf.Stack)
			return card, nil
		}
	}
	return nil, errors.New(Instr.ForbMove)
}

func (gf *GameField) PutCardOnField(card *Card) {
	gf.CardField = append(gf.CardField, card)
	if len(gf.CardField) > 1 {
		sort.SliceStable(gf.CardField,
			func(i, j int) bool { return gf.CardField[i].Cost < gf.CardField[j].Cost })
	}
}

func (gf *GameField) DeleteCardFromField(targetID uint) (*Card, error) {
	for i, card := range gf.CardField {
		if card.ID == targetID {
			gf.CardField = slices.Delete(gf.CardField, i, i+1)
			return card, nil
		}
	}
	return nil, errors.New(Instr.ForbMove)
}

/////////////////
/////////////////

func (t *Table) Execution(targetfield string) error {
	switch targetfield {
	case Field.Assault:
		fallthrough
	case Field.Distant:
		fallthrough
	case Field.Siege:
		{
			executionIDs := t.Players[t.Pm.PasPlr].Fields[targetfield].GetIDsExecution(false, 0)
			for _, ID := range executionIDs {
				card, err := t.Players[t.Pm.PasPlr].Fields[targetfield].DeleteCardFromField(ID)
				if err != nil {
					return err
				}
				t.Players[t.Pm.PasPlr].PutCardToGrave(card)
			}
		}
	default: //Global Execution
		{
			if t.PlayerA.MaxCardScore == t.MaxCardScore {
				t.PlayerA.GlobalExecution(t.MaxCardScore)
			}
			if t.PlayerB.MaxCardScore == t.MaxCardScore {
				t.PlayerB.GlobalExecution(t.MaxCardScore)
			}
		}
	}
	return nil
}

func (pf *PlayerField) GlobalExecution(maxScore uint) error {
	var executionIDs []uint
	if pf.AssaultField.MaxCardScore == maxScore {
		executionIDs = pf.AssaultField.GetIDsExecution(true, maxScore)
		for _, ID := range executionIDs {
			card, err := pf.AssaultField.DeleteCardFromField(ID)
			if err != nil {
				return err
			}
			pf.PutCardToGrave(card)
		}
	}
	if pf.DistantField.MaxCardScore == maxScore {
		executionIDs = pf.DistantField.GetIDsExecution(true, maxScore)
		for _, ID := range executionIDs {
			card, err := pf.DistantField.DeleteCardFromField(ID)
			if err != nil {
				return err
			}
			pf.PutCardToGrave(card)
		}
	}
	if pf.SiegeField.MaxCardScore == maxScore {
		executionIDs = pf.SiegeField.GetIDsExecution(true, maxScore)
		for _, ID := range executionIDs {
			card, err := pf.SiegeField.DeleteCardFromField(ID)
			if err != nil {
				return err
			}
			pf.PutCardToGrave(card)
		}
	}
	return nil
}

//////////////////
//////////////////

func (pf *PlayerField) GetIDsHand() []uint {
	var returnIDs []uint
	for _, card := range pf.Hand {
		returnIDs = append(returnIDs, card.ID)
	}
	return returnIDs
}

func (pf *PlayerField) GetIDsStack() []uint {
	var returnIDs []uint
	for _, card := range pf.Stack {
		returnIDs = append(returnIDs, card.ID)
	}
	return returnIDs
}

func (pf *PlayerField) GetIDsWeather() []uint {
	var returnIDs []uint
	for _, card := range pf.WeatherField {
		returnIDs = append(returnIDs, card.ID)
	}
	return returnIDs
}

func (pf *PlayerField) GetIDsGrave(forHeal bool) []uint {
	var returnIDs []uint
	for _, card := range pf.Grave {
		if forHeal {
			switch {
			case card.Rareness:
				fallthrough
			case card.Role == Role.Decoy:
				fallthrough
			case card.Role == Role.Execution:
				fallthrough
			case card.Role == Role.Horn:
				fallthrough
			case card.Role == Role.Weather:
				continue
			}

		} else {
			returnIDs = append(returnIDs, card.ID)
		}
	}
	return returnIDs
}

func (pf *PlayerField) GetIDsLeaderWeather() []uint {
	var returnIDs []uint
	for _, card := range pf.Stack {
		switch pf.LeaderCard.LeaderBonus {
		case Weather.Frost:
			fallthrough
		case Weather.Fog:
			fallthrough
		case Weather.Rain:
			{
				if card.Name == pf.LeaderCard.LeaderBonus {
					returnIDs = append(returnIDs, card.ID)
					break
				}
			}
		default:
			{
				if card.Role == Role.Weather {
					returnIDs = append(returnIDs, card.ID)
				}
			}
		}
	}
	return returnIDs
}

func (gf *GameField) GetIDsExecution(global bool, maxScore uint) []uint {
	var executionIDs []uint
	if global {
		if gf.MaxCardScore < maxScore {
			return nil
		}
		for _, card := range gf.CardField {
			if card.Score == gf.MaxCardScore && !card.Rareness && card.Role != Role.Decoy {
				executionIDs = append(executionIDs, card.ID)
			}
		}
		return executionIDs
	} else {
		if gf.Score >= 10 {
			for _, card := range gf.CardField {
				if card.Score == gf.MaxCardScore && !card.Rareness && card.Role != Role.Decoy {
					executionIDs = append(executionIDs, card.ID)
				}
			}
			return executionIDs
		}
		return nil
	}
}
