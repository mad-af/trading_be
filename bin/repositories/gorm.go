package repositories

import (
	"fmt"
	"time"
	cf "trading_be/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitPostgre() *gorm.DB {
	var connection = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cf.Env.PostgreHost,
		cf.Env.PostgreUser,
		cf.Env.PostgrePassword,
		cf.Env.PostgreDbName,
		cf.Env.PostgrePort,
		cf.Env.PostgreSslMode,
	)

	var db, err = gorm.Open(postgres.Open(connection), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic("Failed to connect database postgre")
	}

	if err := db.AutoMigrate(ModelTables...); err != nil {
		panic("Migration: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to create pool connection database postgre")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
