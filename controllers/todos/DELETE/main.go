package delete

import (
	"context"
	"errors"
	"go-todo-api/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTodo(ctx *gin.Context) {
 	path := ctx.Param("id")

	if path == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
		"message": errors.New("id is required").Error(),
		})
		return
	}

	if len(path) != 24 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
		"message": "id is invalid",
		})
		return
	}

	id, err := primitive.ObjectIDFromHex(path)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id " + err.Error(),
		})
		return
	}

	filter := bson.M{"_id": id}

  result := database.Collection.FindOneAndDelete(context.TODO(), filter)

  if err := result.Err(); err != nil {
    ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
      "message": err.Error(),
    })
    return
  }

  ctx.IndentedJSON(http.StatusOK, gin.H{
    "message": "Todo deleted successfully",
  })

  return
}
