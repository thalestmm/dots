- Include `sync` command: pulls, commits local changes (with machine identifier), pushes to remote (`dots sync`)
- Use the relative `$HOME/...` path in the dots config file for support accross different machines
- Add support for private repos using the github CLI
- Add specific child commands for `sync`: `pull` and `push` (base cmd runs all)
- Create `add` command, where we setup a new dotfiles dir containing the path (dir or file) to a new app to be tracked.
  This could be a part of the `init` command, asking if the user already has a configuration, etc
- In the help command, include instructions as to how the dotfiles dir should be structured
- Implement `update` command, which fetches the latest `dots` version and updates the local installation
- Add `version` command, which prints the current `dots` version. Add `-v` flag to print the version from the
  root command
