package event_handler

import (
	"context"

	"github.com/ddd/internal/context/log_handler/infra/dispatcher"
	"github.com/ddd/pkg/building_blocks/domain"
)

// UserRegisteredHandler handles the user registered event
func SelectLogFileEventHandler2(ctx context.Context, event domain.Event, dependencies dispatcher.AdditionalDependencies) error {
	return nil
}
