package engine

import (
	"errors"
	"slices"
)

func (t *Table) PutCard(cardID uint, targetfield string, targetID uint) error {
	card := t.Players[t.Pm.ActPlr].PickCardFromHand(cardID)
	var err error
	switch {
	case card == nil:
		fallthrough
	case !card.TargetField[targetfield]:
		{
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
	if err == nil {
		t.Players[t.Pm.ActPlr].DeleteCardFromHand(targetID)
	}
	return err
}

func (t *Table) PutWeatherCard(card *Card) error {

	if card.Name == Weather.Sun {
		t.WeatherFlags.Frost = false
		t.WeatherFlags.Fog = false
		t.WeatherFlags.Rain = false
		if len(t.PlayerA.WeatherField) != 0 {
			t.PlayerA.Grave = append(t.PlayerA.Grave, t.PlayerA.WeatherField...)
			t.PlayerA.WeatherField = nil
		}
		if len(t.PlayerB.WeatherField) != 0 {
			t.PlayerB.Grave = append(t.PlayerB.Grave, t.PlayerB.WeatherField...)
			t.PlayerB.WeatherField = nil
		}
		t.Players[t.Pm.ActPlr].PutCardToGrave(card)
		t.Pm.Instr = Instr.PmSwitch
		return nil
	}

	indexA := slices.IndexFunc(
		t.Players[t.Pm.ActPlr].WeatherField,
		func(c *Card) bool {
			return c.Name == card.Name
		},
	)
	indexP := slices.IndexFunc(
		t.Players[t.Pm.PasPlr].WeatherField,
		func(c *Card) bool {
			return c.Name == card.Name
		},
	)

	switch {
	case indexA >= 0:
		{
			t.Players[t.Pm.ActPlr].PutCardToGrave(card)
		}
	case indexP >= 0:
		{
			oldCard, err := t.Players[t.Pm.PasPlr].DeleteCardFromWeatherField(t.Players[t.Pm.PasPlr].WeatherField[indexP].ID)
			if err != nil {
				return err
			}
			t.Players[t.Pm.PasPlr].PutCardToGrave(oldCard)
			t.Players[t.Pm.ActPlr].PutCardOnWeatherField(card)
		}
	default:
		{
			t.Players[t.Pm.ActPlr].PutCardOnWeatherField(card)
		}
	}

	switch card.Name {
	case Weather.Frost:
		{
			t.WeatherFlags.Frost = true
		}
	case Weather.Fog:
		{
			t.WeatherFlags.Fog = true
		}
	case Weather.Rain:
		{
			t.WeatherFlags.Rain = true
		}
	}

	t.Pm.Instr = Instr.PmSwitch
	return nil
}

func (t *Table) PutDefaultCard(card *Card, targetfield string) error {
	t.Players[t.Pm.ActPlr].Fields[targetfield].PutCardOnField(card)
	t.Pm.Instr = Instr.PmSwitch
	return nil
}

func (t *Table) PutDecoyCard(card *Card, targetfield string, targetID uint) error {

	if targetID != 0 {
		exchanged, index := t.Players[t.Pm.ActPlr].Fields[targetfield].PickCardFromField(targetID)

		switch {
		case exchanged == nil:
			fallthrough
		case card.Rareness:
			fallthrough
		case card.Role == Role.Decoy:
			{

				return errors.New(Instr.ForbMove)
			}
		}

		t.Players[t.Pm.ActPlr].PutCardToHand(exchanged)
		t.Players[t.Pm.ActPlr].Fields[targetfield].CardField = slices.Replace(t.Players[t.Pm.ActPlr].Fields[targetfield].CardField, index, index+1, card)
	} else {
		t.PutDefaultCard(card, targetfield)
	}

	t.Pm.Instr = Instr.PmSwitch
	return nil
}

func (t *Table) PutSpyCard(card *Card, targetfield string) error {
	t.Players[t.Pm.PasPlr].Fields[targetfield].PutCardOnField(card)
	t.Players[t.Pm.ActPlr].PutRandCardFromStackToHand(2)
	t.Pm.Instr = Instr.PmSwitch
	return nil
}

func (t *Table) PutExecutorCard(card *Card, targetfield string) error {
	t.PutDefaultCard(card, targetfield)
	err := t.Execution(targetfield)
	if err != nil {
		return err
	}
	t.Pm.Instr = Instr.PmSwitch
	return nil
}

func (t *Table) PutExecutionCard(card *Card) error {
	t.Players[t.Pm.ActPlr].PutCardToGrave(card)
	err := t.Execution("")
	if err != nil {
		return err
	}
	t.Pm.Instr = Instr.PmSwitch
	return nil
}

func (t *Table) PutHealerCard(card *Card, targetfield string) error {
	t.PutDefaultCard(card, targetfield)
	healIDs := t.Players[t.Pm.ActPlr].GetIDsGrave(true)
	if healIDs != nil {
		t.Pm.IDs = healIDs
		t.Pm.Instr = Instr.HealCard
	} else {
		t.Pm.Instr = Instr.PmSwitch
	}
	return nil
}
