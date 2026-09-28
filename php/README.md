# CloudflarePages PHP SDK



The PHP SDK for the CloudflarePages API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->PagesDeployment()` — with named operations (`list`/`load`/`create`/`update`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/cloudflare-pages-sdk/releases](https://github.com/voxgig-sdk/cloudflare-pages-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'cloudflarepages_sdk.php';

$client = new CloudflarePagesSDK([
    "apikey" => getenv("CLOUDFLARE_PAGES_APIKEY"),
]);
```

### 2. List pagesdeployment records

```php
try {
    // list() returns entity instances; data_get() reads each record.
    $pagesdeployments = $client->PagesDeployment()->list();
    foreach ($pagesdeployments as $record) {
        $item = $record->data_get();
        echo $item["id"] . " " . $item["errors"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load a pagesdeployment

```php
try {
    // load() returns the ENTITY — call data_get() for the PagesDeployment record (throws on error).
    $pagesdeployment = $client->PagesDeployment()->load(["id" => "example_id", "account_id" => "example_account_id", "project_id" => "example_project_id"]);
    print_r($pagesdeployment->data_get());
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 4. Create, update, and remove

```php
// create() returns the ENTITY — call data_get() for the created PagesDeployment record.
$created = $client->PagesDeployment()->create(["account_id" => "example_account_id", "errors" => [], "messages" => [], "result" => [], "success" => true]);

// Remove
$client->PagesDeployment()->remove(["id" => $created->data_get()["id"], "account_id" => "example_account_id", "project_id" => "example_project_id"]);
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $pagesdomains = $client->PagesDomain()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = CloudflarePagesSDK::test([
    "entity" => ["pagesdomain" => ["test01" => ["id" => "test01"]]],
]);

// list() returns entity instances (throws on error);
// call data_get() for the mock record.
$pagesdomain = $client->PagesDomain()->list();
print_r(array_map(fn($item) => $item->data_get(), $pagesdomain));
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new CloudflarePagesSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
CLOUDFLARE_PAGES_TEST_LIVE=TRUE
CLOUDFLARE_PAGES_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### CloudflarePagesSDK

```php
require_once 'cloudflarepages_sdk.php';
$client = new CloudflarePagesSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = CloudflarePagesSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### CloudflarePagesSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `PagesDeployment` | `($data): PagesDeploymentEntity` | Create a PagesDeployment entity instance. |
| `PagesDomain` | `($data): PagesDomainEntity` | Create a PagesDomain entity instance. |
| `PagesPagesAssetsCheckMissing` | `($data): PagesPagesAssetsCheckMissingEntity` | Create a PagesPagesAssetsCheckMissing entity instance. |
| `PagesProject` | `($data): PagesProjectEntity` | Create a PagesProject entity instance. |
| `Upload` | `($data): UploadEntity` | Create an Upload entity instance. |
| `UpsertHash` | `($data): UpsertHashEntity` | Create an UpsertHash entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$pages_deployment = $client->PagesDeployment();`

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
| `errors` | `array` |  |
| `filters` | `array` | Filters to apply to the tail session. |
| `id` | `string` |  |
| `messages` | `array` |  |
| `result` | `array` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the PagesDeployment record (throws on error).
$pages_deployment = $client->PagesDeployment()->load(["id" => "pages_deployment_id", "account_id" => "account_id", "project_id" => "project_id"]);
```

#### Example: List

```php
// list() returns an array of PagesDeployment records (throws on error).
$pages_deployments = $client->PagesDeployment()->list();
```

#### Example: Create

```php
$pages_deployment = $client->PagesDeployment()->create([
    "account_id" => null, // string
    "errors" => null, // array
    "messages" => null, // array
    "result" => null, // array
    "success" => null, // bool
]);
```


### PagesDomain

Create an instance: `$pages_domain = $client->PagesDomain();`

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
| `errors` | `array` |  |
| `id` | `string` |  |
| `messages` | `array` |  |
| `name` | `string` | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `array` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the PagesDomain record (throws on error).
$pages_domain = $client->PagesDomain()->load(["id" => "pages_domain_id", "account_id" => "account_id", "project_id" => "project_id"]);
```

#### Example: List

```php
// list() returns an array of PagesDomain records (throws on error).
$pages_domains = $client->PagesDomain()->list();
```

#### Example: Create

```php
$pages_domain = $client->PagesDomain()->create([
    "account_id" => null, // string
    "project_name" => null, // string
    "errors" => null, // array
    "messages" => null, // array
    "name" => null, // string
    "result" => null, // array
    "success" => null, // bool
]);
```


### PagesPagesAssetsCheckMissing

Create an instance: `$pages_pages_assets_check_missing = $client->PagesPagesAssetsCheckMissing();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `hashes` | `array` | List of file content hashes to check for existence in the asset store. |

#### Example: Create

```php
$pages_pages_assets_check_missing = $client->PagesPagesAssetsCheckMissing()->create([
    "hashes" => null, // array
]);
```


### PagesProject

Create an instance: `$pages_project = $client->PagesProject();`

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
| `build_config` | `array` | Configs for the project build process. |
| `deployment_configs` | `array` | Configs for deployments in a project. |
| `errors` | `array` |  |
| `id` | `string` |  |
| `messages` | `array` |  |
| `name` | `string` | Name for the Pages project. |
| `production_branch` | `string` | Production branch of the project. |
| `result` | `array` |  |
| `source` | `array` | Configs for the project source control. |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the PagesProject record (throws on error).
$pages_project = $client->PagesProject()->load(["id" => "pages_project_id", "account_id" => "account_id"]);
```

#### Example: List

```php
// list() returns an array of PagesProject records (throws on error).
$pages_projects = $client->PagesProject()->list();
```

#### Example: Create

```php
$pages_project = $client->PagesProject()->create([
    "account_id" => null, // string
    "errors" => null, // array
    "messages" => null, // array
    "name" => null, // string
    "production_branch" => null, // string
    "result" => null, // array
    "source" => null, // array
    "success" => null, // bool
]);
```


### Upload

Create an instance: `$upload = $client->Upload();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `array` |  |
| `messages` | `array` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Create

```php
$upload = $client->Upload()->create([
    "errors" => null, // array
    "messages" => null, // array
    "success" => null, // bool
]);
```


### UpsertHash

Create an instance: `$upsert_hash = $client->UpsertHash();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `array` |  |
| `hashes` | `array` | List of file content hashes to register in the asset store. |
| `messages` | `array` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Create

```php
$upsert_hash = $client->UpsertHash()->create([
    "errors" => null, // array
    "hashes" => null, // array
    "messages" => null, // array
    "success" => null, // bool
]);
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── cloudflarepages_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── schema.php                     -- Generated option + entity specs
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`cloudflarepages_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$pagesdomain = $client->PagesDomain();
$pagesdomain->list();

// $pagesdomain->data_get() now returns the pagesdomain data from the last list
// $pagesdomain->match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
