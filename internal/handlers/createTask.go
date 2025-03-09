package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sol1corejz/test-crud/internal/models"
	"github.com/sol1corejz/test-crud/internal/storage"
	"log"
)

func PostTasks(c *fiber.Ctx) error {

	// Парсинг входных данных
	var taskPayload models.Task
	err := json.Unmarshal(c.Body(), &taskPayload)
	if err != nil {
		log.Print("error unmarshalling task payload")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "failed to parse payload data",
		})
	}

	// Создание задачи
	err = storage.DBStorage.CreateTask(taskPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create task",
		})
	}

	// Отправка ответа
	return c.SendStatus(fiber.StatusCreated)
}
