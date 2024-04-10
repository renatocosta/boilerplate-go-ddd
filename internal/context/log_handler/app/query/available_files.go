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
	repo logfile.LogFileRepository
}

func NewAvailableLogFilesHandler(repo logfile.LogFileRepository) AvailableFilesHandler {
	return availableLogFilesHandler{repo: repo}
}

func (h availableLogFilesHandler) Handle(ctx context.Context, query AvailableLogFiles) (*[]logfile.LogFileEntity, error) {
	entities, err := h.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return entities, nil
}
