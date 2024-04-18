package workflow

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ddd/cmd/log_handler/config"
	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/internal/shared"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type WorkFlow struct {
	ctx context.Context
}

func NewWorkFlow(ctx context.Context) shared.WorkFlowable {
	return &WorkFlow{
		ctx: ctx,
	}
}

func (w WorkFlow) StartFrom(event events.LogFileSelected) {

	cfg := config.GetConfig()

	c, err := client.Dial(client.Options{HostPort: cfg.Variable.TemporalHostPort})

	if err != nil {
		panic("Unable to create Temporal client:")
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "log-file-" + event.ID.String(),
		TaskQueue: PlayersKilledTaskQueueName,
	}

	log.Printf("Starting to send file to human file ID: %s", event.ID.String())

	we, err := c.ExecuteWorkflow(w.ctx, options, w.PlayersKilledWorkflow, command.CreateHumanLogFileCommand{ID: event.ID, Content: event.Content})

	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	var result string

	err = we.Get(w.ctx, &result)

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
		MaximumAttempts:        2, // 0 is unlimited retries
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

		var undoOutput string
		undoErr := workflow.ExecuteActivity(ctx, w.Undo, input).Get(ctx, &undoOutput)

		if undoErr != nil {
			return "",
				fmt.Errorf("PlayersKilled: failed to find players killed %v: %v, undo: %v",
					input.Content, humanFileErr, undoErr)
		}

		return "", humanFileErr
	}

	// Find players killed
	inputPlayersKilled := humanFileOutput
	var playersKilledOutput string

	playerKilledErr := workflow.ExecuteActivity(ctx, w.PlayersKilled, inputPlayersKilled).Get(ctx, &playersKilledOutput)

	if playerKilledErr != nil {

		return "",
			fmt.Errorf("PlayersKilled: failed to find players killed %v: %v.",
				input.Content, playerKilledErr)
	}

	result := fmt.Sprintf("Workflow completed (transaction IDs: %s, %s)", humanFileOutput, playersKilledOutput)

	return result, nil

}
