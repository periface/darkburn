package models

import "strings"

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

func starts_with_dot(name string) bool {
	return strings.HasPrefix(name, ".")
}
func (r *Result) Add_file_or_ignore(file FileList) {
	if starts_with_dot(file.Name) {
		return
	}
	switch file.Extension {
	case ".svg":
		r.Files = append(r.Files, file)
	default:
	}
}
