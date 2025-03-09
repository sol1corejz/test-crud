package storage

import (
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sol1corejz/test-crud/configs"
	"github.com/sol1corejz/test-crud/internal/models"
	"log"
	"time"
)

// Storage - интерфейс хранилища данных, предоставляющий методы для работы с пользователями и учетными данными.
// Используется для расширяемости и удобства тестирования.
type Storage interface {
	// CreateTask добавляет новую задачу
	CreateTask(t models.Task) error
	// GetTasks получает список задач
	GetTasks() error
	// UpdateTask обновляет задачу
	UpdateTask(id int, t models.Task) error
	// DeleteTask удаляет задачу
	DeleteTask(id int) error
}

// StorageImpl - реализация интерфейса Storage, использующая базу данных PostgreSQL.
type StorageImpl struct {
	DB *sql.DB
}

// DBStorage - глобальный объект для работы с базой данных.
var DBStorage StorageImpl

// ErrNotFound - ошибка, возвращаемая при отсутствии данных.
var ErrNotFound = errors.New("not found")

// ConnectDB устанавливает соединение с базой данных и создает необходимые таблицы.
func (storage StorageImpl) ConnectDB(cfg *configs.Config) error {
	if cfg.Storage.ConnectionString == "" {
		return errors.New("no connection string provided")
	}

	// Открываем соединение с базой данных PostgreSQL
	db, err := sql.Open("pgx", cfg.Storage.ConnectionString)
	if err != nil {
		log.Fatal(err)
		return err
	}

	storage.DB = db
	DBStorage.DB = db

	// Создаем таблицу tasks, если она отсутствует
	_, err = storage.DB.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
			created_at TIMESTAMP DEFAULT now(),
			updated_at TIMESTAMP DEFAULT now()
		)
	`)
	if err != nil {
		log.Fatal("failed to create tasks table:", err)
		return err
	}

	return nil

}

func (storage StorageImpl) CreateTask(t models.Task) error {
	_, err := storage.DB.Exec(`
		INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)
	`, t.Title, t.Description, t.Status)

	if err != nil {
		log.Print("failed to create task")
		return err
	}

	return nil
}
func (storage StorageImpl) GetTasks() ([]models.RawTask, error) {
	rows, err := storage.DB.Query(`
		SELECT * FROM tasks
	`)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Print("failed to find tasks")
			return nil, ErrNotFound
		}
		log.Print("failed to retrieve tasks", err.Error())
		return nil, err
	}
	defer rows.Close()

	tasks := make([]models.RawTask, 0)
	for rows.Next() {
		var task models.RawTask
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Print("failed to retrieve tasks", err.Error())
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		log.Print("failed to retrieve tasks", err.Error())
		return nil, err
	}

	return tasks, nil
}
func (storage StorageImpl) UpdateTask(id int, t models.Task) error {
	_, err := storage.DB.Exec(`
		UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5
	`, t.Title, t.Description, t.Status, time.Now(), id)

	if err != nil {
		log.Print("failed to save task", err.Error())
		return err
	}

	return nil
}
func (storage StorageImpl) DeleteTask(id int) error {
	_, err := storage.DB.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		log.Print("failed to delete task", err.Error())
		return err
	}
	return nil
}
