# Nightblog

> [!WARNING]  
> This project is currently in heavy development. Expect breaking changes and incomplete features.

Nightblog is a terminal application for managing and editing blog posts. It handles creating and editing blog posts with predefined frontmatter compatible with Astro.

## Requirements

- **Go**: 1.26 or higher
- **Git**: (Optional, for managing remote library)

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
   - `LocalBlogsDir`: Where your blogs are stored (default: `~/Documents/Blogs`).
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
- `new <title>`: Creates a new blog post file with basic frontmatter in your `LocalBlogsDir`.
- `list`: Lists all blog posts in your `LocalBlogsDir`.

## Features
- [x] Create new blog posts with frontmatter.
- [x] List existing blog posts.
- [ ] Edit existing blog posts.
- [ ] Terminal User Interface (TUI) with Bubble Tea (Work in progress).
- [ ] Remote synchronization.

## Project Structure

- `internal/`:
  - `blog/`: Core blog logic (creation, listing, formatting).
  - `cli/`: Command-line interface logic using Cobra.
  - `config/`: Configuration loading, parsing, and management logic.
  - `storage/`: Storage and directory initialization.
  - `tui/`: Planned Terminal User Interface implementation.
  - `ui/`: UI styles and components (Lipgloss).

## Dependencies
- [Cobra](https://github.com/spf13/cobra): CLI framework.
- [Bubble Tea](https://github.com/charmbracelet/bubbletea): TUI framework (In development).
- [Lipgloss](https://github.com/charmbracelet/lipgloss): Terminal styling.

## License

This project is licensed under the [MIT License](LICENSE).
