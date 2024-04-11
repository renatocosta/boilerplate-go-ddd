package match_reporting_main

import (
	"context"
	"fmt"

	"github.com/ddd/internal/context/match_reporting/app/command"
	"github.com/ddd/internal/context/match_reporting/infra/service"
)

func Main(ctx context.Context, rawData [][]string) {

	app := service.NewApplication(ctx)

	findPlayersKilledCommand := command.FindPlayersKilledCommand{Data: rawData}

	resultPlayersKilled, _ := app.Commands.FindPlayersKilled.Handle(ctx, findPlayersKilledCommand)

	fmt.Println(resultPlayersKilled)

}
