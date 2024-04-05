package bus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AggregateRootTestCase struct {
	T              *testing.T
	eventBus       *EventBus
	raisedEvents   []string
	expectedEvents []string
}

func (ag *AggregateRootTestCase) Given(command func()) *AggregateRootTestCase {
	command()
	return ag
}

func (ag *AggregateRootTestCase) When(events []string) *AggregateRootTestCase {
	for _, eventType := range events {
		ag.raisedEvents = append(ag.raisedEvents, eventType)
	}
	return ag
}

func (ag *AggregateRootTestCase) Then(events ...string) *AggregateRootTestCase {
	ag.expectedEvents = events
	return ag
}

func (ag *AggregateRootTestCase) Assert() {
	assert.EqualValues(ag.T, ag.expectedEvents, ag.raisedEvents, "The expected events are not equal.")
}
