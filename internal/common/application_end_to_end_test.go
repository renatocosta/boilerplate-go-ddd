package common

import (
	"context"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	eventsH "github.com/ddd/internal/context/log_handler/domain/model/human_logfile/events"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/internal/context/log_handler/infra/service"
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
	repo.EXPECT().GetAll(gomock.Any()).Return(nil, nil).Times(1)

	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		ag.T.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectCommit()

	app := service.NewApplication(ctx, eventBus, repo, db)
	appM := serviceM.NewApplication(ctx)

	resultLogFile := service.SelectLogFileCommandDispatcher(ctx, &app, support.NewString(pathFile))
	resultHumanLogFile := service.CreateHumanLogFileCommandDispatcher(ctx, &app, resultLogFile)
	rawData := integration.PreSendCommand(resultHumanLogFile)
	serviceM.FindPlayersKilledCommandDispatcher(ctx, &appM, rawData)
}
