#!/usr/bin/env bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Detect OS and set appropriate variables
detect_os() {
    case "$(uname -s)" in
        Darwin*)
            OS="macOS"
            INSTALL_DIR="/usr/local/bin"
            BINARY_NAME="neno"
            ;;
        Linux*)
            OS="Linux"
            INSTALL_DIR="/usr/local/bin"
            BINARY_NAME="neno"
            ;;
        CYGWIN*|MINGW*|MSYS*)
            OS="Windows"
            INSTALL_DIR="$HOME/bin"
            BINARY_NAME="neno.exe"
            ;;
        *)
            OS="Unknown"
            INSTALL_DIR="/usr/local/bin"
            BINARY_NAME="neno"
            ;;
    esac
}

detect_os
INSTALL_PATH="${INSTALL_DIR}/${BINARY_NAME}"

echo -e "${CYAN}╔════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║  NENO - Note Manager Installer        ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${CYAN}Detected OS:${NC} $OS"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}✗ Error: Go is not installed${NC}"
    echo -e "${YELLOW}  Please install Go from https://golang.org/dl/${NC}"
    if [ "$OS" = "macOS" ]; then
        echo -e "${YELLOW}  macOS: brew install go${NC}"
    elif [ "$OS" = "Linux" ]; then
        echo -e "${YELLOW}  Linux: sudo apt install golang-go (Ubuntu/Debian)${NC}"
        echo -e "${YELLOW}         sudo pacman -S go (Arch)${NC}"
    fi
    exit 1
fi

echo -e "${GREEN}✓${NC} Go found: $(go version)"

# Build the application
echo -e "${CYAN}→${NC} Building ${BINARY_NAME}..."
if go build -o "${BINARY_NAME}" .; then
    echo -e "${GREEN}✓${NC} Build successful"
else
    echo -e "${RED}✗ Build failed${NC}"
    exit 1
fi

# Ensure install directory exists and check if we need sudo
if [ ! -d "${INSTALL_DIR}" ]; then
    if [ "$OS" = "Windows" ]; then
        mkdir -p "${INSTALL_DIR}"
        SUDO=""
    else
        SUDO="sudo"
        ${SUDO} mkdir -p "${INSTALL_DIR}"
    fi
fi

# Check if we need sudo
if [ "$OS" = "Windows" ] || [ -w "${INSTALL_DIR}" ]; then
    SUDO=""
else
    SUDO="sudo"
    echo -e "${YELLOW}⚠${NC}  Need sudo privileges to install to ${INSTALL_DIR}"
fi

# Check if already installed
if [ -f "${INSTALL_PATH}" ]; then
    echo -e "${YELLOW}⚠${NC}  ${BINARY_NAME} is already installed at ${INSTALL_PATH}"
    read -p "  Overwrite? (y/N) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${YELLOW}✗ Installation cancelled${NC}"
        rm "${BINARY_NAME}"
        exit 0
    fi
fi

# Install the binary
echo -e "${CYAN}→${NC} Installing ${BINARY_NAME} to ${INSTALL_PATH}..."
if ${SUDO} mv "${BINARY_NAME}" "${INSTALL_PATH}"; then
    ${SUDO} chmod +x "${INSTALL_PATH}"
    echo -e "${GREEN}✓${NC} Installation complete!"
else
    echo -e "${RED}✗ Installation failed${NC}"
    exit 1
fi

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║  Installation successful!              ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${CYAN}Quick Start:${NC}"
echo -e "  ${YELLOW}neno new <title>${NC}    - Create a new note"
echo -e "  ${YELLOW}neno list${NC}           - List all notes"
echo -e "  ${YELLOW}neno tags${NC}           - Show all tags"
echo -e "  ${YELLOW}neno --help${NC}         - See all commands"
echo ""
echo -e "${CYAN}Notes are stored in:${NC} ~/.neno/pages/"
echo ""
if [ "$OS" = "Windows" ]; then
    echo -e "${YELLOW}Note:${NC} Make sure $INSTALL_DIR is in your PATH"
    echo -e "  Add this to your .bashrc or .zshrc:"
    echo -e "  export PATH=\"\$HOME/bin:\$PATH\""
    echo ""
fi
