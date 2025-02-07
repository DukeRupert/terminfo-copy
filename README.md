# Terminfo Copy Utility

A simple command-line utility to copy Terminfo configurations to remote servers. This tool streamlines the process of setting up terminal configurations across multiple machines, particularly useful for users of modern terminal emulators.

## Purpose

When working with modern terminal emulators, having the correct Terminfo configuration on remote servers is crucial for proper terminal behavior. This utility automates the process of copying your local Terminfo configuration to remote servers, helping prevent common issues with terminal compatibility.

## Installation

To install the utility, ensure you have Go installed on your system, then run:

```bash
go install github.com/YOUR_USERNAME/terminfo-copy@latest
```

Or clone the repository and use the provided Makefile:

```bash
git clone https://github.com/YOUR_USERNAME/terminfo-copy.git
cd terminfo-copy

# Run directly
make run

# Or install globally
make install
```

## Usage

After installation, you can run the utility from any directory by typing:

```bash
terminfo-copy
```

If you've installed using `make install`, ensure that your Go bin directory (typically `~/go/bin`) is in your system's PATH.

The tool will:
1. Prompt you for the remote server address (e.g., user@hostname)
2. Extract your local Terminfo configuration
3. Copy and install it on the remote server

## Requirements

- Go 1.19 or later
- SSH access to the remote server
- `infocmp` command available locally
- Write permissions in the Terminfo directory on the remote server

## How It Works

The utility executes the following command sequence:
```bash
infocmp -x | ssh YOUR-SERVER -- tic -x -
```

This extracts your local terminal information using `infocmp` and pipes it through SSH to the remote server, where `tic` installs it in the appropriate location.

## Acknowledgments

This utility was inspired by the Terminfo installation documentation from the [Ghostty terminal emulator](https://ghostty.org/docs/help/terminfo). The approach of using `infocmp` and `tic` for Terminfo transfer was adapted from their documentation.

## License

MIT License - See LICENSE file for details
