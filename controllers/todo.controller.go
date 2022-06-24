package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-todo-apps/configs"
	"go-todo-apps/dto"
	"go-todo-apps/models"
)

func CreateResponseTodo(TodoModel models.Todo) dto.TodoResponse {
	return dto.TodoResponse{
		ID:        TodoModel.ID,
		Name:      TodoModel.Name,
		Check:     TodoModel.Check,
		CreatedAt: TodoModel.CreatedAt,
	}
}

func SerializeResponseSingle(data models.Todo, code int, message string, state bool) fiber.Map {
	return fiber.Map{
		"code":    code,
		"state":   state,
		"message": message,
		"data":    data,
	}
}

func SerializeResponseMany(data []dto.TodoResponse, code int, message string, state bool) fiber.Map {
	return fiber.Map{
		"code":    code,
		"state":   state,
		"message": message,
		"data":    data,
	}
}

func ResponseError(err fiber.Error) fiber.Map {
	return fiber.Map{
		"code":    400,
		"state":   false,
		"message": "Something Error",
		"error":   err,
	}
}

func GetAll(c *fiber.Ctx) error {
	todos := []models.Todo{}

	configs.DB.Find(&todos)

	responseTodos := []dto.TodoResponse{}

	for _, todo := range todos {
		responseTodo := CreateResponseTodo(todo) //Looping to serialize single data to function
		responseTodos = append(responseTodos, responseTodo)
	}

	responses := SerializeResponseMany(responseTodos, 200, "Data Found", true)

	return c.Status(200).JSON(responses)
}

func GetSingle(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if id == 0 {
		return c.Status(400).JSON(ResponseError(fiber.Error{
			Message: "Invalid ID",
			Code:    404,
		}))
	}

	var todo models.Todo
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    404,
			"message": "Data Not Found",
			"state":   false,
		})
	}

	configs.DB.Find(&todo, "id = ?", id)
	responses := SerializeResponseSingle(todo, 200, "Data Found", true)

	return c.Status(200).JSON(responses)
}

func Create(c *fiber.Ctx) error {
	//Get Table value
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	configs.DB.Create(&todo)

	responses := SerializeResponseSingle(todo, 200, "Data Found", true)

	return c.Status(200).JSON(responses)
}

func Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var todo models.Todo

	if id == 0 {
		return c.Status(400).JSON(ResponseError(fiber.Error{
			Message: "Invalid ID",
			Code:    404,
		}))
	}

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    404,
			"message": "Data Not Found",
			"state":   false,
		})
	}

	configs.DB.Find(&todo, "id = ?", id)

	var updatedTodo dto.Todo

	if err := c.BodyParser(&updatedTodo); err != nil {
		return c.Status(400).JSON(ResponseError(fiber.Error{
			Message: err.Error(),
		}))
	}

	todo.Name = updatedTodo.Name
	todo.Check = updatedTodo.Check

	configs.DB.Save(&todo)

	return c.Status(200).JSON(SerializeResponseSingle(todo, 200, "Success Save Data", true))

}

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var todo models.Todo

	if id == 0 {
		return c.Status(400).JSON(ResponseError(fiber.Error{
			Message: "Invalid ID",
			Code:    404,
		}))
	}

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    404,
			"message": "Data Not Found",
			"state":   false,
		})
	}

	configs.DB.Find(&todo, "id = ?", id)

	if err := configs.DB.Delete(&todo); err != nil {
		return c.Status(400).JSON(ResponseError(fiber.Error{
			Message: "Failed to Delete Data",
			Code:    404,
		}))
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"state":   true,
		"message": fmt.Sprintf("Success Data with ID: %s", id),
	})
}
