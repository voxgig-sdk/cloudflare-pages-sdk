# CloudflarePages Lua SDK Reference

Complete API reference for the CloudflarePages Lua SDK.


## CloudflarePagesSDK

### Constructor

```lua
local sdk = require("cloudflare-pages_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `PagesDeployment(data)`

Create a new `PagesDeployment` entity instance. Pass `nil` for no initial data.

#### `PagesDomain(data)`

Create a new `PagesDomain` entity instance. Pass `nil` for no initial data.

#### `PagesPagesAssetsCheckMissing(data)`

Create a new `PagesPagesAssetsCheckMissing` entity instance. Pass `nil` for no initial data.

#### `PagesProject(data)`

Create a new `PagesProject` entity instance. Pass `nil` for no initial data.

#### `Upload(data)`

Create a new `Upload` entity instance. Pass `nil` for no initial data.

#### `UpsertHash(data)`

Create a new `UpsertHash` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## PagesDeploymentEntity

```lua
local pages_deployment = client:PagesDeployment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `table` | Yes |  |
| `filters` | `table` | No | Filters to apply to the tail session. |
| `id` | `string` | No |  |
| `messages` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `success` | `boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PagesDeployment():create({
  account_id = --[[ string ]],
  errors = --[[ table ]],
  messages = --[[ table ]],
  result = --[[ table ]],
  success = --[[ boolean ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PagesDeployment():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PagesDeployment():load({ id = "pages_deployment_id", account_id = "account_id", project_id = "project_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:PagesDeployment():remove({ id = "pages_deployment_id", account_id = "account_id", project_id = "project_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PagesDeploymentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PagesDomainEntity

```lua
local pages_domain = client:PagesDomain(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `table` | Yes |  |
| `id` | `string` | No |  |
| `messages` | `table` | Yes |  |
| `name` | `string` | Yes | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `table` | Yes |  |
| `success` | `boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PagesDomain():create({
  account_id = --[[ string ]],
  project_name = --[[ string ]],
  errors = --[[ table ]],
  messages = --[[ table ]],
  name = --[[ string ]],
  result = --[[ table ]],
  success = --[[ boolean ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PagesDomain():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PagesDomain():load({ id = "pages_domain_id", account_id = "account_id", project_id = "project_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:PagesDomain():remove({ id = "pages_domain_id", account_id = "account_id", project_id = "project_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:PagesDomain():update({
  id = "pages_domain_id",
  account_id = "account_id",
  project_id = "project_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PagesDomainEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PagesPagesAssetsCheckMissingEntity

```lua
local pages_pages_assets_check_missing = client:PagesPagesAssetsCheckMissing(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `hashes` | `table` | Yes | List of file content hashes to check for existence in the asset store. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PagesPagesAssetsCheckMissing():create({
  hashes = --[[ table ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PagesPagesAssetsCheckMissingEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PagesProjectEntity

```lua
local pages_project = client:PagesProject(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_config` | `table` | No | Configs for the project build process. |
| `deployment_configs` | `table` | No | Configs for deployments in a project. |
| `errors` | `table` | Yes |  |
| `id` | `string` | No |  |
| `messages` | `table` | Yes |  |
| `name` | `string` | Yes | Name for the Pages project. |
| `production_branch` | `string` | Yes | Production branch of the project. |
| `result` | `table` | Yes |  |
| `source` | `table` | Yes | Configs for the project source control. |
| `success` | `boolean` | Yes | Whether the API call was successful. |

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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PagesProject():create({
  account_id = --[[ string ]],
  errors = --[[ table ]],
  messages = --[[ table ]],
  name = --[[ string ]],
  production_branch = --[[ string ]],
  result = --[[ table ]],
  source = --[[ table ]],
  success = --[[ boolean ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PagesProject():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PagesProject():load({ id = "pages_project_id", account_id = "account_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:PagesProject():remove({ id = "pages_project_id", account_id = "account_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:PagesProject():update({
  id = "pages_project_id",
  account_id = "account_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PagesProjectEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## UploadEntity

```lua
local upload = client:Upload(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `table` | Yes |  |
| `messages` | `table` | Yes |  |
| `success` | `boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Upload():create({
  errors = --[[ table ]],
  messages = --[[ table ]],
  success = --[[ boolean ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UploadEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## UpsertHashEntity

```lua
local upsert_hash = client:UpsertHash(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `table` | Yes |  |
| `hashes` | `table` | Yes | List of file content hashes to register in the asset store. |
| `messages` | `table` | Yes |  |
| `success` | `boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:UpsertHash():create({
  errors = --[[ table ]],
  hashes = --[[ table ]],
  messages = --[[ table ]],
  success = --[[ boolean ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpsertHashEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
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

