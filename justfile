# List all available commands (default runner)
_help:
    @echo ""
    @just --list

alias r := run

# Run the application
[group('dev')]
run args="-git https://github.com/thalestmm/dots.git -dry-run":
    @go run . {{ args }}

alias b := build

# Build the application binary
[group('dev')]
build:
    @go build -o tmp/main .

alias t := test

# Test the entire application
[group('ci')]
[group('dev')]
test:
    @go test ./...

alias c := commit

# Stage all changes and commit
[group('dev')]
commit msg="chore: update":
    @git add .
    @git commit -m "{{ msg }}"

alias p := push

# Stage all changes, commit and push
[group('dev')]
push msg="chore: update":
    @just c "{{ msg }}"
    @git push

alias f := format

# Format code
[group('ci')]
format:
    @go fmt ./...
    @go vet ./...

# Bump app version in app.json, create new tag and publish binaries with goreleaser
[group('ci')]
release version="patch":
    #!/usr/bin/env bash
    git switch main
    git pull

    set -euo pipefail

    echo ""
    echo "Bumping version ({{ version }})"
    echo ""

    current_version=$(jq -r '.version' app.json)
    echo "Current version:  $current_version"

    case "{{ version }}" in major|minor|patch)

        major=$(echo "$current_version" | cut -d. -f1)
        minor=$(echo "$current_version" | cut -d. -f2)
        patch=$(echo "$current_version" | cut -d. -f3)

        case "{{ version }}" in
            major) major=$((major + 1)); minor=0; patch=0 ;;
            minor) minor=$((minor + 1)); patch=0 ;;
            patch) patch=$((patch + 1)) ;;
        esac

        new_version="$major.$minor.$patch"
        echo "New version:      $new_version"
    esac

    # Update version in app.json
    jq --arg v "$new_version" '.version = $v' app.json > app.json.tmp
    mv app.json.tmp app.json

    git add app.json
    git commit -m "chore: update version in app.json to v$new_version"
    git tag -a v$new_version -m "Release v$new_version"
    git push origin v$new_version

    # Create new release using the GitHub CLI
    # gh release create v$new_version --generate-notes --title "Release v$new_version" --fail-on-no-commits --verify-tag --draft
    git push origin main

    goreleaser -f config/goreleaser.yaml

alias si := script-install

# Run the installation script
[group('dev')]
script-install:
    ./scripts/install.sh

alias gi := go-install

# Install locally with go
[group('dev')]
go-install:
    @go install .
