package workflow

import (
	"context"
	"log"

	"github.com/ddd/internal/context/log_handler/app/command"
	commandM "github.com/ddd/internal/context/match_reporting/app/command"
	"github.com/ddd/pkg/integration"
)

func (w WorkFlow) HumanFile(ctx context.Context, command command.CreateHumanLogFileCommand) ([][]string, error) {
	log.Printf("Sending file to human with %d - ID \n",
		command.ID,
	)

	resultHumanLogFile, err := w.appL.Commands.CreateHumanLogFile.Handle(ctx, command)

	result := integration.PreSendCommand(resultHumanLogFile)

	return result, err
}

func (w WorkFlow) PlayersKilled(ctx context.Context, data [][]string) (string, error) {
	log.Print("Finding players killed")

	findPlayersKilledCommand := commandM.FindPlayersKilledCommand{Data: data}

	resultPlayersKilled, err := w.appM.Commands.FindPlayersKilled.Handle(ctx, findPlayersKilledCommand)

	return resultPlayersKilled, err
}
