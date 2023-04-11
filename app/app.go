package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

const postgresDSN = "user=root password= dbname=gorm host=localhost port=26257 sslmode=disable TimeZone=Europe/Chisinau"

type User struct {
	ID       uuid.UUID
	UserName string

	UpdatedAt   *time.Time
	DeletedByID *uuid.UUID `gorm:"updateOnSoftDelete"`
	DeletedAt   *gorm.DeletedAt
}

func main() {
	log.Println("hello world")
	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//var users []User
	//result := db.Find(&users)
	//if result.Error != nil {
	//	log.Fatalln(result.Error)
	//}
	//
	//log.Println(users)

	var id = "60ea2d8f-4fac-42ba-b363-4536f6a8edf4"

	var user User
	result := db.First(&user, "id = ?", id)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	deletedByID := uuid.New()
	user.DeletedByID = &deletedByID

	log.Println(user)

	log.Println("DELETING......")
	result = db.Delete(&user)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	//log.Println("RESTORING...")
	//db.Model(&User{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil)
}
