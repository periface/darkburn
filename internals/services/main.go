package services

import (
	"darkburn/internals/db"
	"darkburn/internals/models"
	"fmt"
)

func Store_File(file models.FileList) (int64, error) {
	db := db.Get_DB(false)
	fmt.Println("Store File")

	get_items_query := `
    INSERT INTO items (extension, name, absolute_path, created_at)
    VALUES (?, ?, ?, ?)
    `
	result, err := db.Exec(get_items_query, file.Extension, file.Name, file.AbsolutePath, file.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return result.LastInsertId()

}
func Get_Files(input models.DataTable) ([]models.FileList, error) {
	db := db.Get_DB(false)
	query, err := input.GetTableQuery()
	get_items_query := `
    SELECT * FROM items
    `
	rows, err := db.Query(get_items_query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var files []models.FileList
	for rows.Next() {
		var file models.FileList
		err := rows.Scan(&file.Id, &file.Extension, &file.Name, &file.AbsolutePath, &file.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		files = append(files, file)
	}
	println("====================================")
	fmt.Printf("Files: %v", files)
	println("====================================")
	return files, nil

}
