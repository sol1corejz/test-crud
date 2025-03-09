// Package models - пакет для описания моделей
package models

// Task - модель для описания чистых данных задачи, имеющих смысловую нагрузку
// Нужна для парсинга входных данных
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// RawTask - модель для полного описания данных задачи
// Нужна для парсинга данных с бд
type RawTask struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
