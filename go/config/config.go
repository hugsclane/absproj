package config

import (
	"github.com/hugsclane/absproj/go/internal/postgres"
	"github.com/hugsclane/absproj/go/internal/redis"
	"github.com/hugsclane/absproj/go/internal/server"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel zapcore.Level
	Postgres postgres.Config
	Redis    redis.Config
	Server   server.Config
}
