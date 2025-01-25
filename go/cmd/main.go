package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/hugsclane/absproj/go/config"
	"github.com/hugsclane/absproj/go/internal/postgres"
	"go.uber.org/zap"
)

func main() {
	cfg := configFromEnv()

	// app context
	ctx, cancel := context.WithCancel(context.Background())

	// create logger
	lg, err := newLogger()
	if err != nil {
		log.Fatal("failed to start logger")
	}

	// create database connection
	pg, err := postgres.NewDatabase(cfg.Postgres)
	if err != nil {
		lg.Error("failed to connect to postgres", zap.Error(err))
	}

	// create redis connection
	rd, err := redis.NewDatabase(cfg.Redis)
	if err != nil {
		lg.Error("failed to connect to redis", zap.Error(err))
	}

	// create web server
	srv := server.New(lg, postgres, redis)
	go func() {
		if err := srv.Listen(); err != nil {
			lg.Error("application closed unexpectedly with error", zap.Error(err))
		}
	}()

	// wait for an interrupt from the command line
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	if err := srv.Close(); err != nil {
		lg.Error("failed to close pg", zap.Error(err))
	}
	if err := pg.Close(); err != nil {
		lg.Error("failed to close pg", zap.Error(err))
	}
	if err := rd.Close(); err != nil {
		lg.Error("failed to close redis", zap.Error(err))
	}
}

func configFromEnv() *config.Config {

}

func newLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	return cfg.Build()
}
