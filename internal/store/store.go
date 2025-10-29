package store

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var fileMode = 0755

func NotesDir() string {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".neno", "pages")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.FileMode(fileMode))
	}

	return dir
}

func SanitizeFilename(name string) string {
	replacer := strings.NewReplacer(" ", "-", "/", "-", "\\", "-")
	return replacer.Replace(strings.ToLower(name))
}

func ListNotes() ([]string, error) {
	var files []string
	root := NotesDir()
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Ext(path) == ".md" {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func Debug() {
	fmt.Println("Note directory:", NotesDir())
}
