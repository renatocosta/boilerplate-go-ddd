package workflow

import (
	"context"
	"log"
	"time"

	"github.com/ddd/internal/context/log_handler/app"
	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	appM "github.com/ddd/internal/context/match_reporting/app"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type WorkFlowable interface {
	StartFrom(event events.LogFileSelected)
	PlayersKilledWorkflow(ctx workflow.Context, input command.CreateHumanLogFileCommand) (string, error)
	HumanFile(ctx context.Context, command command.CreateHumanLogFileCommand) ([][]string, error)
	PlayersKilled(ctx context.Context, data [][]string) (string, error)
}

type WorkFlow struct {
	appL app.Application
	appM appM.Application
}

func NewWorkFlowFromLogHandler(app app.Application) WorkFlowable {
	return WorkFlow{
		appL: app,
	}
}

func NewWorkFlowFromMatchReporting(app appM.Application) WorkFlowable {
	return WorkFlow{
		appM: app,
	}
}

func (w WorkFlow) StartFrom(event events.LogFileSelected) {

	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "log-file-" + event.ID.String(),
		TaskQueue: PlayersKilledTaskQueueName,
	}

	log.Printf("Starting send file to human file ID: %s", event.ID.String())

	we, err := c.ExecuteWorkflow(context.Background(), options, w.PlayersKilledWorkflow, command.CreateHumanLogFileCommand{ID: event.ID, Content: event.Content})

	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	var result string

	err = we.Get(context.Background(), &result)

	if err != nil {
		log.Fatalln("Unable to get Workflow result:", err)
	}

	log.Println(result)

}

func (w WorkFlow) PlayersKilledWorkflow(ctx workflow.Context, input command.CreateHumanLogFileCommand) (string, error) {

	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        500, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	// Create a human file
	var humanFileOutput [][]string

	humanFileErr := workflow.ExecuteActivity(ctx, w.HumanFile, input).Get(ctx, &humanFileOutput)

	if humanFileErr != nil {
		return "", humanFileErr
	}

	// Find players killed
	inputPlayersKilled := humanFileOutput
	var playersKilledOutput string

	playerKilledErr := workflow.ExecuteActivity(ctx, w.PlayersKilled, inputPlayersKilled).Get(ctx, &playersKilledOutput)

	if playerKilledErr != nil {
		return "", playerKilledErr
	}

	return "", nil

}
