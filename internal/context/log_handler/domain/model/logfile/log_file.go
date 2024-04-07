package logfile

import (
	"bufio"
	"errors"
	"os"
	"time"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/pkg/building_blocks/domain"
	"github.com/ddd/pkg/support"
	"github.com/google/uuid"
)

var (
	ErrUnableToRead = errors.New("Unable to read file")
)

type LogFile interface {
	ExtractFrom(*os.File) ([]string, error)
	Select(rows []string) *LogFileEntity
}

type LogFileEntity struct {
	domain.AggregateRoot
	Path support.String
	File *os.File
}

func NewLogFile(id uuid.UUID, p support.String) LogFile {
	return &LogFileEntity{
		AggregateRoot: domain.AggregateRoot{ID: id},
		Path:          p,
	}
}

func (l *LogFileEntity) ExtractFrom(file *os.File) ([]string, error) {
	defer file.Close()

	var rows []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, ErrUnableToRead
	}

	return rows, nil
}

func (l *LogFileEntity) Select(rows []string) *LogFileEntity {
	//Raise event
	event := domain.Event{
		Type:      events.LogFileSelectedEvent,
		Timestamp: time.Now(),
		Data: events.LogFileSelected{
			ID:      l.ID,
			Content: rows,
			Path:    l.Path,
		},
	}

	l.AggregateRoot.RecordThat(event)

	return l
}
