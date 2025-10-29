package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/victorsvart/neno/internal/store"
	"github.com/victorsvart/neno/internal/styles"
)

var openCmd = &cobra.Command{
	Use:   "open [title]",
	Short: "Open a note for editing in nano",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		filename := store.SanitizeFilename(title)
		path := filepath.Join(store.NotesDir(), filename+".md")

		// Check if the file exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println(styles.Error("", "Note not found: "+title))
			fmt.Println(styles.InfoStyle.Render("  Tip: Use 'neno new " + title + "' to create it"))
			return
		}

		fmt.Println(styles.Info("", "Opening note in nano..."))
		
		// Open the note in nano
		nanoCmd := exec.Command("nano", path)
		nanoCmd.Stdin = os.Stdin
		nanoCmd.Stdout = os.Stdout
		nanoCmd.Stderr = os.Stderr

		if err := nanoCmd.Run(); err != nil {
			fmt.Println(styles.Error("", "Failed to open nano: "+err.Error()))
			return
		}
	},
}
