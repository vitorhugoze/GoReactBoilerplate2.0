package handlers

import (
	"encoding/json"
	"server/models"
	"server/services"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func UpdateUserHandler(ctx *fiber.Ctx) error {

	var reqUsr models.User

	claims, err := utils.GetClaims(ctx.Cookies("token"))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(ctx.Body(), &reqUsr); err != nil {
		return err
	}

	reqUsr.User_id = claims.User.User_id

	if err = services.UpdateUser(&reqUsr); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
