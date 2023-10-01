package handlers

import (
	"encoding/json"
	"server/models"
	"server/services"

	"github.com/gofiber/fiber/v2"
)

func SignUpHandler(ctx *fiber.Ctx) error {

	var user models.User

	json.Unmarshal(ctx.Body(), &user)

	if err := services.CreateUser(&user); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
