// Command mcp-confluence runs a Confluence MCP server that exposes tools for
// reading, creating, and updating Confluence pages using structured content blocks.
// It can also be used as a CLI tool for testing and scripting.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	runtime "github.com/plexusone/omniskill/mcp/server"
	"github.com/plexusone/omnitoken"
	"github.com/spf13/cobra"

	// Register desktop vault providers (1Password, etc.)
	_ "github.com/plexusone/omnivault-desktop"

	"github.com/plexusone/mcp-confluence/confluence"
	confskill "github.com/plexusone/mcp-confluence/skills/confluence"
)

const (
	serverName    = "mcp-confluence"
	serverVersion = "v0.3.0"
)

var (
	// Credential flags (persistent across all commands)
	baseURL         string
	username        string
	apiToken        string
	vaultURI        string
	credentialsName string

	// Output format flag
	outputFormat string

	// search-pages flags
	searchLimit int
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "mcp-confluence",
	Short: "MCP server and CLI for Confluence",
	Long: `An MCP (Model Context Protocol) server for reading, creating, and updating
Confluence pages using structured content blocks.
Can also be used as a CLI tool for testing and scripting.

Running without a subcommand starts the MCP server (default behavior).

Credentials can be provided via:
  - Direct credentials (base URL, username, API token)
  - Vault-backed credentials via omnitoken`,
	Example: `  # Start MCP server (default)
  mcp-confluence --base-url https://example.atlassian.net/wiki \
                 --username user@example.com --api-token your-token

  # CLI: Read a page
  mcp-confluence read-page 12345 --base-url https://example.atlassian.net/wiki \
                                 --username user@example.com --api-token your-token

  # CLI: Search for pages
  mcp-confluence search-pages "type=page AND text~authentication" \
                              --base-url https://example.atlassian.net/wiki \
                              --username user@example.com --api-token your-token`,
	SilenceUsage: true,
	RunE:         runServer,
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the MCP server",
	Long:  "Start the MCP server using stdio transport for communication with MCP clients.",
	RunE:  runServer,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s\n", serverName, serverVersion)
	},
}

// CLI commands for Confluence tools
var readPageCmd = &cobra.Command{
	Use:   "read-page <page-id>",
	Short: "Read a page as structured blocks",
	Long:  "Read a Confluence page and return its content as structured blocks.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runTool("read_page", map[string]any{
			"page_id": args[0],
		})
	},
}

var readPageXHTMLCmd = &cobra.Command{
	Use:   "read-page-xhtml <page-id>",
	Short: "Read a page as raw XHTML",
	Long:  "Read a Confluence page and return its content as raw XHTML.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runTool("read_page_xhtml", map[string]any{
			"page_id": args[0],
		})
	},
}

var searchPagesCmd = &cobra.Command{
	Use:   "search-pages <cql>",
	Short: "Search pages using CQL",
	Long:  "Search for Confluence pages using CQL (Confluence Query Language).",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{
			"cql": args[0],
		}
		if searchLimit > 0 {
			params["limit"] = searchLimit
		}
		return runTool("search_pages", params)
	},
}

var deletePageCmd = &cobra.Command{
	Use:   "delete-page <page-id>",
	Short: "Delete a page",
	Long:  "Delete a Confluence page by its ID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runTool("delete_page", map[string]any{
			"page_id": args[0],
		})
	},
}

func init() {
	// Persistent flags (available to all commands)
	rootCmd.PersistentFlags().StringVar(&baseURL, "base-url", "",
		"Confluence base URL (env: CONFLUENCE_BASE_URL)")
	rootCmd.PersistentFlags().StringVar(&username, "username", "",
		"Confluence username/email (env: CONFLUENCE_USERNAME)")
	rootCmd.PersistentFlags().StringVar(&apiToken, "api-token", "",
		"Confluence API token (env: CONFLUENCE_API_TOKEN)")
	rootCmd.PersistentFlags().StringVar(&vaultURI, "vault", "",
		"vault URI for credentials (env: OMNITOKEN_VAULT_URI)")
	rootCmd.PersistentFlags().StringVar(&credentialsName, "credentials-name", "",
		"name of credentials in vault (env: OMNITOKEN_CREDENTIALS_NAME)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "json",
		"output format: json, pretty (default: json)")

	// search-pages flags
	searchPagesCmd.Flags().IntVar(&searchLimit, "limit", 0, "maximum results to return")

	// Add commands
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)

	// Confluence CLI commands
	rootCmd.AddCommand(readPageCmd)
	rootCmd.AddCommand(readPageXHTMLCmd)
	rootCmd.AddCommand(searchPagesCmd)
	rootCmd.AddCommand(deletePageCmd)
}

// applyEnvDefaults applies environment variable defaults to flags
func applyEnvDefaults() {
	if baseURL == "" {
		baseURL = os.Getenv("CONFLUENCE_BASE_URL")
	}
	if username == "" {
		username = os.Getenv("CONFLUENCE_USERNAME")
	}
	if apiToken == "" {
		apiToken = os.Getenv("CONFLUENCE_API_TOKEN")
	}
	if vaultURI == "" {
		vaultURI = os.Getenv("OMNITOKEN_VAULT_URI")
	}
	if credentialsName == "" {
		credentialsName = os.Getenv("OMNITOKEN_CREDENTIALS_NAME")
	}
	if credentialsName == "" {
		credentialsName = "confluence"
	}
}

// getSkill creates and initializes a Confluence skill with proper credentials
func getSkill(ctx context.Context) (*confskill.Skill, func(), error) {
	applyEnvDefaults()

	hasDirectCreds := baseURL != "" && username != "" && apiToken != ""
	hasVaultCreds := vaultURI != ""

	var client *confluence.Client
	var tokenMgr *omnitoken.TokenManager

	cleanup := func() {}

	if hasVaultCreds {
		var err error
		tokenMgr, err = omnitoken.NewFromVaultURI(vaultURI)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create token manager: %w", err)
		}
		cleanup = func() {
			if err := tokenMgr.Close(); err != nil {
				log.Printf("Warning: failed to close token manager: %v", err)
			}
		}

		httpClient, err := tokenMgr.GetClient(ctx, credentialsName)
		if err != nil {
			cleanup()
			return nil, nil, fmt.Errorf("failed to get authenticated client: %w", err)
		}

		if baseURL == "" {
			cleanup()
			return nil, nil, fmt.Errorf("--base-url is required when using vault credentials")
		}

		client = confluence.NewClient(baseURL, &bearerAuthFromClient{httpClient}, confluence.WithHTTPClient(httpClient))
	} else if hasDirectCreds {
		auth := confluence.BasicAuth{
			Username: username,
			Token:    apiToken,
		}
		client = confluence.NewClient(baseURL, auth)
	} else {
		return nil, nil, fmt.Errorf("credentials required: use --base-url/--username/--api-token or --vault")
	}

	skill := confskill.New(client)
	if err := skill.Init(ctx); err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("failed to initialize Confluence skill: %w", err)
	}

	fullCleanup := func() {
		if err := skill.Close(); err != nil {
			log.Printf("Warning: failed to close Confluence skill: %v", err)
		}
		cleanup()
	}

	return skill, fullCleanup, nil
}

// outputResult outputs the result in the specified format
func outputResult(result any) error {
	var data []byte
	var err error

	switch outputFormat {
	case "pretty":
		data, err = json.MarshalIndent(result, "", "  ")
	default:
		data, err = json.Marshal(result)
	}

	if err != nil {
		return fmt.Errorf("failed to marshal result: %w", err)
	}

	fmt.Println(string(data))
	return nil
}

// runTool runs a tool by name with the given params
func runTool(toolName string, params map[string]any) error {
	ctx := context.Background()

	skill, cleanup, err := getSkill(ctx)
	if err != nil {
		return err
	}
	defer cleanup()

	for _, tool := range skill.Tools() {
		if tool.Name() == toolName {
			result, err := tool.Call(ctx, params)
			if err != nil {
				return err
			}
			return outputResult(result)
		}
	}
	return fmt.Errorf("tool not found: %s", toolName)
}

func runServer(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	skill, cleanup, err := getSkill(ctx)
	if err != nil {
		return err
	}
	defer cleanup()

	rt := runtime.New(&mcp.Implementation{
		Name:    serverName,
		Version: serverVersion,
	}, nil)

	rt.RegisterSkill(skill)

	if err := rt.ServeStdio(ctx); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

// bearerAuthFromClient is a no-op auth method when the HTTP client already has auth.
type bearerAuthFromClient struct {
	client *http.Client
}

func (b *bearerAuthFromClient) Apply(req *http.Request) {
	// No-op: the HTTP client already handles authentication
}
