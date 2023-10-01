package handlers

import (
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthHandler(ctx *fiber.Ctx) error {

	cookies := ctx.Cookies("token")
	if cookies == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	claims, err := utils.GetClaims(cookies)
	if err != nil || claims.Valid() != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.SendStatus(fiber.StatusOK)
}