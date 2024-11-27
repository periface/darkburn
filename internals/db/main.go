package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func create_items_table(db *sql.DB) {
	create_table := `
    CREATE TABLE items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        extension TEXT,
        name TEXT,
        absolute_path TEXT
    );
    `
	_, err := db.Exec(create_table)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Table created successfully..")
}
func create_dir() {
	os.MkdirAll("./darkb_db", os.ModePerm)
}
func create_config_table(db *sql.DB) {
	create_table := `
    CREATE TABLE config (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        first_run INTEGER
    );
    `
	_, err := db.Exec(create_table)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Table created successfully..")
}
func Get_DB() *sql.DB {
	create_dir()
	db, err := sql.Open("sqlite3", "./darkb_db/darkburn.db")
	if err != nil {
		fmt.Println(err)
	}
	create_config_table(db)
	create_items_table(db)
	return db
}
