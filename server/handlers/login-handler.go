package handlers

import (
	"encoding/json"
	"server/models"
	"server/services"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {

	var user models.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return err
	}

	dbUsr, err := services.GetUser(&user, false)
	if err != nil {
		return err
	}

	token, err := utils.GenerateJwt(dbUsr)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: token,
	})

	return c.SendStatus(200)
}
