# GVM - Go Version Manager

GVM is a tool for managing multiple Go versions. It supports Windows, Linux, and macOS systems.

## Features

- Install multiple Go versions
- Switch between different Go versions
- List installed Go versions
- Uninstall unnecessary Go versions

## Installation

Download the appropriate gvm executable for your system and place it in your environment's PATH.


## Usage

```bash
# Install a specific Go version
gvm install 1.24.0

# Switch to a specific Go version
gvm use 1.24.0

# List installed Go versions
gvm list

# Uninstall a specific Go version
gvm uninstall 1.24.0

# Display help information
gvm help
```

## Compile the Project

```bash
go build -o gvm main.go
```

## Core Directory Functionality

- `$GVM_ROOT`：GVM working directory.
- `$GVM_ROOT/versions`：Go version installation directory. Each version of Go is installed in a separate subdirectory.
- `$GVM_ROOT/archives`: Go version archive files.
- `$GVM_ROOT/current`: Go current executable directory.

current directory is a symbolic link pointing to the executable directory of the currently used Go version in the versions directory.

## Env Variables

- `GVM_ROOT`: GVM working directory.

## Configuration

### Linux & MacOS

Add the source lines from the snippet below to the correct profile file (~/.bashrc, ~/.bash_profile, ~/.zshrc, or ~/.profile):

```bash
export GVM_ROOT="$HOME/.gvm" 
export PATH=$PATH:$GVM_ROOT/current/bin
```

### Windows

Add the following lines to your environment variables:

- `GVM_ROOT`: Add the GVM working directory to the environment variables. (default: `%APPDATA%\gvm`)
- `PATH`: Add `%GVM_ROOT%\current\bin` to the PATH variable.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes.

## License

MIT License