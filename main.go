package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/omatheu/go-and-sqlite/database"
	_ "github.com/omatheu/go-and-sqlite/database"
	"github.com/omatheu/go-and-sqlite/handlers"
	_ "github.com/omatheu/go-and-sqlite/handlers"
)

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
	err = db.AutoMigrate(&handlers.User{})
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/users", handlers.CreateUser(db))
	app.Get("/users/:id", handlers.GetUserByID(db))
	app.Put("/users/:id", handlers.UpdateUser(db))
	app.Delete("/users/:id", handlers.DeleteUser(db))

	log.Fatal(app.Listen(":8080"))
}
