package adapters

import (
	"context"
	"database/sql"
	"os"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/pkg/support"
)

type Repo struct {
	DB *sql.DB
}

func (r *Repo) Add(entity logfile.LogFile, ctx context.Context) error {
	fileEntity, _ := entity.(*logfile.LogFileEntity)
	query := "INSERT INTO `log_file` (`id`, `path`) VALUES (?,?)"
	_, err := r.DB.ExecContext(ctx, query, fileEntity.ID, fileEntity.Path.String)
	return err
}

func (r *Repo) ReadFrom(path support.String) (*logfile.LogFileEntity, error) {
	if path.String == "" {
		return nil, logfile.ErrNameEmpty
	}

	file, err := os.Open(path.String)
	if err != nil {
		return nil, logfile.ErrFileNotFound
	}

	return &logfile.LogFileEntity{
		Path: path,
		File: file,
	}, nil
}

func (r *Repo) GetAll(ctx context.Context) (*[]logfile.LogFileEntity, error) {

	rows, err := r.DB.QueryContext(ctx, "SELECT * FROM `log_file`")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var logFiles []logfile.LogFileEntity

	for rows.Next() {

		var logFile logfile.LogFileEntity
		err := rows.Scan(&logFile.ID, &logFile.Path)

		if err != nil {
			return nil, err
		}
		logFiles = append(logFiles, logFile)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &logFiles, nil

}

func NewLogFileRepository(db *sql.DB) logfile.LogFileRepository {
	return &Repo{DB: db}
}
