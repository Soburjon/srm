package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/work/SRM/api/V1/models"
	"github.com/work/SRM/pkg/utils"
	"net/http"
	"strings"
)

// Login method user id va authorize qaytaradi
// @Description user id va authorize qaytaradi
// @Summary user id va authorize qaytaradi
// @Tags register
// @Accept json
// @Produce json
// @Param login body models.RegisterRequest true "login"
// @Success 200 {object} models.LoginResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /register/login/ [POST]
func (a *Api) Login(c *fiber.Ctx) error {
	req := models.RegisterRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !strings.Contains(req.PhoneNumber, "+") ||
		len(req.PhoneNumber) != 13 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "phone number is not correctly filled",
		})
	}
	res, err := a.RegisterService.Login(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(res.UserID, map[string]string{
		"role": res.Role,
	})
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.LoginResponse{
		UserID:    res.UserID,
		Authorize: tokens.Access,
	})
}
