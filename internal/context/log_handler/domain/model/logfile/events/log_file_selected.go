package events

import "github.com/google/uuid"

const LogFileSelectedEvent = "LogFileSelected"

type LogFileSelected struct {
	ID      uuid.UUID
	Content []string
	Path    string
}
