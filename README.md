# ascii-art

ASCII Art Generator - Convert text strings into ASCII art using predefined banner styles (standard, shadow, thinkertoy)

## Features

- 🎨 Three banner styles: standard, shadow, thinkertoy
- ⚡ High performance (sub-millisecond rendering)
- 🧪 100% test coverage on critical packages
- 📦 Zero external dependencies (Go standard library only)
- 🔧 Cross-platform support (Linux, macOS, Windows)
- 🎯 Support for newline characters in input

## Installation

### Prerequisites

- Go 1.22.2 or higher

### Build from source

```bash
# Clone the repository
git clone <repository-url>
cd ascii-art

# Build
make build
# or: go build -o ascii-art .

# Run
./ascii-art "Hello World" standard
```

## Usage

### Run without building

```bash
go run . "text" [banner]
```

### Run with built binary

```bash
./ascii-art "text" [banner]
```

**Arguments**:
- `text`: The text to convert to ASCII art (required)
- `banner`: Banner style - standard, shadow, or thinkertoy (optional, defaults to standard)

### Examples

**Standard banner (default):**
```bash
go run . "Hello"
# or: ./ascii-art "Hello"
```

**Shadow banner:**
```bash
go run . "Hello" shadow
# or: ./ascii-art "Hello" shadow
```

**Thinkertoy banner:**
```bash
go run . "Hello" thinkertoy
# or: ./ascii-art "Hello" thinkertoy
```

**Newline support:**
```bash
go run . "Hello\nWorld"
# or: ./ascii-art "Hello\nWorld"
```

## Development

### Setup

```bash
# Run tests
make test

# Run with coverage
make coverage

# Run linters
make lint

# Format code
make fmt
```

### Project Structure

```
ascii-art/
├── .gitignore                 # Git ignore rules
├── .golangci.yml              # Linter configuration
├── LICENSE                    # Project license
├── Makefile                   # Build automation
├── go.mod                     # Go module file
├── main.go                    # CLI entry point
├── integration_test.go        # End-to-end tests
├── main_test.go               # Unit tests for main package
├── parser/                    # Banner file parsing package
│   ├── banner_parser.go
│   └── parser_test.go
├── renderer/                  # ASCII art rendering package
│   ├── renderer.go
│   ├── renderer_test.go
│   └── renderer_integration_test.go
├── testdata/                  # Banner files and test fixtures
│   ├── standard.txt           # Standard banner
│   ├── shadow.txt             # Shadow banner
│   ├── thinkertoy.txt         # Thinkertoy banner
│   ├── corrupted.txt          # Test fixture: corrupted file
│   ├── empty.txt              # Test fixture: empty file
│   └── oversized.txt          # Test fixture: oversized file
└── Documentation/
    ├── README.md              # This file
    ├── AGENTS.md              # AI agent instructions
    ├── CHANGELOG.md           # Version history
    ├── CONTRIBUTING.md        # Contribution guidelines
    └── PERMISSIONS.md         # Team permissions
```

### Running Tests

```bash
# All tests
make test

# With coverage report
make coverage

# Run benchmarks
make bench
```

### Build Commands

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Build for specific platforms
make build-linux    # Linux (amd64 and arm64)
make build-darwin   # macOS (amd64 and arm64)
make build-windows  # Windows (amd64)
```

## Architecture

The project follows a clean architecture with three main packages:

- **main**: CLI interface and orchestration
- **parser**: Banner file reading and character map building
- **renderer**: Text-to-ASCII-art conversion

## Performance

- **Single word ("Hello")**: ~83 µs (12,000 ops/sec)
- **Sentence (42 chars)**: ~143 µs (7,000 ops/sec)
- **Parser**: 100% coverage
- **Renderer**: 100% coverage
- **Overall coverage**: 79.8%

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

## Documentation

- [AGENTS.md](AGENTS.md) - AI agent instructions
- [CHANGELOG.md](CHANGELOG.md) - Version history
- [CONTRIBUTING.md](CONTRIBUTING.md) - Contribution guidelines

## License

See [LICENSE](LICENSE) file for details.
