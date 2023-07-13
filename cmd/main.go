package main

import (
	"fmt"
	"os"

	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/handlers"
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/repositories"
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/routes"
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	eventRepository := repositories.NewEventRepository(db)

	eventService := services.NewEventService(*eventRepository)

	eventHandler := handlers.NewEventHandler(*eventService)

	routes.SetupRoutes(r, eventHandler)

	r.Run(":8080")
}
