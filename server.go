package main

import "net/http"

func main() {
	router := NewRouter()
	http.ListenAndServe(PORT, router)
}
