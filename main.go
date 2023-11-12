package main

import (
	"go-todo-api/database"
	"go-todo-api/env"
	"go-todo-api/router"
)

func main() {
	env.Init()
	database.Init()
	router.Init()
}
