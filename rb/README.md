# CloudflarePages Ruby SDK



The Ruby SDK for the CloudflarePages API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.PagesDeployment` — with named operations (`list`/`load`/`create`/`update`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/cloudflare-pages-sdk/releases](https://github.com/voxgig-sdk/cloudflare-pages-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "CloudflarePages_sdk"

client = CloudflarePagesSDK.new({
  "apikey" => ENV["CLOUDFLARE_PAGES_APIKEY"],
})
```

### 2. List pagesdeployment records

```ruby
begin
  # list returns an Array of PagesDeployment records — iterate directly.
  pagesdeployments = client.PagesDeployment.list
  pagesdeployments.each do |item|
    puts "#{item["id"]} #{item["errors"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load a pagesdeployment

```ruby
begin
  # load returns the ENTITY — call data_get for the PagesDeployment record (raises on error).
  pagesdeployment = client.PagesDeployment.load({ "id" => "example_id", "account_id" => "example_account_id", "project_id" => "example_project_id" })
  puts pagesdeployment
rescue => err
  warn "load failed: #{err}"
end
```

### 4. Create, update, and remove

```ruby
# create returns the ENTITY — call data_get for the created PagesDeployment record.
created = client.PagesDeployment.create({ "account_id" => "example_account_id", "errors" => [], "messages" => [], "result" => {}, "success" => true })

# Remove
client.PagesDeployment.remove({ "id" => created.data_get["id"], "account_id" => "example_account_id", "project_id" => "example_project_id" })
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  pagesdomains = client.PagesDomain.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = CloudflarePagesSDK.test({
  "entity" => { "pagesdomain" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
pagesdomain = client.PagesDomain.list()
puts pagesdomain
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = CloudflarePagesSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
CLOUDFLARE_PAGES_TEST_LIVE=TRUE
CLOUDFLARE_PAGES_APIKEY=<your-key>
```

Then run:

```bash
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### CloudflarePagesSDK

```ruby
require_relative "CloudflarePages_sdk"
client = CloudflarePagesSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = CloudflarePagesSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### CloudflarePagesSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `PagesDeployment` | `(data) -> PagesDeploymentEntity` | Create a PagesDeployment entity instance. |
| `PagesDomain` | `(data) -> PagesDomainEntity` | Create a PagesDomain entity instance. |
| `PagesPagesAssetsCheckMissing` | `(data) -> PagesPagesAssetsCheckMissingEntity` | Create a PagesPagesAssetsCheckMissing entity instance. |
| `PagesProject` | `(data) -> PagesProjectEntity` | Create a PagesProject entity instance. |
| `Upload` | `(data) -> UploadEntity` | Create an Upload entity instance. |
| `UpsertHash` | `(data) -> UpsertHashEntity` | Create an UpsertHash entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `CloudflarePagesError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### PagesDeployment

| Field | Description |
| --- | --- |
| `errors` |  |
| `filters` | Filters to apply to the tail session. |
| `id` |  |
| `messages` |  |
| `result` |  |
| `success` | Whether the API call was successful. |

Operations: Create, List, Load, Remove.

API path: `/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry`

#### PagesDomain

| Field | Description |
| --- | --- |
| `errors` |  |
| `id` |  |
| `messages` |  |
| `name` | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` |  |
| `success` | Whether the API call was successful. |

Operations: Create, List, Load, Remove, Update.

API path: `/accounts/{account_id}/pages/projects/{project_name}/domains`

#### PagesPagesAssetsCheckMissing

| Field | Description |
| --- | --- |
| `hashes` | List of file content hashes to check for existence in the asset store. |

Operations: Create.

API path: `/pages/assets/check-missing`

#### PagesProject

| Field | Description |
| --- | --- |
| `build_config` | Configs for the project build process. |
| `deployment_configs` | Configs for deployments in a project. |
| `errors` |  |
| `id` |  |
| `messages` |  |
| `name` | Name for the Pages project. |
| `production_branch` | Production branch of the project. |
| `result` |  |
| `source` | Configs for the project source control. |
| `success` | Whether the API call was successful. |

Operations: Create, List, Load, Remove, Update.

API path: `/accounts/{account_id}/pages/projects/{project_name}/purge_build_cache`

#### Upload

| Field | Description |
| --- | --- |
| `errors` |  |
| `messages` |  |
| `success` | Whether the API call was successful. |

Operations: Create.

API path: `/pages/assets/upload`

#### UpsertHash

| Field | Description |
| --- | --- |
| `errors` |  |
| `hashes` | List of file content hashes to register in the asset store. |
| `messages` |  |
| `success` | Whether the API call was successful. |

Operations: Create.

API path: `/pages/assets/upsert-hashes`



## Entities


### PagesDeployment

Create an instance: `pages_deployment = client.PagesDeployment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `Array` |  |
| `filters` | `Array` | Filters to apply to the tail session. |
| `id` | `String` |  |
| `messages` | `Array` |  |
| `result` | `Hash` |  |
| `success` | `Boolean` | Whether the API call was successful. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PagesDeployment record (raises on error).
pages_deployment = client.PagesDeployment.load({ "id" => "pages_deployment_id", "account_id" => "account_id", "project_id" => "project_id" })
```

#### Example: List

```ruby
# list returns an Array of PagesDeployment records (raises on error).
pages_deployments = client.PagesDeployment.list
```

#### Example: Create

```ruby
pages_deployment = client.PagesDeployment.create({
  "account_id" => "example_account_id", # String
  "errors" => [], # Array
  "messages" => [], # Array
  "result" => {}, # Hash
  "success" => true, # Boolean
})
```


### PagesDomain

Create an instance: `pages_domain = client.PagesDomain`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `Array` |  |
| `id` | `String` |  |
| `messages` | `Array` |  |
| `name` | `String` | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `Hash` |  |
| `success` | `Boolean` | Whether the API call was successful. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PagesDomain record (raises on error).
pages_domain = client.PagesDomain.load({ "id" => "pages_domain_id", "account_id" => "account_id", "project_id" => "project_id" })
```

#### Example: List

```ruby
# list returns an Array of PagesDomain records (raises on error).
pages_domains = client.PagesDomain.list
```

#### Example: Create

```ruby
pages_domain = client.PagesDomain.create({
  "account_id" => "example_account_id", # String
  "project_name" => "example_project_name", # String
  "errors" => [], # Array
  "messages" => [], # Array
  "name" => "example_name", # String
  "result" => {}, # Hash
  "success" => true, # Boolean
})
```


### PagesPagesAssetsCheckMissing

Create an instance: `pages_pages_assets_check_missing = client.PagesPagesAssetsCheckMissing`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `hashes` | `Array` | List of file content hashes to check for existence in the asset store. |

#### Example: Create

```ruby
pages_pages_assets_check_missing = client.PagesPagesAssetsCheckMissing.create({
  "hashes" => [], # Array
})
```


### PagesProject

Create an instance: `pages_project = client.PagesProject`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `build_config` | `Hash` | Configs for the project build process. |
| `deployment_configs` | `Hash` | Configs for deployments in a project. |
| `errors` | `Array` |  |
| `id` | `String` |  |
| `messages` | `Array` |  |
| `name` | `String` | Name for the Pages project. |
| `production_branch` | `String` | Production branch of the project. |
| `result` | `Hash` |  |
| `source` | `Hash` | Configs for the project source control. |
| `success` | `Boolean` | Whether the API call was successful. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PagesProject record (raises on error).
pages_project = client.PagesProject.load({ "id" => "pages_project_id", "account_id" => "account_id" })
```

#### Example: List

```ruby
# list returns an Array of PagesProject records (raises on error).
pages_projects = client.PagesProject.list
```

#### Example: Create

```ruby
pages_project = client.PagesProject.create({
  "account_id" => "example_account_id", # String
  "errors" => [], # Array
  "messages" => [], # Array
  "name" => "example_name", # String
  "production_branch" => "example_production_branch", # String
  "result" => {}, # Hash
  "source" => {}, # Hash
  "success" => true, # Boolean
})
```


### Upload

Create an instance: `upload = client.Upload`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `Array` |  |
| `messages` | `Array` |  |
| `success` | `Boolean` | Whether the API call was successful. |

#### Example: Create

```ruby
upload = client.Upload.create({
  "errors" => [], # Array
  "messages" => [], # Array
  "success" => true, # Boolean
})
```


### UpsertHash

Create an instance: `upsert_hash = client.UpsertHash`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `Array` |  |
| `hashes` | `Array` | List of file content hashes to register in the asset store. |
| `messages` | `Array` |  |
| `success` | `Boolean` | Whether the API call was successful. |

#### Example: Create

```ruby
upsert_hash = client.UpsertHash.create({
  "errors" => [], # Array
  "hashes" => [], # Array
  "messages" => [], # Array
  "success" => true, # Boolean
})
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | Test transport |

### test

Test transport.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── CloudflarePages_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── schema.rb                  -- Generated option + entity specs
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`CloudflarePages_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
pagesdomain = client.PagesDomain
pagesdomain.list()

# pagesdomain.data_get now returns the pagesdomain data from the last list
# pagesdomain.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
