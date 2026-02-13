# Program Flowchart

Execution flow from CLI input to ASCII art output. The program has two modes: **Normal** (text only) and **Color** (with ANSI coloring).

```mermaid
flowchart TD
    A["CLI Arguments\nos.Args"] --> B{"hasColorFlag?\n--color= prefix"}

    B -->|No| C["ParseArgs()\ntext, banner"]
    B -->|Yes| D["flagparser.ParseArgs()\nvalidate syntax"]

    C --> E["GetBannerPath()\nbanner file path"]
    D --> F["extractColorArgs()\ncolorSpec, substring,\ntext, banner"]

    E --> G["parser.LoadBanner()\nBanner map"]
    F --> H["color.Parse()\nRGB struct"]

    H --> I["GetBannerPath()\nbanner file path"]
    I --> J["parser.LoadBanner()\nBanner map"]
    J --> K["color.ANSI()\nANSI escape code"]

    G --> L["renderer.RendererASCII()\nASCII art string"]
    L --> M["fmt.Print()\nstdout"]

    K --> N["For each line in text"]
    N --> O["renderer.RendererASCII()\nASCII art lines"]
    O --> P["parser.CharWidths()\ncharacter widths"]
    P --> Q["coloring.ApplyColor()\ncolored ASCII art"]
    Q --> R{"More lines?"}
    R -->|Yes| N
    R -->|No| S["fmt.Print()\nstdout"]

    style B fill:#f39c12,color:#fff
    style R fill:#f39c12,color:#fff
    style M fill:#2ecc71,color:#fff
    style S fill:#2ecc71,color:#fff
```

## Mode Comparison

| Aspect | Normal Mode | Color Mode |
|--------|------------|------------|
| Validation | `ParseArgs()` | `flagparser.ParseArgs()` |
| Color parsing | — | `color.Parse()` → `color.ANSI()` |
| Rendering | Single call | Per-line loop |
| Post-processing | — | `CharWidths()` + `ApplyColor()` |
