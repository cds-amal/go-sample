#!/usr/bin/env bash

# Remove the existing go.sum file
rm ./go.sum >/dev/null 2>&1

export GOPRIVATE=github.com/DIN-center
# Download the dependencies for the specified module
go mod download github.com/DIN-center/din-sc/apps/din-go
go list -m all

# Check if the go.sum file matches the expected go.sum file
if diff go.sum expected.go.sum >/dev/null; then
    # Check if GOPATH is set, if not use $HOME/go
    if [ -z "$GOPATH" ]; then
        GOPATH="$HOME/go"
    fi

    # List the contents of the module path
    # printf "======================================================================================\n"
    # printf "Listing contents of $GOPATH/pkg/mod/github.com/\!d\!i\!n-center/din-sc/apps\n"
    # printf "======================================================================================\n\n"
    # tree $GOPATH/pkg/mod/github.com/\!d\!i\!n-center/din-sc/apps 

    printf "go.sum:\n"
    sed 's/^/    /' ./go.sum  # Add indentation to each line of go.sum
    printf "\nYay... go mod download - did its thing!\n"

else
    echo "Failed to download din-go"
fi
