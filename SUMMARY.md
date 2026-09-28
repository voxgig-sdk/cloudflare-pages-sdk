# Cloudflare Pages

Cloudflare Pages: projects, deployments, domains, and build configurations.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 6 entities and 26 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### [PagesDeployment](docs/api/pages_deployment.html)

Results: Retry deployment response.; Rollback deployment response.; Create deployment tail response.; Create deployment response.; Get deployments response.; Get deployment logs response.; Get deployment info response.; Delete deployment tail response.; Delete deployment response.

SDK operations: `create`, `list`, `load`, `remove`.

Key fields to recognise:

- `filters`: Filters to apply to the tail session.
- `id`: Id of the deployment.
- `result`: A tail session for streaming logs from a Pages deployment.
- `success`: Whether the API call was successful.

### [PagesDomain](docs/api/pages_domain.html)

Results: Add domain response.; Get domains response.; Get domain response.; Delete domain response.; Patch domain response.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `name`: Fully qualified domain name for the Pages project, such as `example.com`.
- `success`: Whether the API call was successful.

### [PagesPagesAssetsCheckMissing](docs/api/pages_pages_assets_check_missing.html)

Results: Check missing response.

SDK operations: `create`.

Key fields to recognise:

- `hashes`: List of file content hashes to check for existence in the asset store.

### [PagesProject](docs/api/pages_project.html)

Results: Purge build cache response.; Connect project source response.; Create project response.; Get projects response.; Upload token response.; Get project response.; Delete project response.; Disconnect project source response.; Update project response.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `build_config`: Configs for the project build process.
- `deployment_configs`: Configs for deployments in a project.
- `id`: ID of the project.
- `name`: Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.
- `production_branch`: Production branch of the project. Used to identify production deployments.

### [Upload](docs/api/upload.html)

Results: Upload asset response.

SDK operations: `create`.

Key fields to recognise:

- `success`: Whether the API call was successful.

### [UpsertHash](docs/api/upsert_hash.html)

Results: Upsert hashes response.

SDK operations: `create`.

Key fields to recognise:

- `hashes`: List of file content hashes to register in the asset store.
- `success`: Whether the API call was successful.

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| [PagesDeployment](docs/api/pages_deployment.html) | `create` | `POST /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry` | Required |
| [PagesDeployment](docs/api/pages_deployment.html) | `create` | `POST /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/rollback` | Required |
| [PagesDeployment](docs/api/pages_deployment.html) | `create` | `POST /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/tails` | Required |
| [PagesDeployment](docs/api/pages_deployment.html) | `create` | `POST /accounts/{account_id}/pages/projects/{project_name}/deployments` | Required |
| [PagesDeployment](docs/api/pages_deployment.html) | `list` | `GET /accounts/{account_id}/pages/projects/{project_name}/deployments` | Required |
| [PagesDeployment](docs/api/pages_deployment.html) | `list` | `GET /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/history/logs` | Required |
| [PagesDeployment](docs/api/pages_deployment.html) | `load` | `GET /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}` | Required |
| [PagesDeployment](docs/api/pages_deployment.html) | `remove` | `DELETE /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/tails/{tail_id}` | Required |
| [PagesDeployment](docs/api/pages_deployment.html) | `remove` | `DELETE /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}` | Required |
| [PagesDomain](docs/api/pages_domain.html) | `create` | `POST /accounts/{account_id}/pages/projects/{project_name}/domains` | Required |
| [PagesDomain](docs/api/pages_domain.html) | `list` | `GET /accounts/{account_id}/pages/projects/{project_name}/domains` | Required |
| [PagesDomain](docs/api/pages_domain.html) | `load` | `GET /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}` | Required |
| [PagesDomain](docs/api/pages_domain.html) | `remove` | `DELETE /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}` | Required |
| [PagesDomain](docs/api/pages_domain.html) | `update` | `PATCH /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}` | Required |
| [PagesPagesAssetsCheckMissing](docs/api/pages_pages_assets_check_missing.html) | `create` | `POST /pages/assets/check-missing` | Required |
| [PagesProject](docs/api/pages_project.html) | `create` | `POST /accounts/{account_id}/pages/projects/{project_name}/purge_build_cache` | Required |
| [PagesProject](docs/api/pages_project.html) | `create` | `POST /accounts/{account_id}/pages/projects/{project_name}/source` | Required |
| [PagesProject](docs/api/pages_project.html) | `create` | `POST /accounts/{account_id}/pages/projects` | Required |
| [PagesProject](docs/api/pages_project.html) | `list` | `GET /accounts/{account_id}/pages/projects` | Required |
| [PagesProject](docs/api/pages_project.html) | `list` | `GET /accounts/{account_id}/pages/projects/{project_name}/upload-token` | Required |
| [PagesProject](docs/api/pages_project.html) | `load` | `GET /accounts/{account_id}/pages/projects/{project_name}` | Required |
| [PagesProject](docs/api/pages_project.html) | `remove` | `DELETE /accounts/{account_id}/pages/projects/{project_name}` | Required |
| [PagesProject](docs/api/pages_project.html) | `remove` | `DELETE /accounts/{account_id}/pages/projects/{project_name}/source` | Required |
| [PagesProject](docs/api/pages_project.html) | `update` | `PATCH /accounts/{account_id}/pages/projects/{project_name}` | Required |
| [Upload](docs/api/upload.html) | `create` | `POST /pages/assets/upload` | Required |
| [UpsertHash](docs/api/upsert_hash.html) | `create` | `POST /pages/assets/upsert-hashes` | Required |

## Connect to the API

- Client API: `https://api.cloudflare.com/client/v4`

The default credential is sent in the `X-Auth-Email` header.

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| [Golang](docs/sdks/go.html) | `go/` | Build from source |
| [Lua](docs/sdks/lua.html) | `lua/` | Build from source |
| [PHP](docs/sdks/php.html) | `php/` | Build from source |
| [Python](docs/sdks/py.html) | `py/` | Build from source |
| [Ruby](docs/sdks/rb.html) | `rb/` | Build from source |
| [TypeScript](docs/sdks/ts.html) | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### [Go CLI](docs/tools/go-cli.html)

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### [Go MCP server](docs/tools/go-mcp.html)

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `cloudflare-pages_list`: List records for an entity. Supported entities: `pages_deployment`, `pages_domain`, `pages_project`.
- `cloudflare-pages_load`: Load one record for an entity. Supported entities: `pages_deployment`, `pages_domain`, `pages_project`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- [`test`](docs/features/test.html): In-memory mock transport for testing without a live server

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the [first-call guide](docs/guides/first-call.html) for the setup sequence.
- Read the [authentication guide](docs/guides/authentication.html) before using protected routes.
- Use the [API reference](docs/api/index.html) for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

