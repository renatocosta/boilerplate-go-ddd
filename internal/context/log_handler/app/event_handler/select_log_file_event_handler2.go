package event_handler

import (
	"context"

	"github.com/ddd/pkg/building_blocks/domain"
	"github.com/ddd/pkg/building_blocks/infra/bus"
)

// UserRegisteredHandler handles the user registered event
func SelectLogFileEventHandler2(ctx context.Context, event domain.Event, dependencies bus.AdditionalDependencies) error {
	return nil
}
