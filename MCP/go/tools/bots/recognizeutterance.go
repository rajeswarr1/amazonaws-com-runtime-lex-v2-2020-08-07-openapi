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

func RecognizeutteranceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/bots/%s/botAliases/%s/botLocales/%s/sessions/%s/utterance#Content-Type", cfg.BaseURL, botId, botAliasId, localeId, sessionId)
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
		if val, ok := args["x-amz-lex-session-state"]; ok {
			req.Header.Set("x-amz-lex-session-state", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-amz-lex-request-attributes"]; ok {
			req.Header.Set("x-amz-lex-request-attributes", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Content-Type"]; ok {
			req.Header.Set("Content-Type", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Response-Content-Type"]; ok {
			req.Header.Set("Response-Content-Type", fmt.Sprintf("%v", val))
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
		var result models.RecognizeUtteranceResponse
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

func CreateRecognizeutteranceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_bots_botId_botAliases_botAliasId_botLocales_localeId_sessions_sessionId_utterance#Content-Type",
		mcp.WithDescription("<p>Sends user input to Amazon Lex V2. You can send text or speech. Clients use this API to send text and audio requests to Amazon Lex V2 at runtime. Amazon Lex V2 interprets the user input using the machine learning model built for the bot.</p> <p>The following request fields must be compressed with gzip and then base64 encoded before you send them to Amazon Lex V2. </p> <ul> <li> <p>requestAttributes</p> </li> <li> <p>sessionState</p> </li> </ul> <p>The following response fields are compressed using gzip and then base64 encoded by Amazon Lex V2. Before you can use these fields, you must decode and decompress them. </p> <ul> <li> <p>inputTranscript</p> </li> <li> <p>interpretations</p> </li> <li> <p>messages</p> </li> <li> <p>requestAttributes</p> </li> <li> <p>sessionState</p> </li> </ul> <p>The example contains a Java application that compresses and encodes a Java object to send to Amazon Lex V2, and a second that decodes and decompresses a response from Amazon Lex V2.</p> <p>If the optional post-fulfillment response is specified, the messages are returned as follows. For more information, see <a href="https://docs.aws.amazon.com/lexv2/latest/dg/API_PostFulfillmentStatusSpecification.html">PostFulfillmentStatusSpecification</a>.</p> <ul> <li> <p> <b>Success message</b> - Returned if the Lambda function completes successfully and the intent state is fulfilled or ready fulfillment if the message is present.</p> </li> <li> <p> <b>Failed message</b> - The failed message is returned if the Lambda function throws an exception or if the Lambda function returns a failed intent state without a message.</p> </li> <li> <p> <b>Timeout message</b> - If you don't configure a timeout message and a timeout, and the Lambda function doesn't return within 30 seconds, the timeout message is returned. If you configure a timeout, the timeout message is returned when the period times out. </p> </li> </ul> <p>For more information, see <a href="https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete.html">Completion message</a>.</p>"),
		mcp.WithString("botId", mcp.Required(), mcp.Description("The identifier of the bot that should receive the request.")),
		mcp.WithString("botAliasId", mcp.Required(), mcp.Description("The alias identifier in use for the bot that should receive the request.")),
		mcp.WithString("localeId", mcp.Required(), mcp.Description("The locale where the session is in use.")),
		mcp.WithString("sessionId", mcp.Required(), mcp.Description("The identifier of the session in use.")),
		mcp.WithString("x-amz-lex-session-state", mcp.Description("<p>Sets the state of the session with the user. You can use this to set the current intent, attributes, context, and dialog action. Use the dialog action to determine the next step that Amazon Lex V2 should use in the conversation with the user.</p> <p>The <code>sessionState</code> field must be compressed using gzip and then base64 encoded before sending to Amazon Lex V2.</p>")),
		mcp.WithString("x-amz-lex-request-attributes", mcp.Description("<p>Request-specific information passed between the client application and Amazon Lex V2 </p> <p>The namespace <code>x-amz-lex:</code> is reserved for special attributes. Don't create any request attributes for prefix <code>x-amz-lex:</code>.</p> <p>The <code>requestAttributes</code> field must be compressed using gzip and then base64 encoded before sending to Amazon Lex V2.</p>")),
		mcp.WithString("Content-Type", mcp.Required(), mcp.Description("<p>Indicates the format for audio input or that the content is text. The header must start with one of the following prefixes:</p> <ul> <li> <p>PCM format, audio data must be in little-endian byte order.</p> <ul> <li> <p>audio/l16; rate=16000; channels=1</p> </li> <li> <p>audio/x-l16; sample-rate=16000; channel-count=1</p> </li> <li> <p>audio/lpcm; sample-rate=8000; sample-size-bits=16; channel-count=1; is-big-endian=false</p> </li> </ul> </li> <li> <p>Opus format</p> <ul> <li> <p>audio/x-cbr-opus-with-preamble;preamble-size=0;bit-rate=256000;frame-size-milliseconds=4</p> </li> </ul> </li> <li> <p>Text format</p> <ul> <li> <p>text/plain; charset=utf-8</p> </li> </ul> </li> </ul>")),
		mcp.WithString("Response-Content-Type", mcp.Description("<p>The message that Amazon Lex V2 returns in the response can be either text or speech based on the <code>responseContentType</code> value.</p> <ul> <li> <p>If the value is <code>text/plain;charset=utf-8</code>, Amazon Lex V2 returns text in the response.</p> </li> <li> <p>If the value begins with <code>audio/</code>, Amazon Lex V2 returns speech in the response. Amazon Lex V2 uses Amazon Polly to generate the speech using the configuration that you specified in the <code>responseContentType</code> parameter. For example, if you specify <code>audio/mpeg</code> as the value, Amazon Lex V2 returns speech in the MPEG format.</p> </li> <li> <p>If the value is <code>audio/pcm</code>, the speech returned is <code>audio/pcm</code> at 16 KHz in 16-bit, little-endian format.</p> </li> <li> <p>The following are the accepted values:</p> <ul> <li> <p>audio/mpeg</p> </li> <li> <p>audio/ogg</p> </li> <li> <p>audio/pcm (16 KHz)</p> </li> <li> <p>audio/* (defaults to mpeg)</p> </li> <li> <p>text/plain; charset=utf-8</p> </li> </ul> </li> </ul>")),
		mcp.WithString("inputStream", mcp.Description("Input parameter: User input in PCM or Opus audio format or text format as described in the <code>requestContentType</code> parameter.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    RecognizeutteranceHandler(cfg),
	}
}
