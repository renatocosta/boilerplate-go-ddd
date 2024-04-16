package app

import (
	"context"

	"github.com/ddd/pkg/building_blocks/domain"
)

type EventHandleable interface {
	Handle(ctx context.Context, event domain.Event) error
}
