package main

import (
	"fmt"
	"github.com/KarlvenK/korm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := korm.NewEngine("sqlite3", "korm.db")
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS use;").Exec()
	_, _ = s.Raw("CREATE TABLE user(name text);").Exec()
	_, _ = s.Raw("CREATE TABLE user(name text);").Exec()
	result, _ := s.Raw("INSERT INTO user(`name`) values(?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
