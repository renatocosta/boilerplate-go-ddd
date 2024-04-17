package testsuite

import (
	"context"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ddd/cmd/log_handler/config"
	"github.com/ddd/internal/context/log_handler/app/command"
	eventsH "github.com/ddd/internal/context/log_handler/domain/model/human_logfile/events"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/internal/context/log_handler/infra/service"
	commandM "github.com/ddd/internal/context/match_reporting/app/command"
	"github.com/ddd/internal/shared/workflow"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"

	_ "github.com/go-sql-driver/mysql"

	serviceM "github.com/ddd/internal/context/match_reporting/infra/service"
	"github.com/ddd/pkg/building_blocks/infra/bus"
	"github.com/ddd/pkg/integration"
	"github.com/ddd/pkg/support"
)

var ag = &bus.AggregateRootTestCase{}
var eventBus = bus.NewEventBus()

func TestShouldBeAbleToRunEndToEndCommandsSuccessfully(t *testing.T) {

	ag.T = t

	ag.
		Given(runEndToEndCommands).
		When(eventBus.RaisedEvents()).
		Then(events.LogFileSelectedEvent, eventsH.HumanLogFileCreatedEvent).
		Assert()
}

func runEndToEndCommands() {

	pathFile := support.GetFilePath("internal/context/log_handler/infra/storage/qgames.log")

	controll := gomock.NewController(ag.T)
	defer controll.Finish()

	var repo = service.NewMockLogFileRepository(controll)

	file, _ := os.Open(pathFile)
	logFileEntity := &logfile.LogFileEntity{
		Path: support.NewString(pathFile),
		File: file,
	}

	repo.EXPECT().ReadFrom(gomock.Any()).Return(logFileEntity, nil).Times(1)
	repo.EXPECT().Add(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		ag.T.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectCommit()

	workFlow := workflow.NewMockWorkFlowable(controll)
	workFlow.EXPECT().StartFrom(gomock.Any()).Times(1)

	cfg := &config.Config{
		Database: db,
		EventBus: eventBus,
		WorkFlow: workFlow,
		Repo:     repo}

	app, _ := service.NewApplication(ctx, cfg)

	selectLogFileCommand := command.SelectLogFileCommand{ID: uuid.New(), Path: support.NewString(pathFile)}
	resultLogFile, err := app.Commands.SelectLogFile.Handle(ctx, selectLogFileCommand)

	if err != nil {
		ag.T.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	appM := serviceM.NewApplication(ctx)

	createHumanLogFileCommand := command.CreateHumanLogFileCommand{ID: uuid.New(), Content: resultLogFile}
	resultHumanLogFile, err := app.Commands.CreateHumanLogFile.Handle(ctx, createHumanLogFileCommand)

	rawData := integration.PreSendCommand(resultHumanLogFile)

	findPlayersKilledCommand := commandM.FindPlayersKilledCommand{Data: rawData}
	_, err = appM.Commands.FindPlayersKilled.Handle(ctx, findPlayersKilledCommand)

	if err != nil {
		ag.T.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

}
