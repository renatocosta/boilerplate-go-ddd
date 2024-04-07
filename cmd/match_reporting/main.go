package match_reporting_main

import (
	"context"
	"fmt"

	"github.com/ddd/internal/context/match_reporting/infra/service"
)

func Main(ctx context.Context, rawData [][]string) {

	app := service.NewApplication(ctx)

	resultPlayersKilled := service.FindPlayersKilledCommandDispatcher(ctx, &app, rawData)

	fmt.Println(resultPlayersKilled)

}
