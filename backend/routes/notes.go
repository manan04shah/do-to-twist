package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/do-to-list/database"
	"github.com/manan04shah/do-to-list/models"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title  string `gorm:"size:255;not null;" json:"title"`
	Body   string `gorm:"not null;" json:"body"`
	UserID uint   `json:"user_id"` // New field to store the associated user's ID
}

func CreateResponseNote(noteModel models.Note) Note {
	return Note{
		ID:     noteModel.ID,
		Title:  noteModel.Title,
		Body:   noteModel.Body,
		UserID: noteModel.UserID,
	}
}

func CreateNote(c *fiber.Ctx) error {
	var note models.Note

	if err := c.BodyParser(&note); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&note)
	responseNote := CreateResponseNote(note)

	return c.Status(200).JSON(responseNote)
}

func GetNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note

	database.Database.Db.First(&note, id)
	responseNote := CreateResponseNote(note)

	return c.Status(200).JSON(responseNote)
}

func UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateNote models.Note

	if err := c.BodyParser(&updateNote); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var note models.Note
	database.Database.Db.First(&note, id)
	if updateNote.Title != "" {
		note.Title = updateNote.Title
	}
	if updateNote.Body != "" {
		note.Body = updateNote.Body
	}
	if updateNote.UserID != 0 {
		note.UserID = updateNote.UserID
	}

	database.Database.Db.Save(&note)

	responseNote := CreateResponseNote(note)

	return c.Status(200).JSON(responseNote)
}

func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note

	database.Database.Db.Delete(&note, id)

	return c.Status(200).JSON(fiber.Map{
		"message": "Note deleted successfully",
	})
}
