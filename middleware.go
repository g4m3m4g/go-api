package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
	"time"

)

func middleware(c *fiber.Ctx) error {
	fmt.Printf("Request: %s\nTime: %s\n", c.OriginalURL(), time.Now())
	return c.Next()
}