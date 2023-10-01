package handlers

import (
	"server/services"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteUserHandler(ctx *fiber.Ctx) error {

	claims, err := utils.GetClaims(ctx.Cookies("token"))
	if err != nil {
		return err
	}

	if err = services.DeleteUser(&claims.User); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
