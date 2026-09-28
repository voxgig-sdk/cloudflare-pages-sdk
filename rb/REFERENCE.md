# CloudflarePages Ruby SDK Reference

Complete API reference for the CloudflarePages Ruby SDK.


## CloudflarePagesSDK

### Constructor

```ruby
require_relative 'CloudflarePages_sdk'

client = CloudflarePagesSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CloudflarePagesSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = CloudflarePagesSDK.test
```


### Instance Methods

#### `PagesDeployment(data = nil)`

Create a new `PagesDeployment` entity instance. Pass `nil` for no initial data.

#### `PagesDomain(data = nil)`

Create a new `PagesDomain` entity instance. Pass `nil` for no initial data.

#### `PagesPagesAssetsCheckMissing(data = nil)`

Create a new `PagesPagesAssetsCheckMissing` entity instance. Pass `nil` for no initial data.

#### `PagesProject(data = nil)`

Create a new `PagesProject` entity instance. Pass `nil` for no initial data.

#### `Upload(data = nil)`

Create a new `Upload` entity instance. Pass `nil` for no initial data.

#### `UpsertHash(data = nil)`

Create a new `UpsertHash` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## PagesDeploymentEntity

```ruby
pages_deployment = client.PagesDeployment
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `Array` | Yes |  |
| `filters` | `Array` | No | Filters to apply to the tail session. |
| `id` | `String` | No |  |
| `messages` | `Array` | Yes |  |
| `result` | `Hash` | Yes |  |
| `success` | `Boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PagesDeployment.create({
  "account_id" => "example_account_id", # String
  "errors" => [], # Array
  "messages" => [], # Array
  "result" => {}, # Hash
  "success" => true, # Boolean
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PagesDeployment.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.PagesDeployment.load({ "id" => "pages_deployment_id", "account_id" => "account_id", "project_id" => "project_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.PagesDeployment.remove({ "id" => "pages_deployment_id", "account_id" => "account_id", "project_id" => "project_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PagesDeploymentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PagesDomainEntity

```ruby
pages_domain = client.PagesDomain
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `Array` | Yes |  |
| `id` | `String` | No |  |
| `messages` | `Array` | Yes |  |
| `name` | `String` | Yes | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `Hash` | Yes |  |
| `success` | `Boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PagesDomain.create({
  "account_id" => "example_account_id", # String
  "project_name" => "example_project_name", # String
  "errors" => [], # Array
  "messages" => [], # Array
  "name" => "example_name", # String
  "result" => {}, # Hash
  "success" => true, # Boolean
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PagesDomain.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.PagesDomain.load({ "id" => "pages_domain_id", "account_id" => "account_id", "project_id" => "project_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.PagesDomain.remove({ "id" => "pages_domain_id", "account_id" => "account_id", "project_id" => "project_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.PagesDomain.update({
  "id" => "pages_domain_id",
  "account_id" => "account_id",
  "project_id" => "project_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PagesDomainEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PagesPagesAssetsCheckMissingEntity

```ruby
pages_pages_assets_check_missing = client.PagesPagesAssetsCheckMissing
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `hashes` | `Array` | Yes | List of file content hashes to check for existence in the asset store. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PagesPagesAssetsCheckMissing.create({
  "hashes" => [], # Array
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PagesPagesAssetsCheckMissingEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PagesProjectEntity

```ruby
pages_project = client.PagesProject
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_config` | `Hash` | No | Configs for the project build process. |
| `deployment_configs` | `Hash` | No | Configs for deployments in a project. |
| `errors` | `Array` | Yes |  |
| `id` | `String` | No |  |
| `messages` | `Array` | Yes |  |
| `name` | `String` | Yes | Name for the Pages project. |
| `production_branch` | `String` | Yes | Production branch of the project. |
| `result` | `Hash` | Yes |  |
| `source` | `Hash` | Yes | Configs for the project source control. |
| `success` | `Boolean` | Yes | Whether the API call was successful. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `build_config` | - | - | - | - | - |
| `deployment_configs` | - | - | - | - | - |
| `errors` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `messages` | - | - | - | - | - |
| `name` | - | - | - | Yes | - |
| `production_branch` | - | - | - | Yes | - |
| `result` | - | - | - | - | - |
| `source` | - | - | - | - | - |
| `success` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PagesProject.create({
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

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PagesProject.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.PagesProject.load({ "id" => "pages_project_id", "account_id" => "account_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.PagesProject.remove({ "id" => "pages_project_id", "account_id" => "account_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.PagesProject.update({
  "id" => "pages_project_id",
  "account_id" => "account_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PagesProjectEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## UploadEntity

```ruby
upload = client.Upload
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `Array` | Yes |  |
| `messages` | `Array` | Yes |  |
| `success` | `Boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Upload.create({
  "errors" => [], # Array
  "messages" => [], # Array
  "success" => true, # Boolean
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `UploadEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## UpsertHashEntity

```ruby
upsert_hash = client.UpsertHash
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `Array` | Yes |  |
| `hashes` | `Array` | Yes | List of file content hashes to register in the asset store. |
| `messages` | `Array` | Yes |  |
| `success` | `Boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.UpsertHash.create({
  "errors" => [], # Array
  "hashes" => [], # Array
  "messages" => [], # Array
  "success" => true, # Boolean
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `UpsertHashEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```ruby
client = CloudflarePagesSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

