package workflow

import (
	"context"
	"log"

	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/infra/adapters"
	"github.com/ddd/internal/context/log_handler/infra/service"
	commandM "github.com/ddd/internal/context/match_reporting/app/command"
	serviceM "github.com/ddd/internal/context/match_reporting/infra/service"
	"github.com/ddd/pkg/integration"
)

func (w WorkFlow) HumanFile(ctx context.Context, command command.CreateHumanLogFileCommand) ([][]string, error) {
	log.Printf("Sending file to human with %d - ID \n",
		command.ID,
	)

	conf, _ := service.NewConfig(ctx)

	app, _ := service.NewApplication(ctx, conf.GetEventBus(), adapters.NewLogFileRepository(conf.GetDB()), conf.GetDB(), conf.GetWorkFlow())

	resultHumanLogFile, err := app.Commands.CreateHumanLogFile.Handle(ctx, command)

	result := integration.PreSendCommand(resultHumanLogFile)

	return result, err
}

func (w WorkFlow) PlayersKilled(ctx context.Context, data [][]string) (string, error) {
	log.Print("Finding players killed")

	findPlayersKilledCommand := commandM.FindPlayersKilledCommand{Data: data}
	appM := serviceM.NewApplication(ctx)

	resultPlayersKilled, err := appM.Commands.FindPlayersKilled.Handle(ctx, findPlayersKilledCommand)

	return resultPlayersKilled, err
}

func (w WorkFlow) Undo(ctx context.Context, command command.CreateHumanLogFileCommand) (string, error) {
	log.Print("Undoing players killed")
	return "", nil
}
