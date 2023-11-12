package get

import (
	"context"
	"errors"
	"go-todo-api/database"
	"go-todo-api/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodo(ctx *gin.Context) {
  path := ctx.Param("id")

  if path == "" {
    ctx.IndentedJSON(http.StatusBadRequest, errors.New("id is required"))
    return
  }

  if len(path) != 24 {
    ctx.IndentedJSON(http.StatusBadRequest, errors.New("id is invalid"))
    return
  }

  id, err := primitive.ObjectIDFromHex(path)

  if err != nil {
    ctx.IndentedJSON(http.StatusBadRequest, err)
    return
  }

  filter := bson.D{{"_id", id}}

  result := database.Collection.FindOne(context.TODO(), filter)

  if err := result.Err(); err != nil {
    ctx.IndentedJSON(http.StatusNotFound, err)
    return
  }

  var todo types.Todo

  if err := result.Decode(&todo); err != nil {
    ctx.IndentedJSON(http.StatusNotFound, err)
    return
  }

  ctx.IndentedJSON(http.StatusOK, todo)

  return
}

func GetAllTodos(ctx *gin.Context) {
  cursor, err := database.Collection.Find(context.TODO(), bson.D{{}})

  if err != nil {
    ctx.IndentedJSON(http.StatusNotFound, err)
    return
  }

  var todos []types.Todo

  if err := cursor.All(context.TODO(), &todos); err != nil {
    ctx.IndentedJSON(http.StatusNotFound, err)
    return
  }

  if len(todos) == 0 {
    ctx.IndentedJSON(http.StatusNoContent, errors.New("no todos found"))
    return
  }

  ctx.IndentedJSON(http.StatusOK, todos)

  return
}