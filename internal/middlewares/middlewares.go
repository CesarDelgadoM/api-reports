package middlewares

import (
	"strconv"

	"github.com/CesarDelgadoM/api-reports/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Validate jwt token of the cookie
func IsAuthenticated(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.SecretKey), nil
	})
	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	return ctx.Next()
}

// Get userid of the cookie
func GetUserId(ctx *fiber.Ctx) (uint, error) {
	cookie := ctx.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.SecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*jwt.StandardClaims)

	id, _ := strconv.Atoi(payload.Subject)

	return uint(id), nil
}
