package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/omatheu/go-and-sqlite/database"
	_ "github.com/omatheu/go-and-sqlite/database"
	_ "github.com/omatheu/go-and-sqlite/docs"
	"github.com/omatheu/go-and-sqlite/handlers"
	"github.com/omatheu/go-and-sqlite/models"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	db, err := database.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	// Get the underlying sql.DB object and defer its closure
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Perform database migration
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/users", handlers.CreateUser(db))
	app.Get("/users/:id", handlers.GetUserByID(db))
	app.Put("/users/:id", handlers.UpdateUser(db))
	app.Delete("/users/:id", handlers.DeleteUser(db))

	log.Fatal(app.Listen(":8080"))
}
