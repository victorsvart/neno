package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/victorsvart/neno/internal/store"
	"github.com/victorsvart/neno/internal/styles"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search notes for a string (uses ripgrep)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		dir := store.NotesDir()

		fmt.Println(styles.Info("", "Searching for: "+styles.HighlightStyle.Render(query)))
		fmt.Println()

		out, err := exec.Command("rg", "--color", "always", "--heading", "--line-number", query, dir).CombinedOutput()
		if err != nil {
			fmt.Println(styles.Warning("", "No matches found"))
			fmt.Println(styles.PathStyle.Render("  Make sure ripgrep is installed: https://github.com/BurntSushi/ripgrep"))
			return
		}

		resultStyle := lipgloss.NewStyle().MarginLeft(2)
		fmt.Println(resultStyle.Render(string(out)))
	},
}
