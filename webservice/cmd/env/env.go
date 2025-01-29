
package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hugsclane/absproj/tree/webservice/webservice/config"
	"github.com/hugsclane/absproj/tree/webservice/webservice/internal/postgres"
	"github.com/hugsclane/absproj/tree/webservice/webservice/internal/redis"
	"github.com/hugsclane/absproj/tree/webservice/webservice/internal/server"
	"go.uber.org/zap/zapcore"
)

// Config - returns config derived from the environment
func Config() (config.Config, error) {
	var pe parserErrors

	cfg := config.Config{
		LogLevel: logLevel("APP_LOGLEVEL", &pe),
		Postgres: postgres.Config{
			Host:     os.Getenv("APP_POSTGRES_HOST"),
			Port:     uint16Env("APP_POSTGRES_PORT", &pe),
			Database: os.Getenv("APP_POSTGRES_DATABASE"),
			User:     os.Getenv("APP_POSTGRES_USER"),
			Password: os.Getenv("APP_POSTGRES_PASSWORD"),
		},
		Redis: redis.Config{
			Host: os.Getenv("APP_REDIS_HOST"),
			Port: intEnv("APP_REDIS_PORT", &pe),
		},
		Server: server.Config{
			Port: intEnv("APP_SERVER_PORT", &pe),
		},
	}

	if len(pe.errs) > 0 {
		msgs := []string{}
		for _, e := range pe.errs {
			msgs = append(msgs, fmt.Sprintf("error parsing env var %s: %s", e.envVar, e.msg))
		}
		return config.Config{}, errors.New(strings.Join(msgs, ","))
	}

	return cfg, nil
}

type parserErrors struct {
	errs []parserError
}

type parserError struct {
	envVar string
	msg    string
}

func (pe *parserErrors) add(envVar, msg string) {
	pe.errs = append(pe.errs, parserError{envVar: envVar, msg: msg})
}

func logLevel(envVar string, pe *parserErrors) zapcore.Level {
	val := os.Getenv(envVar)

	lvl, err := zapcore.ParseLevel(val)
	if err != nil {
		pe.add(envVar, err.Error())
		return 0
	}

	return lvl
}

func intEnv(envVar string, pe *parserErrors) int {
	val := os.Getenv(envVar)
	ret, err := strconv.Atoi(val)
	if err != nil {
		pe.add(envVar, err.Error())
		return 0
	}
	return ret
}

func uint16Env(envVar string, pe *parserErrors) uint16 {
	val := os.Getenv(envVar)
	ret, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		pe.add(envVar, err.Error())
		return 0
	}
	return uint16(ret)
}
