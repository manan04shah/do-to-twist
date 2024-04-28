package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/do-to-list/database"
	"github.com/manan04shah/do-to-list/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber!")
}

func setupRoutes(app *fiber.App) {
	//Welcome endpoint
	app.Get("/", welcome)

	//User routes
	app.Post("/create/user", routes.CreateUser)
	app.Get("/get/user/:id", routes.GetUser)
	app.Put("/update/user/:id", routes.UpdateUser)
	app.Delete("/delete/user/:id", routes.DeleteUser)
	app.Get("/get/allNotesByUser/:id", routes.GetUserNotes)

	//Notes routes
	app.Post("/create/note", routes.CreateNote)
	app.Get("/get/note/:id", routes.GetNote)
	app.Put("/update/note/:id", routes.UpdateNote)
	app.Delete("/delete/note/:id", routes.DeleteNote)

	//Get title
	app.Post("/title", routes.GetTitle)
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
