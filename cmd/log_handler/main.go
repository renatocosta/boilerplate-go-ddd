package main

import (
	"context"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/ddd/internal/context/log_handler/infra/adapters"
	"github.com/ddd/internal/context/log_handler/infra/ports/http"
	"github.com/ddd/internal/context/log_handler/infra/service"
	"github.com/ddd/internal/shared/workflow"
	"github.com/ddd/pkg/support"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	var errorApp string

	ctx, cancel := context.WithCancel(context.Background())
	defer support.ShutdownApp(ctx, cancel, &errorApp)

	conf, err := service.NewConfig(ctx)
	defer conf.GetDB().Close()

	if err != nil {
		errorApp = err.Error()
		cancel()
		return
	}

	app, _ := service.NewApplication(ctx, conf.GetEventBus(), adapters.NewLogFileRepository(conf.GetDB()), conf.GetDB(), workflow.NewWorkFlow(ctx))

	router := gin.Default()
	h := http.HttpServer{App: app}
	http.InitRoutes(&router.RouterGroup, h)
	if err := router.Run(":8888"); err != nil {
		errorApp = err.Error()
		cancel()
	}
}
