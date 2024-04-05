package logfile

import (
	"context"
	"errors"
)

var (
	ErrNameEmpty    = errors.New("Path file can not be empty")
	ErrFileNotFound = errors.New("Path file not found")
)

type LogFileRepository interface {
	Add(entity LogFile, ctx context.Context) error
	ReadFrom(path string) (*LogFileEntity, error)
}
