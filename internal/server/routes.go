package server

import (
	"github.com/dalmazox/todo-api/internal/controllers"
	"github.com/dalmazox/todo-api/internal/db"
	"github.com/dalmazox/todo-api/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	todoService := services.NewTodoService(db.NewRepository())
	todoController := controllers.NewTodoController(todoService)

	r.POST("/v1/todos", todoController.Create)
}
