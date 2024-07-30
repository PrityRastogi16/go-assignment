package db

import (
	"fmt"
	"newAssignment/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type User struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	Email    string `gorm:"unique;not null"`
// 	Password string `gorm:"not null"`
// }

var DB *gorm.DB

func InitDB() {

	//we read our .env file
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbname)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database!")
	} else {
		fmt.Println("Connected to the database!")
		DB = db
		MigrateDB()
	}

}

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Author{}, &models.Blog{}, &models.VerificationToken{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	} else {
		fmt.Println("Database migrated successfully!")
	}
}

// func createTable() {
// 	createUserTable := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		email TEXT NOT NULL UNIQUE,
// 		password TEXT NOT NULL
// 	)`
// 	_, err := DB.Exec(createUserTable)
// 	if err != nil {
// 		panic("Not created user table: " + err.Error())
// 	}

// 	createEventsTable := `
// 	CREATE TABLE IF NOT EXISTS events (
// 		id SERIAL PRIMARY KEY,
// 		name TEXT NOT NULL,
// 		description TEXT NOT NULL,
// 		location TEXT NOT NULL,
// 		dateTime TIMESTAMPTZ NOT NULL,
// 		user_id INTEGER,
// 		FOREIGN KEY(user_id) REFERENCES users(id)
// 	)`
// 	_, err = DB.Exec(createEventsTable)
// 	if err != nil {
// 		panic("Not created events table: " + err.Error())
// 	}

// 	createRegistrationTable := `
// 	CREATE TABLE IF NOT EXISTS registrations (
// 		id SERIAL PRIMARY KEY,
// 		event_id INTEGER NOT NULL,
// 		user_id INTEGER,
// 		FOREIGN KEY(event_id) REFERENCES events(id),
// 		FOREIGN KEY(user_id) REFERENCES users(id)
// 	)`
// 	_, err = DB.Exec(createRegistrationTable)
// 	if err != nil {
// 		panic("Not created registrations table: " + err.Error())
// 	}
// }
