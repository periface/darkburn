package models

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type FileList struct {
	Id           int
	Extension    string
	Name         string
	AbsolutePath string
	CreatedAt    time.Time
}

type Config struct {
	Path string
}

type Result struct {
	Files []FileList
}

func starts_with_dot(name string) bool {
	return strings.HasPrefix(name, ".")
}

var accepted_extensions = []string{".svg", ".dxf"}

func (r *Result) Add_file_or_ignore(file FileList) bool {
	if starts_with_dot(file.Name) {
		return false
	}
	for _, ext := range accepted_extensions {
		if file.Extension == ext {
			r.Files = append(r.Files, file)
			return true
		}
		return false
	}
	return false
}

type DataTable struct {
	Page          int
	PageSize      int
	Total         int
	Sort          string
	Filter        string
	SortColumn    string
	FilterColumn  string
	FilterValue   string
	tableName     string
	columns       []string
	filterColumns []string
}

// SetTableName sets the table name in DB for the query
func (tableInput *DataTable) SetTableName(tableName string) {
	tableInput.tableName = tableName
}

// SetColumns sets the columns to be queried
func (tableInput *DataTable) SetColumns(columns []string) {
	tableInput.columns = columns
}

// SetFilterColumns sets the columns to be filtered
func (tableInput *DataTable) SetFilterColumns(filterColumns []string) {
	tableInput.filterColumns = filterColumns
}
func (tableInput *DataTable) GetTableQuery() (string, error) {
	if tableInput.tableName == "" {
		return "", errors.New("Table name is required")
	}
	if len(tableInput.columns) == 0 {
		return "", errors.New("Columns are required")
	}
	if len(tableInput.filterColumns) == 0 {
		return "", errors.New("Filter columns are required")
	}
	return tableInput.buildTableQuery(), nil
}

func (tableInput *DataTable) buildTableQuery() string {
	pageSize := strconv.Itoa(tableInput.PageSize)
	page := strconv.Itoa(tableInput.Page) // page is 1 based, but offset is 0 based

	query := "SELECT "
	for i, column := range tableInput.columns {
		if i == 0 {
			query += column
		} else {
			query += ", " + column
		}
	}
	query += " FROM " + tableInput.tableName
	if tableInput.FilterColumn != "" && tableInput.FilterValue != "" {
		query += " WHERE " + tableInput.FilterColumn + " = '" + tableInput.FilterValue + "'"
		if tableInput.Filter != "" {
			for i, column := range tableInput.filterColumns {
				if i == 0 {
					query += " AND (" + column + " LIKE '%" + tableInput.Filter + "%'"
				} else {
					query += " OR " + column + " LIKE '%" + tableInput.Filter + "%'"
				}
			}

			query += ")"
		}
	} else if tableInput.Filter != "" {
		for i, column := range tableInput.filterColumns {
			if i == 0 {
				query += " WHERE (" + column + " LIKE '%" + tableInput.Filter + "%'"
			} else {
				query += " OR " + column + " LIKE '%" + tableInput.Filter + "%'"
			}
		}
		query += ")"
	}
	if tableInput.SortColumn != "" {
		query += " ORDER BY " + tableInput.SortColumn + " " + tableInput.Sort
	}
	if pageSize != "" {
		query += " LIMIT " + pageSize
	}
	if page != "" {
		page, _ := strconv.Atoi(page)
		pageSize, _ := strconv.Atoi(pageSize)
		offset := page * pageSize
		query += " OFFSET " + strconv.Itoa(offset)
	}
	return query
}

func (tableInput *DataTable) GetCountQuery() (string, error) {
	return tableInput.countQuery()
}

func (tableInput *DataTable) countQuery() (string, error) {
	query := "SELECT COUNT(*) FROM " + tableInput.tableName

	if tableInput.tableName == "" {
		return "", errors.New("Table name is required")
	}
	if len(tableInput.columns) == 0 {
		return "", errors.New("Columns are required")
	}
	if len(tableInput.filterColumns) == 0 {
		return "", errors.New("Filter columns are required")
	}
	if tableInput.FilterColumn != "" && tableInput.FilterValue != "" {
		query += " WHERE " + tableInput.FilterColumn + " = '" + tableInput.FilterValue + "'"
		if tableInput.Filter != "" {
			for i, column := range tableInput.filterColumns {
				if i == 0 {
					query += " AND (" + column + " LIKE '%" + tableInput.Filter + "%'"
				} else {
					query += " OR " + column + " LIKE '%" + tableInput.Filter + "%'"
				}
			}

			query += ")"
		}
	} else if tableInput.Filter != "" {
		for i, column := range tableInput.filterColumns {
			if i == 0 {
				query += " WHERE (" + column + " LIKE '%" + tableInput.Filter + "%'"
			} else {
				query += " OR " + column + " LIKE '%" + tableInput.Filter + "%'"
			}
		}
		query += ")"
	}
	return query, nil
}

type TableOutput[T any] struct {
	Rows     []T `json:"rows"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}
