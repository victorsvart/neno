package tui

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00D9FF")).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(lipgloss.Color("#626262")).
			MarginBottom(1).
			Padding(0, 1)

	contentStyle = lipgloss.NewStyle().
			Padding(0, 2).
			MarginBottom(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Italic(true).
			Padding(0, 2)

	boxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00D9FF")).
			Padding(1, 2).
			Margin(1, 0)
)

type model struct {
	content string
	title   string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	
	// Title
	b.WriteString(titleStyle.Render("📝 " + m.title))
	b.WriteString("\n\n")
	
	// Content
	b.WriteString(contentStyle.Render(m.content))
	b.WriteString("\n")
	
	// Help text
	b.WriteString(helpStyle.Render("Press 'q' or 'esc' to quit"))
	b.WriteString("\n")

	return boxStyle.Render(b.String())
}

func View(content string, title string) {
	p := tea.NewProgram(model{content: content, title: title})
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
