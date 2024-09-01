package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/RedrikShuhartRed/TaskManager/models"
)

// TaskRepository provides methods to perform CRUD operations on the tasks table in the database.
type TaskRepository struct {
	DB *sqlx.DB
}

// NewTaskRepository creates and returns a new instance of TaskRepository with the provided database connection.
func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

// AddNewTask inserts a new task into the verbatasks table in the database.
func (r *TaskRepository) AddNewTask(ctx context.Context, task *models.Task) error {
	query := `INSERT INTO tasks (title, description, due_date, created_at,updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRow(query, task.Title, task.Description, task.DueDate, task.CreatedAt, task.UpdatedAt).Scan(&task.ID)
	return err
}

// GetAllTasks retrieves all tasks from the verbatasks table in the database.
func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Select(&tasks, "SELECT * FROM tasks")
	return tasks, err
}

// GetTaskByID retrieves a task by its ID from the verbatasks table in the database.
func (r *TaskRepository) GetTaskByID(ctx context.Context, id int) (*models.Task, error) {
	var task models.Task
	err := r.DB.Get(&task, "SELECT * FROM tasks WHERE id=$1", id)
	return &task, err
}

// UpdateTaskByID updates an existing task in the verbatasks table in the database.
func (r *TaskRepository) UpdateTaskByID(ctx context.Context, task *models.Task) error {
	query := `UPDATE tasks SET title=$1, description=$2, due_date=$3, updated_at=$4 WHERE id=$5`
	_, err := r.DB.Exec(query, task.Title, task.Description, task.DueDate, task.UpdatedAt, task.ID)
	return err
}

// DeleteTask deletes a task from the tasks table in the database by its ID.
func (r *TaskRepository) DeleteTask(ctx context.Context, id int) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id=$1", id)
	return err
}
