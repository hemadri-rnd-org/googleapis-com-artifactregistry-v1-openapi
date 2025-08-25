package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/artifact-registry-api/mcp-server/config"
	"github.com/artifact-registry-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Artifactregistry_projects_locations_repositories_files_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["filter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter=%v", val))
		}
		if val, ok := args["orderBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderBy=%v", val))
		}
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/%s/files%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
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
		var result models.ListFilesResponse
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

func CreateArtifactregistry_projects_locations_repositories_files_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_parent_files",
		mcp.WithDescription("Lists files."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The name of the repository whose files will be listed. For example: \"projects/p1/locations/us-central1/repositories/repo1")),
		mcp.WithString("filter", mcp.Description("An expression for filtering the results of the request. Filter rules are case insensitive. The fields eligible for filtering are: * `name` * `owner` An example of using a filter: * `name=\"projects/p1/locations/us-central1/repositories/repo1/files/a/b/*\"` --> Files with an ID starting with \"a/b/\". * `owner=\"projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/1.0\"` --> Files owned by the version `1.0` in package `pkg1`.")),
		mcp.WithString("orderBy", mcp.Description("The field to order the results by.")),
		mcp.WithNumber("pageSize", mcp.Description("The maximum number of files to return.")),
		mcp.WithString("pageToken", mcp.Description("The next_page_token value returned from a previous list request, if any.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Artifactregistry_projects_locations_repositories_files_listHandler(cfg),
	}
}
