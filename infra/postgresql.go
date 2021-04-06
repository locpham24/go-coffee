package infra

import (
	"fmt"
	"github.com/locpham24/go-coffee/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGDB struct {
	*gorm.DB
}

var pgSingleton *PGDB

func InitPostgreSQL(config config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.PostgreSQL.Host,
		config.PostgreSQL.Username,
		config.PostgreSQL.Password,
		config.PostgreSQL.Name,
		config.PostgreSQL.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return err
	}

	pgSingleton = &PGDB{db}

	return nil
}

func GetDB() *PGDB {
	if pgSingleton == nil {
		panic("can not connect database")
	}
	return pgSingleton
}
