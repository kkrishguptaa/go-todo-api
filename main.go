package main

import (
 	"errors"
  "os"
  "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string    `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
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

func getAllTodos(context *gin.Context) {
  context.IndentedJSON(http.StatusOK, todos)
}


func getTodo(context *gin.Context) {
  id := context.Param("id")

  for _, todo := range todos {
    if todo.ID == id {
      context.IndentedJSON(http.StatusOK, todo)
      return
    }
  }

  context.IndentedJSON(http.StatusNotFound, gin.H{"error": errors.New("Todo not found").Error()})
  return
}

func createTodo (context *gin.Context) {
  var newTodo todo

  if err := context.BindJSON(&newTodo); err != nil {
    context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})

    return
  }

  todos = append(todos, newTodo)

  context.IndentedJSON(http.StatusCreated, newTodo)
}

func patchTodo (context *gin.Context) {
  id := context.Param("id")

  var updatedTodo struct {
    Item string `json:"title"`
    Completed bool `json:"completed"`
  }

  if err := context.BindJSON(&updatedTodo); err != nil {
    context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
    return
  }

  var newTodo todo = todo{ID: id, Item: updatedTodo.Item, Completed: updatedTodo.Completed}

  for index, todo := range todos {
    if todo.ID == id {
      todos[index] = newTodo

      context.IndentedJSON(http.StatusOK, newTodo)
      return
    }
  }

  context.IndentedJSON(http.StatusNotFound, gin.H{"error": errors.New("Todo not found").Error()})
  return
}

func deleteTodo (context *gin.Context) {
  id := context.Param("id")

  for index, todo := range todos {
    if todo.ID == id {
      todos = append(todos[:index], todos[index+1:]...)

      context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
      return
    }
  }
}

func main() {
	router := gin.Default()

	router.GET("/todos/all", getAllTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/:id", patchTodo)
	router.DELETE("/todos/:id", deleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	URL := fmt.Sprintf("localhost:%s", port)

	router.Run(URL)
}
