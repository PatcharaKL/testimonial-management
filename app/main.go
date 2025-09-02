package main

import (
	"fmt"
	"log"
	"os"

	"testimonial-management/internal/handler"
	"testimonial-management/internal/repository"
	"testimonial-management/internal/usecases"
	"testimonial-management/pkg"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Environment variables are loaded from the container environment

	dbConfig := pkg.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     getEnvAsInt("DB_PORT", 5432),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable", // or from env if needed
	}

	db, err := pkg.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Database connection established successfully.")

	app := fiber.New()

	// Allow CORS for patchara.dev/testimonial
	app.Use(func(c *fiber.Ctx) error {
		c.Response().Header.Add("Access-Control-Allow-Origin", "https://patchara.dev")
		c.Response().Header.Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Response().Header.Add("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	repo := repository.NewTestimonialRepository(db)
	usecase := usecases.NewTestimonialUsecase(repo)
	handler.RegisterTestimonialRoutes(app, usecase)

	log.Fatal(app.Listen(":3000"))
}

// getEnvAsInt reads an environment variable as int, returns fallback if not set or invalid
func getEnvAsInt(name string, fallback int) int {
	val := os.Getenv(name)
	var i int
	_, err := fmt.Sscanf(val, "%d", &i)
	if err != nil {
		return fallback
	}
	return i
}
