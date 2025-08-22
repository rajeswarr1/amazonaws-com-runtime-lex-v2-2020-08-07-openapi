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

func RecognizetextHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/bots/%s/botAliases/%s/botLocales/%s/sessions/%s/text", cfg.BaseURL, botId, botAliasId, localeId, sessionId)
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
		var result models.RecognizeTextResponse
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

func CreateRecognizetextTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_bots_botId_botAliases_botAliasId_botLocales_localeId_sessions_sessionId_text",
		mcp.WithDescription("<p>Sends user input to Amazon Lex V2. Client applications use this API to send requests to Amazon Lex V2 at runtime. Amazon Lex V2 then interprets the user input using the machine learning model that it build for the bot.</p> <p>In response, Amazon Lex V2 returns the next message to convey to the user and an optional response card to display.</p> <p>If the optional post-fulfillment response is specified, the messages are returned as follows. For more information, see <a href="https://docs.aws.amazon.com/lexv2/latest/dg/API_PostFulfillmentStatusSpecification.html">PostFulfillmentStatusSpecification</a>.</p> <ul> <li> <p> <b>Success message</b> - Returned if the Lambda function completes successfully and the intent state is fulfilled or ready fulfillment if the message is present.</p> </li> <li> <p> <b>Failed message</b> - The failed message is returned if the Lambda function throws an exception or if the Lambda function returns a failed intent state without a message.</p> </li> <li> <p> <b>Timeout message</b> - If you don't configure a timeout message and a timeout, and the Lambda function doesn't return within 30 seconds, the timeout message is returned. If you configure a timeout, the timeout message is returned when the period times out. </p> </li> </ul> <p>For more information, see <a href="https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete.html">Completion message</a>.</p>"),
		mcp.WithString("botId", mcp.Required(), mcp.Description("The identifier of the bot that processes the request.")),
		mcp.WithString("botAliasId", mcp.Required(), mcp.Description("The alias identifier in use for the bot that processes the request.")),
		mcp.WithString("localeId", mcp.Required(), mcp.Description("The locale where the session is in use.")),
		mcp.WithString("sessionId", mcp.Required(), mcp.Description("The identifier of the user session that is having the conversation.")),
		mcp.WithObject("sessionState", mcp.Description("Input parameter: The state of the user's session with Amazon Lex V2.")),
		mcp.WithString("text", mcp.Required(), mcp.Description("Input parameter: The text that the user entered. Amazon Lex V2 interprets this text.")),
		mcp.WithObject("requestAttributes", mcp.Description("Input parameter: <p>Request-specific information passed between the client application and Amazon Lex V2 </p> <p>The namespace <code>x-amz-lex:</code> is reserved for special attributes. Don't create any request attributes with the prefix <code>x-amz-lex:</code>.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    RecognizetextHandler(cfg),
	}
}
