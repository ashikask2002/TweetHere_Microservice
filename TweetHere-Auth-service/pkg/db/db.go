package db

import (
	"Tweethere-Auth/pkg/config"
	"Tweethere-Auth/pkg/domain"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBname, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.UserOTPLogin{})
	db.AutoMigrate(&domain.FollowingRequest{})
	db.AutoMigrate(&domain.Followings{})
	db.AutoMigrate(&domain.Followers{})

	CheckAndCreateAdmin(db)
	return db, dbErr

}

func CheckAndCreateAdmin(db *gorm.DB){
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0{
		password :="tweethere"
		hashedpassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err != nil{
			return
		}
		admin := domain.Admin{
			ID: 1,
			Firstname: "tweethere",
			Lastname: "admin",
			Email: "tweethere@gmail.com",
			Password: string(hashedpassword),
		}
		db.Create(admin)
	}
}
