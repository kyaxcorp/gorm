package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const postgresDSN = "user=root password= dbname=gorm host=localhost port=26257 sslmode=disable TimeZone=Europe/Chisinau"

type User struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;not null;->;default:gen_random_uuid()"`

	DeletedByID *uuid.UUID `gorm:"type:uuid;updateOnSoftDelete"`
	DeletedAt   *gorm.DeletedAt
}

func main() {
	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	var id = "60ea2d8f-4fac-42ba-b363-4536f6a8edf4"

	// Retrieve the user
	var user User
	result := db.First(&user, "id = ?", id)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	// Set the user who will delete the record
	deletedByID := uuid.New()
	user.DeletedByID = &deletedByID

	// Delete
	result = db.Delete(&user)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
}
