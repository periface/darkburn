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
	if database == nil {
		database = db.Get_DB()
	} else {
		fmt.Println("Database already exists")
	}
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
			result.Add_file_or_ignore(file)
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
			println("Search in: ", folder_path)
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
			}
			result.Add_file_or_ignore(file)
		}
	}
	return result
}
