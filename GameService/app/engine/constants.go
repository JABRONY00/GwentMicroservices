package engine

type role struct {
	Decoy     string
	Execution string
	Executor  string
	Healer    string
	Horn      string
	Leader    string
	Spy       string
	Weather   string
}

type field struct {
	Assault string
	Distant string
	Siege   string
}

type weather struct {
	Any   string
	Sun   string
	Frost string
	Fog   string
	Rain  string
}

type player struct {
	A   string
	B   string
	Act string
	Pas string
}

type race struct {
	Nrth string
	Nilf string
	Sctl string
	Mstr string
}

type instruction struct {
	ForbMove  string
	HealCard  string
	LBonus    string
	Meta      string
	Move      string
	Pass      string
	PmSwitch  string
	PutCard   string
	PickCard  string
	PickWCard string
	Refresh   string
	StealCard string
	Toss1     string
	Toss2     string
	Wait      string
}

type bonus struct {
	Boost     string
	Horn      string
	LeaderAct string
	LeaderPas string
	Squad     string
}

var Instr = instruction{
	ForbMove:  "forbidden-move",
	HealCard:  "heal-card",
	LBonus:    "leader-bonus",
	Meta:      "meta",
	Move:      "move",
	Pass:      "pass",
	PmSwitch:  "permission-switch",
	PutCard:   "put-card",
	PickCard:  "pick-card",
	PickWCard: "pick-weather-card",
	Refresh:   "refresh-table",
	StealCard: "steal-card",
	Toss1:     "toss1",
	Toss2:     "toss2",
	Wait:      "wait",
}

var Role = role{
	Decoy:     "decoy",
	Execution: "execution",
	Executor:  "executor",
	Healer:    "healer",
	Horn:      "horn",
	Leader:    "leader",
	Spy:       "spy",
	Weather:   "weather",
}

var Field = field{
	Assault: "assault-field",
	Distant: "distant-field",
	Siege:   "siege-field",
}

var Weather = weather{
	Any:   "any-weather",
	Sun:   "sun",
	Frost: "frost",
	Fog:   "fog",
	Rain:  "rain",
}

var Player = player{
	A:   "PlayerA",
	B:   "PlayerB",
	Act: "ActivePlayer",
	Pas: "PassivePlayer",
}

var Race = race{
	Nrth: "nrth",
	Nilf: "nilf",
	Sctl: "sctl",
	Mstr: "mstr",
}

var Bonus = bonus{
	Boost:     "boost",
	Horn:      "horn",
	LeaderAct: "leader-act",
	LeaderPas: "leader-pas",
	Squad:     "squad",
}
