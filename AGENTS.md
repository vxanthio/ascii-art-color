# AGENTS.md

Instructions for AI coding agents working on the ascii-art-color project.

## Project Overview

This is a Go CLI application that converts text to ASCII art using three banner styles (standard, shadow, thinkertoy), with optional ANSI 24-bit color support for full text or specific substrings. The project uses only Go standard library with zero external dependencies.

## Setup Commands

```bash
# No external dependencies to install (Go standard library only)

# Build the project
make build
# or: go build -o ascii-art ./cmd/ascii-art

# Run the application (normal mode)
cd cmd/ascii-art && go run . "Hello" standard

# Run the application (color mode)
cd cmd/ascii-art && go run . --color=red "Hello"
cd cmd/ascii-art && go run . --color=red He "Hello"
```

## Testing Commands

```bash
# Run all tests
make test
# or: go test -v ./...

# Generate coverage report
make coverage
# or: go test -coverprofile=coverage.out ./...
```

## Code Style and Conventions

### Go Standards
- Follow standard Go formatting: `gofmt` and `goimports`
- Use `golangci-lint` with configuration in `.golangci.yml`
- Run `make check` before committing (runs fmt, vet, lint)

### Code Organization
- **Package documentation**: Every package must have a doc comment
- **Function documentation**: Exported functions must have Parameters/Returns sections
- **Constants over magic numbers**: Define const for all numeric literals
- **Error wrapping**: Use `fmt.Errorf` with `%w` for error chains
- **Error messages**: Lowercase, no ending punctuation (Go style guide ST1005)

### Naming Conventions
- Package names: lowercase, single word (parser, renderer, color, coloring, flagparser)
- Exported functions: PascalCase (RenderText, BuildCharacterMap)
- Unexported functions: camelCase (renderLine, validateInput)
- Test functions: TestFunctionName_Scenario

### Best Practices
- Use `strings.Builder` for efficient string concatenation
- Use `bufio.Scanner` for line-by-line file reading

## Testing Standards

### Test Organization
- Use table-driven tests for multiple scenarios
- Test files: `*_test.go` in same package
- Integration tests: `integration_test.go` in main package

### Coverage Requirements
- Aim for >90% overall coverage
- 100% coverage on parser and renderer packages (critical)
- main() function coverage optional (tested via integration)

### Test Patterns
```go
func TestFunctionName_Scenario(t *testing.T) {
    input := "test data"

    result, err := FunctionName(input)

    if err != nil {
        t.Errorf("Expected no error, got: %v", err)
    }
    // ... more assertions
}
```

## Project Structure

```
ascii-art-color/
├── .gitignore                 # Git ignore rules
├── .golangci.yml              # Linter configuration
├── LICENSE                    # Project license
├── Makefile                   # Build automation
├── go.mod                     # Go module file (no external deps)
├── AGENTS.md                  # This file
├── CHANGELOG.md               # Version history
├── CONTRIBUTING.md            # Contribution guidelines
├── PERMISSIONS.md             # Team permissions
├── README.md                  # User documentation
├── cmd/
│   └── ascii-art/
│       ├── main.go            # CLI entry point
│       ├── main_test.go       # Unit tests for main package
│       ├── integration_test.go # End-to-end tests
│       └── testdata/          # Banner files and test fixtures
│           ├── standard.txt
│           ├── shadow.txt
│           ├── thinkertoy.txt
│           ├── corrupted.txt  # Test fixture
│           ├── empty.txt      # Test fixture
│           └── oversized.txt  # Test fixture
└── internal/
    ├── color/                 # Color specification parsing
    │   ├── color.go
    │   └── color_test.go
    ├── coloring/              # ANSI color application to ASCII art
    │   ├── coloring.go
    │   └── coloring_test.go
    ├── flagparser/            # CLI argument validation
    │   ├── flagparser.go
    │   └── flagparser_test.go
    ├── parser/                # Banner file parsing
    │   ├── banner_parser.go
    │   └── parser_test.go
    └── renderer/              # ASCII art rendering
        ├── renderer.go
        └── renderer_test.go
```

## Security Considerations

### File Operations
- All file paths must be validated before use
- Banner files are read-only, loaded from `testdata/` directory
- Use `#nosec G304` annotation ONLY when path is validated (not user input)

### Input Validation
- Validate all user input before processing
- Support only ASCII characters 32-126 (printable characters)
- Return error for unsupported characters (do not silently skip)

### Error Handling
- Never expose internal paths in error messages
- Wrap errors with context using `fmt.Errorf`
- Use proper exit codes: 0 (success), 1 (usage), 2 (banner error), 3 (render error), 4 (color error)

## Commit Message Format

Use Conventional Commits format:

```
<type>(<scope>): <description>

[optional body]
```

**Types**: `feat`, `fix`, `docs`, `test`, `refactor`, `perf`, `chore`, `build`

**Scopes**: `parser`, `renderer`, `main`, `color`, `coloring`, `flagparser`, `docs`, `build`, `tests`

**Example**:
```
feat(parser): add validation for banner file format

Added check to ensure banner file has exactly 855 lines before parsing.
Prevents crash on malformed banner files.
```

## Build and Release

### Local Build
```bash
make build          # Current platform
make build-all      # All platforms (Linux, macOS, Windows)
```

### Version Management
- Use semantic versioning (MAJOR.MINOR.PATCH)
- Version info injected via Makefile ldflags from git tags
- Update CHANGELOG.md for all releases

## Performance Guidelines

### Optimization Rules
- DO use `strings.Builder` for string concatenation
- DO preallocate slices when size is known
- DON'T add complexity for hypothetical performance gains

## Common Tasks

### Adding a New Banner Style
1. Add banner file to `cmd/ascii-art/testdata/<name>.txt`
2. Update `GetBannerPath()` in `cmd/ascii-art/main.go` to recognize new name
3. Add integration test in `cmd/ascii-art/integration_test.go`
4. Update README.md with new banner style
5. Update CHANGELOG.md

### Fixing a Bug
1. Write failing test that reproduces the bug
2. Fix the bug
3. Verify test now passes
4. Run full test suite: `make test`
5. Run linters: `make lint`
6. Update CHANGELOG.md if user-facing

### Adding a Feature
1. Discuss approach (architectural decision)
2. Write tests first (TDD)
3. Implement feature
4. Update documentation (README, inline docs)
5. Update CHANGELOG.md

## DO NOT

- Add external dependencies (use only Go standard library)
- Modify banner files in `testdata/`
- Skip tests or reduce coverage
- Commit without running `make check`
- Use deprecated Go features
- Add TODOs or FIXMEs without GitHub issues

## Quality Checklist

Before committing, ensure:
- [ ] Code is formatted: `make fmt`
- [ ] Tests pass: `make test`
- [ ] Linters pass: `make lint`
- [ ] Coverage maintained or improved
- [ ] Documentation updated
- [ ] CHANGELOG.md updated (if applicable)
- [ ] Conventional commit message used

---

*This file follows the [AGENTS.md](https://agents.md/) open standard for guiding AI coding agents.*
