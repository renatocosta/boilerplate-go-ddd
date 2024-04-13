package service

import (
	"context"

	"github.com/ddd/internal/context/match_reporting/app"
	"github.com/ddd/internal/context/match_reporting/app/command"
	players_killed "github.com/ddd/internal/context/match_reporting/domain/model/player_killed"
	"github.com/ddd/internal/context/match_reporting/domain/model/player_killed/state"
	"github.com/ddd/internal/shared/workflow"
	_ "github.com/go-sql-driver/mysql"
)

func NewApplication(ctx context.Context) (app.Application, workflow.WorkFlowable) {

	playerStates := players_killed.NewPlayer(state.NewKillPlayer(), state.NewDeathPlayer())
	playerKilled := players_killed.NewPlayerKilled(playerStates)

	app := app.Application{
		Commands: app.Commands{
			FindPlayersKilled: command.NewFindPlayersKilledHandler(playerKilled),
		},
	}

	workflow := workflow.NewWorkFlowFromMatchReporting(app)

	return app, workflow
}
