package shared

import (
	"context"

	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"go.temporal.io/sdk/workflow"
)

type WorkFlowable interface {
	StartFrom(event events.LogFileSelected)
	PlayersKilledWorkflow(ctx workflow.Context, input command.CreateHumanLogFileCommand) (string, error)
	HumanFile(ctx context.Context, command command.CreateHumanLogFileCommand) ([][]string, error)
	PlayersKilled(ctx context.Context, data [][]string) (string, error)
	Undo(ctx context.Context, command command.CreateHumanLogFileCommand) (string, error)
}
