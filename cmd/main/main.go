package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	// Миграация
	err = migrateDB()
	if err != nil {
		log.Fatal(err)
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

func migrateDB() error {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:12345678@localhost:5432/tasks?sslmode=disable",
	)
	if err != nil {
		return err
	}

	// Применение всех миграций
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Миграции успешно применены!")
	return nil
}
