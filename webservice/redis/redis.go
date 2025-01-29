// create and interface for the rest of the application to interact with redis
package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/hugsclane/absproj/tree/webservice/internal/model"
	"go.uber.org/zap"
)

var (
	ErrNotExists = fmt.Errorf("the row does not exist")
)

type Config struct {
	Host string
	Port int
}

type Database struct {
	lg   *zap.Logger
	pool *redis.Pool
}

func NewDatabase(lg *zap.Logger, cfg Config) (*Database, error) {
	return &Database{
		lg: lg,
		pool: &redis.Pool{
			MaxIdle:         10,
			MaxActive:       100,
			MaxConnLifetime: time.Minute,
			DialContext: func(ctx context.Context) (redis.Conn, error) {
				return redis.DialContext(ctx, "tcp", cfg.Host+":"+strconv.Itoa(cfg.Port))
			},
		},
	}, nil
}

func (d *Database) closeLog(c redis.Conn) {
	if err := c.Close(); err != nil {
		d.lg.Error("failed to close redis connection", zap.Error(err))
	}
}

// GetDataSet - gets a dataset
func (d *Database) GetDataSet(ctx context.Context, dataSetKey string) (*model.DataSet, error) {
	c, err := d.pool.GetContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get redis conn: %w", err)
	}
	defer d.closeLog(c)

	var ds model.DataSet
	b, err := redis.Bytes(c.Do("GET", dataSetKey))
	if err != nil {
		if err == redis.ErrNil {
			return nil, ErrNotExists
		}
		return nil, fmt.Errorf("failed to get dataset from redis: %w", err)
	}

	err = json.Unmarshal(b, &ds)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal dataset: %w", err)
	}

	return &ds, nil
}

// SetDataSet  - set a dataset
func (d *Database) AddDataSet(ctx context.Context, dataSet *model.DataSet) error {
	c, err := d.pool.GetContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to get redis conn: %w", err)
	}
	defer d.closeLog(c)

	b, err := json.Marshal(dataSet)
	if err != nil {
		return fmt.Errorf("failed to marshal dataset: %w", err)
	}

	_, err = c.Do("SET", dataSet.Key, b)
	if err != nil {
		return fmt.Errorf("failed to set dataset in redis: %w", err)
	}
	return nil
}

func (d *Database) Close() error {
	return d.pool.Close()
}
