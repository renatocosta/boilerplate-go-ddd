package app

import (
	"context"

	"github.com/ddd/pkg/building_blocks/domain"
)

type EventHandlerList struct {
	EventName string
	Handlers  []EventHandleable
}

func ListenEvents(ctx context.Context, eventChan <-chan domain.Event, eventHandlers []EventHandlerList) {

	for {
		select {
		case event := <-eventChan:

			for _, handler := range eventHandlers {

				if handler.EventName == event.Type {

					for _, handlerFunc := range handler.Handlers {
						handlerFunc.Handle(ctx, event)
					}

				}
			}

		}

	}

}
