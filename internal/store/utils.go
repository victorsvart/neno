package store

import (
	"fmt"
	"os"
)

func EnsureDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}
}
