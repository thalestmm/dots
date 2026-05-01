# Stage all changes and [c]ommit
[group('dev')]
c msg="chore: update":
    @git add .
    @git commit -m "{{ msg }}"
