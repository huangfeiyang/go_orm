package main

import (
	"database/sql"
	"log"
	// _"github.com/mattn/go-sqlite3"
	_"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = "root"
	ip = "127.0.0.1"
	dbName = "local"
	port = "3306"
)

var DB *sql.DB

func main() {
	dbPath := strings.Join([]string{userName, ":", password, "@tcp(",ip,":",port,")/", dbName, "?charset=utf8"}, "")

	DB, dbErr :=sql.Open("mysql", dbPath)
	if dbErr != nil {
		log.Println("err")
		log.Println(dbErr)
	}
	log.Println(DB)
	// db, _ := sql.Open("sqlite3", "gee.db")
	// defer func() { _ = db.Close() }()

	// _, _ = db.Exec("DROP TABLE IF EXISTS User;")
	// _, _ = db.Exec("CREATE TABLE User(Name text);")

	// result, err := db.Exec("INSERT INTO User(`Name`) VALUES (?), (?)", "Tom", "Sam")

	// if err != nil {
	// 	affected, _ := result.RowsAffected()
	// 	log.Println(affected)
	// }

	// row := db.QueryRow("SELECT Name FROM User LIMIT 1")

	// var name string
	// if err := row.Scan(&name); err == nil {
	// 	log.Println(name)
	// }
}