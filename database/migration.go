package database

import (
	"log"

	"github.com/arwahyu01/go-jwt/app/models/user"
)

func AutoMigrateTables() {
	err := DB.AutoMigrate(
		&user.Users{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully!")
}
