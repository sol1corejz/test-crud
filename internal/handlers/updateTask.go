package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sol1corejz/test-crud/internal/models"
	"github.com/sol1corejz/test-crud/internal/storage"
	"log"
	"strconv"
)

func UpdateTask(c *fiber.Ctx) error {

	// Получение id задачи из параметров запроса
	ID := c.Params("id")
	parsedID, err := strconv.Atoi(ID)
	if err != nil {
		log.Print("failed to parse ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse ID",
		})
	}

	// Парсинг входных данных
	var taskPayload models.Task
	err = json.Unmarshal(c.Body(), &taskPayload)
	if err != nil {
		log.Print("error unmarshalling task payload")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "failed to parse payload data",
		})
	}

	// Обновление задачи
	err = storage.DBStorage.UpdateTask(parsedID, taskPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update task",
		})
	}

	// Отправка ответа
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "task updated",
	})
}
