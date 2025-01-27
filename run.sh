#!/bin/bash

#Build the main.go
go build -o leetcode main.go
echo "Builded executable leetcode"

# Define the source and destination paths
SOURCE_DIR="$(pwd)"
# Print the directory
echo "The source directory is: $SOURCE_DIR"

DEST_DIR="/usr/local/bin"

# Files to copy
EXECUTABLE="leetcode"
JSON_FILE="leetcode.json"

# Check if the source files exist
if [ ! -f "$SOURCE_DIR/$EXECUTABLE" ] || [ ! -f "$SOURCE_DIR/$JSON_FILE" ]; then
    echo "Error: One or both source files not found."
    exit 1
fi

# Remove the existing files from /usr/local/bin
echo "Removing existing files from /usr/local/bin/..."
sudo rm -f "$DEST_DIR/$EXECUTABLE" "$DEST_DIR/$JSON_FILE"

# Copy the new files to /usr/local/bin
echo "Copying new files to /usr/local/bin/..."
sudo cp "$SOURCE_DIR/$EXECUTABLE" "$DEST_DIR/"
sudo cp "$SOURCE_DIR/$JSON_FILE" "$DEST_DIR/"

# Verify that the files were copied successfully
if [ -f "$DEST_DIR/$EXECUTABLE" ] && [ -f "$DEST_DIR/$JSON_FILE" ]; then
    echo "Files successfully copied to /usr/local/bin/"
else
    echo "Error: Failed to copy the files."
    exit 1
fi

# Set execute permissions for the executable
echo "Setting execute permissions for the leetcode executable..."
sudo chmod +x "$DEST_DIR/$EXECUTABLE"

# Confirm everything is set up
echo "Installation completed. You can now run 'leetcode -max_pagiantion' from anywhere."
