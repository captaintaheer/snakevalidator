// main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"snakevalidator/routes"
)

// var db *sql.DB

// const (
// 	dbHost     = "Host"
// 	dbPort     = "5432"
// 	dbUser     = "postgres"
// 	dbPassword = "pass"
// 	dbName     = "snake"
// )

// func init() {
// 	var err error
// 	// Connection string format: "user=username dbname=snake_validator sslmode=disable"
// 	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
// 		dbUser, dbPassword, dbHost, dbPort, dbName)

// 	db, err = sql.Open("postgres", connectionString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	r := gin.Default()

	routes.DefineRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
