# Installation

## Requirements

- Go 1.24 or later
- A Confluence Cloud or Data Center instance
- Confluence API token (Cloud) or credentials (Data Center)

## Install from Source

```bash
go install github.com/plexusone/mcp-confluence/cmd/mcp-confluence@latest
```

## Build from Source

```bash
git clone https://github.com/plexusone/mcp-confluence.git
cd mcp-confluence
go build ./cmd/mcp-confluence
```

## Verify Installation

```bash
mcp-confluence version
```

## As a Library

You can also use the packages directly:

```bash
go get github.com/plexusone/mcp-confluence
```

### Available Packages

| Package | Description |
|---------|-------------|
| `storage` | IR types, render, parse, validate for Confluence Storage Format |
| `confluence` | REST API client with IR integration |
| `mcpserver` | MCP server implementation |
