package event_handler

import (
	"context"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/internal/shared"
	"github.com/ddd/pkg/building_blocks/app"
	"github.com/ddd/pkg/building_blocks/domain"
)

type SelectLogFileEventIntegrationHandler struct {
	WorkFlow shared.WorkFlowable
}

func NewSelectLogFileEventIntegrationHandler(workFlow shared.WorkFlowable) app.EventHandleable {
	return &SelectLogFileEventIntegrationHandler{WorkFlow: workFlow}
}

func (d SelectLogFileEventIntegrationHandler) Handle(ctx context.Context, event domain.Event) error {

	d.WorkFlow.StartFrom(event.Data.(events.LogFileSelected))

	return nil
}
