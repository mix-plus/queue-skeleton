package svc

import (
	"github.com/hibiken/asynq"
	"github.com/mix-plus/queue-skeleton/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
	}
}
