#!/usr/bin/env bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Installation variables
BINARY_NAME="neno"
INSTALL_PATH="/usr/local/bin/${BINARY_NAME}"
DATA_DIR="${HOME}/.neno"

echo -e "${CYAN}╔════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║  NENO - Uninstaller                    ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════╝${NC}"
echo ""

# Check if neno is installed
if [ ! -f "${INSTALL_PATH}" ]; then
    echo -e "${RED}✗ ${BINARY_NAME} is not installed at ${INSTALL_PATH}${NC}"
    exit 1
fi

echo -e "${YELLOW}This will remove:${NC}"
echo -e "  • Binary: ${INSTALL_PATH}"
echo ""
echo -e "${YELLOW}⚠  Your notes in ${DATA_DIR} will NOT be deleted${NC}"
echo ""

read -p "Continue with uninstallation? (y/N) " -n 1 -r
echo

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${YELLOW}✗ Uninstallation cancelled${NC}"
    exit 0
fi

# Remove the binary
echo -e "${CYAN}→${NC} Removing ${BINARY_NAME}..."

# Check if we need sudo
if [ -w "/usr/local/bin" ]; then
    SUDO=""
else
    SUDO="sudo"
fi

if ${SUDO} rm -f "${INSTALL_PATH}"; then
    echo -e "${GREEN}✓${NC} Binary removed"
else
    echo -e "${RED}✗ Failed to remove binary${NC}"
    exit 1
fi

echo ""
echo -e "${GREEN}✓ Uninstallation complete!${NC}"
echo ""
echo -e "${CYAN}Your notes are still available at: ${DATA_DIR}${NC}"
echo -e "${YELLOW}To remove your notes too, run:${NC} rm -rf ${DATA_DIR}"
echo ""

