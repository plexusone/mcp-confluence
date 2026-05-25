# Setup

## Get Your Confluence Credentials

### Confluence Cloud

1. Go to [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens)
2. Create a new API token
3. Note your email address (username)
4. Note your Confluence URL (e.g., `https://example.atlassian.net/wiki`)

### Confluence Data Center

Use your regular username and password, or configure an application link.

## Configure Credentials

### Option 1: Environment Variables

```bash
export CONFLUENCE_BASE_URL="https://example.atlassian.net/wiki"
export CONFLUENCE_USERNAME="user@example.com"
export CONFLUENCE_API_TOKEN="your-api-token"
```

### Option 2: Command-Line Flags

```bash
mcp-confluence --base-url https://example.atlassian.net/wiki \
               --username user@example.com \
               --api-token your-api-token
```

### Option 3: Vault-Backed Credentials

For production use, store credentials in a vault:

```bash
# 1Password
export OP_SERVICE_ACCOUNT_TOKEN="ops_..."
mcp-confluence --vault op://MyVault --credentials-name confluence \
               --base-url https://example.atlassian.net/wiki

# Bitwarden
export BW_ACCESS_TOKEN="..."
export BW_ORGANIZATION_ID="..."
mcp-confluence --vault bw://org-id --credentials-name confluence \
               --base-url https://example.atlassian.net/wiki

# Keeper
export KSM_TOKEN="US:..."
mcp-confluence --vault keeper:// --credentials-name confluence \
               --base-url https://example.atlassian.net/wiki
```

See [Credentials](../configuration/credentials.md) for detailed configuration options.

## Test Your Setup

```bash
# Test with CLI mode - search for pages
mcp-confluence search-pages "type=page" --limit 5
```
