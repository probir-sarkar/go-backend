package database

import "log"

func Migrate() {
	err := DB.AutoMigrate(&ContactForm{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
	log.Println("Database schema migrated successfully!")
}
