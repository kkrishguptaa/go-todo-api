package router

import (
	"go-todo-api/controllers"

	"github.com/gin-gonic/gin"
)

func Init() {
  r := gin.Default()

  r.GET("/todos", controllers.GetAllTodos)
  r.GET("/todos/:id", controllers.GetTodo)
  r.POST("/todos", controllers.CreateTodo)
  r.PATCH("/todos/:id", controllers.PatchTodo)
  r.DELETE("/todos/:id", controllers.DeleteTodo)

  r.GET("/", controllers.Default(r.Routes()))
  r.Static("/favicon.ico", "../assets/favicon.ico")

  r.Run()
}
