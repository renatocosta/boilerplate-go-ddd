package service

import (
	"context"
	"database/sql"

	"github.com/ddd/internal/context/log_handler/app"
	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/app/event_handler"
	"github.com/ddd/internal/context/log_handler/app/query"
	"github.com/ddd/internal/context/log_handler/domain/model/human_logfile"
	eventsH "github.com/ddd/internal/context/log_handler/domain/model/human_logfile/events"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/pkg/building_blocks/domain"
	"github.com/ddd/pkg/building_blocks/infra/bus"
	"github.com/google/uuid"
)

func NewApplication(ctx context.Context, eventBus *bus.EventBus, logFileRepo logfile.LogFileRepository, db *sql.DB) app.Application {

	eventChan := make(chan domain.Event)

	subscriberLogFileSelected := events.LogFileSelectedEvent
	eventBus.Subscribe(subscriberLogFileSelected, eventChan)

	subscriberHumanLogFileCreated := eventsH.HumanLogFileCreatedEvent
	eventBus.Subscribe(subscriberHumanLogFileCreated, eventChan)

	eventHandlers := []bus.EventHandlerList{
		bus.EventHandlerList{
			EventName: subscriberLogFileSelected,
			Handlers: []bus.EventHandlerFunc{
				event_handler.SelectLogFileEventHandler,
				event_handler.SelectLogFileEventHandler2,
			},
		},
		bus.EventHandlerList{
			EventName: subscriberHumanLogFileCreated,
			Handlers: []bus.EventHandlerFunc{
				event_handler.CreateHumanFileEventHandler,
			},
		},
	}

	additionalDependencies := bus.NewAdditionalDependencies(logFileRepo)

	go bus.HandleEvent(ctx, eventChan, eventHandlers, additionalDependencies)

	return app.Application{
		Commands: app.Commands{
			SelectLogFile:      command.NewSelectLogFileHandler(eventBus, logFileRepo, db),
			CreateHumanLogFile: command.NewCreateHumanLogFileHandler(eventBus),
		},
		Queries: app.Queries{
			LogFiles: query.NewAvailableLogFilesHandler(),
		},
	}
}

func SelectLogFileCommandDispatcher(ctx context.Context, app *app.Application, pathFile string) []string {
	selectLogFileCommand := command.SelectLogFileCommand{ID: uuid.New(), Path: pathFile}
	resultLogFile, err := app.Commands.SelectLogFile.Handle(ctx, selectLogFileCommand)
	if err != nil {
		panic(err.Error())
	}
	return resultLogFile
}

func CreateHumanLogFileCommandDispatcher(ctx context.Context, app *app.Application, resultLogFile []string) []human_logfile.HumanLogFileRowable {
	createHumanLogFileCommand := command.CreateHumanLogFileCommand{ID: uuid.New(), Content: resultLogFile}
	resultHumanLogFile, err := app.Commands.CreateHumanLogFile.Handle(ctx, createHumanLogFileCommand)
	if err != nil {
		panic(err.Error())
	}
	return resultHumanLogFile
}
