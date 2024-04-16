package service

import (
	"context"
	"database/sql"

	"github.com/ddd/internal/shared"
	"github.com/ddd/pkg/building_blocks/infra/bus"
)

type ConfigProvider interface {
	GetDB() *sql.DB
	GetEventBus() *bus.EventBus
	GetWorkFlow() shared.WorkFlowable
}

type Config struct {
	ctx      context.Context
	Db       *sql.DB
	Workflow shared.WorkFlowable
	EventBus *bus.EventBus
}

func (c *Config) GetDB() *sql.DB {
	return c.Db
}

func (c *Config) GetEventBus() *bus.EventBus {
	return c.EventBus
}

func (c *Config) GetWorkFlow() shared.WorkFlowable {
	return c.Workflow
}

func NewConfig(ctx context.Context) (ConfigProvider, error) {
	c := &Config{}
	db, err := GetDb()
	if err != nil {
		return c, err
	}

	c.ctx = ctx
	c.Db = db
	c.EventBus = bus.NewEventBus()

	return c, nil
}
