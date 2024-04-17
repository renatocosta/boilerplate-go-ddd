package main

import (
	"context"
	_ "net/http/pprof"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/ddd/cmd/log_handler/config"
	"github.com/ddd/internal/context/log_handler/infra/ports/http"
	"github.com/ddd/internal/context/log_handler/infra/service"
	"github.com/ddd/internal/shared/workflow"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Start(ctx, workflow.NewWorkFlow(ctx))
	if err != nil {
		panic(err)
	}
	defer cfg.Close()

	app, _ := service.NewApplication(ctx, cfg)

	router := gin.Default()
	h := http.HttpServer{App: app}
	http.InitRoutes(&router.RouterGroup, h)
	if err := router.Run(":8888"); err != nil {
		//cancel()
	}
}
