package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/omatheu/go-and-sqlite/models"
	"gorm.io/gorm"
)

// HTTPError represents an error response
type HTTPError struct {
	Error string `json:"error"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users [post]
func CreateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(HTTPError{Error: err.Error()})
		}
		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(HTTPError{Error: err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /users/{id} [get]
func GetUserByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(HTTPError{Error: "Invalid user ID"})
		}
		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(HTTPError{Error: "User not found"})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update a user by ID with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/{id} [put]
func UpdateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(HTTPError{Error: "Invalid user ID"})
		}
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(HTTPError{Error: err.Error()})
		}
		user.ID = uint(id)
		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(HTTPError{Error: err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} HTTPError
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/{id} [delete]
func DeleteUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(HTTPError{Error: "Invalid user ID"})
		}
		if err := db.Delete(&models.User{}, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(HTTPError{Error: err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(HTTPError{Error: "User deleted"})
	}
}
