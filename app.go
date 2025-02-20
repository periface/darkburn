package main

import (
	"context"
	"darkburn/internals/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"darkburn/internals/models"
	"darkburn/internals/services"
	"github.com/atotto/clipboard"
)

const MAIN_DIR = "."

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

var appConfig models.Config

func build_app_config() {

	// read config.json file
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		appConfig = models.Config{}
		appConfig.Path = MAIN_DIR
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&appConfig)
	if err != nil {
		fmt.Println(err)
	}
	if appConfig.Path == "" {
		appConfig.Path = MAIN_DIR
	}
}

var database *sql.DB

func build_app_db() {
	database = db.Init_DB(true)
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	fmt.Println("App started")
	build_app_config()
	build_app_db()
	// test sql
	result, err := database.Exec("SELECT * FROM items")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	a.ctx = ctx
}
func search_in_folder(folder string, result *models.Result) {
	folder_files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range folder_files {
		if file.IsDir() {
			folder_path := filepath.Join(folder, file.Name())
			search_in_folder(folder_path, result)
		} else {
			file_info, err := file.Info()
			if err != nil {
				fmt.Println(err)
			}
			extension := strings.ToLower(filepath.Ext(file_info.Name()))
			file := models.FileList{
				Extension:    extension,
				Name:         file_info.Name(),
				AbsolutePath: filepath.Join(folder, file_info.Name()),
			}
			stored := result.Add_file_or_ignore(file)

			if stored {
				fmt.Println("File:", file.Name)
				fmt.Println("Stored:", stored)
				id, err := services.Store_File(file)

				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Stored file with id:", id)
			}
		}
	}
}

func (a *App) CopyToClipboard(text string) {

	file_data, err := os.ReadFile(text)
	if err != nil {
		fmt.Println(err)
	}
	clipboard.WriteAll(string(file_data))
	time.Sleep(time.Second)
}

func (a *App) OpenInExplorer(text string) {
	folder_path := filepath.Join(MAIN_DIR, text)
	absolute_path, err := filepath.Abs(folder_path)
	if err == nil {
	} else {
		fmt.Println(err)
	}
	cmd := exec.Command("explorer", "/select,", absolute_path)
	cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
func (a *App) GetFiles() []models.FileList {
	input := models.DataTable{}
	input.SetColumns([]string{"id", "extension", "name", "absolute_path", "created_at"})
	input.SetFilterColumns([]string{"name"})
	files, err := services.Get_Files(input)
	if err != nil {
		fmt.Println(err)
	}
	return files
}
func (a *App) StartApp() models.Result {
	result := models.Result{
		Files: []models.FileList{},
	}
	files, err := os.ReadDir(appConfig.Path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if file.IsDir() {
			folder_path := filepath.Join(appConfig.Path, file.Name())
			search_in_folder(folder_path, &result)
		} else {
			file_info, err := file.Info()
			if err != nil {
				fmt.Println(err)
			}
			extension := strings.ToLower(filepath.Ext(file_info.Name()))
			file := models.FileList{
				Extension:    extension,
				Name:         file_info.Name(),
				AbsolutePath: filepath.Join(MAIN_DIR, file_info.Name()),
				CreatedAt:    file_info.ModTime(),
			}
			stored := result.Add_file_or_ignore(file)
			fmt.Println("File:", file.Name)
			fmt.Println("Stored:", stored)

			if stored {
				id, err := services.Store_File(file)

				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Stored file with id:", id)
			}
		}
	}
	// order by created_at
	for i := 0; i < len(result.Files); i++ {
		for j := i + 1; j < len(result.Files); j++ {
			if result.Files[i].CreatedAt.Before(result.Files[j].CreatedAt) {
				result.Files[i], result.Files[j] = result.Files[j], result.Files[i]
			}
		}
	}
	return result
}
