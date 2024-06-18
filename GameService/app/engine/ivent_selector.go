package engine

import (
	"errors"
	"sort"
)

func (t *Table) MoveRouter(reqBody RequestData) error {
	var err error
	switch reqBody.Instr {
	case Instr.Pass:
		{
			t.Pm.Instr = Instr.Pass
		}
	case Instr.LBonus:
		{
			err = t.LeaderBonusActive(reqBody.CardID)
		}
	case Instr.PutCard:
		{
			err = t.PutCard(reqBody.CardID, reqBody.TargetField, reqBody.TargetID)
		}
	case Instr.HealCard:
	case Instr.Toss1:
		fallthrough
	case Instr.Toss2:
		{
			err = t.TossCard(reqBody.CardID)
		}
	default:
		{
			return errors.New(Instr.ForbMove)
		}
	}
	if err != nil {
		return err
	}
	t.TableScoreCounter()
	if len(t.Players[t.Pm.ActPlr].Hand) == 0 && !t.Players[t.Pm.ActPlr].LeaderFlag {
		t.Pm.Instr = Instr.Pass
	}
	switch t.Pm.Instr {
	case Instr.Pass:
		{
			t.Pass()
		}
	case Instr.PmSwitch:
		{
			t.PermissionSwitch()
		}
	}
	return nil
}

//////////// HANDLERS

func (t *Table) LeaderBonusActive(cardID uint) error { ///Expected
	if !t.Players[t.Pm.ActPlr].LeaderFlag {
		return errors.New(Instr.ForbMove)
	}
	card := t.Players[t.Pm.ActPlr].LeaderCard
	var err error
	switch card.CardBonus.LeaderBonus {
	case Weather.Sun:
		fallthrough
	case Weather.Frost:
		fallthrough
	case Weather.Fog:
		fallthrough
	case Weather.Rain:
		fallthrough
	case Weather.Any:
		{
			err = t.LeaderWeather()
		}

	case "horn-assault":
		fallthrough
	case "horn-distant":
		fallthrough
	case "horn-siege":
		{
			err = t.LeaderHorn()
		}

	case "kill-assault":
		{
			err = t.Execution(Field.Assault)
		}
	case "kill-siege":
		{
			err = t.Execution(Field.Siege)
		}
	case "kill-distant":
		{
			err = t.Execution(Field.Distant)
		}
	case "show-cards":
		{

		}
	case "Skillrestriction":
		{
			t.PlayerA.LeaderFlag = false
			t.PlayerB.LeaderFlag = false
		}
	case "steal-card":
		{

		}
	case "exchange-card":
		{
			err = t.LeaderExchange()
		}
	}

	return err
}

func (t *Table) PutCard(cardID uint, targetfield string, targetID uint) error {
	card, err := t.Players[t.Pm.ActPlr].DeleteCardFromHand(cardID)
	switch {
	case err != nil:
		fallthrough
	case !card.TargetField[targetfield]:
		{
			t.Players[t.Pm.ActPlr].PutCardToHand(card)
			return errors.New(Instr.ForbMove)
		}
	}
	switch card.Role {
	case Role.Weather: // Complete
		{
			err = t.PutWeatherCard(card)
		}
	case Role.Decoy:
		{
			err = t.PutDecoyCard(card, targetfield, targetID)
		}
	case Role.Spy: // Complete
		{
			err = t.PutSpyCard(card, targetfield)
		}
	case Role.Executor: // Complete
		{
			err = t.PutExecutorCard(card, targetfield)
		}
	case Role.Execution: // ! Complete
		{
			err = t.PutExecutionCard(card)
		}
	case Role.Healer: //Expected
		{
			err = t.PutHealerCard(card, targetfield)
		}
	default: // Complete
		{
			err = t.PutDefaultCard(card, targetfield)
		}
	}
	return err
}

func (t *Table) HealCard(cardID uint, targetfield string) error {
	card, err := t.Players[t.Pm.ActPlr].DeleteCardFromGrave(cardID)
	switch {
	case err != nil:
		fallthrough
	case !card.TargetField[targetfield]:
		{
			return errors.New(Instr.ForbMove)
		}
	}
	switch card.Role {
	case Role.Healer:
		{
			err = t.PutHealerCard(card, targetfield)
		}
	case Role.Executor:
		{
			err = t.PutExecutorCard(card, targetfield)
		}
	default:
		{
			err = t.PutDefaultCard(card, targetfield)
		}
	}
	return err
}

func (t *Table) TossCard(cardID uint) error {
	card, err := t.Players[t.Pm.ActPlr].DeleteCardFromHand(cardID)
	if err != nil {
		return err
	}
	t.Players[t.Pm.ActPlr].PutCardToGrave(card)
	switch {
	case t.Pm.Instr == Instr.Toss1:
		{
			t.Pm.Instr = Instr.Toss2
			t.Pm.IDs = t.Players[t.Pm.ActPlr].GetIDsHand()
		}
	case t.Pm.Instr == Instr.Toss2:
		{
			t.Pm.Instr = Instr.PickCard
			t.Pm.IDs = t.Players[t.Pm.ActPlr].GetIDsHand()
		}
	}
	return nil
}

func (t *Table) PickCard(cardID uint) error {
	card, err := t.Players[t.Pm.ActPlr].DeleteCardFromStack(cardID)
	if err != nil {
		return err
	}
	t.Players[t.Pm.ActPlr].PutCardToHand(card)
	if len(t.Players[t.Pm.ActPlr].Hand) > 1 {
		sort.SliceStable(t.Players[t.Pm.ActPlr].Hand,
			func(i, j int) bool {
				return t.Players[t.Pm.ActPlr].Hand[i].Cost < t.Players[t.Pm.ActPlr].Hand[j].Cost
			})
	}
	return nil
}

func (t *Table) PickWeatherCard(cardID uint) error {
	card, err := t.Players[t.Pm.ActPlr].DeleteCardFromStack(cardID)
	if err != nil {
		return err
	}
	err = t.PutWeatherCard(card)
	return err
}

func (t *Table) StealCard() {
	if len(t.Players[t.Pm.ActPlr].Hand) > 1 {
		sort.SliceStable(t.Players[t.Pm.ActPlr].Hand, func(i, j int) bool {
			return t.Players[t.Pm.ActPlr].Hand[i].Cost < t.Players[t.Pm.ActPlr].Hand[j].Cost
		})
	}
}

func (t *Table) MetaResponse(player string, cardID uint) {
	pointer := t.CardByID(cardID)
	if pointer != nil {
		t.Players[player].Conn.Mut.Lock()
		t.Players[player].Conn.WriteJSON(ResponseData{Instr: "meta", Data: *pointer})
		t.Players[player].Conn.Mut.Unlock()
		return
	}
}
