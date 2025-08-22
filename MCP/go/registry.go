package main

import (
	"github.com/amazon-lex-runtime-v2/mcp-server/config"
	"github.com/amazon-lex-runtime-v2/mcp-server/models"
	tools_bots "github.com/amazon-lex-runtime-v2/mcp-server/tools/bots"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_bots.CreateDeletesessionTool(cfg),
		tools_bots.CreateGetsessionTool(cfg),
		tools_bots.CreatePutsessionTool(cfg),
		tools_bots.CreateRecognizetextTool(cfg),
		tools_bots.CreateRecognizeutteranceTool(cfg),
	}
}
