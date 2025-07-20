#!/bin/bash
set -e

# iplookup installation script
# https://github.com/gofika/iplookup

REPO="gofika/iplookup"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="iplookup"

# Detect system architecture
get_arch() {
    local arch=$(uname -m)
    case $arch in
        x86_64) echo "x86_64" ;;
        amd64) echo "x86_64" ;;
        aarch64) echo "arm64" ;;
        arm64) echo "arm64" ;;
        armv7l) echo "armv7" ;;
        armv6l) echo "armv6" ;;
        i686) echo "i386" ;;
        i386) echo "i386" ;;
        *) echo "Unsupported architecture: $arch" >&2; exit 1 ;;
    esac
}

# Detect operating system
get_os() {
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    case $os in
        linux) echo "linux" ;;
        darwin) echo "darwin" ;;
        freebsd) echo "freebsd" ;;
        openbsd) echo "openbsd" ;;
        *) echo "Unsupported operating system: $os" >&2; exit 1 ;;
    esac
}

# Get latest version
get_latest_version() {
    curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
}

# Download and install
install_iplookup() {
    local version=$(get_latest_version)
    local os=$(get_os)
    local arch=$(get_arch)

    echo "Installing iplookup ${version} for ${os}/${arch}..."

    # Build download URL
    local filename="iplookup_${version#v}_${os}_${arch}"
    if [ "$os" = "linux" ] || [ "$os" = "freebsd" ] || [ "$os" = "openbsd" ]; then
        filename="${filename}.tar.gz"
    else
        filename="${filename}.tar.gz"
    fi

    local download_url="https://github.com/${REPO}/releases/download/${version}/${filename}"

    # Create temporary directory
    local tmp_dir=$(mktemp -d)
    trap "rm -rf $tmp_dir" EXIT

    # Download file
    echo "Downloading ${download_url}..."
    curl -sL "$download_url" -o "$tmp_dir/$filename"

    # Extract file
    echo "Extracting..."
    tar -xzf "$tmp_dir/$filename" -C "$tmp_dir"

    # Check if sudo is needed
    if [ -w "$INSTALL_DIR" ]; then
        mv "$tmp_dir/$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
        chmod +x "$INSTALL_DIR/$BINARY_NAME"
    else
        echo "Administrator privileges required to install to $INSTALL_DIR"
        sudo mv "$tmp_dir/$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
        sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
    fi

    echo "iplookup has been successfully installed to $INSTALL_DIR/$BINARY_NAME"
    echo ""
    echo "Usage:"
    echo "  iplookup 8.8.8.8"
    echo ""
    echo "Show help:"
    echo "  iplookup -h"
}

# Check dependencies
check_dependencies() {
    local deps=("curl" "tar")
    for dep in "${deps[@]}"; do
        if ! command -v "$dep" &> /dev/null; then
            echo "Error: $dep is required but not installed" >&2
            exit 1
        fi
    done
}

# Main function
main() {
    echo "=== iplookup Installation Script ==="
    echo ""

    check_dependencies
    install_iplookup
}

main "$@"