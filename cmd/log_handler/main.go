package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ddd/pkg/building_blocks/infra/bus"
	"github.com/ddd/pkg/integration"
	"github.com/ddd/pkg/support"

	"github.com/ddd/internal/context/log_handler/infra/adapters"
	"github.com/ddd/internal/context/log_handler/infra/service"
)

var eventBus = bus.NewEventBus()

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	ctx := context.Background()

	db, err := service.GetDb()
	defer func() {
		db.Close()
	}()

	if err != nil {
		panic(err.Error())
	}
	app := service.NewApplication(ctx, eventBus, adapters.NewLogFileRepository(db), db)

	pathFile := support.GetFilePath("internal/context/log_handler/infra/storage/qgames.log")
	resultLogFile := service.SelectLogFileCommandDispatcher(ctx, &app, support.NewString(pathFile))

	resultHumanLogFile := service.CreateHumanLogFileCommandDispatcher(ctx, &app, resultLogFile)

	rawData := integration.PreSendCommand(resultHumanLogFile)

	integration.Dispatch(ctx, rawData)
}
