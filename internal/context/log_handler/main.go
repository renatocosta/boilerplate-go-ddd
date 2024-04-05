package main

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ddd/pkg/building_blocks/infra/bus"
	"github.com/ddd/pkg/integration"
	"github.com/ddd/pkg/support"

	"github.com/ddd/internal/context/log_handler/infra/adapters"
	"github.com/ddd/internal/context/log_handler/infra/service"
	serviceM "github.com/ddd/internal/context/match_reporting/infra/service"
)

var eventBus = bus.NewEventBus()

func main() {

	ctx := context.Background()

	db, err := service.GetDb()
	defer func() {
		db.Close()
	}()

	if err != nil {
		panic(err.Error())
	}
	app := service.NewApplication(ctx, eventBus, adapters.NewLogFileRepository(db), db)
	appM := serviceM.NewApplication(ctx)

	pathFile := support.GetFilePath("/../../tmp/qgames.log")
	resultLogFile := service.SelectLogFileCommandDispatcher(ctx, &app, pathFile)

	resultHumanLogFile := service.CreateHumanLogFileCommandDispatcher(ctx, &app, resultLogFile)

	rawData := integration.PreSendCommand(resultHumanLogFile)
	resultPlayersKilled := serviceM.FindPlayersKilledCommandDispatcher(ctx, &appM, rawData)

	fmt.Println(resultPlayersKilled)

}
