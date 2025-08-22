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

func DeletesessionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result models.DeleteSessionResponse
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

func CreateDeletesessionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_bots_botId_botAliases_botAliasId_botLocales_localeId_sessions_sessionId",
		mcp.WithDescription("<p>Removes session information for a specified bot, alias, and user ID. </p> <p>You can use this operation to restart a conversation with a bot. When you remove a session, the entire history of the session is removed so that you can start again.</p> <p>You don't need to delete a session. Sessions have a time limit and will expire. Set the session time limit when you create the bot. The default is 5 minutes, but you can specify anything between 1 minute and 24 hours.</p> <p>If you specify a bot or alias ID that doesn't exist, you receive a <code>BadRequestException.</code> </p> <p>If the locale doesn't exist in the bot, or if the locale hasn't been enables for the alias, you receive a <code>BadRequestException</code>.</p>"),
		mcp.WithString("botId", mcp.Required(), mcp.Description("The identifier of the bot that contains the session data.")),
		mcp.WithString("botAliasId", mcp.Required(), mcp.Description("The alias identifier in use for the bot that contains the session data.")),
		mcp.WithString("localeId", mcp.Required(), mcp.Description("The locale where the session is in use.")),
		mcp.WithString("sessionId", mcp.Required(), mcp.Description("The identifier of the session to delete.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletesessionHandler(cfg),
	}
}
