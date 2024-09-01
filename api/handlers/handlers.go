package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/RedrikShuhartRed/TaskManager/api/service"
	"github.com/RedrikShuhartRed/TaskManager/models"
)

var (
	errTaskNotFound = errors.New("Task not found, check ID")
)

// Handler handles HTTP requests for tasks.
type Handler struct {
	service *service.TaskService
}

// NewHandler create new copy of Hadler struct.
func NewHandler(service *service.TaskService) *Handler {
	return &Handler{
		service: service,
	}
}

func handleError(c *gin.Context, statusCode int, err error, logMessage string) {
	zap.S().Errorf(logMessage, err)
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

// AddNewTask adds a new task to the database.
func (h *Handler) AddNewTask(c *gin.Context) {

	var input models.Input

	err := c.ShouldBindJSON(&input)
	if err != nil {
		handleError(c, http.StatusBadRequest, err, "error decoding request body: %v")
		return
	}
	task, err := h.service.AddNewTask(c.Request.Context(), input.Title, input.Description, input.DueDate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err, "error insert into verbatasks: %v")
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, task)
}

// GetAllTasks retrieves all tasks from the database.
func (h *Handler) GetAllTasks(c *gin.Context) {

	tasks, err := h.service.GetAllTasks(c.Request.Context())
	if err != nil {
		handleError(c, http.StatusInternalServerError, err, "error iget tasks: %v")
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a task from the database by its ID.
func (h *Handler) GetTaskByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err, "error get id task, id not ints: %v")
		return
	}

	task, err := h.service.GetTaskByID(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(c, http.StatusNotFound, errTaskNotFound, "error getting task, task not found: %v")
			return
		}
		handleError(c, http.StatusInternalServerError, err, "error getting task: %v")
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, task)

}

// DeleteTask delete a task from the database by its ID.
func (h *Handler) DeleteTask(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err, "error get id task, id not int: %v")
		return
	}

	err = h.service.DeleteTask(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(c, http.StatusNotFound, errTaskNotFound, "error getting task, task not found: %v")
			return
		}
		handleError(c, http.StatusInternalServerError, err, "error delete tasks: %v")
		return
	}
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusNoContent)

}

// UpdateTaskByID update a task from the database by its ID.
func (h *Handler) UpdateTaskByID(c *gin.Context) {

	var input models.Input

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err, "error get id task, id not int: %v")
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		handleError(c, http.StatusBadRequest, err, "error decoding request body: %v")
		return
	}

	task, err := h.service.UpdateTaskByID(c.Request.Context(), id, input.Title, input.Description, input.DueDate)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(c, http.StatusNotFound, errTaskNotFound, "error getting task, task not found: %v")
			return
		}
		handleError(c, http.StatusInternalServerError, err, "error getting task: %v")
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, task)
}
