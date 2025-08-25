package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/artifact-registry-api/mcp-server/config"
	"github.com/artifact-registry-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Artifactregistry_projects_locations_repositories_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["repositoryId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("repositoryId=%v", val))
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
		// Create properly typed request body using the generated schema
		var requestBody models.Repository
		
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
		url := fmt.Sprintf("%s/v1/%s/repositories%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Operation
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

func CreateArtifactregistry_projects_locations_repositories_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_parent_repositories",
		mcp.WithDescription("Creates a repository. The returned Operation will finish once the repository has been created. Its response will be the created Repository."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The name of the parent resource where the repository will be created.")),
		mcp.WithString("repositoryId", mcp.Description("Required. The repository id to use for this repository.")),
		mcp.WithString("mode", mcp.Description("Input parameter: Optional. The mode of the repository.")),
		mcp.WithBoolean("disallowUnspecifiedMode", mcp.Description("Input parameter: Optional. If this is true, aunspecified repo type will be treated as error. Is used for new repo types that don't have any specific fields. Right now is used by AOSS team when creating repos for customers.")),
		mcp.WithObject("dockerConfig", mcp.Description("Input parameter: DockerRepositoryConfig is docker related repository details. Provides additional configuration details for repositories of the docker format type.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. The time when the repository was created.")),
		mcp.WithObject("labels", mcp.Description("Input parameter: Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.")),
		mcp.WithBoolean("satisfiesPzs", mcp.Description("Input parameter: Output only. If set, the repository satisfies physical zone separation.")),
		mcp.WithString("name", mcp.Description("Input parameter: The name of the repository, for example: `projects/p1/locations/us-central1/repositories/repo1`.")),
		mcp.WithObject("remoteRepositoryConfig", mcp.Description("Input parameter: Remote repository configuration.")),
		mcp.WithBoolean("cleanupPolicyDryRun", mcp.Description("Input parameter: Optional. If true, the cleanup pipeline is prevented from deleting versions in this repository.")),
		mcp.WithString("description", mcp.Description("Input parameter: The user-provided description of the repository.")),
		mcp.WithObject("virtualRepositoryConfig", mcp.Description("Input parameter: Virtual repository configuration.")),
		mcp.WithObject("mavenConfig", mcp.Description("Input parameter: MavenRepositoryConfig is maven related repository details. Provides additional configuration details for repositories of the maven format type.")),
		mcp.WithString("format", mcp.Description("Input parameter: Optional. The format of packages that are stored in the repository.")),
		mcp.WithString("kmsKeyName", mcp.Description("Input parameter: The Cloud KMS resource name of the customer managed encryption key that's used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.")),
		mcp.WithString("sizeBytes", mcp.Description("Input parameter: Output only. The size, in bytes, of all artifact storage in this repository. Repositories that are generally available or in public preview use this to calculate storage costs.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. The time when the repository was last updated.")),
		mcp.WithObject("cleanupPolicies", mcp.Description("Input parameter: Optional. Cleanup policies for this repository. Cleanup policies indicate when certain package versions can be automatically deleted. Map keys are policy IDs supplied by users during policy creation. They must unique within a repository and be under 128 characters in length.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Artifactregistry_projects_locations_repositories_createHandler(cfg),
	}
}
