package main

import (
	"flag"
	"solo/internal/delivery/http"
	"solo/internal/factory"
	"solo/pkg/logger"
	"strconv"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "9000", "Config file path")
	flag.Parse()

	portInt, err := strconv.Atoi(port)
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	initObject, _ := factory.NewAPI(portInt)
	cmd, _ := http.NewHttp(initObject.API, portInt)
	cmd.Run()
}
