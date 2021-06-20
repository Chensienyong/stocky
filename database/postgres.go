package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db     *gorm.DB
}

// Option holds all necessary options for database.
type Option struct {
	User         string
	Password     string
	Host         string
	Port         string
	Database     string
	Timezone     string
}

func NewPostgres(opt Option) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", opt.Host, opt.User, opt.Password, opt.Database, opt.Port, opt.Timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &Postgres{}, err
	}

	return &Postgres{Db: db}, nil
}
