package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/ddd/cmd/log_handler/config"
	httpPort "github.com/ddd/internal/context/log_handler/infra/ports/http"
	"github.com/ddd/internal/context/log_handler/infra/service"
	"github.com/ddd/internal/shared/workflow"
	"github.com/ddd/pkg/support"
)

func main() {

	ctx := context.Background()
	cfg, err := config.Start(ctx, workflow.NewWorkFlow(ctx))
	support.PanicOnError(err, "")

	defer cfg.Close()
	app, _ := service.NewApplication(ctx, cfg)

	h := httpPort.HttpServer{App: app}

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	httpPort.InitRoutes(r, h)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), loggedRouter))

}
