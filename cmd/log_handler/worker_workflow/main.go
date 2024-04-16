package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ddd/internal/shared"
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

	initWorkerWorkFlow(workflow.NewWorkFlow(ctx))
}

func initWorkerWorkFlow(wf shared.WorkFlowable) {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, workflow.PlayersKilledTaskQueueName, worker.Options{
		Identity: "Log Handler",
	})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(wf.PlayersKilledWorkflow)
	w.RegisterActivity(wf.HumanFile)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
