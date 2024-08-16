package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/omatheu/go-and-sqlite/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func GetUserByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func UpdateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		user.ID = uint(id)
		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func DeleteUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		if err := db.Delete(&models.User{}, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted"})
	}
}
