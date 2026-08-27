#!/bin/bash

# Clear the terminal screen
clear
echo "=================================================="
echo "      WELCOME TO YOUR NEW GO & TONIC PROJECT      "
echo "=================================================="
echo ""

# Ask for the new project name
echo "What would you like to name your new Gin & Tonic project?"
echo "(Example: my-awesome-app)"
read -p "Module Name: " NEW_NAME

# If just hit enter, show an error and stop
if [ -z "$NEW_NAME" ]; then
    echo ""
    echo "Error: Project name cannot be empty."
    exit 1
fi

OLD_NAME="gin-and-tonic"

echo ""
echo "Customizing ..."

# Find and replace all import strings across your Go files safely
if [[ "$OSTYPE" == "darwin"* ]]; then
    # Mac command layout
    find . -type f -name "*.go" -exec sed -i '' "s|${OLD_NAME}|${NEW_NAME}|g" {} +
    sed -i '' "s|module ${OLD_NAME}|module ${NEW_NAME}|g" ./go.mod
else
    # Linux / Git Bash command layout
    find . -type f -name "*.go" -exec sed -i "s|${OLD_NAME}|${NEW_NAME}|g" {} +
    sed -i "s|module ${OLD_NAME}|module ${NEW_NAME}|g" ./go.mod
fi

# Run go mod tidy
echo "Tidy up Go modules..."
go mod tidy

# Success feedback
echo ""
echo "=================================================="
echo "      SUCCESS! Your project is ready to go!       "
echo "=================================================="
echo "  All internal imports changed to: \"$NEW_NAME\". "
echo "  Run 'go run main.go' to boot your server!       "
echo "=================================================="
