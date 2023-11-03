package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDatabase() (err error) {
	dsn := "host=localhost user=bob password=admin dbname=jwtgo port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
