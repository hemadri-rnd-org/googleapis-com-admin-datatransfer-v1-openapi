package main

import (
	"github.com/admin-sdk-api/mcp-server/config"
	"github.com/admin-sdk-api/mcp-server/models"
	tools_applications "github.com/admin-sdk-api/mcp-server/tools/applications"
	tools_transfers "github.com/admin-sdk-api/mcp-server/tools/transfers"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_applications.CreateDatatransfer_applications_listTool(cfg),
		tools_applications.CreateDatatransfer_applications_getTool(cfg),
		tools_transfers.CreateDatatransfer_transfers_listTool(cfg),
		tools_transfers.CreateDatatransfer_transfers_insertTool(cfg),
		tools_transfers.CreateDatatransfer_transfers_getTool(cfg),
	}
}
