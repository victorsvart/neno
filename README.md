# NENO - Never Enough Notes Organizer

<p align="center">
  <strong>A beautiful, markdown-first note organizer for your terminal</strong>
</p>

## ✨ Features

- 📝 **Create and organize** markdown notes
- 🎨 **Beautiful TUI** with syntax highlighting and styling
- 🔍 **Powerful search** using ripgrep
- 🏷️  **Tag system** for categorizing notes
- 📋 **Interactive list** view for easy note browsing
- ✏️  **Nano integration** for quick editing
- 🎯 **Simple and fast** - designed for efficiency

## 🚀 Quick Start

### Installation

**One-line install (from Git):**
```bash
git clone https://github.com/victorsvart/neno.git && cd neno && ./install.sh
```

**Or using curl (web installer):**
```bash
curl -sSL https://raw.githubusercontent.com/victorsvart/neno/main/web-install.sh | bash
```

**Or using Make:**
```bash
git clone https://github.com/victorsvart/neno.git
cd neno
make install
```

**Or manual installation:**
```bash
git clone https://github.com/victorsvart/neno.git
cd neno
./install.sh
```

See [INSTALL.md](INSTALL.md) for more installation options and troubleshooting.

### Basic Usage

```bash
# Create a new note
neno new "My First Note"

# List all notes (interactive)
neno list

# Show a note
neno show "My First Note"

# Edit a note
neno open "My First Note"

# Search notes
neno search "keyword"

# Show all tags
neno tags

# Show notes by tag
neno tags misc
```

## 📚 Commands

| Command | Description |
|---------|-------------|
| `neno new <title>` | Create a new note |
| `neno list` | Interactive list of all notes |
| `neno show <title>` | Display a note in the viewer |
| `neno open <title>` | Open a note for editing in nano |
| `neno search <query>` | Search through all notes |
| `neno tags` | List all available tags |
| `neno tags <tag>` | List notes with a specific tag |

## 📝 Note Format

Notes are stored as markdown files in `~/.neno/pages/` with metadata:

```markdown
# Note Title

---
Created:: 2024-10-29
Tags:: [misc, important]
---

Your note content here...
```

## 🎨 Interactive List

The `neno list` command provides an interactive interface:

- **Enter** - Show the selected note
- **e** - Edit the selected note in nano
- **/** - Filter/search notes
- **↑/↓** - Navigate notes
- **q** - Quit

## 🛠️ Development

```bash
# Build
make build

# Run tests
make test

# Format code
make fmt

# Install locally
make install

# See all commands
make help
```

## 📦 Project Structure

```
neno/
├── cmd/                # Command definitions
│   ├── list.go        # Interactive list command
│   ├── new.go         # Create notes
│   ├── open.go        # Edit notes
│   ├── search.go      # Search functionality
│   ├── show.go        # Display notes
│   ├── tags.go        # Tag management
│   └── root.go        # Root command
├── internal/
│   ├── config/        # Configuration
│   ├── store/         # Note storage and parsing
│   ├── styles/        # Visual styling
│   └── tui/           # Terminal UI components
├── main.go            # Entry point
├── install.sh         # Installation script
├── uninstall.sh       # Uninstallation script
├── Makefile           # Build automation
└── INSTALL.md         # Detailed installation guide
```

## 🎯 Requirements

- **Go** 1.20 or higher
- **ripgrep** (optional, for search functionality)
- **nano** (for editing notes)

## 🔧 Uninstallation

```bash
# Using the uninstall script
./uninstall.sh

# Or using Make
make uninstall
```

**Note:** This removes the binary but keeps your notes in `~/.neno/pages/`

## 🤝 Contributing

Contributions are welcome! Feel free to:

- Report bugs
- Suggest features
- Submit pull requests

## 📄 License

This project is open source and available under the MIT License.

## 🙏 Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling

---

<p align="center">
  Made with ❤️ for note-takers everywhere
</p>

