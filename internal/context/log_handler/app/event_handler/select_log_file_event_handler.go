package event_handler

import (
	"context"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/pkg/building_blocks/app"
	"github.com/ddd/pkg/building_blocks/domain"
)

type SelectLogFileEventHandler struct {
	LogFileRepo logfile.LogFileRepository
}

func NewSelectLogFileEventHandler(logFileRepo logfile.LogFileRepository) app.EventHandleable {
	return &SelectLogFileEventHandler{LogFileRepo: logFileRepo}
}

func (s SelectLogFileEventHandler) Handle(ctx context.Context, event domain.Event) error {
	return s.LogFileRepo.Add(
		logfile.NewLogFile(event.Data.(events.LogFileSelected).ID, event.Data.(events.LogFileSelected).Path),
		ctx,
	)
}
