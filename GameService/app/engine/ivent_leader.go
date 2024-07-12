package engine

import "errors"

func (t *Table) LeaderWeather() error {
	switch t.Players[t.Pm.ActPlr].LeaderCard.Bonuses[Bonus.LeaderAct] {
	case Weather.Sun:
		{
			if !(t.WeatherFlags.Frost || t.WeatherFlags.Fog || t.WeatherFlags.Rain) {
				return errors.New(Instr.ForbMove)
			}
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
		}
	case Weather.Any:
		{
			t.Pm.IDs = t.Players[t.Pm.ActPlr].GetIDsLeaderWeather()
			if t.Pm.IDs == nil {
				return errors.New(Instr.ForbMove)
			}
			t.Pm.Instr = Instr.PickWCard
		}
	case Weather.Frost:
		fallthrough
	case Weather.Fog:
		fallthrough
	case Weather.Rain:
		{
			t.Pm.IDs = t.Players[t.Pm.ActPlr].GetIDsLeaderWeather()
			if t.Pm.IDs == nil {
				return errors.New(Instr.ForbMove)
			}
			card, err := t.Players[t.Pm.ActPlr].DeleteCardFromStack(t.Pm.IDs[0])
			if err != nil {
				return err
			}
			err = t.PutWeatherCard(card)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *Table) LeaderExchange() error {
	switch {
	case len(t.Players[t.Pm.ActPlr].Hand) < 2:
		fallthrough
	case len(t.Players[t.Pm.ActPlr].Stack) < 1:
		{
			return errors.New(Instr.ForbMove)
		}
	}
	t.Pm.IDs = t.Players[t.Pm.ActPlr].GetIDsHand()
	t.Pm.Instr = Instr.Toss1
	return nil
}

func (t *Table) LeaderHorn() error {
	return nil
}
