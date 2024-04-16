package event_handler

import (
	"context"

	"github.com/ddd/pkg/building_blocks/app"
	"github.com/ddd/pkg/building_blocks/domain"
)

type CreateHumanLogFileEventHandler struct {
}

func NewCreateHumanLogFileEventHandler() app.EventHandleable {
	return &CreateHumanLogFileEventHandler{}
}

func (c CreateHumanLogFileEventHandler) Handle(ctx context.Context, event domain.Event) error {
	return nil
}
