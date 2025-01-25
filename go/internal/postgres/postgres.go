package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

type Config struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

var (
	ErrNotExists   = fmt.Errorf("the row does not exist")
	ErrMoreThanOne = fmt.Errorf("there is more than one row")
)

func NewDatabase(c Config) (*Database, error) {
	pool, err := pgxpool.NewWithConfig(context.Background(), &pgxpool.Config{
		ConnConfig: &pgx.ConnConfig{
			Config: pgconn.Config{
				Host:     c.Host,
				Port:     c.Port,
				Database: c.Database,
				User:     c.User,
				Password: c.Password,
			},
		},

		// give us some conns for throughput
		MinConns: 10,
		MaxConns: 1000,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool %w", err)
	}
	return &Database{
		pool: pool,
	}, nil
}

// GetDataSet - gets a dataset
func (d *Database) GetDataSet(ctx context.Context, dataSetKey string) (*DataSet, error) {
	rows, err := d.pool.Query(ctx, `SELECT key, data FROM data_set WHERE key = $1`, dataSetKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get the data set %w", err)
	}
	defer rows.Close()

	var ds DataSet
	err = scanOne(rows, func(rows pgx.Rows) error {
		return rows.Scan(&ds.Key, &ds.Data)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to scan data set %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan data set %w", err)
	}

	return &ds, nil
}

func scanOne(rows pgx.Rows, scanFunc func(rows pgx.Rows) error) error {
	var ct int
	for rows.Next() {
		if ct > 0 {
			return ErrMoreThanOne
		}
		if err := scanFunc(rows); err != nil {
			return err
		}
		ct++
	}
	if ct == 0 {
		return ErrNotExists
	}
	return nil
}

// AddDataSet  - adds a dataset
func (d *Database) AddDataSet(ctx context.Context, dataSet *DataSet) error {
	_, err := d.pool.Exec(ctx, `INSERT INTO data_sets (
		key,
		data
	) VALUES (
	  $1,
	  $2
	)`, dataSet.Key, dataSet.Data)

	if err != nil {
		return fmt.Errorf("failed to est data set %w", err)
	}

	return nil
}
