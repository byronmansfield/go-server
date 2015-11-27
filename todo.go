package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	var err error

	// connect to db
	db := dbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")

	todo := make([]Todo, 0)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		var id int
		var name string
		var completed bool
		var due time.Time

		if err := rows.Scan(&id, &name, &completed, &due); err != nil {
			panic(err)
		}
		todo = append(todo, Todo{Id: id, Name: name, Completed: completed, Due: due})
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	js, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	var err error

	// connect to db
	db := dbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")

	todo := make([]Todo, 0)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		var id int
		var name string
		var completed bool
		var due time.Time

		if err := rows.Scan(&id, &name, &completed, &due); err != nil {
			panic(err)
		}
		todo = append(todo, Todo{Id: id, Name: name, Completed: completed, Due: due})
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	js, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// connect to db
	db := dbConnect()
	defer db.Close()

	// Insert data to database
	stmt, err := db.Prepare("INSERT INTO tasks(name, completed, due) VALUES ($1, $2, $3)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(todo.Name, todo.Completed, todo.Due)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}

}
