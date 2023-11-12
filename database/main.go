package database

import (
	"context"
	"go-todo-api/env"
	"go-todo-api/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var Database *mongo.Database
var Collection *mongo.Collection

func Init() {
  getClient()
  getDatabase()
  getCollection()
}

func getClient () {
  options.ServerAPI(options.ServerAPIVersion1)
  bsonOpts := &options.BSONOptions {
    UseJSONStructTags: true,
    OmitZeroStruct: true,
  }

	client, err := mongo.Connect(
	context.TODO(),
	options.Client().
	  SetBSONOptions(bsonOpts).
	  ApplyURI(env.Get("DATABASE_URL")),
	)

	if err != nil {
		panic(err)
	}

	println("connected to cluster")

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
     panic(err)
  }

  println("pinged cluster")

	Client = client
}

func getDatabase() {
	db := Client.Database(env.Get("DATABASE_NAME"))

	println("connected to database " + env.Get("DATABASE_NAME"))

	Database = db
}

func getCollection() {
  collection := Database.Collection("Todos")

  println("connected to collection Todos")

  Collection = collection
}

func ParseDatabaseTodo(r *bson.M) types.Todo {
  return types.Todo{
    ID: (*r)["_id"].(primitive.ObjectID).Hex(),
    Item: (*r)["title"].(string),
    Completed: (*r)["completed"].(bool),
  }
}

func ParseDatabaseTodos(r *[]bson.M) []types.Todo {
  var todos []types.Todo

  for _, result := range *r {
    todos = append(todos, ParseDatabaseTodo(&result))
  }

  return todos
}
