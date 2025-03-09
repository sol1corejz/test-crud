package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sol1corejz/test-crud/configs"
	"github.com/sol1corejz/test-crud/internal/handlers"
	"github.com/sol1corejz/test-crud/internal/storage"
	"log"
)

func main() {

	// Загрузка конфига
	config, err := configs.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Print("Ошибка загрузки конфига:", err)
		return
	}

	// Подключение к бд
	err = storage.DBStorage.ConnectDB(config)
	if err != nil {
		log.Fatal(err)
	}

	// Создание сервера
	app := fiber.New()

	// Регистрация маршрутов
	app.Get("/tasks", handlers.GetTasks)
	app.Post("/tasks", handlers.PostTasks)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	// Запуск сервера
	err = app.Listen(config.Server.Address)
	if err != nil {
		log.Fatal(err)
	}
}
