package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Subscriptions_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		resourceGroupNameVal, ok := args["resourceGroupName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceGroupName"), nil
		}
		resourceGroupName, ok := resourceGroupNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceGroupName"), nil
		}
		serviceNameVal, ok := args["serviceName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: serviceName"), nil
		}
		serviceName, ok := serviceNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: serviceName"), nil
		}
		sidVal, ok := args["sid"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: sid"), nil
		}
		sid, ok := sidVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: sid"), nil
		}
		subscriptionIdVal, ok := args["subscriptionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: subscriptionId"), nil
		}
		subscriptionId, ok := subscriptionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: subscriptionId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api-version=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.SubscriptionUpdateParameters
		
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
		url := fmt.Sprintf("%s/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/subscriptions/%s%s", cfg.BaseURL, resourceGroupName, serviceName, sid, subscriptionId, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["If-Match"]; ok {
			req.Header.Set("If-Match", fmt.Sprintf("%v", val))
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
		var result interface{}
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

func CreateSubscriptions_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_subscriptions_subscriptionId_resourceGroups_resourceGroupName_providers_Microsoft_ApiManagement_service_serviceName_subscriptions_sid",
		mcp.WithDescription("Updates the details of a subscription specified by its identifier."),
		mcp.WithString("resourceGroupName", mcp.Required(), mcp.Description("The name of the resource group.")),
		mcp.WithString("serviceName", mcp.Required(), mcp.Description("The name of the API Management service.")),
		mcp.WithString("sid", mcp.Required(), mcp.Description("Subscription entity Identifier. The entity represents the association between a user and a product in API Management.")),
		mcp.WithString("If-Match", mcp.Required(), mcp.Description("ETag of the Subscription Entity. ETag should match the current entity state from the header response of the GET request or it should be * for unconditional update.")),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("Version of the API to be used with the client request.")),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.")),
		mcp.WithString("state", mcp.Description("Input parameter: Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.")),
		mcp.WithString("stateComment", mcp.Description("Input parameter: Comments describing subscription state change by the administrator.")),
		mcp.WithString("userId", mcp.Description("Input parameter: User identifier path: /users/{uid}")),
		mcp.WithString("expirationDate", mcp.Description("Input parameter: New subscription expiration date.")),
		mcp.WithString("name", mcp.Description("Input parameter: Subscription name.")),
		mcp.WithString("primaryKey", mcp.Description("Input parameter: Primary subscription key.")),
		mcp.WithString("productId", mcp.Description("Input parameter: Product identifier path: /products/{productId}")),
		mcp.WithString("secondaryKey", mcp.Description("Input parameter: Secondary subscription key.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Subscriptions_updateHandler(cfg),
	}
}
