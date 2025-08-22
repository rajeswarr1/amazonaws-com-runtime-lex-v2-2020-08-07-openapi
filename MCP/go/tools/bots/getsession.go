package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazon-lex-runtime-v2/mcp-server/config"
	"github.com/amazon-lex-runtime-v2/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetsessionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		botIdVal, ok := args["botId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: botId"), nil
		}
		botId, ok := botIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: botId"), nil
		}
		botAliasIdVal, ok := args["botAliasId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: botAliasId"), nil
		}
		botAliasId, ok := botAliasIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: botAliasId"), nil
		}
		localeIdVal, ok := args["localeId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: localeId"), nil
		}
		localeId, ok := localeIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: localeId"), nil
		}
		sessionIdVal, ok := args["sessionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: sessionId"), nil
		}
		sessionId, ok := sessionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: sessionId"), nil
		}
		url := fmt.Sprintf("%s/bots/%s/botAliases/%s/botLocales/%s/sessions/%s", cfg.BaseURL, botId, botAliasId, localeId, sessionId)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.GetSessionResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetsessionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_bots_botId_botAliases_botAliasId_botLocales_localeId_sessions_sessionId",
		mcp.WithDescription("<p>Returns session information for a specified bot, alias, and user.</p> <p>For example, you can use this operation to retrieve session information for a user that has left a long-running session in use.</p> <p>If the bot, alias, or session identifier doesn't exist, Amazon Lex V2 returns a <code>BadRequestException</code>. If the locale doesn't exist or is not enabled for the alias, you receive a <code>BadRequestException</code>.</p>"),
		mcp.WithString("botId", mcp.Required(), mcp.Description("The identifier of the bot that contains the session data.")),
		mcp.WithString("botAliasId", mcp.Required(), mcp.Description("The alias identifier in use for the bot that contains the session data.")),
		mcp.WithString("localeId", mcp.Required(), mcp.Description("The locale where the session is in use.")),
		mcp.WithString("sessionId", mcp.Required(), mcp.Description("The identifier of the session to return.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsessionHandler(cfg),
	}
}
