package middlewares

import (
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {

	cookies := ctx.Cookies("token")
	if cookies == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	claims, err := utils.GetClaims(cookies)
	if err != nil || claims.Valid() != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	if token, err := utils.GenerateJwt(&claims.User); err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	} else {
		ctx.Cookie(&fiber.Cookie{
			Name:  "token",
			Value: token,
		})
	}

	return ctx.Next()
}
