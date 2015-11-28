package main

import (
	"fmt"
	"github.com/byronmansfield/go-server/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// base endpoint
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome from Go-Server!")
}

func main() {

	// Instantiate new router
	router := httprouter.New()

	// Get Todo Api instance
	ta := api.NewTodoApi()

	// Todo resources
	router.GET("/todo", ta.TodoIndex)

	router.GET("/todo/:name", ta.GetTodo)

	router.POST("/user", ta.CreateTodo)

	router.DELETE("/user/:name", ta.DeleteTodo)

	// start server
	http.ListenAndServe(PORT, router)
}
