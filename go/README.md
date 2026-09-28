# CloudflarePages Golang SDK



The Golang SDK for the CloudflarePages API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.PagesDeployment(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/cloudflare-pages-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/cloudflare-pages-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/cloudflare-pages-sdk/go=../cloudflare-pages-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/cloudflare-pages-sdk/go"
)

func main() {
    client := sdk.NewCloudflarePagesSDK(map[string]any{
        "apikey": os.Getenv("CLOUDFLARE_PAGES_APIKEY"),
    })

    // List pagesDeployment records — the value is the array of records itself.
    pagesDeployments, err := client.PagesDeployment(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range pagesDeployments.([]any) {
        fmt.Println(item)
    }

    // Load a single pagesDeployment — the value is the loaded record.
    pagesDeployment, err := client.PagesDeployment(nil).Load(map[string]any{"id": "example_id", "account_id": "example_account_id", "project_id": "example_project_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(pagesDeployment)

    // Create a pagesDeployment.
    created, err := client.PagesDeployment(nil).Create(map[string]any{"account_id": "example_account_id", "errors": []any{}, "messages": []any{}, "result": map[string]any{}, "success": true}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)

    // Remove a pagesDeployment.
    removed, err := client.PagesDeployment(nil).Remove(map[string]any{"id": "example_id", "account_id": "example_account_id", "project_id": "example_project_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(removed)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
pagesdomains, err := client.PagesDomain(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = pagesdomains
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

pagesDomain, err := client.PagesDomain(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(pagesDomain) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewCloudflarePagesSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
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
cd go && go test ./test/...
```


## Reference

### NewCloudflarePagesSDK

```go
func NewCloudflarePagesSDK(options map[string]any) *CloudflarePagesSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *CloudflarePagesSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### CloudflarePagesSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `PagesDeployment` | `(data map[string]any) CloudflarePagesEntity` | Create a PagesDeployment entity instance. |
| `PagesDomain` | `(data map[string]any) CloudflarePagesEntity` | Create a PagesDomain entity instance. |
| `PagesPagesAssetsCheckMissing` | `(data map[string]any) CloudflarePagesEntity` | Create a PagesPagesAssetsCheckMissing entity instance. |
| `PagesProject` | `(data map[string]any) CloudflarePagesEntity` | Create a PagesProject entity instance. |
| `Upload` | `(data map[string]any) CloudflarePagesEntity` | Create an Upload entity instance. |
| `UpsertHash` | `(data map[string]any) CloudflarePagesEntity` | Create an UpsertHash entity instance. |

### Entity interface (CloudflarePagesEntity)

All entities implement the `CloudflarePagesEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    pagesDeployment, err := client.PagesDeployment(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // pagesDeployment is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### PagesDeployment

| Field | Description |
| --- | --- |
| `"errors"` |  |
| `"filters"` | Filters to apply to the tail session. |
| `"id"` |  |
| `"messages"` |  |
| `"result"` |  |
| `"success"` | Whether the API call was successful. |

Operations: Create, List, Load, Remove.

API path: `/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry`

#### PagesDomain

| Field | Description |
| --- | --- |
| `"errors"` |  |
| `"id"` |  |
| `"messages"` |  |
| `"name"` | Fully qualified domain name for the Pages project, such as `example.com`. |
| `"result"` |  |
| `"success"` | Whether the API call was successful. |

Operations: Create, List, Load, Remove, Update.

API path: `/accounts/{account_id}/pages/projects/{project_name}/domains`

#### PagesPagesAssetsCheckMissing

| Field | Description |
| --- | --- |
| `"hashes"` | List of file content hashes to check for existence in the asset store. |

Operations: Create.

API path: `/pages/assets/check-missing`

#### PagesProject

| Field | Description |
| --- | --- |
| `"build_config"` | Configs for the project build process. |
| `"deployment_configs"` | Configs for deployments in a project. |
| `"errors"` |  |
| `"id"` |  |
| `"messages"` |  |
| `"name"` | Name for the Pages project. |
| `"production_branch"` | Production branch of the project. |
| `"result"` |  |
| `"source"` | Configs for the project source control. |
| `"success"` | Whether the API call was successful. |

Operations: Create, List, Load, Remove, Update.

API path: `/accounts/{account_id}/pages/projects/{project_name}/purge_build_cache`

#### Upload

| Field | Description |
| --- | --- |
| `"errors"` |  |
| `"messages"` |  |
| `"success"` | Whether the API call was successful. |

Operations: Create.

API path: `/pages/assets/upload`

#### UpsertHash

| Field | Description |
| --- | --- |
| `"errors"` |  |
| `"hashes"` | List of file content hashes to register in the asset store. |
| `"messages"` |  |
| `"success"` | Whether the API call was successful. |

Operations: Create.

API path: `/pages/assets/upsert-hashes`



## Entities


### PagesDeployment

Create an instance: `pagesDeployment := client.PagesDeployment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `[]any` |  |
| `filters` | `[]any` | Filters to apply to the tail session. |
| `id` | `string` |  |
| `messages` | `[]any` |  |
| `result` | `map[string]any` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```go
pagesDeployment, err := client.PagesDeployment(nil).Load(map[string]any{"id": "pages_deployment_id", "account_id": "account_id", "project_id": "project_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(pagesDeployment) // the loaded record
```

#### Example: List

```go
pagesDeployments, err := client.PagesDeployment(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(pagesDeployments) // the array of records
```

#### Example: Create

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


### PagesDomain

Create an instance: `pagesDomain := client.PagesDomain(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `[]any` |  |
| `id` | `string` |  |
| `messages` | `[]any` |  |
| `name` | `string` | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `map[string]any` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```go
pagesDomain, err := client.PagesDomain(nil).Load(map[string]any{"id": "pages_domain_id", "account_id": "account_id", "project_id": "project_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(pagesDomain) // the loaded record
```

#### Example: List

```go
pagesDomains, err := client.PagesDomain(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(pagesDomains) // the array of records
```

#### Example: Create

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


### PagesPagesAssetsCheckMissing

Create an instance: `pagesPagesAssetsCheckMissing := client.PagesPagesAssetsCheckMissing(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `hashes` | `[]any` | List of file content hashes to check for existence in the asset store. |

#### Example: Create

```go
result, err := client.PagesPagesAssetsCheckMissing(nil).Create(map[string]any{
    "hashes": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PagesProject

Create an instance: `pagesProject := client.PagesProject(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `build_config` | `map[string]any` | Configs for the project build process. |
| `deployment_configs` | `map[string]any` | Configs for deployments in a project. |
| `errors` | `[]any` |  |
| `id` | `string` |  |
| `messages` | `[]any` |  |
| `name` | `string` | Name for the Pages project. |
| `production_branch` | `string` | Production branch of the project. |
| `result` | `map[string]any` |  |
| `source` | `map[string]any` | Configs for the project source control. |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```go
pagesProject, err := client.PagesProject(nil).Load(map[string]any{"id": "pages_project_id", "account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(pagesProject) // the loaded record
```

#### Example: List

```go
pagesProjects, err := client.PagesProject(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(pagesProjects) // the array of records
```

#### Example: Create

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


### Upload

Create an instance: `upload := client.Upload(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `[]any` |  |
| `messages` | `[]any` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Create

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


### UpsertHash

Create an instance: `upsertHash := client.UpsertHash(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `[]any` |  |
| `hashes` | `[]any` | List of file content hashes to register in the asset store. |
| `messages` | `[]any` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Create

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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/cloudflare-pages-sdk/go/
├── cloudflare-pages.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/cloudflare-pages-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
pagesdomain := client.PagesDomain(nil)
pagesdomain.List(nil, nil)

// pagesdomain.Data() now returns the pagesdomain data from the last list
// pagesdomain.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
