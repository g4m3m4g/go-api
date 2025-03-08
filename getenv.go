package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
)

func getEnv(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}