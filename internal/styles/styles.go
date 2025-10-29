package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	primary   = lipgloss.Color("#00D9FF")
	success   = lipgloss.Color("#00FF87")
	error     = lipgloss.Color("#FF5F87")
	warning   = lipgloss.Color("#FFD700")
	muted     = lipgloss.Color("#626262")
	
	// Success message style
	SuccessStyle = lipgloss.NewStyle().
		Foreground(success).
		Bold(true).
		PaddingLeft(1)
	
	// Error message style
	ErrorStyle = lipgloss.NewStyle().
		Foreground(error).
		Bold(true).
		PaddingLeft(1)
	
	// Info message style
	InfoStyle = lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		PaddingLeft(1)
	
	// Warning message style
	WarningStyle = lipgloss.NewStyle().
		Foreground(warning).
		Bold(true).
		PaddingLeft(1)
	
	// Path style
	PathStyle = lipgloss.NewStyle().
		Foreground(muted).
		Italic(true)
	
	// Title style
	TitleStyle = lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		Underline(true)
	
	// Box style for content
	BoxStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(primary).
		Padding(1, 2)
	
	// Highlight style
	HighlightStyle = lipgloss.NewStyle().
		Foreground(warning).
		Bold(true)
)

// Success formats a success message
func Success(icon, message string) string {
	return SuccessStyle.Render("✓ " + icon + " " + message)
}

// Error formats an error message
func Error(icon, message string) string {
	return ErrorStyle.Render("✗ " + icon + " " + message)
}

// Info formats an info message
func Info(icon, message string) string {
	return InfoStyle.Render("→ " + icon + " " + message)
}

// Warning formats a warning message
func Warning(icon, message string) string {
	return WarningStyle.Render("⚠ " + icon + " " + message)
}

