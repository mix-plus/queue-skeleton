package svc

import (
	"fmt"
	"github.com/mix-plus/queue-skeleton/internal/config"

	"github.com/hibiken/asynq"
)

func newAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.RedisConf.Addr, Password: c.RedisConf.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
		},
	)
}
