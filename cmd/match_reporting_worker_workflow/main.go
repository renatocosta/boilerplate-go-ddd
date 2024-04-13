package main

import (
	"context"
	"log"

	"github.com/ddd/internal/context/match_reporting/infra/service"
	"github.com/ddd/internal/shared/workflow"
	"github.com/ddd/pkg/support"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	var errorApp string

	ctx, cancel := context.WithCancel(context.Background())
	defer support.ShutdownApp(ctx, cancel, &errorApp)

	_, workFlow := service.NewApplication(ctx)

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
	w.RegisterActivity(wf.PlayersKilled)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
