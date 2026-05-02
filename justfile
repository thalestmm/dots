# [R]un
[group('dev')]
r args="":
    @go run . {{ args }}

# [B]uild
[group('dev')]
b:
    @go build -o tmp/main .

# [W]atch
[group('dev')]
w:
    @air -c config/air.toml

# Stage all changes and [c]ommit
[group('dev')]
c msg="chore: update": fmt
    @git add .
    @git commit -m "{{ msg }}"

# Format code
[group('ci')]
fmt:
    @go fmt ./...
