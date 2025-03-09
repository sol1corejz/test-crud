// Package handlers - пакет содержащий реализациую обработчиков маршрутов
package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sol1corejz/test-crud/internal/storage"
)

// GetTasks - функция получения задач
func GetTasks(c *fiber.Ctx) error {

	// Получение задач
	tasks, err := storage.DBStorage.GetTasks()
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "tasks not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch tasks",
		})
	}

	// Отправка ответа
	return c.JSON(fiber.Map{
		"tasks": tasks,
	})
}
