package service

import (
	"context"
	"database/sql"

	"github.com/ddd/internal/context/log_handler/app"
	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/app/event_handler"
	"github.com/ddd/internal/context/log_handler/app/query"
	eventsH "github.com/ddd/internal/context/log_handler/domain/model/human_logfile/events"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	eventHandlerInfra "github.com/ddd/internal/context/log_handler/infra/event_handler"
	"github.com/ddd/pkg/building_blocks/domain"
	"github.com/ddd/pkg/building_blocks/infra/bus"
)

func NewApplication(ctx context.Context, eventBus *bus.EventBus, logFileRepo logfile.LogFileRepository, db *sql.DB) (app.Application, func()) {

	eventChan := make(chan domain.Event)

	subscriberLogFileSelected := events.LogFileSelectedEvent
	eventBus.Subscribe(subscriberLogFileSelected, eventChan)

	subscriberHumanLogFileCreated := eventsH.HumanLogFileCreatedEvent
	eventBus.Subscribe(subscriberHumanLogFileCreated, eventChan)

	eventHandlers := []eventHandlerInfra.EventHandlerList{
		eventHandlerInfra.EventHandlerList{
			EventName: subscriberLogFileSelected,
			Handlers: []eventHandlerInfra.EventHandlerFunc{
				event_handler.SelectLogFileEventHandler,
				event_handler.SelectLogFileEventHandler2,
			},
		},
		eventHandlerInfra.EventHandlerList{
			EventName: subscriberHumanLogFileCreated,
			Handlers: []eventHandlerInfra.EventHandlerFunc{
				event_handler.CreateHumanFileEventHandler,
			},
		},
	}

	additionalDependencies := eventHandlerInfra.NewAdditionalDependencies(logFileRepo)

	go eventHandlerInfra.HandleEvent(ctx, eventChan, eventHandlers, additionalDependencies)

	return app.Application{
			Commands: app.Commands{
				SelectLogFile:      command.NewSelectLogFileHandler(eventBus, logFileRepo, db),
				CreateHumanLogFile: command.NewCreateHumanLogFileHandler(eventBus),
			},
			Queries: app.Queries{
				LogFiles: query.NewAvailableLogFilesHandler(logFileRepo),
			},
		},
		func() {
			//_ = closeDbConnection()
			//_ = closeGrpcConnection()
		}
}
