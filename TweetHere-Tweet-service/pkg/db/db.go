package db

import (
	"fmt"
	"tweet-service/pkg/config"
	"tweet-service/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBname, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(&domain.Post{})
	db.AutoMigrate(&domain.Like{})
	db.AutoMigrate(&domain.BookMark{})
	db.AutoMigrate(&domain.Comment{})

	return db, dbErr
}
