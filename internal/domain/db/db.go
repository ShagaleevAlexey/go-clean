package db

import (
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


// TODO: Gorm field is a public property - this is not ISP 👎
type DB struct {
	Gorm *gorm.DB
}

func NewDB(config IDBConfig, l *logrus.Logger) (*DB, error) {
	db, err := gorm.Open(config.Dialect(), config.ConnectionString())
	if err != nil {
		return nil, err
	}

	db.SetLogger(l)
	//Gorm.LogMode(true)

	return &DB{Gorm: db}, nil
}

func (db *DB) CreateSession() *Session {
	return NewSession(db.Gorm)
}

func (db *DB) Close() error {
	return db.Gorm.Close()
}

// TODO: Gorm field is a public property - this is not ISP 👎
type Session struct {
	Gorm *gorm.DB
}

func NewSession(gorm *gorm.DB) *Session {
	return &Session{Gorm: gorm.Begin()}
}

func (s *Session) BeginTransaction(value interface{}) *Session {
	return NewSession(s.Gorm)
}

func (s *Session) DefferWithRollback() {
	if r := recover(); r != nil {
		s.Gorm.Rollback()
	}
}
