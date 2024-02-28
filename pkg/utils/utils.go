package utils

import "github.com/gofiber/fiber/v2"

const (
	SecretKey string = "reports-project-key"
)

func GetRequestURI(ctx *fiber.Ctx) string {
	return string(ctx.Request().RequestURI())
}
