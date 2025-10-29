package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"sort"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/victorsvart/neno/internal/store"
	"github.com/victorsvart/neno/internal/styles"
	"github.com/victorsvart/neno/internal/tui"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

type noteItem struct {
	note *store.Note
}

func (i noteItem) Title() string { return i.note.Title }
func (i noteItem) Description() string {
	tags := ""
	if len(i.note.Tags) > 0 {
		tags = " • Tags: " + fmt.Sprintf("%v", i.note.Tags)
	}
	return i.note.Created.Format("2006-01-02") + tags
}
func (i noteItem) FilterValue() string { return i.note.Title }

type listModel struct {
	list     list.Model
	selected *store.Note
	action   string // "show" or "edit"
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(noteItem)
			if ok {
				m.selected = i.note
				return m, tea.Quit
			}
		case "e":
			// Edit mode
			i, ok := m.list.SelectedItem().(noteItem)
			if ok {
				m.selected = i.note
				m.action = "edit"
				return m, tea.Quit
			}
		case "s":
			// Show mode
			i, ok := m.list.SelectedItem().(noteItem)
			if ok {
				m.selected = i.note
				m.action = "show"
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m listModel) View() string {
	helpText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Italic(true).
		Render("\n[enter] show • [e] edit • [/] filter • [q] quit")
	return docStyle.Render(m.list.View() + helpText)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all notes with interactive selection",
	Run: func(cmd *cobra.Command, args []string) {
		notes, err := store.GetAllNotes()
		if err != nil {
			fmt.Println(styles.Error("", "Could not load notes: "+err.Error()))
			return
		}

		if len(notes) == 0 {
			fmt.Println(styles.Warning("", "No notes found"))
			fmt.Println(styles.InfoStyle.Render("  Create your first note with: neno new <title>"))
			return
		}

		// Sort notes by created date (newest first)
		sort.Slice(notes, func(i, j int) bool {
			return notes[i].Created.After(notes[j].Created)
		})

		items := make([]list.Item, len(notes))
		for i, note := range notes {
			items[i] = noteItem{note: note}
		}

		delegate := list.NewDefaultDelegate()
		delegate.SetSpacing(1)

		l := list.New(items, delegate, 0, 0)
		l.Title = "📝 Your Notes"
		l.SetShowHelp(true)

		m := listModel{list: l, action: "show"}
		p := tea.NewProgram(m, tea.WithAltScreen())

		finalModel, err := p.Run()
		if err != nil {
			fmt.Println(styles.Error("", "Error running list: "+err.Error()))
			return
		}

		if m, ok := finalModel.(listModel); ok && m.selected != nil {
			if m.action == "edit" {
				// Open in nano
				nanoCmd := exec.Command("nano", m.selected.Path)
				nanoCmd.Stdin = os.Stdin
				nanoCmd.Stdout = os.Stdout
				nanoCmd.Stderr = os.Stderr
				if err := nanoCmd.Run(); err != nil {
					fmt.Println(styles.Error("", "Failed to open nano: "+err.Error()))
				}
			} else {
				// Show the note
				content, err := os.ReadFile(m.selected.Path)
				if err != nil {
					fmt.Println(styles.Error("", "Could not read note: "+err.Error()))
					return
				}
				// Import tui package
				showNoteInteractive(string(content), m.selected.Title)
			}
		}
	},
}

// Helper function to show note
func showNoteInteractive(content, title string) {
	tui.View(content, title)
}
