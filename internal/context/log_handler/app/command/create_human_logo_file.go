package command

import (
	"context"

	"github.com/ddd/internal/context/log_handler/domain/model/human_logfile"
	"github.com/ddd/internal/context/log_handler/domain/services"
	"github.com/ddd/pkg/building_blocks/app"
	"github.com/ddd/pkg/building_blocks/infra/bus"
	"github.com/google/uuid"
)

type CreateHumanLogFileCommand struct {
	ID      uuid.UUID
	Content []string
}

type CreateHumanLogFileHandler app.CommandHandler[CreateHumanLogFileCommand, []human_logfile.HumanLogFileRowable]

type createHumanLogFileHandler struct {
	eventBus *bus.EventBus
}

func NewCreateHumanLogFileHandler(eventBus *bus.EventBus) CreateHumanLogFileHandler {
	return createHumanLogFileHandler{eventBus: eventBus}
}

func (h createHumanLogFileHandler) Handle(ctx context.Context, cmd CreateHumanLogFileCommand) ([]human_logfile.HumanLogFileRowable, error) {

	humanLogFile := human_logfile.NewHumanLogFile(cmd.ID)

	for _, row := range cmd.Content {
		rowMapper := services.NewHumanRowMapper()
		rowMap := rowMapper.Map(row)

		if len(rowMap) > 0 {
			humanLogFile.AddRow(
				human_logfile.NewHumanLogFileRow(
					rowMap["who_killed"],
					rowMap["who_died"],
					rowMap["means_of_death"],
				),
			)
		}
	}

	humanlogFile := humanLogFile.Create()

	for _, event := range humanlogFile.Events {
		h.eventBus.Publish(event)
	}

	return humanLogFile.GetRows(), nil
}
