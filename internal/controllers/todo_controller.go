package controllers

import (
	"net/http"

	"github.com/dalmazox/todo-api/internal/dto"
	"github.com/dalmazox/todo-api/internal/services"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService *services.TodoService
}

func NewTodoController(todoService *services.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

// @Summary 		Create a new TODO
// @Description 	This endpoint creates a new TODO
// @Tags 			Todo
// @Accept 			json
// @Produces		json
// @Param 			request body dto.CreateTodoRequest true "Request Body"
// @Success 		201
// @Failure 		400 {object} HttpError
// @Failure 		500 {object} HttpError
// @Router 			/v1/todos [post]
func (t *TodoController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var request *dto.CreateTodoRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := t.todoService.CreateTodo(ctx, request.Description, request.Done); err != nil {
		NewError(c, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusCreated)
}
