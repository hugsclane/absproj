package redis

type Config struct {
}

type Database struct {
}

func NewDatabase(cfg Config) (*Database, error) {
	return &Database{}, nil
}
