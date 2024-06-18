package engine

func (t *Table) RefreshTable() {
	t.PlayerA.Conn.Mut.Lock()
	t.PlayerA.Conn.WriteJSON(ResponseData{Instr: Instr.Refresh, Data: t.GetTableInfo(t.PlayerA.Name)})
	t.PlayerA.Conn.Mut.Unlock()
	t.PlayerB.Conn.Mut.Lock()
	t.PlayerB.Conn.WriteJSON(ResponseData{Instr: Instr.Refresh, Data: t.GetTableInfo(t.PlayerB.Name)})
	t.PlayerB.Conn.Mut.Unlock()
}

func (t *Table) GetTableInfo(name string) TableShort {
	var (
		ts     TableShort
		player string
		enemy  string
	)
	switch {
	case name == t.PlayerA.Name:
		{
			player = t.PlayerA.Name
			enemy = t.PlayerB.Name
		}
	case name == t.PlayerB.Name:
		{
			player = t.PlayerB.Name
			enemy = t.PlayerA.Name
		}
	}
	if t.Winner == player {
		ts.Winner = true
	}
	if t.Pm.ActPlr == player {
		ts.ActivePlayer = true
	} else {
		ts.ActivePlayer = false
	}
	ts.WinTokens.Player = t.Players[player].WinTokens
	ts.WinTokens.Enemy = t.Players[enemy].WinTokens

	ts.WeatherField.Frost = t.WeatherFlags.Frost
	ts.WeatherField.Fog = t.WeatherFlags.Fog
	ts.WeatherField.Rain = t.WeatherFlags.Rain
	ts.WeatherField.Cards = t.Players[player].GetIDsWeather()
	ts.WeatherField.Cards = append(ts.WeatherField.Cards, t.Players[enemy].GetIDsWeather()...)

	ts.Assaultfield.Player, ts.Assaultfield.HornPlayer = t.Players[player].AssaultField.GetCardsForRefresher()
	ts.Distantfield.Player, ts.Distantfield.HornPlayer = t.Players[player].DistantField.GetCardsForRefresher()
	ts.SiegeField.Player, ts.SiegeField.HornPlayer = t.Players[player].SiegeField.GetCardsForRefresher()
	ts.Assaultfield.Enemy, ts.Assaultfield.HornEnemy = t.Players[enemy].AssaultField.GetCardsForRefresher()
	ts.Distantfield.Enemy, ts.Distantfield.HornEnemy = t.Players[enemy].DistantField.GetCardsForRefresher()
	ts.SiegeField.Enemy, ts.SiegeField.HornEnemy = t.Players[enemy].SiegeField.GetCardsForRefresher()

	ts.Gravefield.Player = t.Players[player].GetIDsGrave(false)
	ts.Gravefield.Enemy = t.Players[enemy].GetIDsGrave(false)

	ts.Stack.Player = t.Players[player].GetIDsStack()
	ts.Stack.Enemy = uint(len(t.Players[enemy].Stack))

	ts.Hand.Player = t.Players[player].GetIDsHand()
	ts.Hand.Enemy = uint(len(t.Players[enemy].Hand))

	ts.Leaderfield.Player = t.Players[player].LeaderFlag
	ts.Leaderfield.Enemy = t.Players[enemy].LeaderFlag

	return ts
}

func (gf *GameField) GetCardsForRefresher() ([]CardShort, bool) {
	var cshort []CardShort

	for _, card := range gf.CardField {
		var cs CardShort
		cs.ID = card.ID
		cs.Score = card.Score
		switch {
		case card.Score == card.Cost:
			{
				cs.Status = "black"
			}
		case card.Score > card.Cost:
			{
				cs.Status = "green"
			}
		case card.Score < card.Cost:
			{
				cs.Status = "red"
			}
		}
		cshort = append(cshort, cs)
	}
	var horn bool
	if gf.HornField != nil {
		horn = true
	}
	return cshort, horn
}
