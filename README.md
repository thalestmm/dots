# dots •○•

Pure `go` implementation of a simple `dotfiles` manager for synchronization accross multiple environments.

Available for Linux and MacOS systems.

Example setup for a target directory:

```bash
target
├── a
│   └── .config
│       └── dot
│           └── a.json
└── b
    └── .config
        └── dot
            └── b.json
```

In this case, by running `dot -dir target`, the final outcome will be the `~/.config/dot` directory, containing the `a.json` and `b.json` symlinked files.


---
*No AI was used during the development of this tool.*
