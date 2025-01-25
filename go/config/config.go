package config

import (
	"github.com/hugsclane/absproj/go/internal/app"
	"github.com/hugsclane/absproj/go/internal/postgres"
	"github.com/hugsclane/absproj/go/internal/redis"
)

type Config struct {
	Postgres postgres.Config
	Redis    redis.Config
	App      app.Config
}
