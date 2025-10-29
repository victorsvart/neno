package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/victorsvart/neno/internal/store"
	"github.com/victorsvart/neno/internal/styles"
	"github.com/victorsvart/neno/internal/tui"
)

var showCmd = &cobra.Command{
	Use:   "show [title]",
	Short: "Show a note in terminal view mode",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		filename := store.SanitizeFilename(title)
		path := filepath.Join(store.NotesDir(), filename+".md")

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(styles.Error("", "Note not found: "+title))
			fmt.Println(styles.InfoStyle.Render("  Tip: Use 'neno new " + title + "' to create it"))
			return
		}

		tui.View(string(content), title)
	},
}

