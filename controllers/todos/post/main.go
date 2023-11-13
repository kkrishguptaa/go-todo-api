package post

import (
	"context"
	"go-todo-api/database"
	"go-todo-api/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(ctx *gin.Context) {
  var todo types.TodoWithoutID

  if err := ctx.ShouldBindJSON(&todo); err != nil {
    ctx.IndentedJSON(http.StatusBadRequest, gin.H{
      "message": err.Error(),
    })
    return
  }

  result, err := database.Collection.InsertOne(context.TODO(), todo)

  if err != nil {
    ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
      "message": err.Error(),
    })
    return
  }

  ctx.IndentedJSON(http.StatusCreated, gin.H{
    "id": result.InsertedID,
  })

  return
}
