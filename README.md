# Nightblog

A simple terminal application for managing nightails' blogs.

## Overview

Nightblog is a command-line utility written in Go designed to help manage blog content locally. It handles configuration management, ensures the local library directory exists, and allows you to quickly create new blog posts with predefined frontmatter.

## Requirements

- **Go**: 1.26 or higher
- **Git**: (Optional, for managing remote library)

## Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd nightblog
   ```

2. **Install dependencies**:
   This project uses only the Go standard library, so no external dependencies are required.
   ```bash
   go mod download
   ```

3. **Configuration**:
   The application automatically creates a default configuration file upon the first run at:
   `~/.config/nightblog/config.json`

   You can manually edit this file to change settings:
   - `LocalLibraryPath`: Where your blogs are stored (default: `~/Documents/Blogs`).
   - `RemoteLibraryURL`: URL for remote blog synchronization (TODO).
   - `DefaultTextEditor`: Your preferred editor (default: `nvim`).

## Usage

### Run the application
To run the application directly:
```bash
go run main.go [command]
```

### Build the application
To compile the binary:
```bash
go build -o nightblog main.go
```

### Available Commands
- `new <title>`: Creates a new blog post file with basic frontmatter in your `LocalLibraryPath`.

## Scripts
- TODO: Add build/test scripts (e.g., Makefile).

## Environment Variables
- None currently used. Configuration is handled via `config.json`.

## Tests
- TODO: Implement unit and integration tests.

## Project Structure

- `main.go`: Application entry point and initialization logic.
- `internal/cli/`: Command-line interface logic and command implementations.
  - `cli.go`: Main CLI router.
  - `cmd_new.go`: Implementation of the `new` command.
- `internal/config/`: Configuration loading, parsing, and management logic.
- `internal/library/`: Library directory verification and creation utilities.
- `go.mod`: Go module definition.

## License
- TODO: Add license information.
