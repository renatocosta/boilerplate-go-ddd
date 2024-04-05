package adapters

import (
	"context"
	"database/sql"
	"os"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
)

type Repo struct {
	DB *sql.DB
}

func (r *Repo) Add(entity logfile.LogFile, ctx context.Context) error {
	fileEntity, _ := entity.(*logfile.LogFileEntity)
	query := "INSERT INTO `log_file` (`id`, `path`) VALUES (?,?)"
	_, err := r.DB.ExecContext(ctx, query, fileEntity.ID, fileEntity.Path)
	return err
}

func (r *Repo) ReadFrom(path string) (*logfile.LogFileEntity, error) {
	if path == "" {
		return nil, logfile.ErrNameEmpty
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, logfile.ErrFileNotFound
	}

	return &logfile.LogFileEntity{
		Path: path,
		File: file,
	}, nil
}

func NewLogFileRepository(db *sql.DB) logfile.LogFileRepository {
	return &Repo{DB: db}
}
