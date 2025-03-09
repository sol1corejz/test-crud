package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sol1corejz/test-crud/internal/storage"
	"log"
	"strconv"
)

// DeleteTask - функция для удаления задачи
func DeleteTask(c *fiber.Ctx) error {

	// Получение id задачи из параметров запроса
	ID := c.Params("id")
	parsedID, err := strconv.Atoi(ID)
	if err != nil {
		log.Print("failed to parse ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse ID",
		})
	}

	// Удаление задачи
	err = storage.DBStorage.DeleteTask(parsedID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot delete task",
		})
	}

	// Отправка ответа
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "task deleted",
	})
}
