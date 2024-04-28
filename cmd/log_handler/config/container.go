package config

import (
	"context"
	"database/sql"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/internal/context/log_handler/infra/adapters"
	"github.com/ddd/internal/shared"
	"github.com/ddd/pkg/building_blocks/infra/bus"
	"github.com/ddd/pkg/support"
)

type Config struct {
	Database *sql.DB
	EventBus *bus.EventBus
	WorkFlow shared.WorkFlowable
	Repo     logfile.LogFileRepository
}

func Start(ctx context.Context, workFlow shared.WorkFlowable) (*Config, error) {

	cfg := GetConfig()

	err := cfg.DatabaseOpen()
	if err != nil {
		return cfg, err
	}

	cfg.Repo = adapters.NewLogFileRepository(cfg.Database)
	cfg.EventBus = bus.NewEventBus()
	cfg.WorkFlow = workFlow

	//cfg.Log = lognative.NewLogNative()

	return cfg, nil
}

func GetConfig() *Config {
	err := Load()
	support.PanicOnError(err, "Unable to load envs.")

	return &Config{}
}

func (r Config) Close() {
	r.Database.Close()
}
