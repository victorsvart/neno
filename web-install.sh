#!/usr/bin/env bash

# NENO Web Installer
# Usage: curl -sSL https://raw.githubusercontent.com/victorsvart/neno/main/web-install.sh | bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Configuration
REPO_URL="https://github.com/victorsvart/neno.git"
BINARY_NAME="neno"

# Detect OS and set appropriate install directory
detect_os() {
    case "$(uname -s)" in
        Darwin*)
            OS="macOS"
            INSTALL_DIR="/usr/local/bin"
            ;;
        Linux*)
            OS="Linux"
            INSTALL_DIR="/usr/local/bin"
            ;;
        CYGWIN*|MINGW*|MSYS*)
            OS="Windows"
            INSTALL_DIR="$HOME/bin"
            BINARY_NAME="neno.exe"
            ;;
        *)
            OS="Unknown"
            INSTALL_DIR="/usr/local/bin"
            ;;
    esac
}

detect_os
INSTALL_PATH="${INSTALL_DIR}/${BINARY_NAME}"
TEMP_DIR=$(mktemp -d 2>/dev/null || mktemp -d -t 'neno')

echo -e "${CYAN}╔════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║  NENO - Web Installer                  ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${CYAN}Detected OS:${NC} $OS"
echo ""

# Cleanup on exit
cleanup() {
    if [ -d "$TEMP_DIR" ]; then
        rm -rf "$TEMP_DIR"
    fi
}
trap cleanup EXIT

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}✗ Error: Go is not installed${NC}"
    echo -e "${YELLOW}  Please install Go from https://golang.org/dl/${NC}"
    if [ "$OS" = "macOS" ]; then
        echo -e "${YELLOW}  macOS: brew install go${NC}"
    elif [ "$OS" = "Linux" ]; then
        echo -e "${YELLOW}  Linux: Visit https://go.dev/doc/install${NC}"
    fi
    exit 1
fi

echo -e "${GREEN}✓${NC} Go found: $(go version)"

# Check if git is installed
if ! command -v git &> /dev/null; then
    echo -e "${RED}✗ Error: Git is not installed${NC}"
    echo -e "${YELLOW}  Please install Git first${NC}"
    exit 1
fi

echo -e "${GREEN}✓${NC} Git found: $(git --version)"

# Clone the repository
echo -e "${CYAN}→${NC} Cloning repository..."
if git clone --depth=1 "$REPO_URL" "$TEMP_DIR" &> /dev/null; then
    echo -e "${GREEN}✓${NC} Repository cloned"
else
    echo -e "${RED}✗ Failed to clone repository${NC}"
    echo -e "${YELLOW}  Make sure the repository URL is correct:${NC}"
    echo -e "${YELLOW}  $REPO_URL${NC}"
    exit 1
fi

# Build the application
echo -e "${CYAN}→${NC} Building ${BINARY_NAME}..."
cd "$TEMP_DIR"
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
    CURRENT_VERSION=$(${INSTALL_PATH} --version 2>/dev/null || echo "unknown")
    echo -e "${YELLOW}  Current version: ${CURRENT_VERSION}${NC}"
    read -p "  Overwrite? (y/N) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${YELLOW}✗ Installation cancelled${NC}"
        exit 0
    fi
fi

# Install the binary
echo -e "${CYAN}→${NC} Installing ${BINARY_NAME} to ${INSTALL_PATH}..."
if ${SUDO} install -m 755 "${BINARY_NAME}" "${INSTALL_PATH}"; then
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
echo -e "${CYAN}To uninstall:${NC} ${SUDO} rm ${INSTALL_PATH}"
echo ""

