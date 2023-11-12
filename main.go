package main

import (
	"context"
	"errors"
	"net/http"
	"os"

	"gitlab.com/xkrishguptaa/go-todo-api/-/tree/main/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

type Route struct {
}

var todos = []todo{
	{
		ID:        "1",
		Item:      "Learn Go",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Make a REST API",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Add GET, POST, PATCH",
		Completed: false,
	},
}

func getAllTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todos)
}

func getTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, todo := range todos {
		if todo.ID == id {
			ctx.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": errors.New("Todo not found").Error()})
	return
}

func createTodo(ctx *gin.Context) {
	var newTodo todo

	if err := ctx.BindJSON(&newTodo); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})

		return
	}

	todos = append(todos, newTodo)

	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func patchTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTodo struct {
		Item      string `json:"title"`
		Completed bool   `json:"completed"`
	}

	if err := ctx.BindJSON(&updatedTodo); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	var newTodo todo = todo{ID: id, Item: updatedTodo.Item, Completed: updatedTodo.Completed}

	for index, todo := range todos {
		if todo.ID == id {
			todos[index] = newTodo

			ctx.IndentedJSON(http.StatusOK, newTodo)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": errors.New("Todo not found").Error()})
	return
}

func deleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)

			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}
}

func defaultHandler(routes gin.RoutesInfo) gin.HandlerFunc {
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

func main() {
	util.init()
	godotenv.Load()

	router := gin.Default()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoUri := os.Getenv("DATABASE_URL")
	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	router.GET("/todos/all", getAllTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/:id", patchTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.GET("/", defaultHandler(router.Routes()))

	router.StaticFile("/favicon.ico", "./favicon.ico")

	router.Run()
}
