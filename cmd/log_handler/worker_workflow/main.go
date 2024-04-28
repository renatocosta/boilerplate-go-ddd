package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ddd/internal/shared"
	"github.com/ddd/internal/shared/workflow"
	"github.com/ddd/pkg/support"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	initWorkerWorkFlow(workflow.NewWorkFlow(ctx))
}

func initWorkerWorkFlow(wf shared.WorkFlowable) {
	c, err := client.Dial(client.Options{HostPort: os.Getenv("TEMPORAL_HOST_PORT")})
	support.PanicOnError(err, "Unable to create Temporal client.")
	defer c.Close()

	w := worker.New(c, workflow.PlayersKilledTaskQueueName, worker.Options{
		Identity: "Log Handler",
	})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(wf.PlayersKilledWorkflow)
	w.RegisterActivity(wf.HumanFile)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	support.PanicOnError(err, "Unable to start Worker.")
}
