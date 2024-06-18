package engine

import "errors"

func (t *Table) PutWeatherCard(card *Card) error {
	switch card.Name {
	case Weather.Frost:
		{
			if t.WeatherFlags.Frost {
				t.Players[t.Pm.ActPlr].PutCardToGrave(card)
			}
			t.WeatherFlags.Frost = true
			t.Players[t.Pm.ActPlr].PutCardOnWeatherField(card)
		}
	case Weather.Fog:
		{
			if t.WeatherFlags.Fog {
				t.Players[t.Pm.ActPlr].PutCardToGrave(card)
			}
			t.WeatherFlags.Fog = true
			t.Players[t.Pm.ActPlr].PutCardOnWeatherField(card)
		}
	case Weather.Rain:
		{
			if t.WeatherFlags.Rain {
				t.Players[t.Pm.ActPlr].PutCardToGrave(card)
			}
			t.WeatherFlags.Rain = true
			t.Players[t.Pm.ActPlr].PutCardOnWeatherField(card)
		}
	case Weather.Sun:
		{
			t.WeatherFlags.Frost = false
			t.WeatherFlags.Fog = false
			t.WeatherFlags.Rain = false
			if len(t.PlayerA.WeatherField) != 0 {
				t.PlayerA.Grave = append(t.PlayerA.Grave, t.PlayerA.WeatherField...)
			}
			if len(t.PlayerB.WeatherField) != 0 {
				t.PlayerB.Grave = append(t.PlayerB.Grave, t.PlayerB.WeatherField...)
			}
			t.Players[t.Pm.ActPlr].PutCardToGrave(card)
		}
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

func (t *Table) PutDecoyCard(card *Card, targetfield string, targetID uint) error {
	t.PutDefaultCard(card, targetfield)
	if targetID != 0 {
		card, err := t.Players[t.Pm.ActPlr].Fields[targetfield].DeleteCardFromField(targetID)
		switch {
		case err != nil:
			fallthrough
		case card.Rareness:
			fallthrough
		case card.Role == Role.Decoy:
			{
				return errors.New(Instr.ForbMove)
			}
		}
		t.Players[t.Pm.ActPlr].PutCardToHand(card)
	}
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

func (t *Table) PutDefaultCard(card *Card, targetfield string) error {
	t.Players[t.Pm.ActPlr].Fields[targetfield].PutCardOnField(card)
	t.Pm.Instr = Instr.PmSwitch
	return nil
}
