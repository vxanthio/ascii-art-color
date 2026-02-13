# Architecture Overview

High-level view of the system architecture. Packages are grouped by responsibility layer.

```mermaid
flowchart LR
    subgraph CLI["CLI Layer"]
        main["main\n(cmd/ascii-art)"]
    end

    subgraph Input["Input Processing"]
        flagparser["flagparser\nCLI validation"]
        color["color\nColor parsing"]
    end

    subgraph Core["Core Engine"]
        parser["parser\nBanner loading"]
        renderer["renderer\nASCII rendering"]
    end

    subgraph Output["Output Processing"]
        coloring["coloring\nANSI color application"]
    end

    main -->|"validates args"| flagparser
    main -->|"parses color spec"| color
    main -->|"loads banner file"| parser
    main -->|"renders text"| renderer
    main -->|"applies color"| coloring

    style CLI fill:#4a90d9,color:#fff
    style Input fill:#7b68ee,color:#fff
    style Core fill:#2ecc71,color:#fff
    style Output fill:#e67e22,color:#fff
```

## Package Responsibilities

| Layer | Package | Responsibility |
|-------|---------|---------------|
| CLI | `main` | Orchestrates all packages, handles I/O |
| Input | `flagparser` | Validates CLI argument structure |
| Input | `color` | Parses color specs (named, hex, RGB) into RGB values |
| Core | `parser` | Reads banner files, builds character maps |
| Core | `renderer` | Converts text to ASCII art using banner maps |
| Output | `coloring` | Applies ANSI color codes to rendered ASCII art |

## Key Design Decisions

- **Zero inter-package dependencies** — all packages depend only on the Go standard library
- **Main as orchestrator** — `main` is the only package that imports other project packages
- **Stateless packages** — all functions are pure transformations (no global state, no side effects except file I/O in parser)
