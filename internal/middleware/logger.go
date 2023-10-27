package middleware

import (
	"github.com/gofiber/fiber/v2"
	"shorturl/internal/pkg/logging"
)

func Logger(c *fiber.Ctx) error {
	err := c.Next()

	path := string(c.Request().URI().Path())
	method := c.Method()
	code := c.Response().StatusCode()

	logging.Log.Infoln(path, method, code)

	return err
}
