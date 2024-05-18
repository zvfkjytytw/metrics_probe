package main

import (
	"context"
	"flag"

	metricsApp "github.com/zvfkjytytw/metrics_probe/internal/server/app"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "./conf/server.yaml", "Server config file")
	flag.Parse()

	config, err := metricsApp.GetConfig(configFile)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := metricsApp.NewApp(config)
	if err != nil {
		panic(err)
	}

	err = app.Run(ctx)
	if err != nil {
		panic(err)
	}
}
