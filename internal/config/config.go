package config

import (
	"os"
	"path/filepath"
)

var filemode = 0755

type Config struct {
	Editor string `yaml:"editor"`
}

func DefaultConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".neno", "config.yaml")
}

func EnsureConfig() {
	path := DefaultConfigPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(path), os.FileMode(filemode))
		os.WriteFile(path, []byte("editor: nvim\n"), 0644)
	}
}
