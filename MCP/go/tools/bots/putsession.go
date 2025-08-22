package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-lex-runtime-v2/mcp-server/config"
	"github.com/amazon-lex-runtime-v2/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func PutsessionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/bots/%s/botAliases/%s/botLocales/%s/sessions/%s", cfg.BaseURL, botId, botAliasId, localeId, sessionId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["ResponseContentType"]; ok {
			req.Header.Set("ResponseContentType", fmt.Sprintf("%v", val))
		}

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
		var result models.PutSessionResponse
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

func CreatePutsessionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_bots_botId_botAliases_botAliasId_botLocales_localeId_sessions_sessionId",
		mcp.WithDescription("Creates a new session or modifies an existing session with an Amazon Lex V2 bot. Use this operation to enable your application to set the state of the bot."),
		mcp.WithString("botId", mcp.Required(), mcp.Description("The identifier of the bot that receives the session data.")),
		mcp.WithString("botAliasId", mcp.Required(), mcp.Description("The alias identifier of the bot that receives the session data.")),
		mcp.WithString("localeId", mcp.Required(), mcp.Description("The locale where the session is in use.")),
		mcp.WithString("sessionId", mcp.Required(), mcp.Description("The identifier of the session that receives the session data.")),
		mcp.WithString("ResponseContentType", mcp.Description("<p>The message that Amazon Lex V2 returns in the response can be either text or speech depending on the value of this parameter. </p> <ul> <li> <p>If the value is <code>text/plain; charset=utf-8</code>, Amazon Lex V2 returns text in the response.</p> </li> </ul>")),
		mcp.WithArray("messages", mcp.Description("Input parameter: A list of messages to send to the user. Messages are sent in the order that they are defined in the list.")),
		mcp.WithObject("requestAttributes", mcp.Description("Input parameter: <p>Request-specific information passed between Amazon Lex V2 and the client application.</p> <p>The namespace <code>x-amz-lex:</code> is reserved for special attributes. Don't create any request attributes with the prefix <code>x-amz-lex:</code>.</p>")),
		mcp.WithObject("sessionState", mcp.Required(), mcp.Description("Input parameter: The state of the user's session with Amazon Lex V2.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PutsessionHandler(cfg),
	}
}
