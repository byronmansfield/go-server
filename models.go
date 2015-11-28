package main

type Todo struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Completed     bool   `json:"completed"`
	Created       []byte `json:"created"`
	Updated       []byte `json:"updated"`
	Due           []byte `json:"due"`
	DateCompleted []byte `json:"date_completed"`
}

type Todos []Todo

// apiError define structure of API error
type apiError struct {
	Tag     string `json:"-"`
	Error   error  `json:"-"`
	Message string `json:"error"`
	Code    int    `json:"code"`
}
