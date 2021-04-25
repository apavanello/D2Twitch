package main

type Match struct {
	Match      MatchClass `json:"match"`
	Teams      []Team     `json:"teams"`
	Buildings  []Building `json:"buildings"`
	GraphData  GraphData  `json:"graph_data"`
	DeltaFrame bool       `json:"delta_frame"`
}

type Building struct {
	Team      int64 `json:"team"`
	Heading   int64 `json:"heading"`
	Type      int64 `json:"type"`
	Lane      int64 `json:"lane"`
	Tier      int64 `json:"tier"`
	X         int64 `json:"x"`
	Y         int64 `json:"y"`
	Destroyed bool  `json:"destroyed"`
}

type GraphData struct {
	GraphGold []int64 `json:"graph_gold"`
}

type MatchClass struct {
	ServerSteamID string `json:"server_steam_id"`
	Matchid       string `json:"matchid"`
	Timestamp     int64  `json:"timestamp"`
	GameTime      int64  `json:"game_time"`
	GameMode      int64  `json:"game_mode"`
	LeagueID      int64  `json:"league_id"`
	LeagueNodeID  int64  `json:"league_node_id"`
	GameState     int64  `json:"game_state"`
	Picks         []Ban  `json:"picks"`
	Bans          []Ban  `json:"bans"`
}

type Ban struct {
	Hero int64 `json:"hero"`
	Team int64 `json:"team"`
}

type Team struct {
	TeamNumber  int64    `json:"team_number"`
	TeamID      int64    `json:"team_id"`
	TeamName    string   `json:"team_name"`
	TeamTag     string   `json:"team_tag"`
	TeamLogo    string   `json:"team_logo"`
	Score       int64    `json:"score"`
	NetWorth    int64    `json:"net_worth"`
	TeamLogoURL string   `json:"team_logo_url"`
	Players     []Player `json:"players"`
}

type Player struct {
	Accountid    int64  `json:"accountid"`
	Playerid     int64  `json:"playerid"`
	Name         string `json:"name"`
	Team         int64  `json:"team"`
	Heroid       int64  `json:"heroid"`
	Level        int64  `json:"level"`
	KillCount    int64  `json:"kill_count"`
	DeathCount   int64  `json:"death_count"`
	AssistsCount int64  `json:"assists_count"`
	DeniesCount  int64  `json:"denies_count"`
	LhCount      int64  `json:"lh_count"`
	Gold         int64  `json:"gold"`
	X            int64  `json:"x"`
	Y            int64  `json:"y"`
	NetWorth     int64  `json:"net_worth"`
}
