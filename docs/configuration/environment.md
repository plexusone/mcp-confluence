# Environment Variables

All command-line flags can be set via environment variables.

## Available Variables

### Credential Configuration

| Variable | Flag | Description |
|----------|------|-------------|
| `CONFLUENCE_BASE_URL` | `--base-url` | Confluence instance URL |
| `CONFLUENCE_USERNAME` | `--username` | Confluence username/email |
| `CONFLUENCE_API_TOKEN` | `--api-token` | Confluence API token |
| `OMNITOKEN_VAULT_URI` | `--vault` | Vault URI for credentials |
| `OMNITOKEN_CREDENTIALS_NAME` | `--credentials-name` | Name of credentials in vault |

### Vault Provider Authentication

| Variable | Description |
|----------|-------------|
| `OP_SERVICE_ACCOUNT_TOKEN` | 1Password service account token |
| `BW_ACCESS_TOKEN` | Bitwarden access token |
| `BW_ORGANIZATION_ID` | Bitwarden organization ID |
| `KSM_TOKEN` | Keeper token (format: `REGION:TOKEN`) |
| `KSM_CONFIG` | Keeper config (base64-encoded JSON) |

## Precedence

Command-line flags take precedence over environment variables.

```bash
# Environment variable is used
export CONFLUENCE_BASE_URL=https://example.atlassian.net/wiki
mcp-confluence
# Uses: https://example.atlassian.net/wiki

# Flag overrides environment
export CONFLUENCE_BASE_URL=https://example.atlassian.net/wiki
mcp-confluence --base-url https://other.atlassian.net/wiki
# Uses: https://other.atlassian.net/wiki
```

## Examples

### Direct Credentials

```bash
export CONFLUENCE_BASE_URL="https://example.atlassian.net/wiki"
export CONFLUENCE_USERNAME="user@example.com"
export CONFLUENCE_API_TOKEN="your-api-token"
mcp-confluence
```

### 1Password

```bash
export OP_SERVICE_ACCOUNT_TOKEN="ops_..."
export OMNITOKEN_VAULT_URI=op://MyVault
export OMNITOKEN_CREDENTIALS_NAME=confluence
export CONFLUENCE_BASE_URL="https://example.atlassian.net/wiki"
mcp-confluence
```

### Bitwarden

```bash
export BW_ACCESS_TOKEN="..."
export BW_ORGANIZATION_ID="..."
export OMNITOKEN_VAULT_URI=bw://org-id
export OMNITOKEN_CREDENTIALS_NAME=confluence
export CONFLUENCE_BASE_URL="https://example.atlassian.net/wiki"
mcp-confluence
```

### Keeper

```bash
export KSM_TOKEN="US:..."
export OMNITOKEN_VAULT_URI=keeper://
export OMNITOKEN_CREDENTIALS_NAME=confluence
export CONFLUENCE_BASE_URL="https://example.atlassian.net/wiki"
mcp-confluence
```

## Shell Configuration

### Bash/Zsh

Add to `~/.bashrc` or `~/.zshrc`:

```bash
# Confluence MCP Server credentials
export CONFLUENCE_BASE_URL="https://example.atlassian.net/wiki"
export CONFLUENCE_USERNAME="user@example.com"
export CONFLUENCE_API_TOKEN="your-api-token"
```

### Fish

Add to `~/.config/fish/config.fish`:

```fish
set -gx CONFLUENCE_BASE_URL "https://example.atlassian.net/wiki"
set -gx CONFLUENCE_USERNAME "user@example.com"
set -gx CONFLUENCE_API_TOKEN "your-api-token"
```
