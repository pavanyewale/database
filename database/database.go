package database

import "context"

type Database interface {
	Ping(ctx context.Context) error
}

type DBName string

const (
	PostgreSQL DBName = "PostgreSQL"
)

type Config struct {
	Name DBName
}

func NewDatabase(conf *Config) Database {

	return nil
}
