package config

import (
	"github.com/hugsclane/absproj/go/internal/postgres"
	"github.com/hugsclane/absproj/go/internal/redis"
	"github.com/hugsclane/absproj/go/internal/server"
)

type Config struct {
	Postgres postgres.Config
	Redis    redis.Config
	Server   server.Config
}
