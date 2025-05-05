package main

import (
	"context"
	"flag"
	"go.uber.org/zap"
	_ "net/http/pprof"
	"solo/internal/factory"
	"time"

	"solo/pkg"

	"solo/config"
	"solo/pkg/logger"
)

var configFile string

func main() {
	// init flag
	flag.StringVar(&configFile, "config-file", "env/development.yml", "Config file path")
	flag.Parse()

	cnf, err := config.ReadConfig(configFile)
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	logger.GetLoggerInstanceFromContext(context.Background()).Info("ReadConfig",
		zap.Any("cfg", cnf),
	)

	err = cnf.Verify()
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	object, err := factory.NewMiner(cnf)
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	taskWatcher := object.Miner
goto_here:
	verifed := taskWatcher.Verify()
	if !verifed {
		_, _, err := taskWatcher.MakeVerify()
		if err != nil {
			logger.AtLog.Error(err)
		}

		time.Sleep(time.Second * pkg.TimeToWating)
		goto goto_here
	}

	done := make(chan bool)

	// get and process tasks
	// taskWatcher.RejoinForMinting(ctx)
	ctx := context.Background()
	go taskWatcher.GetPendingTasks(ctx)

	go taskWatcher.ExecuteTasks(ctx)

	<-done
}
