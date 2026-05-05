# dots •○•

Pure `go` implementation of a simple `dotfiles` manager for synchronization accross multiple environments.

Available for Linux and MacOS systems.

## Installation

```bash
go install github.com/thalestmm/dots@latest
```

Run `dots` !

## Setup

Example structure for your `dotfiles` repository:

```bash
.dotfiles
├── a
│   └── .config
│       └── dots
│           └── a.json
└── b
    └── .config
        └── dots
            └── b.json
```

<<<<<<< HEAD
In this case, by running `dots -dir target`, the final outcome will be the `~/.config/dots` directory, containing the `a.json` and `b.json` symlinked files.
=======
In this case, the final outcome will be the `~/.config/dots` directory, containing the `a.json` and `b.json` symlinked files.
>>>>>>> 840933b (chore: update)

---
*No AI was used during the development of this tool.*
