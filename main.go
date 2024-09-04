package main

import (
	"github.com/employee/database"
	"github.com/employee/handler"

	"fmt"
	"log"

	//"net/http"
	"os"

	"github.com/employee/model"
	"github.com/employee/repository"
	"github.com/employee/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := database.NewConnection(config)
/////
	if err != nil {
		log.Fatal("could not load the data base")
	}
	err = model.MigrateEmployee(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	// Create an instance of EmployeeService with the database connection
	//fmt.Println("db", db)

	repo := &repository.EmployeeRepository{DB: db}
	//fmt.Println("repo", repo)
	employeeService := &service.EmployeeService{Repo: repo}

	// Create a new Fiber app
	app := fiber.New()

	// Register employee handler with the employeeService
	handler.RegisterEmployeeHandler(app, employeeService)

	// Serve Swagger UI
	//app.Static("/swagger", "../swagger") // Assuming your Swagger UI files are in the "swagger" directory

	// Start the server
	err = app.Listen(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
