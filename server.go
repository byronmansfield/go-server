package main

import (
	"github.com/byronmansfield/go-server/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	// Instantiate new router
	router := httprouter.New()

	// Get Todo Api instance
	ta := api.NewTodoApi()

	// index
	router.GET("/", ta.Index)

	// Todo resources
	router.GET("/todo", ta.GetTodos)

	router.GET("/todo/:id", ta.GetTodoById)

	router.POST("/todo", ta.CreateTodo)

	router.DELETE("/todo/:name", ta.DeleteTodo)

	// start server
	http.ListenAndServe(PORT, router)
}
