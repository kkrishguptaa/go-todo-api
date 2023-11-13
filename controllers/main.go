package controllers

import (
	delete "go-todo-api/controllers/todos/DELETE"
	get "go-todo-api/controllers/todos/GET"
	patch "go-todo-api/controllers/todos/PATCH"
	post "go-todo-api/controllers/todos/POST"
	"net/http"

	"github.com/gin-gonic/gin"
)

var GetTodo = get.GetTodo
var GetAllTodos = get.GetAllTodos

var CreateTodo = post.CreateTodo

var PatchTodo = patch.PatchTodo

var DeleteTodo = delete.DeleteTodo

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
