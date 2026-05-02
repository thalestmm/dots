# [R]un
[group('dev')]
r args="-dir example":
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



    goreleaser -f config/goreleaser.yaml --snapshot --clean
