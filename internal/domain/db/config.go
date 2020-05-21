package db

import (
	"fmt"
	"github.com/caarlos0/env"
)

type IDBConfig interface {
	Dialect() string
	ConnectionString() string
}

type DBSQLiteConfig struct {
	Path string `json:"path,omitempty" env:"DB_FILE_PATH"`
}

func NewDBSQLiteConfig() (*DBSQLiteConfig, error) {
	c := &DBSQLiteConfig{}
	if err := c.Import(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *DBSQLiteConfig) Import() error {
	if err := env.Parse(c); err != nil {
		return err
	}

	return nil
}

func (c *DBSQLiteConfig) Dialect() string {
	return "sqlite3"
}

func (c *DBSQLiteConfig) ConnectionString() string {
	return c.Path
}

type DBPostgreSQLConfig struct {
	Host     string `json:"host,omitempty" env:"DB_HOST" envDefault:"0.0.0.0"`
	Port     int    `json:"port,omitempty" env:"DB_PORT" envDefault:"5432"`
	User     string `json:"user,omitempty" env:"DB_USER" envDefault:"guest"`
	Password string `json:"password,omitempty" env:"DB_PASSWORD" envDefault:""`
	DBName   string `json:"dbname,omitempty" env:"DB_NAME" envDefault:"go-clean"`
}

func NewDBPostgreSQLConfigFromEnv() (*DBPostgreSQLConfig, error) {
	c := &DBPostgreSQLConfig{}
	if err := c.Import(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *DBPostgreSQLConfig) Import() error {
	if err := env.Parse(c); err != nil {
		return err
	}

	return nil
}

func (c *DBPostgreSQLConfig) Dialect() string {
	return "postgres"
}

func (c *DBPostgreSQLConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DBName)
}
