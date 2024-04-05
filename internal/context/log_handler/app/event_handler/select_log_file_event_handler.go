package event_handler

import (
	"context"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/pkg/building_blocks/domain"
	"github.com/ddd/pkg/building_blocks/infra/bus"
)

// UserRegisteredHandler handles the user registered event
func SelectLogFileEventHandler(ctx context.Context, event domain.Event, dependencies bus.AdditionalDependencies) error {
	return dependencies.LogFileRepo.Add(
		logfile.NewLogFile(event.Data.(events.LogFileSelected).ID, event.Data.(events.LogFileSelected).Path),
		ctx,
	)
}
