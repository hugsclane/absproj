package config

import (
	"github.com/hugsclane/absproj/tree/webservice/webservice/internal/postgres"
	"github.com/hugsclane/absproj/tree/webservice/webservice/internal/redis"
	"github.com/hugsclane/absproj/tree/webservice/webservice/internal/server"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel zapcore.Level
	Postgres postgres.Config
	Redis    redis.Config
	Server   server.Config
}
