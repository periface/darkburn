package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type FileList struct {
	Extension    string
	Name         string
	AbsolutePath string
}

type Config struct {
	Path string
}

type Result struct {
	Files []FileList
}

const MAIN_DIR = "."

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

var appConfig Config

func build_app_config() {

	// read config.json file
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		appConfig = Config{}
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

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	fmt.Println("App started")
	build_app_config()
	a.ctx = ctx
}
func search_in_folder(folder string, result *Result) {
	folder_files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range folder_files {
		if file.IsDir() {
			folder_path := filepath.Join(folder, file.Name())
			println("Search in: ", folder_path)
			search_in_folder(folder_path, result)
		} else {
			file_info, err := file.Info()
			if err != nil {
				fmt.Println(err)
			}
			extension := strings.ToLower(filepath.Ext(file_info.Name()))
			file := FileList{
				Extension:    extension,
				Name:         file_info.Name(),
				AbsolutePath: filepath.Join(folder, file_info.Name()),
			}
			result.add_file_or_ignore(file)
		}
	}
}

func (r *Result) add_file_or_ignore(file FileList) {
	if starts_with_dot(file.Name) {
		return
	}
	switch file.Extension {
	case ".svg":
		r.Files = append(r.Files, file)
	default:
	}
}
func (a *App) CopyToClipboard(text string) {
	fmt.Println("Copy to clipboard: ", text)

	file_data, err := os.ReadFile(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File data: ", file_data)

	clipboard.WriteAll(string(file_data))
	time.Sleep(time.Second)
}

func (a *App) OpenInExplorer(text string) {
	fmt.Println("Open folder: ", text)
	folder_path := filepath.Join(MAIN_DIR, text)
	absolute_path, err := filepath.Abs(folder_path)
	if err == nil {
        fmt.Println("Absolute path: ", absolute_path)
    } else {
        fmt.Println(err)
	}
	// remove after last slash
	fmt.Println("Open folder: ", absolute_path)
	cmd := exec.Command("explorer", "/select,", absolute_path)
	cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func starts_with_dot(name string) bool {
	return strings.HasPrefix(name, ".")
}
func (a *App) GetFiles() Result {
	result := Result{
		Files: []FileList{},
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
			file := FileList{
				Extension:    extension,
				Name:         file_info.Name(),
				AbsolutePath: filepath.Join(MAIN_DIR, file_info.Name()),
			}
			result.add_file_or_ignore(file)
		}
	}
	return result
}
