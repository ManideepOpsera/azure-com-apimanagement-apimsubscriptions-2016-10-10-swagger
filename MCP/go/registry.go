package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_subscriptions "github.com/apimanagementclient/mcp-server/tools/subscriptions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_subscriptions.CreateSubscriptions_regeneratesecondarykeyTool(cfg),
		tools_subscriptions.CreateSubscriptions_listTool(cfg),
		tools_subscriptions.CreateSubscriptions_deleteTool(cfg),
		tools_subscriptions.CreateSubscriptions_getTool(cfg),
		tools_subscriptions.CreateSubscriptions_updateTool(cfg),
		tools_subscriptions.CreateSubscriptions_createorupdateTool(cfg),
		tools_subscriptions.CreateSubscriptions_regenerateprimarykeyTool(cfg),
	}
}
