package service

import (
	"context"
	"time"

	"github.com/RedrikShuhartRed/TaskManager/api/repository"
	"github.com/RedrikShuhartRed/TaskManager/models"
)

// TaskService is a service layer for solving management tasks.
type TaskService struct {
	repo *repository.TaskRepository
}

// NewTaskService creates a new instance of TaskService.
func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// AddNewTask adds a new task to the database.
func (s *TaskService) AddNewTask(ctx context.Context, title, description, dueDate string) (*models.Task, error) {
	due, err := time.Parse(time.RFC3339, dueDate)
	if err != nil {
		return nil, err
	}

	task := &models.Task{
		Title:       title,
		Description: description,
		DueDate:     due,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err = s.repo.AddNewTask(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetAllTasks retrieves all tasks from the database.
func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetAllTasks(ctx)
}

// GetTaskByID retrieves a task by its ID.
func (s *TaskService) GetTaskByID(ctx context.Context, id int) (*models.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}

// UpdateTaskByID updates a task by its ID.
func (s *TaskService) UpdateTaskByID(ctx context.Context, id int, title, description, dueDate string) (*models.Task, error) {
	task, err := s.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	due, err := time.Parse(time.RFC3339, dueDate)
	if err != nil {
		return nil, err
	}

	task.Title = title
	task.Description = description
	task.DueDate = due
	task.UpdatedAt = time.Now()

	err = s.repo.UpdateTaskByID(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask deletes a task by its ID from the database.
func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.repo.DeleteTask(ctx, id)
}
