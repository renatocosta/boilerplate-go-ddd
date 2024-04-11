package service

import (
	"context"

	"github.com/ddd/internal/context/match_reporting/app"
	"github.com/ddd/internal/context/match_reporting/app/command"
	players_killed "github.com/ddd/internal/context/match_reporting/domain/model/player_killed"
	"github.com/ddd/internal/context/match_reporting/domain/model/player_killed/state"
	_ "github.com/go-sql-driver/mysql"
)

func NewApplication(ctx context.Context) app.Application {

	playerStates := players_killed.NewPlayer(state.NewKillPlayer(), state.NewDeathPlayer())
	playerKilled := players_killed.NewPlayerKilled(playerStates)

	return app.Application{
		Commands: app.Commands{
			FindPlayersKilled: command.NewFindPlayersKilledHandler(playerKilled),
		},
	}
}
