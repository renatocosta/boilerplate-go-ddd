package query

import (
	"context"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/pkg/building_blocks/app"
)

type AvailableLogFiles struct {
}

type AvailableFilesHandler app.QueryHandler[AvailableLogFiles, *[]logfile.LogFileEntity]

type availableLogFilesHandler struct {
}

func NewAvailableLogFilesHandler() AvailableFilesHandler {
	return availableLogFilesHandler{}
}

func (h availableLogFilesHandler) Handle(ctx context.Context, query AvailableLogFiles) (*[]logfile.LogFileEntity, error) {
	return nil, nil
}
