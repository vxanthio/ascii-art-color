# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Parser package for reading and parsing ASCII art banner files
  - `ReadBannerFile()` function to read banner files
  - `ParseCharacters()` function to parse 95 ASCII characters (32-126)
  - `BuildCharacterMap()` convenience function
  - Comprehensive error handling with wrapped errors
  - Security annotations for file operations
- Renderer package for converting text to ASCII art
  - `RenderText()` main rendering function with newline support
  - `renderLine()` for rendering single lines
  - `validateInput()` for character validation
  - Efficient string building with `strings.Builder`
  - Empty string optimization
- Comprehensive test suite
  - 15 parser unit tests with 100% coverage
  - 12 renderer unit tests with 100% coverage
  - Integration tests for end-to-end functionality
  - Table-driven tests for multiple scenarios
  - 79.8% overall test coverage
- Benchmark suite with 16 benchmarks
  - Parser benchmarks (5 tests)
  - Renderer benchmarks (11 tests)
  - Performance metrics documentation
- Professional Makefile with 30+ targets
  - Quality control targets (fmt, vet, lint, check)
  - Development targets (run, build, install)
  - Testing targets (test, coverage, bench)
  - Cross-compilation for Linux, macOS, Windows
  - CI/CD targets (ci, pre-commit)
  - Utility targets (tidy, version, help, clean)
- Comprehensive documentation
  - Professional README.md with usage examples
  - AGENTS.md for AI coding agents
  - CONTRIBUTING.md with development guidelines
  - PERMISSIONS.md for team workflow
- golangci-lint configuration
  - 20+ enabled linters
  - Comprehensive code quality checks
  - Test-specific exclusions

### Changed
- Applied Go best practices throughout codebase
  - Package-level documentation
  - Constants for magic numbers
  - Error wrapping with `fmt.Errorf`
  - Lowercase error messages (Go style guide)
- Modern Go 1.22+ features
  - Range over int syntax
  - Latest package organization

### Fixed
- All linting errors and warnings
- Test coverage for edge cases
- Code formatting consistency

### Performance
- Sub-millisecond rendering for typical use cases
  - Single word ("Hello"): ~83 µs (12,000 ops/sec)
  - Sentence (42 chars): ~143 µs (7,000 ops/sec)
  - Paragraph (139 chars): ~506 µs (2,000 ops/sec)
- Parser optimizations
  - ReadBannerFile: ~42 µs
  - ParseCharacters: ~30 µs
  - BuildCharacterMap: ~78 µs
- Renderer optimizations
  - Empty string: 78.8 ns
  - Linear scaling with text length
  - Efficient memory allocation

## [1.0.0] - YYYY-MM-DD

### Added
- Initial release
- Support for three banner styles (standard, shadow, thinkertoy)
- Command-line interface for text to ASCII art conversion
- Newline support in input text
- Cross-platform compatibility
- Zero external dependencies

---

## Release Guidelines

### Version Format
This project uses [Semantic Versioning](https://semver.org/):
- **MAJOR** version for incompatible API changes
- **MINOR** version for added functionality (backwards compatible)
- **PATCH** version for backwards compatible bug fixes

### Release Types

#### Major Release (X.0.0)
- Breaking changes to CLI interface
- Incompatible changes to banner file format
- Removal of features or flags

#### Minor Release (0.X.0)
- New banner styles
- New CLI flags or options (backwards compatible)
- Performance improvements
- New features

#### Patch Release (0.0.X)
- Bug fixes
- Documentation updates
- Security patches
- Minor performance improvements

### Unreleased Section
The `[Unreleased]` section tracks changes that are committed but not yet released:
- Use for active development
- Move to appropriate version section on release
- Keep organized by change type

### Change Categories

Changes should be grouped into these categories:

- **Added** - New features
- **Changed** - Changes to existing functionality
- **Deprecated** - Soon-to-be removed features
- **Removed** - Removed features
- **Fixed** - Bug fixes
- **Security** - Security vulnerability fixes
- **Performance** - Performance improvements

### Example Entry Format

```markdown
## [1.2.0] - 2024-01-15

### Added
- New "graffiti" banner style (#42)
- Support for color output with --color flag (#38)
- Verbose mode with -v flag for debugging (#45)

### Changed
- Improved error messages for invalid characters (#40)
- Updated help text with more examples (#43)

### Fixed
- Fixed rendering issue with consecutive newlines (#41)
- Corrected alignment for wide characters (#44)

### Performance
- 20% faster rendering for long texts (#39)
- Reduced memory allocations in parser (#46)
```

---

## Links

- [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
- [Semantic Versioning](https://semver.org/spec/v2.0.0.html)
- [Conventional Commits](https://www.conventionalcommits.org/)
