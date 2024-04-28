package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/do-to-list/database"
	"github.com/manan04shah/do-to-list/models"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"size:255;not null;" json:"name"`
	Email string `gorm:"size:255;not null;uniqueIndex" json:"email"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	database.Database.Db.First(&user, id)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateUser models.User

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	database.Database.Db.First(&user, id)
	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		user.Password = updateUser.Password
	}
	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	database.Database.Db.Delete(&user, id)

	return c.Status(200).JSON("User deleted successfully")
}

// Get all notes of a user
func GetUserNotes(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	database.Database.Db.Preload("Notes").First(&user, id)

	return c.Status(200).JSON(user)
}
