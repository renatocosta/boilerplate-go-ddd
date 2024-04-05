package app

import "github.com/ddd/internal/context/match_reporting/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	FindPlayersKilled command.FindPlayersKilledHandler
}
