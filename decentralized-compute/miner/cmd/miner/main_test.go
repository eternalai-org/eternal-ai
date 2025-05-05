package main

import (
	"context"
	"go.uber.org/zap"
	"os"
	"solo/config"
	"solo/internal/factory"
	"solo/internal/model"
	"solo/internal/port"
	"solo/pkg/logger"
	"testing"
)

func logTest(t *testing.T, resp interface{}, err error) {
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(resp)
}

func createMiner() (port.IMiner, error) {
	// init flag
	configFile = os.Getenv("CNF_FILE_TEST")

	cnf, err := config.ReadConfig(configFile)
	if err != nil {
		logger.AtLog.Fatal(err)
		return nil, err
	}

	logger.GetLoggerInstanceFromContext(context.Background()).Info("ReadConfig",
		zap.Any("cfg", cnf),
	)

	err = cnf.Verify()
	if err != nil {
		logger.AtLog.Fatal(err)
		return nil, err
	}

	obj, err := factory.NewMiner(cnf)
	if err != nil {
		logger.AtLog.Fatal(err)
		return nil, err
	}

	return obj.Miner, nil
}

func TestInferChatCompletions(t *testing.T) {
	taskWatcher, err := createMiner()
	if err != nil {
		logTest(t, nil, err)
		return
	}

	ctx := context.Background()
	prompt := "{\n  \"model\": \"SentientAGI/Dobby-Mini-Unhinged-Llama-3.1-8B\",\n  \"stream\": true,\n  \"messages\": [\n    \n    {\n      \"role\": \"user\",\n      \"content\": \"hello!! a quick question. Do you know BITCOIN?\"\n    }\n  ],\n  \"max_tokens\": 4096\n}"
	//seed := pkg.CreateSeed(prompt, fmt.Sprintf("%d", time.Now().UTC().Unix()))

	qt := taskWatcher.GetTaskQueue()

	qt <- &model.Task{
		TaskID:       "1",
		AssignmentID: "1",
		Params:       prompt,
	}

	taskWatcher.ExecuteTasks(ctx)

}
