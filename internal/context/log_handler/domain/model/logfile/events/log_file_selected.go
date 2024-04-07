package events

import (
	"github.com/ddd/pkg/support"
	"github.com/google/uuid"
)

const LogFileSelectedEvent = "LogFileSelected"

type LogFileSelected struct {
	ID      uuid.UUID
	Content []string
	Path    support.String
}
