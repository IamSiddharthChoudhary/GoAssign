package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-Id")
		if requestID == "" {
			requestID = uuid.NewString()
		}

		c.Set("X-Request-Id", requestID)
		c.Locals("request_id", requestID)

		return c.Next()
	}
}
