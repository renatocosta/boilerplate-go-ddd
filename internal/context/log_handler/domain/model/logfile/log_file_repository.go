package logfile

import (
	"context"
	"errors"

	"github.com/ddd/pkg/support"
)

var (
	ErrNameEmpty    = errors.New("Path file can not be empty")
	ErrFileNotFound = errors.New("Path file not found")
)

type LogFileRepository interface {
	Add(entity LogFile, ctx context.Context) error
	ReadFrom(path support.String) (*LogFileEntity, error)
	GetAll(ctx context.Context) (*[]LogFileEntity, error)
}
