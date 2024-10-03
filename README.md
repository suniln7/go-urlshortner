# URL Shortener with YAML and Map-Based Path Redirection

This project demonstrates a simple URL shortening service built with Go. It supports redirection using both predefined map-based paths and dynamic paths from a YAML configuration file.

## Features

- **Redirects users to specific URLs** based on path matching.
- **Supports reading redirection paths from a YAML file**.
- **Falls back to predefined map-based paths** when no YAML is provided or in case of errors.
- Gracefully handles **missing or invalid YAML files**.
  
## How it Works

The application uses two main handlers to manage URL redirections:

1. **`MapHandler`**:  
   Handles URL redirection based on a predefined map of paths to URLs. If the requested path exists in the map, the user is redirected to the corresponding URL. If the path is not found, the fallback handler is used.

2. **`YAMLHandler`**:  
   This handler reads paths and URLs from a YAML file at runtime. If the requested path exists in the YAML file, the user is redirected to the specified URL. If the YAML file is missing or contains errors, the `MapHandler` is used as a fallback.

## How to Use

### 1. YAML Path Configuration

You can specify a YAML file for path redirection via a command-line flag `-yaml`. If no file is provided, the app defaults to a `default.yaml` file in the current directory.

Example:

```bash
go run main.go -yaml=custom_paths.yaml
```

### Prerequisites

- [Go](https://golang.org/doc/install) installed.
- Install the necessary Go packages:

```bash
go get gopkg.in/yaml.v2
