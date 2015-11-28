package api

import (
	"encoding/json"
	"github.com/byronmansfield/go-server/database"
	"github.com/byronmansfield/go-server/helpers"
	"github.com/byronmansfield/go-server/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TodoApi struct{}

func NewTodoApi() *TodoApi {
	return &TodoApi{}
}

// base todo endpoint
// returns array of todos
func (ta TodoApi) TodoIndex(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error

	// connect to db
	db := database.DbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")

	todo := make([]models.Todo, 0)

	helpers.CheckErr(err)

	for rows.Next() {

		var id int
		var name string
		var description string
		var completed bool
		var created []byte
		var updated []byte
		var due []byte
		var date_completed []byte

		if err := rows.Scan(&id, &name, &description, &completed, &created, &updated, &due, &date_completed); err != nil {
			panic(err)
		}
		todo = append(todo, models.Todo{
			Id:            id,
			Name:          name,
			Description:   description,
			Completed:     completed,
			Created:       created,
			Updated:       updated,
			Due:           due,
			DateCompleted: date_completed,
		})
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

// Get a todo by name
func (ta TodoApi) GetTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error

	// connect to db
	db := database.DbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT name FROM tasks where name = ?", p.ByName("name"))

	todo := make([]models.Todo, 0)

	helpers.CheckErr(err)

	for rows.Next() {

		var id int
		var name string
		var description string
		var completed bool
		var created []byte
		var updated []byte
		var due []byte
		var date_completed []byte

		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		todo = append(todo, models.Todo{
			Id:            id,
			Name:          name,
			Description:   description,
			Completed:     completed,
			Created:       created,
			Updated:       updated,
			Due:           due,
			DateCompleted: date_completed,
		})
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

// create a todo
func (ta TodoApi) CreateTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// stub a new todo
	todo := models.Todo{}

	// populate the todo with the data from post req
	json.NewDecoder(r.Body).Decode(&todo)

	// connect to db
	db := database.DbConnect()
	defer db.Close()

	// Insert data to database
	stmt, err := db.Prepare("INSERT INTO tasks(name, description, due) VALUES ($1, $2, $3)")
	helpers.CheckErr(err)

	_, err = stmt.Exec(todo.Name, todo.Description, todo.Due)
	helpers.CheckErr(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

}

// delete a todo
func (ta TodoApi) DeleteTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// connect to db
	db := database.DbConnect()
	defer db.Close()

	// Delete todo from database
	stmt, err := db.Prepare("DELETE FROM tasks where name=?")
	helpers.CheckErr(err)

	_, err = stmt.Exec(p.ByName("name"))
	helpers.CheckErr(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
