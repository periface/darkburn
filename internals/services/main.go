package services

import (
	"darkburn/internals/db"
	"darkburn/internals/models"
	"fmt"
)

func Store_File(file models.FileList) (int64, error) {
	db := db.Get_DB(false)
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
func Get_Files(filter string) ([]models.FileList, error) {
	db := db.Get_DB(false)
	get_items_query := `
    SELECT * FROM items
    `
	if filter != "" {
		get_items_query += " WHERE name LIKE '%" + filter + "%'"
	}
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
	return files, nil

}
