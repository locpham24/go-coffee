package model

import "time"

type User struct {
	tableName   struct{} `sql:"users,alias:users" pg:",discard_unknown_columns"`
	Id          int
	PhoneNumber string
	Password    string
	CreatedAt   time.Time
	CreatedBy   int
	UpdatedBy   int
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
