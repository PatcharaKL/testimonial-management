package handler

import (
	"testimonial-management/internal/dto"
	"testimonial-management/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

func RegisterTestimonialRoutes(app *fiber.App, usecase *usecases.TestimonialUsecase) {
	app.Post("/testimonial", func(c *fiber.Ctx) error {
		var req dto.CreateTestimonialRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		if req.FullName == "" || req.Testimonial == "" {
			return c.Status(400).JSON(fiber.Map{"error": "full_name and testimonial required"})
		}

		if err := usecase.CreateTestimonial(&req); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save testimonial"})
		}
		return c.Status(201).JSON(fiber.Map{"message": "Testimonial created"})
	})

	app.Get("/testimonial", func(c *fiber.Ctx) error {
		testimonials, err := usecase.GetAllTestimonials()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch testimonials"})
		}
		return c.JSON(testimonials)
	})
}
