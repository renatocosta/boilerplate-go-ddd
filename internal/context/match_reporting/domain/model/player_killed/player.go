package players_killed

import (
	"github.com/ddd/internal/context/match_reporting/domain/model/player_killed/state"
)

type Player struct {
	killedPlayer state.PlayerState
	deadPlayer   state.PlayerState
	players      Players
}

func NewPlayer(killedPlayer, deadPlayer state.PlayerState) Playable {
	return &Player{killedPlayer: killedPlayer, deadPlayer: deadPlayer, players: Players{}}
}

func (p *Player) KillUp(match Matchable) {
	playerInfo, ok := p.players[match.GetPlayerWhoKilled()]

	if !ok {
		p.players[match.GetPlayerWhoKilled()] = PlayerInfo{
			WhoKilled: match.GetPlayerWhoKilled(),
			WhoDied:   match.GetPlayerWhoDied(),
		}
		return
	}
	p.killedPlayer.ComputeKills(playerInfo.Kills)
	p.players[match.GetPlayerWhoKilled()] = PlayerInfo{
		Kills:     p.killedPlayer.GetKills(),
		WhoKilled: match.GetPlayerWhoKilled(),
		WhoDied:   match.GetPlayerWhoDied(),
	}
}

func (p *Player) KillDown(match Matchable) {
	playerInfo, ok := p.players[match.GetPlayerWhoDied()]

	if !ok {
		p.players[match.GetPlayerWhoKilled()] = PlayerInfo{
			WhoKilled: match.GetPlayerWhoKilled(),
			WhoDied:   match.GetPlayerWhoDied(),
		}
	}

	p.deadPlayer.ComputeKills(playerInfo.Kills)
	p.players[match.GetPlayerWhoDied()] = PlayerInfo{
		Kills:     p.deadPlayer.GetKills(),
		WhoKilled: match.GetPlayerWhoKilled(),
		WhoDied:   match.GetPlayerWhoDied(),
	}
}

func (p *Player) GetPlayers() Players {
	return p.players
}
