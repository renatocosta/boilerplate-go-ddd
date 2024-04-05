package human_logfile

import (
	"time"

	"github.com/ddd/internal/context/log_handler/domain/model/human_logfile/events"
	"github.com/ddd/pkg/building_blocks/domain"
	"github.com/google/uuid"
)

type HumanLogFile interface {
	AddRow(row HumanLogFileRowable)
	GetTotalKills() int
	GetRows() []HumanLogFileRowable
	Create() *HumanLogFileEntity
}

func NewHumanLogFile(id uuid.UUID) HumanLogFile {
	return &HumanLogFileEntity{
		AggregateRoot: domain.AggregateRoot{ID: id},
	}
}

type HumanLogFileEntity struct {
	domain.AggregateRoot
	Rows []HumanLogFileRowable
}

func (h *HumanLogFileEntity) AddRow(row HumanLogFileRowable) {

	row.Validation()
	h.Rows = append(h.Rows, HumanLogFileRow{
		playerWhoKilled: row.GetPlayerWhoKilled(),
		playerWhoDied:   row.GetPlayerWhoDied(),
		meanOfDeath:     row.GetMeanOfDeath(),
	})
}

func (h *HumanLogFileEntity) GetTotalKills() int {
	return len(h.Rows)
}

func (h *HumanLogFileEntity) GetRows() []HumanLogFileRowable {
	return h.Rows
}

func (h *HumanLogFileEntity) Create() *HumanLogFileEntity {

	//Raise event
	event := domain.Event{
		Type:      events.HumanLogFileCreatedEvent,
		Timestamp: time.Now(),
		Data:      events.HumanLogFileCreated{},
	}

	h.AggregateRoot.RecordThat(event)

	return h
}
