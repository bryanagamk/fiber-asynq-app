package main

import (
	"fiber-asynq-app/database"
	"fiber-asynq-app/jobs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	database.ConnectRedis()

	app := fiber.New()

	app.Post("/users", func(c *fiber.Ctx) error {
		type Request struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		err := jobs.EnqueueCreateUserTask(req.Name, req.Email)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to enqueue job"})
		}

		return c.Status(202).JSON(fiber.Map{"message": "User creation job enqueued"})
	})

	log.Println("Server running on :3000")
	app.Listen(":3000")
}
