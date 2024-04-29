package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ddd/cmd/log_handler/config"
	httpPort "github.com/ddd/internal/context/log_handler/infra/ports/http"
	"github.com/ddd/internal/context/log_handler/infra/service"
	"github.com/ddd/internal/shared/workflow"
	"github.com/ddd/pkg/support"
	"github.com/gorilla/mux"
)

func main() {
	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	ctxWorkflow := context.Background()

	cfg, err := config.Start(ctx, workflow.NewWorkFlow(ctxWorkflow))
	support.PanicOnError(err, "")

	defer cfg.Close()
	app, _ := service.NewApplication(ctx, cfg)

	h := httpPort.HttpServer{App: app}

	r := mux.NewRouter()
	httpPort.InitRoutes(r, h)

	srv := &http.Server{
		Addr:    ":8181",
		Handler: r,
	}

	go func() {
		log.Println("Starting server...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	support.ShutdownHandler(ctx, cancel, srv)

}
