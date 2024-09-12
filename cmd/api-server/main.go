package main

import (
	"fmt"
	"os"
	"time"

	config "github.com/hugoaguirre/go-ms-template/internal"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/hugoaguirre/go-ms-template/cmd/api-server/server"
)

const (
	APP_RETRY_TIMER_IN_SECONDS = 30 * time.Second
)

func main() {
	log.Infof("hello !")
	log.SetReportCaller(true)
	app := getApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func getApp() *cli.App {
	return &cli.App{
		Name:     "gouh",
		Usage:    "gouh gouh",
		Action:   StartServer,
		Flags:    []cli.Flag{},
		Commands: []*cli.Command{},
	}
}

func StartServer(c *cli.Context) error {
	config, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config from env: %w", err)
	}

	for {
		err := server.StartServer(config)
		if err != nil {
			log.Println(err)
			time.Sleep(APP_RETRY_TIMER_IN_SECONDS)
		}
	}
}
