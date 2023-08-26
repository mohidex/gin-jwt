package main

import (
	"fmt"
	"log"

	"github.com/mohidex/identity-service/auth"
	"github.com/mohidex/identity-service/config"
	"github.com/mohidex/identity-service/db"
	"github.com/mohidex/identity-service/models"
	"github.com/mohidex/identity-service/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	conf := config.NewConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos",
		conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBPort)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	if err = gormDB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}
	pgDB := db.NewPgDB(gormDB)
	jwtAuth := auth.NewJWTAuthenticator(conf.JWTPrivateKey, conf.JWTTTL)
	r := server.NewServer(pgDB, jwtAuth)
	if err := r.Start(":5000"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
