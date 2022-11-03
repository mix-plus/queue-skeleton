package main

import (
	"context"
	"flag"
	"github.com/mix-plus/core/conf"
	"github.com/mix-plus/core/di"
	"github.com/mix-plus/queue-skeleton/internal/config"
	"github.com/mix-plus/queue-skeleton/internal/logic"
	"github.com/mix-plus/queue-skeleton/internal/svc"
	"os"
)

var configFile = flag.String("f", "etc/config.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		di.Zap().Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
