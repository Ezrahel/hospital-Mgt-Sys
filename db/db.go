package db

import (
	"hospital-mgt-sys/doctors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() {
	var err error

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect")

	}

	db.AutoMigrate(&doctors.Doctor{})

}
