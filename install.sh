#!/bin/bash

# Avatar CLI - install.sh
# Quickly installs the avtr binary for Linux and macOS

set -e

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then
  ARCH="arm64"
else
  echo "Unsupported architecture: $ARCH"
  exit 1
fi

if [[ "$OS" == "darwin" ]]; then
  FILE="avtr-darwin-$ARCH"
elif [[ "$OS" == "linux" ]]; then
  FILE="avtr-linux-$ARCH"
else
  echo "Unsupported OS: $OS"
  exit 1
fi

# The repo is github.com/Jayakrishna20/file-converter based on the prompt instructions
REPO="Jayakrishna20/file-converter"
DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/$FILE"

echo "Downloading Avatar CLI ($FILE)..."
curl -sL "$DOWNLOAD_URL" -o avtr

echo "Making avtr executable..."
chmod +x avtr

echo "Installing avtr to /usr/local/bin..."
if [[ "$EUID" -ne 0 ]]; then
  sudo mv avtr /usr/local/bin/
else
  mv avtr /usr/local/bin/
fi

echo "✨ Avatar installed successfully! Run 'avtr --help' to get started."
