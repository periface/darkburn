package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func table_exists(db *sql.DB, table_name string) bool {
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", table_name)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	return rows.Next()
}
func create_items_table(db *sql.DB) {
	if table_exists(db, "items") {
		return
	}
	create_table := `
    CREATE TABLE items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        extension TEXT,
        name TEXT,
        absolute_path TEXT,
        created_at DATETIME
    );
    `
	_, err := db.Exec(create_table)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Table created successfully..")
}
func create_dir() {
	if _, err := os.Stat("./darkb_db"); os.IsNotExist(err) {
		os.MkdirAll("./darkb_db", os.ModePerm)
	}
}
func create_config_table(db *sql.DB) {
	if table_exists(db, "config") {
		return
	}
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
func Init_DB(reset_table bool) *sql.DB {
	if reset_table {
		os.Remove("./darkb_db/darkburn.db")
	}
	create_dir()
	db, err := sql.Open("sqlite3", "./darkb_db/darkburn.db")
	if err != nil {
		fmt.Println(err)
	}
	create_config_table(db)
	create_items_table(db)
	return db
}
func Get_DB(reset_table bool) *sql.DB {
	if database == nil {
		database = Init_DB(reset_table)
	}
	return database
}
