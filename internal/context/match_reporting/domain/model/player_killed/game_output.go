package players_killed

import (
	"sort"
)

type GameEntity struct {
	PlayersIn  Players
	PlayersOut []PlayerInfo
}

type GameInfo struct {
	TotalKills int64            `json:"total_kills"`
	Players    []string         `json:"players"`
	Kills      map[string]int64 `json:"kills"`
}

func NewGame(p Players) Game {
	delete(p, WorldKillerLabel)
	return &GameEntity{PlayersIn: p}
}

type Game interface {
	Sort()
	TotalKills() int64
	PlayerNames() []string
	PlayerKills() map[string]int64
}

func (g *GameEntity) Sort() {

	for _, player := range g.PlayersIn {
		g.PlayersOut = append(g.PlayersOut, player)
	}

	sort.SliceStable(g.PlayersOut, func(i, j int) bool {
		return g.PlayersOut[i].Kills > g.PlayersOut[j].Kills
	})
}

func (g *GameEntity) TotalKills() int64 {
	totalKills := int64(0)

	for _, player := range g.PlayersOut {
		totalKills += player.Kills
	}
	return totalKills
}

func (g *GameEntity) PlayerNames() []string {
	names := make([]string, 0, len(g.PlayersOut))
	for _, player := range g.PlayersOut {
		names = append(names, player.WhoKilled)
	}
	return names
}

func (g *GameEntity) PlayerKills() map[string]int64 {
	kills := make(map[string]int64)
	for _, player := range g.PlayersOut {
		kills[player.WhoKilled] = player.Kills
	}
	return kills
}
