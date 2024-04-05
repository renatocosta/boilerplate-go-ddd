package bus

import (
	"github.com/ddd/pkg/building_blocks/domain"
)

// EventBus represents the event bus that handles event subscription and dispatching
type EventBus struct {
	Subscribers  map[string][]chan<- domain.Event
	eventsRaised []string
}

// NewEventBus creates a new instance of the event bus
func NewEventBus() *EventBus {
	return &EventBus{
		Subscribers: make(map[string][]chan<- domain.Event),
	}
}

// Subscribe adds a new subscriber for a given event type
func (eb *EventBus) Subscribe(eventType string, subscriber chan<- domain.Event) {
	eb.Subscribers[eventType] = append(eb.Subscribers[eventType], subscriber)

}

func (eb *EventBus) Publish(event domain.Event) {
	subscribers := eb.Subscribers[event.Type]

	for _, subscriber := range subscribers {
		subscriber <- event
		eb.eventsRaised = append(eb.eventsRaised, event.Type)
	}

}

func (eb *EventBus) RaisedEvents() []string {
	return eb.eventsRaised
}
