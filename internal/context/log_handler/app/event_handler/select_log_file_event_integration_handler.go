package event_handler

import (
	"context"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/internal/context/log_handler/infra/event_handler"
	"github.com/ddd/pkg/building_blocks/domain"
)

func SelectLogFileEventHandlerIntegration(ctx context.Context, event domain.Event, dependencies event_handler.AdditionalDependencies) error {

	dependencies.WorkFlow.StartFrom(event.Data.(events.LogFileSelected))

	return nil
}
