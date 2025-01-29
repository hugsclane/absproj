package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/hugsclane/absproj/tree/webservice/webservice/config"
	"github.com/hugsclane/absproj/tree/webservice/webservice/internal/postgres"
	"github.com/hugsclane/absproj/tree/webservice/webservice/internal/redis"
	"github.com/hugsclane/absproj/tree/webservice/webservice/internal/server"
	"go.uber.org/zap/zapcore"
)

func main() {
	// create logger
	lg, lvl, err := newLogger()
	if err != nil {
		log.Fatal("failed to start logger")
	}

	// parse config
	cfg, err := env.Config()
	if err != nil {
		lg.Error("failed to parse config", zap.Error(err))
	}
	lvl.SetLevel(cfg.LogLevel)

	// create database connection
	pg, err := postgres.NewDatabase(cfg.Postgres)
	if err != nil {
		lg.Error("failed to connect to postgres", zap.Error(err))
	}

	// create redis connection
	rd, err := redis.NewDatabase(lg, cfg.Redis)
	if err != nil {
		lg.Error("failed to connect to redis", zap.Error(err))
	}

	// create web server
	srv := server.NewServer(lg, pg, rd, cfg.Server)
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

func newLogger() (*zap.Logger, zap.AtomicLevel, error) {
	cfg := zap.NewProductionConfig()

	lg, err := cfg.Build()
	if err != nil {
		return nil, cfg.Level, err
	}

	return lg, cfg.Level, nil
}
