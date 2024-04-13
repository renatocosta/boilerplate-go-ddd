package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ddd/internal/context/log_handler/infra/adapters"
	"github.com/ddd/internal/context/log_handler/infra/service"
	"github.com/ddd/internal/shared/workflow"
	"github.com/ddd/pkg/building_blocks/infra/bus"
	"github.com/ddd/pkg/support"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var eventBus = bus.NewEventBus()

func main() {

	var errorApp string

	ctx, cancel := context.WithCancel(context.Background())
	defer support.ShutdownApp(ctx, cancel, &errorApp)

	db, err := service.GetDb()
	defer func() {
		db.Close()
	}()

	if err != nil {
		errorApp = err.Error()
		cancel()
		return
	}

	_, workFlow, _ := service.NewApplication(ctx, eventBus, adapters.NewLogFileRepository(db), db)

	initWorkerWorkFlow(workFlow)
}

func initWorkerWorkFlow(wf workflow.WorkFlowable) {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, workflow.PlayersKilledTaskQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(wf.PlayersKilledWorkflow)
	w.RegisterActivity(wf.HumanFile)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
