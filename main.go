package main

import (
	"github.com/gofiber/fiber/v2"
	"main.go/app"
	"main.go/configs"
	"main.go/repository"
	"main.go/services"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()

	dbClient := configs.Collection(configs.DB, "todos")
	TodoRepositoryDB := repository.NewTodoRepositoryDb(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDB)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen(":8080")
}
