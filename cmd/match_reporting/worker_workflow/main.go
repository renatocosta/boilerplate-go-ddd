package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/ddd/cmd/log_handler/config"
	"github.com/ddd/internal/shared"
	"github.com/ddd/internal/shared/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	initWorkerWorkFlow(workflow.NewWorkFlow(ctx))
}

func initWorkerWorkFlow(wf shared.WorkFlowable) {
	cfg := config.GetConfig()

	c, err := client.Dial(client.Options{HostPort: cfg.Variable.TemporalHostPort})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, workflow.PlayersKilledTaskQueueName, worker.Options{
		Identity: "Match Reporting",
		// Other options if needed
	})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(wf.PlayersKilledWorkflow)
	w.RegisterActivity(wf.PlayersKilled)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
