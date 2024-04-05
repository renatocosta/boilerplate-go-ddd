package app

import (
	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SelectLogFile      command.SelectLogFileHandler
	CreateHumanLogFile command.CreateHumanLogFileHandler
}

type Queries struct {
	LogFiles query.AvailableFilesHandler
}
