# URL Shortener in Go

This repository contains a simple URL shortener service written in Go. It allows you to map specific paths to external URLs and perform redirections based on those mappings. You can define the URL mappings in a **hardcoded map** or via a **YAML file**.

## Features

- **Path-based URL redirection**: Shorten URLs by redirecting specific paths to their corresponding external URLs.
- **Fallback support**: Redirects to a default handler if no matching path is found.
- **YAML configuration**: Easily add URL mappings through a YAML file for more flexible configuration.

## How It Works

1. **Map Handler**: This handler redirects a predefined list of URL paths to corresponding external URLs.
2. **YAML Handler**: If no match is found in the hardcoded map, this handler checks a list of paths specified in a YAML file and redirects accordingly.
3. **Default Handler**: For paths not matched in either the map or YAML file, the server falls back to a default handler that simply returns a "Hello, world!" message.

## How to Run

### Prerequisites

- [Go](https://golang.org/doc/install) installed.
- Install the necessary Go packages:

```bash
go get gopkg.in/yaml.v2