package database

import (
	"bytes"
	"database/sql"
	"github.com/byronmansfield/go-server/models"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type DataBase struct {
	Name     string
	Table    string
	Host     string
	Port     string
	User     string
	Password string
}

var dbinfo DataBase

// get database info from exported environment variables
func getDbInfo() {
	dbinfo.Name = os.Getenv("DB_NAME")
	dbinfo.Table = os.Getenv("DB_TABLE")
	dbinfo.Host = os.Getenv("DB_HOST")
	dbinfo.Port = os.Getenv("DB_PORT")
	dbinfo.User = os.Getenv("DB_USER")
	dbinfo.Password = os.Getenv("DB_PASSWORD")
}

// build the db connection as a string
func buildConnection(c DataBase) string {
	var buffer bytes.Buffer
	buffer.WriteString(c.User)
	buffer.WriteString(":")
	buffer.WriteString(c.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(c.Host)
	buffer.WriteString(":")
	buffer.WriteString(c.Port)
	buffer.WriteString(")/")
	buffer.WriteString(c.Name)
	return buffer.String()
}

// connect to database
func DbConnect() *sql.DB {

	getDbInfo()
	dbConnection := buildConnection(dbinfo)
	db, err := sql.Open("mysql", dbConnection)
	models.CheckErr(err)

	return db
}
