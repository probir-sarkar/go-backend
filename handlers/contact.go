package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/probir-sarkar/go-backend/database"
)

var validate = validator.New()

type ContactForm struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message" validate:"required"`
	Origin  string `json:"origin" validate:"required"`
}

func handleValidationError(err error) fiber.Map {
	// Handle validation errors
	var validationErrors []string
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, err.Field()+" is invalid: "+err.Tag())
	}
	return fiber.Map{
		"error":   "Validation failed",
		"details": validationErrors,
	}
}

func SubmitContactForm(c *fiber.Ctx) error {
	var request ContactForm

	// Parse the JSON body into the struct
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate the request body
	if err := validate.Struct(request); err != nil {
		// Use reusable function to handle validation errors
		return c.Status(fiber.StatusBadRequest).JSON(handleValidationError(err))
	}

	// Insert into the database
	result := database.DB.Create(&request)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save contact form record",
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Contact form submitted successfully!",
		"data":    request,
	})
}
