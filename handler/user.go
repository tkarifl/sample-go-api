package handler

import (
	"errors"
	"fmt"
	"manualVuln/database"
	"manualVuln/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Get all users query
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User
	db.Find(&users)
	return c.JSON(fiber.Map{"status": "success", "message": "All users", "data": users})
}

// Get user by given id query
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	err := db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with given ID", "data": nil})
	} else if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Trouble with getting the user", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

// Create new user query
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": nil})
	} else if err := db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

// Update user query
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	newUserData := new(model.User)
	if err := c.BodyParser(newUserData); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't update user", "data": nil})
	}
	if err := db.First(&user, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with given ID", "data": nil})
	} else if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Trouble with getting the user", "data": nil})
	}
	err := db.Model(&user).Updates(
		map[string]interface {
		}{"name": newUserData.Name,
			"phone": newUserData.Phone,
			"mail":  newUserData.Mail,
		}).Error
	if err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't update user", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Updated user", "data": nil})
}

// Delete User query
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	err := db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with given ID", "data": nil})
	} else if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Trouble with getting the user", "data": nil})
	}
	if err := db.Delete(&user).Error; err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete given user", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
