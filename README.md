# Nightblog

> [!WARNING]  
> This project is currently in heavy development. Expect breaking changes and incomplete features.

Nightblog is a terminal application for managing and editing blog posts. It currently uses SQLite for storage and is transitioning from a file-based system.

## Requirements

- **Go**: 1.26 or higher
- **SQLite**: (Managed by the application)

## Setup

1. **Clone the repository**:
   ```bash
   git clone github.com/nightails/nightblog 
   cd nightblog
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Configuration**:
   The application automatically creates a default configuration file upon the first run at:
   `~/.config/nightblog/config.json`

   You can manually edit this file to change settings:
   - `DatabasePath`: Path to the SQLite database (default: `~/.cache/nightblog/blogs.db`).
   - `RemoteBlogsURL`: URL for remote blog synchronization (currently unused).
   - `Editor`: Your preferred editor (default: `nvim`).

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
- `new <title>`: Creates a new blog post entry in the database.
- `list`: Lists all blog posts stored in the database.

## Features
- [x] SQLite storage integration.
- [ ] Create and list blog posts via CLI (In transition).
- [ ] Terminal User Interface (TUI) with Bubble Tea (Work in progress).
- [ ] Remote synchronization.

## Project Structure

- `internal/`:
  - `blog/`: Core blog logic (creation, listing, formatting).
  - `cli/`: Command-line interface logic using Cobra.
  - `config/`: Configuration loading, parsing, and management logic.
  - `storage/`: Database initialization and SQLC-generated queries.
  - `tui/`: Planned Terminal User Interface implementation.
  - `ui/`: UI styles and components (Lipgloss).

## Dependencies
- [Cobra](https://github.com/spf13/cobra): CLI framework.
- [Bubble Tea](https://github.com/charmbracelet/bubbletea): TUI framework (In development).
- [Lipgloss](https://github.com/charmbracelet/lipgloss): Terminal styling.
- [sqlc](https://sqlc.dev/): Type-safe SQL generator.
- [modernc.org/sqlite](https://modernc.org/sqlite): CGO-free SQLite driver.

## License

This project is licensed under the [MIT License](LICENSE).
