package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/victorsvart/neno/internal/store"
	"github.com/victorsvart/neno/internal/styles"
)

var fileMode = 0644

var newCmd = &cobra.Command{
	Use:   "new [title]",
	Short: "Create a new note",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		filename := store.SanitizeFilename(title)
		path := filepath.Join(store.NotesDir(), filename+".md")

		content := fmt.Sprintf(`# %s

---
Created:: %s
Tags:: [misc]
---

`, title, time.Now().Format("2006-01-02"))

		if err := os.WriteFile(path, []byte(content), os.FileMode(fileMode)); err != nil {
			fmt.Println(styles.Error("", "Could not create note: "+err.Error()))
			return
		}

		fmt.Println(styles.Success("", "Note created successfully!"))
		fmt.Println(styles.PathStyle.Render("  " + path))
	},
}
