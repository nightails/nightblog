# Blog Manager MVP Plan

## MVP Features
- **Create New Blog**: A command (`new`) to create a new blog post file and open it in a terminal-based editor (vim/neovim).
- **Edit Existing Blog**: A command (`edit`) to open an existing blog post file in the configured editor.
- **List Blogs**: A command (`list`) to display all blog posts in the library.

## Configuration
The application will use a configuration file (e.g., `config.json` or `config.yaml`) with the following settings:
- `blogs_dir`: The path to the directory where blog files are stored.
- `editor`: The preferred terminal editor (e.g., `vim`, `nvim`).

## CLI Design
The app will be invoked via the terminal with the following commands:
- `nightblog new <title>`: Creates a new markdown file in `blogs_dir` and opens it.
- `nightblog edit <title>`: Opens an existing blog post from `blogs_dir` for editing.
- `nightblog list`: Lists all files in `blogs_dir`.

## Go Implementation Structure

### 1. Project Setup
- [x] Initialize the Go module: `go mod init nightblog`.
- [x] Use standard library packages: `os`, `path/filepath`, `os/exec`, `encoding/json`.

### 2. Configuration Management
- [x] **Struct Definition**: Create a `Config` struct with fields for `BlogsDir` and `Editor`.
- [x] **Config Path**: Use `os.UserConfigDir()` to find the standard configuration directory for the OS.
- [x] **Persistence**: Use `encoding/json` to read from and write to the config file.
- [x] **Initialization**: If the config doesn't exist, create it with sensible defaults (e.g., `editor: vim`).

### 3. CLI Command Handling
- [ ] **Dispatching**: Use `os.Args` to identify the subcommand (`new`, `edit`, `list`).
- [ ] **Argument Parsing**: Extract the blog title for `new` and `edit` commands from `os.Args[2:]`.

### 4. Feature Implementation
- [ ] **List Command**:
  - [ ] Use `os.ReadDir` on the configured `BlogsDir`.
  - [ ] Iterate and print filenames (optionally filtering for `.md` files).
- [ ] **New/Edit Command**:
  - [ ] **Path Resolution**: Join `BlogsDir` with the filename using `filepath.Join`.
  - [ ] **File Creation**: For `new`, check if the file exists; if not, create it using `os.Create`.
  - [ ] **Terminal Editor Integration**:
    - [ ] Use `os/exec.Command` to start the configured editor with the file path.
    - [ ] **Important**: Connect `cmd.Stdin`, `cmd.Stdout`, and `cmd.Stderr` to the current process's standard streams to allow the interactive editor to take over the terminal.
    - [ ] Use `cmd.Run()` to wait for the editor to close.
