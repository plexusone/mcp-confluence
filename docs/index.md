# Confluence MCP Server

An MCP server for Confluence with safe handling of Confluence Storage Format (XHTML).

## The Problem

When AI assistants interact with Confluence via MCP servers, they often corrupt pages - especially tables - because:

1. LLMs generate Markdown or HTML5 instead of Confluence Storage XHTML
2. Tables require specific structure (`<tbody>`, no `<thead>`)
3. Macros need `ac:` namespaces
4. Round-tripping through incorrect formats causes data loss

## The Solution

This library provides:

- **Structured IR (Intermediate Representation)**: Work with Go types instead of raw XHTML
- **Safe Rendering**: IR to valid Storage XHTML with proper structure
- **Validation**: Catch forbidden tags, missing `<tbody>`, etc. before API calls
- **MCP Server**: Tools that accept structured JSON, never raw XHTML

## Features

- **8 MCP tools** for reading, creating, updating, and searching Confluence pages
- **Structured content blocks** - JSON instead of raw XHTML
- **Safe table handling** - Proper `<tbody>` structure, no `<thead>`
- **Macro support** - Status badges, info panels, code blocks
- **Vault-backed credentials** - 1Password, Bitwarden, Keeper support

## Available Tools

| Tool | Description |
|------|-------------|
| `confluence_read_page` | Read a page as structured blocks |
| `confluence_read_page_xhtml` | Read a page as raw Storage Format XHTML |
| `confluence_update_page` | Update a page with structured blocks |
| `confluence_update_page_xhtml` | Update a page with raw Storage Format XHTML |
| `confluence_create_page` | Create a new page with structured blocks |
| `confluence_create_table` | Create a table block from structured data |
| `confluence_delete_page` | Delete a page |
| `confluence_search_pages` | Search pages using CQL |

## Quick Start

```bash
# Install
go install github.com/plexusone/mcp-confluence/cmd/mcp-confluence@latest

# Configure credentials
export CONFLUENCE_BASE_URL="https://example.atlassian.net/wiki"
export CONFLUENCE_USERNAME="user@example.com"
export CONFLUENCE_API_TOKEN="your-api-token"

# Run as MCP server
mcp-confluence

# Or use CLI mode
mcp-confluence read-page 12345
```

## Next Steps

- [Installation](getting-started/installation.md) - Install the server
- [Setup](getting-started/setup.md) - Configure your credentials
- [Quick Start](getting-started/quickstart.md) - Start using the tools
- [Tools Reference](tools/overview.md) - Detailed tool documentation
- [Storage Format](storage-format/overview.md) - Understanding block types
