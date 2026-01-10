# Contributing to ASCII Art Generator

Thank you for your interest in contributing to the ascii-art project! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Coding Standards](#coding-standards)
- [Testing Guidelines](#testing-guidelines)
- [Commit Message Format](#commit-message-format)
- [Pull Request Process](#pull-request-process)
- [Project Structure](#project-structure)
- [Common Tasks](#common-tasks)

## Code of Conduct

This project is part of the Zone01 curriculum. We expect all contributors to:

- Be respectful and constructive in discussions
- Focus on what is best for the project and community
- Show empathy towards other contributors
- Accept constructive criticism gracefully

## Getting Started

### Prerequisites

- Go 1.22.2 or higher
- make (optional but recommended)
- golangci-lint v2+ (for code quality checks)

### Installation

1. **Fork and clone the repository**
   ```bash
   git clone https://github.com/yourusername/ascii-art.git
   cd ascii-art
   ```

2. **Verify Go installation**
   ```bash
   go version  # Should be 1.22.2 or higher
   ```

3. **Run tests to ensure everything works**
   ```bash
   make test
   ```

4. **Install golangci-lint v2 (required for team compatibility)**
   ```bash
   go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
   ```

   Verify installation:
   ```bash
   golangci-lint --version  # Should show v2.x.x
   ```

## Development Workflow

### 1. Create a Feature Branch

```bash
git checkout -b feature/your-feature-name
# or: git checkout -b fix/bug-description
```

Branch naming conventions:
- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation updates
- `test/` - Test improvements
- `refactor/` - Code refactoring
- `perf/` - Performance improvements

### 2. Make Your Changes

Follow the [Coding Standards](#coding-standards) section below.

### 3. Run Quality Checks

Before committing, always run:

```bash
make check    # Runs fmt, vet, and lint
make test     # Runs all tests
make coverage # Generates coverage report
```

Or run individually:
```bash
make fmt      # Format code
make vet      # Run go vet
make lint     # Run golangci-lint
```

### 4. Commit Your Changes

Follow the [Commit Message Format](#commit-message-format) section.

```bash
git add .
git commit -m "feat(parser): add support for new banner style"
```

### 5. Push and Create Pull Request

```bash
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub.

## Coding Standards

### Go Style Guide

This project follows the official Go style guide and best practices:

1. **Formatting**
   - Use `gofmt` for formatting (enforced)
   - Use `goimports` for import organization (auto-adds/removes imports, organizes them)
   - Line length: 120 characters max

2. **Naming Conventions**
   - Package names: lowercase, single word (`parser`, `renderer`)
   - Exported identifiers: PascalCase (`RenderText`, `BuildCharacterMap`)
   - Unexported identifiers: camelCase (`renderLine`, `validateInput`)
   - Constants: PascalCase or ALL_CAPS for clarity

3. **Documentation**
   - Every package must have a package comment
   - Exported functions must have comments
   - Comment format: `// FunctionName does...`

4. **Error Handling**
   - Always check errors
   - Wrap errors with context: `fmt.Errorf("context: %w", err)`
   - Error messages: lowercase, no ending punctuation

5. **Code Organization**
   - Use constants for magic numbers
   - Group related functionality
   - Keep functions focused and small

### Project-Specific Rules

1. **No External Dependencies**
   - Only use Go standard library
   - Exception: Development tools (linters, etc.)

2. **Security**
   - Validate all file paths before use (never trust user input directly)
   - Use `#nosec` annotations only when path is validated (e.g., through banner name mapping)
   - Never expose internal paths in error messages (keep errors generic)

3. **Performance**
   - Use `strings.Builder` for string concatenation
   - Preallocate slices when size is known
   - Write benchmarks for performance-critical code

## Testing Guidelines

### Test Coverage Requirements

- **Overall**: Aim for >90% coverage
- **Parser package**: Must maintain 100% coverage
- **Renderer package**: Must maintain 100% coverage
- **Main package**: Integration tests required

### Writing Tests

1. **Test File Naming**
   - Unit tests: `*_test.go` (same package)
   - Benchmark tests: `*_bench_test.go`
   - Integration tests: `integration_test.go`

2. **Test Function Naming**
   ```go
   func TestFunctionName_Scenario(t *testing.T)
   func BenchmarkFunctionName(b *testing.B)
   ```

3. **Test Structure**
   ```go
   func TestParseArgs_ValidInput(t *testing.T) {
       args := []string{"prog", "Hello", "standard"}

       text, banner, err := ParseArgs(args)

       if err != nil {
           t.Errorf("Expected no error, got: %v", err)
       }
       if text != "Hello" {
           t.Errorf("Expected 'Hello', got: %q", text)
       }
   }
   ```

4. **Table-Driven Tests**
   Use for testing multiple scenarios:
   ```go
   func TestFunction_MultipleScenarios(t *testing.T) {
       tests := []struct {
           name     string
           input    string
           expected string
           wantErr  bool
       }{
           {"valid input", "test", "result", false},
           {"invalid input", "", "", true},
       }

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               result, err := Function(tt.input)
               if (err != nil) != tt.wantErr {
                   t.Errorf("wantErr %v, got %v", tt.wantErr, err)
               }
               if result != tt.expected {
                   t.Errorf("expected %q, got %q", tt.expected, result)
               }
           })
       }
   }
   ```

### Running Tests

```bash
# All tests
make test

# With coverage
make coverage
make coverage-view  # Opens HTML report

# Specific package
go test ./parser -v

# Specific test
go test -run TestParseBannerFile
```

## Commit Message Format

We use [Conventional Commits](https://www.conventionalcommits.org/) format:

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `test`: Adding or updating tests
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `chore`: Maintenance tasks
- `build`: Build system changes

### Scopes

- `parser`: Parser package
- `renderer`: Renderer package
- `main`: Main package
- `tests`: Test-related
- `docs`: Documentation
- `build`: Build/tooling

### Examples

```
feat(parser): add support for UTF-8 characters

Add UTF-8 character support while maintaining ASCII compatibility.
```

```
fix(parser): handle truncated banner files gracefully

Fixed crash when banner file has fewer than expected lines.
```

```
docs(readme): update installation instructions

Added instructions for installing via go install.
```

**Note:** Footer lines like "Closes #42" are optional and only used when you have linked Jira issues or GitHub issues to track features/bugs. For this project, they're typically not needed unless explicitly using issue tracking.

## Pull Request Process

### Before Submitting

1. **Ensure all checks pass**
   ```bash
   make check  # fmt, vet, lint
   make test   # all tests
   ```

2. **Update documentation**
   - Update README.md if adding features
   - Update CHANGELOG.md
   - Add inline code comments

3. **Check coverage**
   ```bash
   make coverage
   # Ensure coverage didn't decrease
   ```

### PR Template

When creating a PR, include:

1. **Description**
   - What does this PR do?
   - Why is this change needed?

2. **Type of Change**
   - [ ] Bug fix
   - [ ] New feature
   - [ ] Breaking change
   - [ ] Documentation update

3. **Checklist**
   - [ ] Tests pass (`make test`)
   - [ ] Linters pass (`make lint`)
   - [ ] Code formatted (`make fmt`)
   - [ ] Documentation updated
   - [ ] CHANGELOG.md updated

4. **Testing**
   - How was this tested?
   - What test cases were added?

### Review Process

1. At least one approval required
2. No merge conflicts
3. Code follows style guide
4. All tests pass locally
5. Tests demonstrate functionality

## Project Structure

```
ascii-art/
├── .gitignore             # Git ignore rules
├── .golangci.yml          # Linter configuration
├── LICENSE                # Project license
├── Makefile               # Build automation
├── go.mod                 # Go module file
├── main.go                # CLI entry point
├── integration_test.go    # End-to-end tests
├── main_test.go           # Unit tests for main
├── parser/                # Banner file parser
│   ├── banner_parser.go
│   └── parser_test.go
├── renderer/              # ASCII art renderer
│   ├── renderer.go
│   └── renderer_test.go
├── testdata/              # Banner files and test fixtures
│   ├── standard.txt       # Standard banner
│   ├── shadow.txt         # Shadow banner
│   ├── thinkertoy.txt     # Thinkertoy banner
│   ├── corrupted.txt      # Test fixture
│   ├── empty.txt          # Test fixture
│   └── oversized.txt      # Test fixture
└── Documentation/
    ├── README.md          # User documentation
    ├── AGENTS.md          # AI agent instructions
    ├── CHANGELOG.md       # Version history
    ├── CONTRIBUTING.md    # This file
    └── PERMISSIONS.md     # Team permissions
```

## Common Tasks

### Adding a New Feature

1. **Discuss the feature**
   - Open an issue first for significant features
   - Discuss implementation approach

2. **Write tests first (TDD)**
   ```bash
   # Create test file
   touch feature_test.go
   # Write failing tests
   # Run: make test (should fail)
   ```

3. **Implement the feature**
   ```bash
   # Write implementation
   # Run: make test (should pass)
   ```

4. **Update documentation**
   - README.md
   - CHANGELOG.md (add entry under "Added" section)
   - Inline comments

### Fixing a Bug

1. **Reproduce the bug**
   - Write a failing test that demonstrates the bug

2. **Fix the bug**
   - Make minimal changes to fix the issue

3. **Verify the fix**
   ```bash
   make test
   make lint
   ```

4. **Update CHANGELOG.md**
   - Add entry under "Fixed" section

## Questions or Issues?

- **Bugs**: Open an issue with reproduction steps
- **Features**: Open an issue to discuss before implementing
- **Questions**: Open a discussion or issue

## License

This project is part of the Zone01 curriculum.

---

Thank you for contributing to ascii-art! 🎨
