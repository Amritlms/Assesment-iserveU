package models

import (
	"fmt"
	"log"
	"os"
	"time"
	"errors"
	"github.com/akhil/go-bookstore/api/config"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

type User struct{
	ID        	uint32    `gorm:"primary_key;unique" json:"id"`
	Name  	string    `gorm:"size:255;not null;" json:"name"`
	CreatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func init(){
	// loads values from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	config.Connect(os.Getenv("DB_USER"),  os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() (*User,error){
	err := db.Where("id = ?", u.ID).First(&User{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		db.NewRecord(u)
		db.Create(&u)
		return u, nil
	} else {
		return &User{}, errors.New("User Already Exists")
	}
}