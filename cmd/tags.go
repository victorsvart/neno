package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/victorsvart/neno/internal/store"
	"github.com/victorsvart/neno/internal/styles"
)

var (
	tagStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true)
	
	noteListStyle = lipgloss.NewStyle().
		PaddingLeft(4).
		Foreground(lipgloss.Color("#00D9FF"))
)

var tagsCmd = &cobra.Command{
	Use:   "tags [tag]",
	Short: "List all tags or notes with a specific tag",
	Long:  "If no tag is provided, lists all available tags. If a tag is provided, lists all notes with that tag.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// List all tags
			listAllTags()
		} else {
			// List notes by specific tag
			tag := strings.Join(args, " ")
			listNotesByTag(tag)
		}
	},
}

func listAllTags() {
	tags, err := store.GetAllTags()
	if err != nil {
		fmt.Println(styles.Error("", "Could not load tags: "+err.Error()))
		return
	}

	if len(tags) == 0 {
		fmt.Println(styles.Warning("", "No tags found"))
		return
	}

	sort.Strings(tags)

	fmt.Println(styles.TitleStyle.Render("🏷  Available Tags"))
	fmt.Println()

	for _, tag := range tags {
		notes, _ := store.GetNotesByTag(tag)
		count := len(notes)
		fmt.Printf("  %s %s\n", 
			tagStyle.Render("• "+tag),
			styles.PathStyle.Render(fmt.Sprintf("(%d note%s)", count, plural(count))),
		)
	}

	fmt.Println()
	fmt.Println(styles.InfoStyle.Render("  Usage: neno tags <tag> to see notes with a specific tag"))
}

func listNotesByTag(tag string) {
	notes, err := store.GetNotesByTag(tag)
	if err != nil {
		fmt.Println(styles.Error("", "Could not load notes: "+err.Error()))
		return
	}

	if len(notes) == 0 {
		fmt.Println(styles.Warning("", fmt.Sprintf("No notes found with tag '%s'", tag)))
		return
	}

	// Sort by created date (newest first)
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Created.After(notes[j].Created)
	})

	fmt.Println(styles.TitleStyle.Render(fmt.Sprintf("🏷  Notes tagged with '%s'", tag)))
	fmt.Println()

	for _, note := range notes {
		fmt.Printf("  %s\n", noteListStyle.Render("• "+note.Title))
		fmt.Printf("    %s\n", styles.PathStyle.Render(note.Created.Format("2006-01-02")))
	}

	fmt.Println()
	fmt.Printf("%s\n", styles.InfoStyle.Render(fmt.Sprintf("  Found %d note%s", len(notes), plural(len(notes)))))
	fmt.Println(styles.InfoStyle.Render("  Use 'neno show <title>' to view a note"))
}

func plural(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}

