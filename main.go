package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-todo-apps/configs"
	"go-todo-apps/controllers"
	"log"
)

func main() {
	//Initiate Database
	configs.InitiateDB()
	config := fiber.Config{
		ServerHeader: "Masihkasar Server",
	}
	app := fiber.New(config)

	app.Get("/api/v1/todo", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("🖕 %s", "welcome to simple ToDo Apps")

		response := fiber.Map{
			"code":    100,
			"message": msg,
			"state":   true,
		}

		return c.Status(200).JSON(response)
	})

	app.Get("/api/v1/todo/get-all", controllers.GetAll)
	app.Get("/api/v1/todo/single/:id", controllers.GetSingle)
	app.Post("/api/v1/todo/create", controllers.Create)
	app.Post("/api/v1/todo/update/:id", controllers.Update)
	app.Get("/api/v1/todo/delete/:id", controllers.Delete)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    400,
			"message": "Route not Found",
			"state":   false,
		})
	})

	//PORT := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	//fmt.Println("Server running on PORT:", os.Getenv("SERVER_PORT"))
	log.Fatal(app.Listen(":3000"))
}
