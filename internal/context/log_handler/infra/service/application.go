package service

import (
	"context"
	"database/sql"

	"github.com/ddd/internal/context/log_handler/app"
	"github.com/ddd/internal/shared"

	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/app/event_handler"
	"github.com/ddd/internal/context/log_handler/app/query"
	eventsH "github.com/ddd/internal/context/log_handler/domain/model/human_logfile/events"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	appB "github.com/ddd/pkg/building_blocks/app"
	"github.com/ddd/pkg/building_blocks/domain"
	"github.com/ddd/pkg/building_blocks/infra/bus"
)

func NewApplication(ctx context.Context,
	eventBus *bus.EventBus,
	logFileRepo logfile.LogFileRepository,
	db *sql.DB,
	workFlow shared.WorkFlowable,
) (app.Application, func()) {

	eventChan := make(chan domain.Event, 1)

	eventBus.Subscribe(events.LogFileSelectedEvent, eventChan)

	eventBus.Subscribe(eventsH.HumanLogFileCreatedEvent, eventChan)

	eventHandlers := []appB.EventHandlerList{
		appB.EventHandlerList{
			EventName: events.LogFileSelectedEvent,
			Handlers: []appB.EventHandleable{
				event_handler.NewSelectLogFileEventHandler(logFileRepo),
				event_handler.NewSelectLogFileEventIntegrationHandler(workFlow),
			},
		},

		appB.EventHandlerList{
			EventName: eventsH.HumanLogFileCreatedEvent,
			Handlers: []appB.EventHandleable{
				event_handler.NewCreateHumanLogFileEventHandler(),
			},
		},
	}

	go appB.ListenEvents(ctx, eventChan, eventHandlers)

	return app.Application{
			Commands: app.Commands{
				SelectLogFile:      command.NewSelectLogFileHandler(eventBus, logFileRepo, db),
				CreateHumanLogFile: command.NewCreateHumanLogFileHandler(eventBus),
			},
			Queries: app.Queries{
				LogFiles: query.NewAvailableLogFilesHandler(logFileRepo),
			},
		}, func() {

			//_ = closeDbConnection()
			//_ = closeGrpcConnection()
		}
}
