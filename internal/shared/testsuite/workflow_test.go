package testsuite

import (
	"context"
	"errors"
	"testing"

	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/shared/workflow"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_SuccessfulFindPalyersKilledWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	ctx, _ := context.WithCancel(context.Background())
	resultLogFile := []string{
		"Kill: 4 3 7: Zeh killed Isgalamido by MOD_ROCKET_SPLASH",
		"Kill: 2 5 7: Dono da Bola killed Assasinu Credi by MOD_ROCKET_SPLASH",
		"Item: 2 weapon_rocketlauncher"}

	input := command.CreateHumanLogFileCommand{ID: uuid.New(), Content: resultLogFile}
	workflow := workflow.NewWorkFlow(ctx)

	// Mock activity implementation
	env.OnActivity(workflow.HumanFile, mock.Anything, input).Return(nil, nil)
	env.OnActivity(workflow.PlayersKilled, mock.Anything, input).Return("", nil)

	env.ExecuteWorkflow(workflow.PlayersKilledWorkflow, input)
	require.True(t, env.IsWorkflowCompleted())
	//require.NoError(t, env.GetWorkflowError())
}

func Test_HumanFileFailedWorkflow(t *testing.T) {

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	ctx, _ := context.WithCancel(context.Background())
	resultLogFile := []string{
		"Kill: 4 3 7: Zeh killed Isgalamido by MOD_ROCKET_SPLASH",
		"Kill: 2 5 7: Dono da Bola killed Assasinu Credi by MOD_ROCKET_SPLASH",
		"Item: 2 weapon_rocketlauncher"}

	input := command.CreateHumanLogFileCommand{ID: uuid.New(), Content: resultLogFile}
	workflow := workflow.NewWorkFlow(ctx)

	// Mock activity implementation
	env.OnActivity(workflow.HumanFile, mock.Anything, input).Return(nil, nil)
	env.OnActivity(workflow.PlayersKilled, mock.Anything, input).Return("", nil)
	env.OnActivity(workflow.Undo, mock.Anything, input).Return("", errors.New("unable to create human log file"))

	env.ExecuteWorkflow(workflow.PlayersKilledWorkflow, input)
	require.True(t, env.IsWorkflowCompleted())
	require.Error(t, env.GetWorkflowError())
}
