package command

import (
	"context"
	"database/sql"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/pkg/building_blocks/app"
	"github.com/ddd/pkg/building_blocks/infra/bus"
	"github.com/ddd/pkg/support"
	"github.com/google/uuid"
)

type SelectLogFileCommand struct {
	ID   uuid.UUID
	Path support.String
}

type SelectLogFileHandler app.CommandHandler[SelectLogFileCommand, []string]

type selectLogFileHandler struct {
	repo     logfile.LogFileRepository
	db       *sql.DB
	eventBus *bus.EventBus
}

func NewSelectLogFileHandler(eventBus *bus.EventBus, repo logfile.LogFileRepository, db *sql.DB) SelectLogFileHandler {
	return selectLogFileHandler{
		repo:     repo,
		db:       db,
		eventBus: eventBus,
	}
}

func (h selectLogFileHandler) Handle(ctx context.Context, cmd SelectLogFileCommand) ([]string, error) {

	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	logFileRepo, err := h.repo.ReadFrom(cmd.Path)

	if err != nil {
		return []string{}, err
	}

	h.repo.GetAll(ctx)

	logFile := logfile.NewLogFile(cmd.ID, logFileRepo.Path)

	lines, err := logFile.ExtractFrom(logFileRepo.File)

	if err != nil {
		return []string{}, err
	}

	logFile_ := logFile.Select(lines)

	for _, event := range logFile_.Events {
		h.eventBus.Publish(event)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return lines, nil
}
