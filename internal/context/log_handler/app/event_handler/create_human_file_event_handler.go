package event_handler

import (
	"context"

	"github.com/ddd/internal/context/log_handler/infra/event_handler"
	"github.com/ddd/pkg/building_blocks/domain"
)

// UserRegisteredHandler handles the user registered event
func CreateHumanFileEventHandler(ctx context.Context, event domain.Event, dependencies event_handler.AdditionalDependencies) error {
	return nil
}
