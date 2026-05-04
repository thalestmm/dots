# [R]un
[group('dev')]
r args="-git https://github.com/thalestmm/dots.git -dry-run":
    @go run . {{ args }}

# [B]uild
[group('dev')]
b:
    @go build -o tmp/main .

# [T]est
[group('ci')]
[group('dev')]
t:
    @go test ./...

# [W]atch
[group('dev')]
w:
    @air -c config/air.toml

# Stage all changes and [c]ommit
[group('dev')]
c msg="chore: update": f
    @git add .
    @git commit -m "{{ msg }}"

# Stage all changes, commit and [p]ush
[group('dev')]
p msg="chore: update":
    @just c "{{ msg }}"
    @git push

# [F]ormat code
[group('ci')]
f:
    @go fmt ./...
    @go vet ./...

# Bump app version in app.json, create new tag and publish binaries with goreleaser
[group('ci')]
release version="patch": f b t
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
    gh release create v$new_version --generate-notes --title "Release v$new_version" --fail-on-no-commits --verify-tag --draft
    git push origin main

    # goreleaser -f config/goreleaser.yaml --snapshot --clean
