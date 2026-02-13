# Class Diagram

Package relationships, exported types, and function signatures. All internal packages are independent — only `main` imports them.

```mermaid
classDiagram
    class main {
        +main()
        +ParseArgs(args []string) (string, string, error)
        +GetBannerPath(banner string) (string, error)
        -runColorMode(args []string)
        -hasColorFlag(args []string) bool
        -extractColorArgs(args []string) (string, string, string, string, error)
    }

    class parser {
        <<package>>
        +LoadBanner(path string) (Banner, error)
        +CharWidths(text string, banner Banner) []int
    }

    class Banner {
        <<type alias>>
        map~rune, []string~
    }

    class renderer {
        <<package>>
        +RendererASCII(input string, banner map~rune, []string~) (string, error)
    }

    class color {
        <<package>>
        +Parse(colorSpec string) (RGB, error)
        +ANSI(rgb RGB) string
    }

    class RGB {
        <<struct>>
        +R uint8
        +G uint8
        +B uint8
    }

    class coloring {
        <<package>>
        +ApplyColor(asciiArt []string, text string, substring string, colorCode string, charWidths []int) []string
        +Reset string
    }

    class flagparser {
        <<package>>
        +ParseArgs(args []string) error
    }

    main --> parser : loads banners
    main --> renderer : renders text
    main --> color : parses colors
    main --> coloring : applies colors
    main --> flagparser : validates args
    parser --> Banner : returns
    color --> RGB : returns
    parser ..> Banner : defines
    color ..> RGB : defines
```

## Dependency Rules

- `main` depends on all five internal packages
- No internal package imports another internal package
- All packages depend only on the Go standard library
- This ensures packages can be tested, reused, and maintained independently
