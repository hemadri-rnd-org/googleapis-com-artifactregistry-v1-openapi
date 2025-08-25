package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// KfpArtifact represents the KfpArtifact schema from the OpenAPI specification
type KfpArtifact struct {
	Version string `json:"version,omitempty"` // The version associated with the KFP artifact. Must follow the Semantic Versioning standard.
	Name string `json:"name,omitempty"` // Output only. Resource name of the KFP artifact. Since users don't directly interact with this resource, the name will be derived from the associated version. For example, when version = ".../versions/sha256:abcdef...", the name will be ".../kfpArtifacts/sha256:abcdef...".
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	Name string `json:"name,omitempty"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.
	Response map[string]interface{} `json:"response,omitempty"` // The normal, successful response of the operation. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
	Done bool `json:"done,omitempty"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Name string `json:"name,omitempty"` // The name of the tag, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/tags/tag1". If the package part contains slashes, the slashes are escaped. The tag part can only have characters in [a-zA-Z0-9\-._~:@], anything else must be URL encoded.
	Version string `json:"version,omitempty"` // The name of the version the tag refers to, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/sha256:5243811" If the package or version ID parts contain slashes, the slashes are escaped.
}

// UploadGoModuleRequest represents the UploadGoModuleRequest schema from the OpenAPI specification
type UploadGoModuleRequest struct {
}

// ImportYumArtifactsErrorInfo represents the ImportYumArtifactsErrorInfo schema from the OpenAPI specification
type ImportYumArtifactsErrorInfo struct {
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Gcssource ImportYumArtifactsGcsSource `json:"gcsSource,omitempty"` // Google Cloud Storage location where the artifacts currently reside.
}

// ImportGoogetArtifactsGcsSource represents the ImportGoogetArtifactsGcsSource schema from the OpenAPI specification
type ImportGoogetArtifactsGcsSource struct {
	Uris []string `json:"uris,omitempty"` // Cloud Storage paths URI (e.g., `gs://my_bucket/my_object`).
	Usewildcards bool `json:"useWildcards,omitempty"` // Supports URI wildcards for matching multiple objects from a single URI.
}

// UploadYumArtifactResponse represents the UploadYumArtifactResponse schema from the OpenAPI specification
type UploadYumArtifactResponse struct {
	Yumartifacts []YumArtifact `json:"yumArtifacts,omitempty"` // The Yum artifacts updated.
}

// YumArtifact represents the YumArtifact schema from the OpenAPI specification
type YumArtifact struct {
	Packagename string `json:"packageName,omitempty"` // Output only. The yum package name of the artifact.
	Packagetype string `json:"packageType,omitempty"` // Output only. An artifact is a binary or source package.
	Architecture string `json:"architecture,omitempty"` // Output only. Operating system architecture of the artifact.
	Name string `json:"name,omitempty"` // Output only. The Artifact Registry resource name of the artifact.
}

// ImportGoogetArtifactsResponse represents the ImportGoogetArtifactsResponse schema from the OpenAPI specification
type ImportGoogetArtifactsResponse struct {
	Errors []ImportGoogetArtifactsErrorInfo `json:"errors,omitempty"` // Detailed error info for packages that were not imported.
	Googetartifacts []GoogetArtifact `json:"googetArtifacts,omitempty"` // The GooGet artifacts updated.
}

// YumRepository represents the YumRepository schema from the OpenAPI specification
type YumRepository struct {
	Publicrepository GoogleDevtoolsArtifactregistryV1RemoteRepositoryConfigYumRepositoryPublicRepository `json:"publicRepository,omitempty"` // Publicly available Yum repositories constructed from a common repository base and a custom repository path.
}

// MavenArtifact represents the MavenArtifact schema from the OpenAPI specification
type MavenArtifact struct {
	Updatetime string `json:"updateTime,omitempty"` // Output only. Time the artifact was updated.
	Version string `json:"version,omitempty"` // Version of this artifact.
	Artifactid string `json:"artifactId,omitempty"` // Artifact ID for the artifact.
	Createtime string `json:"createTime,omitempty"` // Output only. Time the artifact was created.
	Groupid string `json:"groupId,omitempty"` // Group ID for the artifact. Example: com.google.guava
	Name string `json:"name,omitempty"` // Required. registry_location, project_id, repository_name and maven_artifact forms a unique artifact For example, "projects/test-project/locations/us-west4/repositories/test-repo/mavenArtifacts/ com.google.guava:guava:31.0-jre", where "us-west4" is the registry_location, "test-project" is the project_id, "test-repo" is the repository_name and "com.google.guava:guava:31.0-jre" is the maven artifact.
	Pomuri string `json:"pomUri,omitempty"` // Required. URL to access the pom file of the artifact. Example: us-west4-maven.pkg.dev/test-project/test-repo/com/google/guava/guava/31.0/guava-31.0.pom
}

// AptArtifact represents the AptArtifact schema from the OpenAPI specification
type AptArtifact struct {
	Controlfile string `json:"controlFile,omitempty"` // Output only. Contents of the artifact's control metadata file.
	Name string `json:"name,omitempty"` // Output only. The Artifact Registry resource name of the artifact.
	Packagename string `json:"packageName,omitempty"` // Output only. The Apt package name of the artifact.
	Packagetype string `json:"packageType,omitempty"` // Output only. An artifact is a binary or source package.
	Architecture string `json:"architecture,omitempty"` // Output only. Operating system architecture of the artifact.
	Component string `json:"component,omitempty"` // Output only. Repository component of the artifact.
}

// ImportAptArtifactsMetadata represents the ImportAptArtifactsMetadata schema from the OpenAPI specification
type ImportAptArtifactsMetadata struct {
}

// UploadGoogetArtifactRequest represents the UploadGoogetArtifactRequest schema from the OpenAPI specification
type UploadGoogetArtifactRequest struct {
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// ImportAptArtifactsResponse represents the ImportAptArtifactsResponse schema from the OpenAPI specification
type ImportAptArtifactsResponse struct {
	Aptartifacts []AptArtifact `json:"aptArtifacts,omitempty"` // The Apt artifacts imported.
	Errors []ImportAptArtifactsErrorInfo `json:"errors,omitempty"` // Detailed error info for packages that were not imported.
}

// ListMavenArtifactsResponse represents the ListMavenArtifactsResponse schema from the OpenAPI specification
type ListMavenArtifactsResponse struct {
	Mavenartifacts []MavenArtifact `json:"mavenArtifacts,omitempty"` // The maven artifacts returned.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The token to retrieve the next page of artifacts, or empty if there are no more artifacts to return.
}

// GoogleDevtoolsArtifactregistryV1File represents the GoogleDevtoolsArtifactregistryV1File schema from the OpenAPI specification
type GoogleDevtoolsArtifactregistryV1File struct {
	Owner string `json:"owner,omitempty"` // The name of the Package or Version that owns this file, if any.
	Sizebytes string `json:"sizeBytes,omitempty"` // The size of the File in bytes.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the File was last updated.
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the File was created.
	Fetchtime string `json:"fetchTime,omitempty"` // Output only. The time when the last attempt to refresh the file's data was made. Only set when the repository is remote.
	Hashes []Hash `json:"hashes,omitempty"` // The hashes of the file content.
	Name string `json:"name,omitempty"` // The name of the file, for example: "projects/p1/locations/us-central1/repositories/repo1/files/a%2Fb%2Fc.txt". If the file ID part contains slashes, they are escaped.
}

// PythonRepository represents the PythonRepository schema from the OpenAPI specification
type PythonRepository struct {
	Publicrepository string `json:"publicRepository,omitempty"` // One of the publicly available Python repositories supported by Artifact Registry.
}

// ImportAptArtifactsRequest represents the ImportAptArtifactsRequest schema from the OpenAPI specification
type ImportAptArtifactsRequest struct {
	Gcssource ImportAptArtifactsGcsSource `json:"gcsSource,omitempty"` // Google Cloud Storage location where the artifacts currently reside.
}

// UsernamePasswordCredentials represents the UsernamePasswordCredentials schema from the OpenAPI specification
type UsernamePasswordCredentials struct {
	Passwordsecretversion string `json:"passwordSecretVersion,omitempty"` // The Secret Manager key version that holds the password to access the remote repository. Must be in the format of `projects/{project}/secrets/{secret}/versions/{version}`.
	Username string `json:"username,omitempty"` // The username to access the remote repository.
}

// ProjectSettings represents the ProjectSettings schema from the OpenAPI specification
type ProjectSettings struct {
	Legacyredirectionstate string `json:"legacyRedirectionState,omitempty"` // The redirection state of the legacy repositories in this project.
	Name string `json:"name,omitempty"` // The name of the project's settings. Always of the form: projects/{project-id}/projectSettings In update request: never set In response: always set
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Format string `json:"format,omitempty"` // Optional. The format of packages that are stored in the repository.
	Kmskeyname string `json:"kmsKeyName,omitempty"` // The Cloud KMS resource name of the customer managed encryption key that's used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
	Sizebytes string `json:"sizeBytes,omitempty"` // Output only. The size, in bytes, of all artifact storage in this repository. Repositories that are generally available or in public preview use this to calculate storage costs.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the repository was last updated.
	Cleanuppolicies map[string]interface{} `json:"cleanupPolicies,omitempty"` // Optional. Cleanup policies for this repository. Cleanup policies indicate when certain package versions can be automatically deleted. Map keys are policy IDs supplied by users during policy creation. They must unique within a repository and be under 128 characters in length.
	Mode string `json:"mode,omitempty"` // Optional. The mode of the repository.
	Disallowunspecifiedmode bool `json:"disallowUnspecifiedMode,omitempty"` // Optional. If this is true, aunspecified repo type will be treated as error. Is used for new repo types that don't have any specific fields. Right now is used by AOSS team when creating repos for customers.
	Dockerconfig DockerRepositoryConfig `json:"dockerConfig,omitempty"` // DockerRepositoryConfig is docker related repository details. Provides additional configuration details for repositories of the docker format type.
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the repository was created.
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
	Satisfiespzs bool `json:"satisfiesPzs,omitempty"` // Output only. If set, the repository satisfies physical zone separation.
	Name string `json:"name,omitempty"` // The name of the repository, for example: `projects/p1/locations/us-central1/repositories/repo1`.
	Remoterepositoryconfig RemoteRepositoryConfig `json:"remoteRepositoryConfig,omitempty"` // Remote repository configuration.
	Cleanuppolicydryrun bool `json:"cleanupPolicyDryRun,omitempty"` // Optional. If true, the cleanup pipeline is prevented from deleting versions in this repository.
	Description string `json:"description,omitempty"` // The user-provided description of the repository.
	Virtualrepositoryconfig VirtualRepositoryConfig `json:"virtualRepositoryConfig,omitempty"` // Virtual repository configuration.
	Mavenconfig MavenRepositoryConfig `json:"mavenConfig,omitempty"` // MavenRepositoryConfig is maven related repository details. Provides additional configuration details for repositories of the maven format type.
}

// ListRepositoriesResponse represents the ListRepositoriesResponse schema from the OpenAPI specification
type ListRepositoriesResponse struct {
	Repositories []Repository `json:"repositories,omitempty"` // The repositories returned.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The token to retrieve the next page of repositories, or empty if there are no more repositories to return.
}

// TestIamPermissionsRequest represents the TestIamPermissionsRequest schema from the OpenAPI specification
type TestIamPermissionsRequest struct {
	Permissions []string `json:"permissions,omitempty"` // The set of permissions to check for the `resource`. Permissions with wildcards (such as `*` or `storage.*`) are not allowed. For more information see [IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).
}

// ListFilesResponse represents the ListFilesResponse schema from the OpenAPI specification
type ListFilesResponse struct {
	Files []GoogleDevtoolsArtifactregistryV1File `json:"files,omitempty"` // The files returned.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The token to retrieve the next page of files, or empty if there are no more files to return.
}

// UploadKfpArtifactMetadata represents the UploadKfpArtifactMetadata schema from the OpenAPI specification
type UploadKfpArtifactMetadata struct {
}

// DockerRepository represents the DockerRepository schema from the OpenAPI specification
type DockerRepository struct {
	Publicrepository string `json:"publicRepository,omitempty"` // One of the publicly available Docker repositories supported by Artifact Registry.
}

// Hash represents the Hash schema from the OpenAPI specification
type Hash struct {
	TypeField string `json:"type,omitempty"` // The algorithm used to compute the hash value.
	Value string `json:"value,omitempty"` // The hash value.
}

// ListTagsResponse represents the ListTagsResponse schema from the OpenAPI specification
type ListTagsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The token to retrieve the next page of tags, or empty if there are no more tags to return.
	Tags []Tag `json:"tags,omitempty"` // The tags returned.
}

// UploadKfpArtifactRequest represents the UploadKfpArtifactRequest schema from the OpenAPI specification
type UploadKfpArtifactRequest struct {
	Tags []string `json:"tags,omitempty"` // Tags to be created with the version.
	Description string `json:"description,omitempty"` // Description of the package version.
}

// Location represents the Location schema from the OpenAPI specification
type Location struct {
	Displayname string `json:"displayName,omitempty"` // The friendly name for this location, typically a nearby city name. For example, "Tokyo".
	Labels map[string]interface{} `json:"labels,omitempty"` // Cross-service attributes for the location. For example {"cloud.googleapis.com/region": "us-east1"}
	Locationid string `json:"locationId,omitempty"` // The canonical id for this location. For example: `"us-east1"`.
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata. For example the available capacity at the given location.
	Name string `json:"name,omitempty"` // Resource name for the location, which may vary between implementations. For example: `"projects/example-project/locations/us-east1"`
}

// ListLocationsResponse represents the ListLocationsResponse schema from the OpenAPI specification
type ListLocationsResponse struct {
	Locations []Location `json:"locations,omitempty"` // A list of locations that matches the specified filter in the request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
}

// PythonPackage represents the PythonPackage schema from the OpenAPI specification
type PythonPackage struct {
	Name string `json:"name,omitempty"` // Required. registry_location, project_id, repository_name and python_package forms a unique package name:`projects//locations//repository//pythonPackages/`. For example, "projects/test-project/locations/us-west4/repositories/test-repo/pythonPackages/ python_package:1.0.0", where "us-west4" is the registry_location, "test-project" is the project_id, "test-repo" is the repository_name and python_package:1.0.0" is the python package.
	Packagename string `json:"packageName,omitempty"` // Package for the artifact.
	Updatetime string `json:"updateTime,omitempty"` // Output only. Time the package was updated.
	Uri string `json:"uri,omitempty"` // Required. URL to access the package. Example: us-west4-python.pkg.dev/test-project/test-repo/python_package/file-name-1.0.0.tar.gz
	Version string `json:"version,omitempty"` // Version of this package.
	Createtime string `json:"createTime,omitempty"` // Output only. Time the package was created.
}

// UploadAptArtifactMediaResponse represents the UploadAptArtifactMediaResponse schema from the OpenAPI specification
type UploadAptArtifactMediaResponse struct {
	Operation Operation `json:"operation,omitempty"` // This resource represents a long-running operation that is the result of a network API call.
}

// UploadGoModuleMetadata represents the UploadGoModuleMetadata schema from the OpenAPI specification
type UploadGoModuleMetadata struct {
}

// Version represents the Version schema from the OpenAPI specification
type Version struct {
	Name string `json:"name,omitempty"` // The name of the version, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/art1". If the package or version ID parts contain slashes, the slashes are escaped.
	Relatedtags []Tag `json:"relatedTags,omitempty"` // Output only. A list of related tags. Will contain up to 100 tags that reference this version.
	Updatetime string `json:"updateTime,omitempty"` // The time when the version was last updated.
	Createtime string `json:"createTime,omitempty"` // The time when the version was created.
	Description string `json:"description,omitempty"` // Optional. Description of the version, as specified in its metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Output only. Repository-specific Metadata stored against this version. The fields returned are defined by the underlying repository-specific resource. Currently, the resources could be: DockerImage MavenArtifact
}

// AptRepository represents the AptRepository schema from the OpenAPI specification
type AptRepository struct {
	Publicrepository GoogleDevtoolsArtifactregistryV1RemoteRepositoryConfigAptRepositoryPublicRepository `json:"publicRepository,omitempty"` // Publicly available Apt repositories constructed from a common repository base and a custom repository path.
}

// ListDockerImagesResponse represents the ListDockerImagesResponse schema from the OpenAPI specification
type ListDockerImagesResponse struct {
	Dockerimages []DockerImage `json:"dockerImages,omitempty"` // The docker images returned.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The token to retrieve the next page of artifacts, or empty if there are no more artifacts to return.
}

// VirtualRepositoryConfig represents the VirtualRepositoryConfig schema from the OpenAPI specification
type VirtualRepositoryConfig struct {
	Upstreampolicies []UpstreamPolicy `json:"upstreamPolicies,omitempty"` // Policies that configure the upstream artifacts distributed by the Virtual Repository. Upstream policies cannot be set on a standard repository.
}

// UploadYumArtifactMetadata represents the UploadYumArtifactMetadata schema from the OpenAPI specification
type UploadYumArtifactMetadata struct {
}

// ImportYumArtifactsRequest represents the ImportYumArtifactsRequest schema from the OpenAPI specification
type ImportYumArtifactsRequest struct {
	Gcssource ImportYumArtifactsGcsSource `json:"gcsSource,omitempty"` // Google Cloud Storage location where the artifacts currently reside.
}

// ListPythonPackagesResponse represents the ListPythonPackagesResponse schema from the OpenAPI specification
type ListPythonPackagesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The token to retrieve the next page of artifacts, or empty if there are no more artifacts to return.
	Pythonpackages []PythonPackage `json:"pythonPackages,omitempty"` // The python packages returned.
}

// ListVersionsResponse represents the ListVersionsResponse schema from the OpenAPI specification
type ListVersionsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The token to retrieve the next page of versions, or empty if there are no more versions to return.
	Versions []Version `json:"versions,omitempty"` // The versions returned.
}

// ImportYumArtifactsMetadata represents the ImportYumArtifactsMetadata schema from the OpenAPI specification
type ImportYumArtifactsMetadata struct {
}

// UploadKfpArtifactMediaResponse represents the UploadKfpArtifactMediaResponse schema from the OpenAPI specification
type UploadKfpArtifactMediaResponse struct {
	Operation Operation `json:"operation,omitempty"` // This resource represents a long-running operation that is the result of a network API call.
}

// BatchDeleteVersionsMetadata represents the BatchDeleteVersionsMetadata schema from the OpenAPI specification
type BatchDeleteVersionsMetadata struct {
	Failedversions []string `json:"failedVersions,omitempty"` // The versions the operation failed to delete.
}

// ImportYumArtifactsResponse represents the ImportYumArtifactsResponse schema from the OpenAPI specification
type ImportYumArtifactsResponse struct {
	Errors []ImportYumArtifactsErrorInfo `json:"errors,omitempty"` // Detailed error info for packages that were not imported.
	Yumartifacts []YumArtifact `json:"yumArtifacts,omitempty"` // The yum artifacts imported.
}

// UploadAptArtifactRequest represents the UploadAptArtifactRequest schema from the OpenAPI specification
type UploadAptArtifactRequest struct {
}

// GoogleDevtoolsArtifactregistryV1RemoteRepositoryConfigYumRepositoryPublicRepository represents the GoogleDevtoolsArtifactregistryV1RemoteRepositoryConfigYumRepositoryPublicRepository schema from the OpenAPI specification
type GoogleDevtoolsArtifactregistryV1RemoteRepositoryConfigYumRepositoryPublicRepository struct {
	Repositorybase string `json:"repositoryBase,omitempty"` // A common public repository base for Yum.
	Repositorypath string `json:"repositoryPath,omitempty"` // A custom field to define a path to a specific repository from the base.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Details []map[string]interface{} `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of google.rpc.Code.
}

// ListNpmPackagesResponse represents the ListNpmPackagesResponse schema from the OpenAPI specification
type ListNpmPackagesResponse struct {
	Npmpackages []NpmPackage `json:"npmPackages,omitempty"` // The npm packages returned.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The token to retrieve the next page of artifacts, or empty if there are no more artifacts to return.
}

// CleanupPolicyMostRecentVersions represents the CleanupPolicyMostRecentVersions schema from the OpenAPI specification
type CleanupPolicyMostRecentVersions struct {
	Keepcount int `json:"keepCount,omitempty"` // Minimum number of versions to keep.
	Packagenameprefixes []string `json:"packageNamePrefixes,omitempty"` // List of package name prefixes that will apply this rule.
}

// Binding represents the Binding schema from the OpenAPI specification
type Binding struct {
	Role string `json:"role,omitempty"` // Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`. For an overview of the IAM roles and permissions, see the [IAM documentation](https://cloud.google.com/iam/docs/roles-overview). For a list of the available pre-defined roles, see [here](https://cloud.google.com/iam/docs/understanding-roles).
	Condition Expr `json:"condition,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
	Members []string `json:"members,omitempty"` // Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workforce identity pool. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/group/{group_id}`: All workforce identities in a group. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All workforce identities with a specific attribute value. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/*`: All identities in a workforce identity pool. * `principal://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workload identity pool. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/group/{group_id}`: A workload identity pool group. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All identities in a workload identity pool with a certain attribute. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/*`: All identities in a workload identity pool. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `deleted:principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: Deleted single identity in a workforce identity pool. For example, `deleted:principal://iam.googleapis.com/locations/global/workforcePools/my-pool-id/subject/my-subject-attribute-value`.
}

// GoogetArtifact represents the GoogetArtifact schema from the OpenAPI specification
type GoogetArtifact struct {
	Architecture string `json:"architecture,omitempty"` // Output only. Operating system architecture of the artifact.
	Name string `json:"name,omitempty"` // Output only. The Artifact Registry resource name of the artifact.
	Packagename string `json:"packageName,omitempty"` // Output only. The GooGet package name of the artifact.
}

// ImportAptArtifactsGcsSource represents the ImportAptArtifactsGcsSource schema from the OpenAPI specification
type ImportAptArtifactsGcsSource struct {
	Uris []string `json:"uris,omitempty"` // Cloud Storage paths URI (e.g., gs://my_bucket//my_object).
	Usewildcards bool `json:"useWildcards,omitempty"` // Supports URI wildcards for matching multiple objects from a single URI.
}

// UpstreamCredentials represents the UpstreamCredentials schema from the OpenAPI specification
type UpstreamCredentials struct {
	Usernamepasswordcredentials UsernamePasswordCredentials `json:"usernamePasswordCredentials,omitempty"` // Username and password credentials.
}

// GoModule represents the GoModule schema from the OpenAPI specification
type GoModule struct {
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the Go module is created.
	Name string `json:"name,omitempty"` // The resource name of a Go module.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the Go module is updated.
	Version string `json:"version,omitempty"` // The version of the Go module. Must be a valid canonical version as defined in https://go.dev/ref/mod#glos-canonical-version.
}

// ImportGoogetArtifactsMetadata represents the ImportGoogetArtifactsMetadata schema from the OpenAPI specification
type ImportGoogetArtifactsMetadata struct {
}

// UploadAptArtifactResponse represents the UploadAptArtifactResponse schema from the OpenAPI specification
type UploadAptArtifactResponse struct {
	Aptartifacts []AptArtifact `json:"aptArtifacts,omitempty"` // The Apt artifacts updated.
}

// UpstreamPolicy represents the UpstreamPolicy schema from the OpenAPI specification
type UpstreamPolicy struct {
	Repository string `json:"repository,omitempty"` // A reference to the repository resource, for example: `projects/p1/locations/us-central1/repositories/repo1`.
	Id string `json:"id,omitempty"` // The user-provided ID of the upstream policy.
	Priority int `json:"priority,omitempty"` // Entries with a greater priority value take precedence in the pull order.
}

// GoogleDevtoolsArtifactregistryV1RemoteRepositoryConfigAptRepositoryPublicRepository represents the GoogleDevtoolsArtifactregistryV1RemoteRepositoryConfigAptRepositoryPublicRepository schema from the OpenAPI specification
type GoogleDevtoolsArtifactregistryV1RemoteRepositoryConfigAptRepositoryPublicRepository struct {
	Repositorypath string `json:"repositoryPath,omitempty"` // A custom field to define a path to a specific repository from the base.
	Repositorybase string `json:"repositoryBase,omitempty"` // A common public repository base for Apt.
}

// CleanupPolicyCondition represents the CleanupPolicyCondition schema from the OpenAPI specification
type CleanupPolicyCondition struct {
	Tagstate string `json:"tagState,omitempty"` // Match versions by tag status.
	Versionnameprefixes []string `json:"versionNamePrefixes,omitempty"` // Match versions by version name prefix. Applied on any prefix match.
	Newerthan string `json:"newerThan,omitempty"` // Match versions newer than a duration.
	Olderthan string `json:"olderThan,omitempty"` // Match versions older than a duration.
	Packagenameprefixes []string `json:"packageNamePrefixes,omitempty"` // Match versions by package prefix. Applied on any prefix match.
	Tagprefixes []string `json:"tagPrefixes,omitempty"` // Match versions by tag prefix. Applied on any prefix match.
}

// Policy represents the Policy schema from the OpenAPI specification
type Policy struct {
	Bindings []Binding `json:"bindings,omitempty"` // Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Etag string `json:"etag,omitempty"` // `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Version int `json:"version,omitempty"` // Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
}

// BatchDeleteVersionsRequest represents the BatchDeleteVersionsRequest schema from the OpenAPI specification
type BatchDeleteVersionsRequest struct {
	Names []string `json:"names,omitempty"` // Required. The names of the versions to delete. A maximum of 10000 versions can be deleted in a batch.
	Validateonly bool `json:"validateOnly,omitempty"` // If true, the request is performed without deleting data, following AIP-163.
}

// UploadGoogetArtifactMediaResponse represents the UploadGoogetArtifactMediaResponse schema from the OpenAPI specification
type UploadGoogetArtifactMediaResponse struct {
	Operation Operation `json:"operation,omitempty"` // This resource represents a long-running operation that is the result of a network API call.
}

// MavenRepository represents the MavenRepository schema from the OpenAPI specification
type MavenRepository struct {
	Publicrepository string `json:"publicRepository,omitempty"` // One of the publicly available Maven repositories supported by Artifact Registry.
}

// UploadYumArtifactMediaResponse represents the UploadYumArtifactMediaResponse schema from the OpenAPI specification
type UploadYumArtifactMediaResponse struct {
	Operation Operation `json:"operation,omitempty"` // This resource represents a long-running operation that is the result of a network API call.
}

// Expr represents the Expr schema from the OpenAPI specification
type Expr struct {
	Location string `json:"location,omitempty"` // Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Title string `json:"title,omitempty"` // Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Description string `json:"description,omitempty"` // Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Expression string `json:"expression,omitempty"` // Textual representation of an expression in Common Expression Language syntax.
}

// ImportAptArtifactsErrorInfo represents the ImportAptArtifactsErrorInfo schema from the OpenAPI specification
type ImportAptArtifactsErrorInfo struct {
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Gcssource ImportAptArtifactsGcsSource `json:"gcsSource,omitempty"` // Google Cloud Storage location where the artifacts currently reside.
}

// ListPackagesResponse represents the ListPackagesResponse schema from the OpenAPI specification
type ListPackagesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The token to retrieve the next page of packages, or empty if there are no more packages to return.
	Packages []Package `json:"packages,omitempty"` // The packages returned.
}

// ImportGoogetArtifactsErrorInfo represents the ImportGoogetArtifactsErrorInfo schema from the OpenAPI specification
type ImportGoogetArtifactsErrorInfo struct {
	Gcssource ImportGoogetArtifactsGcsSource `json:"gcsSource,omitempty"` // Google Cloud Storage location where the artifacts currently reside.
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
}

// NpmRepository represents the NpmRepository schema from the OpenAPI specification
type NpmRepository struct {
	Publicrepository string `json:"publicRepository,omitempty"` // One of the publicly available Npm repositories supported by Artifact Registry.
}

// RemoteRepositoryConfig represents the RemoteRepositoryConfig schema from the OpenAPI specification
type RemoteRepositoryConfig struct {
	Dockerrepository DockerRepository `json:"dockerRepository,omitempty"` // Configuration for a Docker remote repository.
	Mavenrepository MavenRepository `json:"mavenRepository,omitempty"` // Configuration for a Maven remote repository.
	Npmrepository NpmRepository `json:"npmRepository,omitempty"` // Configuration for a Npm remote repository.
	Pythonrepository PythonRepository `json:"pythonRepository,omitempty"` // Configuration for a Python remote repository.
	Upstreamcredentials UpstreamCredentials `json:"upstreamCredentials,omitempty"` // The credentials to access the remote repository.
	Yumrepository YumRepository `json:"yumRepository,omitempty"` // Configuration for a Yum remote repository.
	Aptrepository AptRepository `json:"aptRepository,omitempty"` // Configuration for an Apt remote repository.
	Description string `json:"description,omitempty"` // The description of the remote source.
}

// TestIamPermissionsResponse represents the TestIamPermissionsResponse schema from the OpenAPI specification
type TestIamPermissionsResponse struct {
	Permissions []string `json:"permissions,omitempty"` // A subset of `TestPermissionsRequest.permissions` that the caller is allowed.
}

// UploadGoogetArtifactResponse represents the UploadGoogetArtifactResponse schema from the OpenAPI specification
type UploadGoogetArtifactResponse struct {
	Googetartifacts []GoogetArtifact `json:"googetArtifacts,omitempty"` // The GooGet artifacts updated.
}

// ImportYumArtifactsGcsSource represents the ImportYumArtifactsGcsSource schema from the OpenAPI specification
type ImportYumArtifactsGcsSource struct {
	Uris []string `json:"uris,omitempty"` // Cloud Storage paths URI (e.g., gs://my_bucket//my_object).
	Usewildcards bool `json:"useWildcards,omitempty"` // Supports URI wildcards for matching multiple objects from a single URI.
}

// NpmPackage represents the NpmPackage schema from the OpenAPI specification
type NpmPackage struct {
	Version string `json:"version,omitempty"` // Version of this package.
	Createtime string `json:"createTime,omitempty"` // Output only. Time the package was created.
	Name string `json:"name,omitempty"` // Required. registry_location, project_id, repository_name and npm_package forms a unique package For example, "projects/test-project/locations/us-west4/repositories/test-repo/npmPackages/ npm_test:1.0.0", where "us-west4" is the registry_location, "test-project" is the project_id, "test-repo" is the repository_name and npm_test:1.0.0" is the npm package.
	Packagename string `json:"packageName,omitempty"` // Package for the artifact.
	Tags []string `json:"tags,omitempty"` // Tags attached to this package.
	Updatetime string `json:"updateTime,omitempty"` // Output only. Time the package was updated.
}

// SetIamPolicyRequest represents the SetIamPolicyRequest schema from the OpenAPI specification
type SetIamPolicyRequest struct {
	Policy Policy `json:"policy,omitempty"` // An Identity and Access Management (IAM) policy, which specifies access controls for Google Cloud resources. A `Policy` is a collection of `bindings`. A `binding` binds one or more `members`, or principals, to a single `role`. Principals can be user accounts, service accounts, Google groups, and domains (such as G Suite). A `role` is a named list of permissions; each `role` can be an IAM predefined role or a user-created custom role. For some types of Google Cloud resources, a `binding` can also specify a `condition`, which is a logical expression that allows access to a resource only if the expression evaluates to `true`. A condition can add constraints based on attributes of the request, the resource, or both. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies). **JSON example:** ``` { "bindings": [ { "role": "roles/resourcemanager.organizationAdmin", "members": [ "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role": "roles/resourcemanager.organizationViewer", "members": [ "user:eve@example.com" ], "condition": { "title": "expirable access", "description": "Does not grant access after Sep 2020", "expression": "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag": "BwWWja0YfJA=", "version": 3 } ``` **YAML example:** ``` bindings: - members: - user:mike@example.com - group:admins@example.com - domain:google.com - serviceAccount:my-project-id@appspot.gserviceaccount.com role: roles/resourcemanager.organizationAdmin - members: - user:eve@example.com role: roles/resourcemanager.organizationViewer condition: title: expirable access description: Does not grant access after Sep 2020 expression: request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA= version: 3 ``` For a description of IAM and its features, see the [IAM documentation](https://cloud.google.com/iam/docs/).
}

// DockerRepositoryConfig represents the DockerRepositoryConfig schema from the OpenAPI specification
type DockerRepositoryConfig struct {
	Immutabletags bool `json:"immutableTags,omitempty"` // The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Createtime string `json:"createTime,omitempty"` // The time when the package was created.
	Displayname string `json:"displayName,omitempty"` // The display name of the package.
	Name string `json:"name,omitempty"` // The name of the package, for example: `projects/p1/locations/us-central1/repositories/repo1/packages/pkg1`. If the package ID part contains slashes, the slashes are escaped.
	Updatetime string `json:"updateTime,omitempty"` // The time when the package was last updated. This includes publishing a new version of the package.
}

// MavenRepositoryConfig represents the MavenRepositoryConfig schema from the OpenAPI specification
type MavenRepositoryConfig struct {
	Versionpolicy string `json:"versionPolicy,omitempty"` // Version policy defines the versions that the registry will accept.
	Allowsnapshotoverwrites bool `json:"allowSnapshotOverwrites,omitempty"` // The repository with this flag will allow publishing the same snapshot versions.
}

// UploadGoogetArtifactMetadata represents the UploadGoogetArtifactMetadata schema from the OpenAPI specification
type UploadGoogetArtifactMetadata struct {
}

// UploadYumArtifactRequest represents the UploadYumArtifactRequest schema from the OpenAPI specification
type UploadYumArtifactRequest struct {
}

// UploadAptArtifactMetadata represents the UploadAptArtifactMetadata schema from the OpenAPI specification
type UploadAptArtifactMetadata struct {
}

// CleanupPolicy represents the CleanupPolicy schema from the OpenAPI specification
type CleanupPolicy struct {
	Id string `json:"id,omitempty"` // The user-provided ID of the cleanup policy.
	Mostrecentversions CleanupPolicyMostRecentVersions `json:"mostRecentVersions,omitempty"` // CleanupPolicyMostRecentVersions is an alternate condition of a CleanupPolicy for retaining a minimum number of versions.
	Action string `json:"action,omitempty"` // Policy action.
	Condition CleanupPolicyCondition `json:"condition,omitempty"` // CleanupPolicyCondition is a set of conditions attached to a CleanupPolicy. If multiple entries are set, all must be satisfied for the condition to be satisfied.
}

// DockerImage represents the DockerImage schema from the OpenAPI specification
type DockerImage struct {
	Uploadtime string `json:"uploadTime,omitempty"` // Time the image was uploaded.
	Uri string `json:"uri,omitempty"` // Required. URL to access the image. Example: us-west4-docker.pkg.dev/test-project/test-repo/nginx@sha256:e9954c1fc875017be1c3e36eca16be2d9e9bccc4bf072163515467d6a823c7cf
	Buildtime string `json:"buildTime,omitempty"` // The time this image was built. This field is returned as the 'metadata.buildTime' field in the Version resource. The build time is returned to the client as an RFC 3339 string, which can be easily used with the JavaScript Date constructor.
	Imagesizebytes string `json:"imageSizeBytes,omitempty"` // Calculated size of the image. This field is returned as the 'metadata.imageSizeBytes' field in the Version resource.
	Mediatype string `json:"mediaType,omitempty"` // Media type of this image, e.g. "application/vnd.docker.distribution.manifest.v2+json". This field is returned as the 'metadata.mediaType' field in the Version resource.
	Name string `json:"name,omitempty"` // Required. registry_location, project_id, repository_name and image id forms a unique image name:`projects//locations//repository//dockerImages/`. For example, "projects/test-project/locations/us-west4/repositories/test-repo/dockerImages/ nginx@sha256:e9954c1fc875017be1c3e36eca16be2d9e9bccc4bf072163515467d6a823c7cf", where "us-west4" is the registry_location, "test-project" is the project_id, "test-repo" is the repository_name and "nginx@sha256:e9954c1fc875017be1c3e36eca16be2d9e9bccc4bf072163515467d6a823c7cf" is the image's digest.
	Tags []string `json:"tags,omitempty"` // Tags attached to this image.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the docker image was last updated.
}

// VPCSCConfig represents the VPCSCConfig schema from the OpenAPI specification
type VPCSCConfig struct {
	Name string `json:"name,omitempty"` // The name of the project's VPC SC Config. Always of the form: projects/{projectID}/locations/{location}/vpcscConfig In update request: never set In response: always set
	Vpcscpolicy string `json:"vpcscPolicy,omitempty"` // The project per location VPC SC policy that defines the VPC SC behavior for the Remote Repository (Allow/Deny).
}

// ImportGoogetArtifactsRequest represents the ImportGoogetArtifactsRequest schema from the OpenAPI specification
type ImportGoogetArtifactsRequest struct {
	Gcssource ImportGoogetArtifactsGcsSource `json:"gcsSource,omitempty"` // Google Cloud Storage location where the artifacts currently reside.
}

// OperationMetadata represents the OperationMetadata schema from the OpenAPI specification
type OperationMetadata struct {
}

// UploadGoModuleMediaResponse represents the UploadGoModuleMediaResponse schema from the OpenAPI specification
type UploadGoModuleMediaResponse struct {
	Operation Operation `json:"operation,omitempty"` // This resource represents a long-running operation that is the result of a network API call.
}
