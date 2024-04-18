package main

import (
	"context"
	"log"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/ddd/cmd/log_handler/config"
	"github.com/ddd/internal/context/log_handler/infra/ports/http"
	"github.com/ddd/internal/context/log_handler/infra/service"
	"github.com/ddd/internal/shared/workflow"
)

func main() {

	ctx := context.Background()
	cfg, err := config.Start(ctx, workflow.NewWorkFlow(ctx))
	if err != nil {
		log.Fatalln(err)
	}
	defer cfg.Close()

	app, _ := service.NewApplication(ctx, cfg)

	router := gin.Default()
	h := http.HttpServer{App: app}
	http.InitRoutes(&router.RouterGroup, h)

	if err := router.Run(":8181"); err != nil {
	}
}
