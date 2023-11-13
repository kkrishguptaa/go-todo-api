package controllers

import (
	"go-todo-api/controllers/todos/del"
	"go-todo-api/controllers/todos/get"
	"go-todo-api/controllers/todos/patch"
	"go-todo-api/controllers/todos/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

var GetTodo = get.GetTodo
var GetAllTodos = get.GetAllTodos

var CreateTodo = post.CreateTodo

var PatchTodo = patch.PatchTodo

var DeleteTodo = del.DeleteTodo

func Default(routes gin.RoutesInfo) gin.HandlerFunc {
	return func(context *gin.Context) {
		var routesJSON []gin.H

		for _, route := range routes {
			routesJSON = append(routesJSON, gin.H{
				"method":  route.Method,
				"path":    route.Path,
				"handler": route.Handler,
			})
		}

		context.IndentedJSON(http.StatusOK,
			gin.H{"routes": routesJSON},
		)
	}
}
