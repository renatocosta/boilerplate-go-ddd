package domain

import "github.com/google/uuid"

type AggregateRoot struct {
	ID     uuid.UUID
	Events []Event `json:"-"`
}

func (a *AggregateRoot) RecordThat(event Event) {
	a.Events = append(a.Events, event)
}
