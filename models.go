package main

import "time"

type Todo struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"name"`
	Completed   bool      `json:"completed"`
	Due         time.Time `json:"due"`
}

type Todos []Todo

// apiError define structure of API error
type apiError struct {
	Tag     string `json:"-"`
	Error   error  `json:"-"`
	Message string `json:"error"`
	Code    int    `json:"code"`
}
