# Claude Desktop Configuration

Configure Claude Desktop to use the Confluence MCP Server.

## Configuration File Location

| OS | Path |
|----|------|
| macOS | `~/Library/Application Support/Claude/claude_desktop_config.json` |
| Windows | `%APPDATA%\Claude\claude_desktop_config.json` |
| Linux | `~/.config/Claude/claude_desktop_config.json` |

## Basic Configuration

### With Direct Credentials

```json
{
  "mcpServers": {
    "confluence": {
      "command": "/path/to/mcp-confluence",
      "env": {
        "CONFLUENCE_BASE_URL": "https://example.atlassian.net/wiki",
        "CONFLUENCE_USERNAME": "user@example.com",
        "CONFLUENCE_API_TOKEN": "your-api-token"
      }
    }
  }
}
```

### With 1Password

```json
{
  "mcpServers": {
    "confluence": {
      "command": "/path/to/mcp-confluence",
      "env": {
        "CONFLUENCE_BASE_URL": "https://example.atlassian.net/wiki",
        "OP_SERVICE_ACCOUNT_TOKEN": "ops_...",
        "OMNITOKEN_VAULT_URI": "op://MyVault",
        "OMNITOKEN_CREDENTIALS_NAME": "confluence"
      }
    }
  }
}
```

### With Bitwarden

```json
{
  "mcpServers": {
    "confluence": {
      "command": "/path/to/mcp-confluence",
      "env": {
        "CONFLUENCE_BASE_URL": "https://example.atlassian.net/wiki",
        "BW_ACCESS_TOKEN": "...",
        "BW_ORGANIZATION_ID": "...",
        "OMNITOKEN_VAULT_URI": "bw://org-id",
        "OMNITOKEN_CREDENTIALS_NAME": "confluence"
      }
    }
  }
}
```

### With Keeper

```json
{
  "mcpServers": {
    "confluence": {
      "command": "/path/to/mcp-confluence",
      "env": {
        "CONFLUENCE_BASE_URL": "https://example.atlassian.net/wiki",
        "KSM_TOKEN": "US:...",
        "OMNITOKEN_VAULT_URI": "keeper://",
        "OMNITOKEN_CREDENTIALS_NAME": "confluence"
      }
    }
  }
}
```

## Environment Variables Reference

| Variable | Description |
|----------|-------------|
| `CONFLUENCE_BASE_URL` | Confluence instance URL |
| `CONFLUENCE_USERNAME` | Confluence username/email |
| `CONFLUENCE_API_TOKEN` | Confluence API token |
| `OMNITOKEN_VAULT_URI` | Vault URI (e.g., `op://MyVault`) |
| `OMNITOKEN_CREDENTIALS_NAME` | Credential name in vault (default: `confluence`) |
| `OP_SERVICE_ACCOUNT_TOKEN` | 1Password service account token |
| `BW_ACCESS_TOKEN` | Bitwarden access token |
| `BW_ORGANIZATION_ID` | Bitwarden organization ID |
| `KSM_TOKEN` | Keeper token (format: `REGION:TOKEN`) |

## Multiple Servers

You can run multiple MCP servers alongside Confluence:

```json
{
  "mcpServers": {
    "confluence": {
      "command": "/path/to/mcp-confluence",
      "env": {
        "CONFLUENCE_BASE_URL": "https://example.atlassian.net/wiki",
        "CONFLUENCE_USERNAME": "user@example.com",
        "CONFLUENCE_API_TOKEN": "token"
      }
    },
    "google": {
      "command": "/path/to/mcp-google",
      "env": {
        "GOOGLE_CREDENTIALS_FILE": "/path/to/service-account.json"
      }
    },
    "aha": {
      "command": "/path/to/mcp-aha",
      "env": {
        "AHA_DOMAIN": "mycompany",
        "AHA_API_TOKEN": "token"
      }
    }
  }
}
```

## Troubleshooting

### Server Not Starting

Check the Claude Desktop logs:

- macOS: `~/Library/Logs/Claude/`
- Windows: `%APPDATA%\Claude\logs\`

Common issues:

1. **Binary not found**: Verify the path is correct
2. **Credentials not found**: Check environment variables
3. **Permission denied**: Ensure the binary is executable (`chmod +x`)
4. **Invalid URL**: Check `CONFLUENCE_BASE_URL` format

### Verifying Configuration

Test the server manually:

```bash
# Should start and wait for input (Ctrl+C to exit)
/path/to/mcp-confluence --base-url https://example.atlassian.net/wiki \
                        --username user@example.com \
                        --api-token your-token
```

### JSON Syntax Errors

Validate your JSON:

```bash
# On macOS/Linux
cat ~/Library/Application\ Support/Claude/claude_desktop_config.json | python3 -m json.tool
```

## Available Tools in Claude

Once configured, you can ask Claude to:

- "Read Confluence page 12345"
- "Search for pages about authentication"
- "Create a new page in TEAM space"
- "Update page 12345 with a status table"
- "Delete page 67890"
