# Blogs Manager

A simple terminal application for managing nightails' blogs.

## Overview

Blogs Manager is a command-line utility written in Go designed to help manage blog content locally. It handles configuration management and ensures the local library directory exists.

## Requirements

- **Go**: 1.26 or higher
- **Git**: (Optional, for managing remote library)

## Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd blogs_manager
   ```

2. **Install dependencies**:
   This project uses only the Go standard library, so no external dependencies are required.
   ```bash
   go mod download
   ```

3. **Configuration**:
   The application automatically creates a default configuration file at:
   `~/.config/blogs_manager/config.json`

   You can manually edit this file to change settings:
   - `LocalLibraryPath`: Where your blogs are stored (default: `~/Documents/Blogs`).
   - `RemoteLibraryURL`: URL for remote blog synchronization (TODO).
   - `DefaultTextEditor`: Your preferred editor (default: `nvim`).

## Usage

### Run the application
To run the application directly:
```bash
go run main.go
```

### Build the application
To compile the binary:
```bash
go build -o blogs_manager main.go
```

## Scripts
- TODO: Add build/test scripts (e.g., Makefile).

## Environment Variables
- None currently used. Configuration is handled via `config.json`.

## Tests
- TODO: Implement unit and integration tests.

## Project Structure

- `main.go`: Application entry point.
- `internal/config/`: Configuration loading and management logic.
- `internal/library/`: Library directory verification and creation.
- `go.mod`: Go module definition.

## License
- TODO: Add license information.