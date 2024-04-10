package main

import (
	"context"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/ddd/pkg/building_blocks/infra/bus"
	"github.com/ddd/pkg/support"

	"github.com/ddd/internal/context/log_handler/infra/adapters"
	"github.com/ddd/internal/context/log_handler/infra/ports/http"
	"github.com/ddd/internal/context/log_handler/infra/service"
)

var eventBus = bus.NewEventBus()

func main() {

	errorApp := ""

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		support.ShutdownApp(ctx, cancel, errorApp)
	}()

	db, err := service.GetDb()
	defer func() {
		db.Close()
	}()

	if err != nil {
		errorApp = err.Error()
		cancel()
		return
	}

	app, _ := service.NewApplication(ctx, eventBus, adapters.NewLogFileRepository(db), db)

	router := gin.Default()
	h := http.HttpServer{App: app}
	http.InitRoutes(&router.RouterGroup, h)
	if err := router.Run(":8888"); err != nil {
		errorApp = err.Error()
		cancel()
	}

}
