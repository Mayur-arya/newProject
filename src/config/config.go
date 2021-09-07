package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//ConnectDB function: Make database connection
func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("dbUser")
	password := os.Getenv("dbPassword")
	databaseName := os.Getenv("dbName")
	databaseHost := os.Getenv("dbHost")
	dbPort := os.Getenv("dbPort")
	port, err := strconv.Atoi(dbPort)
	/*if err == nil {
		fmt.Println(i1)
	}
	fmt.Println(err)*/
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, databaseHost, port, databaseName)
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	} else {
		fmt.Println("Connection successful")
	}
	//fmt.Println(db)
	// close db when not in use

	// Migrate the schema
	// db.AutoMigrate(&models.User{})
	return db
}
