package event_handler

import (
	"context"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/pkg/building_blocks/domain"
)

type AdditionalDependencies struct {
	LogFileRepo logfile.LogFileRepository
}

func NewAdditionalDependencies(logFileRepo logfile.LogFileRepository) AdditionalDependencies {
	return AdditionalDependencies{LogFileRepo: logFileRepo}
}

type EventHandlerList struct {
	EventName string
	Handlers  []EventHandlerFunc
}

type EventHandlerFunc func(context.Context, domain.Event, AdditionalDependencies) error

func HandleEvent(ctx context.Context, eventChan <-chan domain.Event, eventHandlers []EventHandlerList, additionalDependencies AdditionalDependencies) {

	for {
		select {
		case event := <-eventChan:
			for _, handler := range eventHandlers {
				if handler.EventName == event.Type {

					for _, handlerFunc := range handler.Handlers {
						handlerFunc(ctx, event, additionalDependencies)
					}

				}
			}

		}

	}

}
