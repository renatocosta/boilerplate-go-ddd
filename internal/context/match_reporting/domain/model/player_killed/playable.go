package players_killed

type PlayerInfo struct {
	WhoKilled, WhoDied string
	Kills              int64
}

type Players map[string]PlayerInfo

type Playable interface {
	KillUp(match Matchable)
	KillDown(match Matchable)
	GetPlayers() Players
}
