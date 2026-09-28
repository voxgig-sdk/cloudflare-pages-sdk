# CloudflarePages Golang SDK Reference

Complete API reference for the CloudflarePages Golang SDK.


## CloudflarePagesSDK

### Constructor

```go
func NewCloudflarePagesSDK(options map[string]any) *CloudflarePagesSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *CloudflarePagesSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *CloudflarePagesSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `PagesDeployment(data map[string]any) CloudflarePagesEntity`

Create a new `PagesDeployment` entity instance. Pass `nil` for no initial data.

#### `PagesDomain(data map[string]any) CloudflarePagesEntity`

Create a new `PagesDomain` entity instance. Pass `nil` for no initial data.

#### `PagesPagesAssetsCheckMissing(data map[string]any) CloudflarePagesEntity`

Create a new `PagesPagesAssetsCheckMissing` entity instance. Pass `nil` for no initial data.

#### `PagesProject(data map[string]any) CloudflarePagesEntity`

Create a new `PagesProject` entity instance. Pass `nil` for no initial data.

#### `Upload(data map[string]any) CloudflarePagesEntity`

Create a new `Upload` entity instance. Pass `nil` for no initial data.

#### `UpsertHash(data map[string]any) CloudflarePagesEntity`

Create a new `UpsertHash` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## PagesDeploymentEntity

```go
pagesDeployment := client.PagesDeployment(nil)
fmt.Println(pagesDeployment.GetName()) // "pages_deployment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `[]any` | Yes |  |
| `filters` | `[]any` | No | Filters to apply to the tail session. |
| `id` | `string` | No |  |
| `messages` | `[]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PagesDeployment(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PagesDeployment(nil).Load(map[string]any{"id": "pages_deployment_id", "account_id": "account_id", "project_id": "project_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PagesDeployment(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "errors": []any{},
    "messages": []any{},
    "result": map[string]any{},
    "success": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.PagesDeployment(nil).Remove(map[string]any{"id": "pages_deployment_id", "account_id": "account_id", "project_id": "project_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PagesDeploymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PagesDomainEntity

```go
pagesDomain := client.PagesDomain(nil)
fmt.Println(pagesDomain.GetName()) // "pages_domain"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `[]any` | Yes |  |
| `id` | `string` | No |  |
| `messages` | `[]any` | Yes |  |
| `name` | `string` | Yes | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `map[string]any` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PagesDomain(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PagesDomain(nil).Load(map[string]any{"id": "pages_domain_id", "account_id": "account_id", "project_id": "project_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PagesDomain(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "project_name": "example_project_name",
    "errors": []any{},
    "messages": []any{},
    "name": "example_name",
    "result": map[string]any{},
    "success": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.PagesDomain(nil).Update(map[string]any{
    "id": "pages_domain_id",
    "account_id": "account_id",
    "project_id": "project_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.PagesDomain(nil).Remove(map[string]any{"id": "pages_domain_id", "account_id": "account_id", "project_id": "project_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PagesDomainEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PagesPagesAssetsCheckMissingEntity

```go
pagesPagesAssetsCheckMissing := client.PagesPagesAssetsCheckMissing(nil)
fmt.Println(pagesPagesAssetsCheckMissing.GetName()) // "pages_pages_assets_check_missing"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `hashes` | `[]any` | Yes | List of file content hashes to check for existence in the asset store. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PagesPagesAssetsCheckMissing(nil).Create(map[string]any{
    "hashes": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PagesPagesAssetsCheckMissingEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PagesProjectEntity

```go
pagesProject := client.PagesProject(nil)
fmt.Println(pagesProject.GetName()) // "pages_project"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_config` | `map[string]any` | No | Configs for the project build process. |
| `deployment_configs` | `map[string]any` | No | Configs for deployments in a project. |
| `errors` | `[]any` | Yes |  |
| `id` | `string` | No |  |
| `messages` | `[]any` | Yes |  |
| `name` | `string` | Yes | Name for the Pages project. |
| `production_branch` | `string` | Yes | Production branch of the project. |
| `result` | `map[string]any` | Yes |  |
| `source` | `map[string]any` | Yes | Configs for the project source control. |
| `success` | `bool` | Yes | Whether the API call was successful. |

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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PagesProject(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PagesProject(nil).Load(map[string]any{"id": "pages_project_id", "account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PagesProject(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "errors": []any{},
    "messages": []any{},
    "name": "example_name",
    "production_branch": "example_production_branch",
    "result": map[string]any{},
    "source": map[string]any{},
    "success": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.PagesProject(nil).Update(map[string]any{
    "id": "pages_project_id",
    "account_id": "account_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.PagesProject(nil).Remove(map[string]any{"id": "pages_project_id", "account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PagesProjectEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## UploadEntity

```go
upload := client.Upload(nil)
fmt.Println(upload.GetName()) // "upload"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `[]any` | Yes |  |
| `messages` | `[]any` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Upload(nil).Create(map[string]any{
    "errors": []any{},
    "messages": []any{},
    "success": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `UploadEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## UpsertHashEntity

```go
upsertHash := client.UpsertHash(nil)
fmt.Println(upsertHash.GetName()) // "upsert_hash"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `[]any` | Yes |  |
| `hashes` | `[]any` | Yes | List of file content hashes to register in the asset store. |
| `messages` | `[]any` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.UpsertHash(nil).Create(map[string]any{
    "errors": []any{},
    "hashes": []any{},
    "messages": []any{},
    "success": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `UpsertHashEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```go
client := sdk.NewCloudflarePagesSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
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

