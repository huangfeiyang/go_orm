package main

import (
	"fmt"
	"geeorm"
	// "geeorm/log"

	_"github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "~/mydata/gee.db")
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
    _, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
    _, _ = s.Raw("CREATE TABLE User(Name text);").Exec()

	result, _ := s.Raw("INSERT INTO User(`name`) VALUES (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}