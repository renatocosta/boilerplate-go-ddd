package service

import (
	"context"

	"github.com/ddd/cmd/log_handler/config"
	"github.com/ddd/internal/context/log_handler/app"

	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/app/event_handler"
	"github.com/ddd/internal/context/log_handler/app/query"
	eventsH "github.com/ddd/internal/context/log_handler/domain/model/human_logfile/events"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	appB "github.com/ddd/pkg/building_blocks/app"
	"github.com/ddd/pkg/building_blocks/domain"
)

func NewApplication(ctx context.Context, cfg *config.Config) (app.Application, func()) {

	eventChan := make(chan domain.Event, 1)

	cfg.EventBus.Subscribe(events.LogFileSelectedEvent, eventChan)

	cfg.EventBus.Subscribe(eventsH.HumanLogFileCreatedEvent, eventChan)

	eventHandlers := []appB.EventHandlerList{
		appB.EventHandlerList{
			EventName: events.LogFileSelectedEvent,
			Handlers: []appB.EventHandleable{
				event_handler.NewSelectLogFileEventHandler(cfg.Repo),
				event_handler.NewSelectLogFileEventIntegrationHandler(cfg.WorkFlow),
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
				SelectLogFile:      command.NewSelectLogFileHandler(cfg.EventBus, cfg.Repo, cfg.Database),
				CreateHumanLogFile: command.NewCreateHumanLogFileHandler(cfg.EventBus),
			},
			Queries: app.Queries{
				LogFiles: query.NewAvailableLogFilesHandler(cfg.Repo),
			},
		}, func() {

			//_ = closeDbConnection()
			//_ = closeGrpcConnection()
		}
}
