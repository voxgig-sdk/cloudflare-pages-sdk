# CloudflarePages PHP SDK Reference

Complete API reference for the CloudflarePages PHP SDK.


## CloudflarePagesSDK

### Constructor

```php
require_once __DIR__ . '/cloudflarepages_sdk.php';

$client = new CloudflarePagesSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CloudflarePagesSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = CloudflarePagesSDK::test();
```


### Instance Methods

#### `PagesDeployment($data = null)`

Create a new `PagesDeploymentEntity` instance. Pass `null` for no initial data.

#### `PagesDomain($data = null)`

Create a new `PagesDomainEntity` instance. Pass `null` for no initial data.

#### `PagesPagesAssetsCheckMissing($data = null)`

Create a new `PagesPagesAssetsCheckMissingEntity` instance. Pass `null` for no initial data.

#### `PagesProject($data = null)`

Create a new `PagesProjectEntity` instance. Pass `null` for no initial data.

#### `Upload($data = null)`

Create a new `UploadEntity` instance. Pass `null` for no initial data.

#### `UpsertHash($data = null)`

Create a new `UpsertHashEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): CloudflarePagesUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## PagesDeploymentEntity

```php
$pages_deployment = $client->PagesDeployment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `array` | Yes |  |
| `filters` | `array` | No | Filters to apply to the tail session. |
| `id` | `string` | No |  |
| `messages` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PagesDeployment()->create([
  "account_id" => null, // string
  "errors" => null, // array
  "messages" => null, // array
  "result" => null, // array
  "success" => null, // bool
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PagesDeployment()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PagesDeployment()->load(["id" => "pages_deployment_id", "account_id" => "account_id", "project_id" => "project_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->PagesDeployment()->remove(["id" => "pages_deployment_id", "account_id" => "account_id", "project_id" => "project_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PagesDeploymentEntity`

Create a new `PagesDeploymentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PagesDomainEntity

```php
$pages_domain = $client->PagesDomain();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `array` | Yes |  |
| `id` | `string` | No |  |
| `messages` | `array` | Yes |  |
| `name` | `string` | Yes | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `array` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PagesDomain()->create([
  "account_id" => null, // string
  "project_name" => null, // string
  "errors" => null, // array
  "messages" => null, // array
  "name" => null, // string
  "result" => null, // array
  "success" => null, // bool
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PagesDomain()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PagesDomain()->load(["id" => "pages_domain_id", "account_id" => "account_id", "project_id" => "project_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->PagesDomain()->remove(["id" => "pages_domain_id", "account_id" => "account_id", "project_id" => "project_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->PagesDomain()->update([
  "id" => "pages_domain_id",
  "account_id" => "account_id",
  "project_id" => "project_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PagesDomainEntity`

Create a new `PagesDomainEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PagesPagesAssetsCheckMissingEntity

```php
$pages_pages_assets_check_missing = $client->PagesPagesAssetsCheckMissing();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `hashes` | `array` | Yes | List of file content hashes to check for existence in the asset store. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PagesPagesAssetsCheckMissing()->create([
  "hashes" => null, // array
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PagesPagesAssetsCheckMissingEntity`

Create a new `PagesPagesAssetsCheckMissingEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PagesProjectEntity

```php
$pages_project = $client->PagesProject();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_config` | `array` | No | Configs for the project build process. |
| `deployment_configs` | `array` | No | Configs for deployments in a project. |
| `errors` | `array` | Yes |  |
| `id` | `string` | No |  |
| `messages` | `array` | Yes |  |
| `name` | `string` | Yes | Name for the Pages project. |
| `production_branch` | `string` | Yes | Production branch of the project. |
| `result` | `array` | Yes |  |
| `source` | `array` | Yes | Configs for the project source control. |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PagesProject()->create([
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

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PagesProject()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PagesProject()->load(["id" => "pages_project_id", "account_id" => "account_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->PagesProject()->remove(["id" => "pages_project_id", "account_id" => "account_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->PagesProject()->update([
  "id" => "pages_project_id",
  "account_id" => "account_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PagesProjectEntity`

Create a new `PagesProjectEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## UploadEntity

```php
$upload = $client->Upload();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `array` | Yes |  |
| `messages` | `array` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Upload()->create([
  "errors" => null, // array
  "messages" => null, // array
  "success" => null, // bool
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): UploadEntity`

Create a new `UploadEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## UpsertHashEntity

```php
$upsert_hash = $client->UpsertHash();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `array` | Yes |  |
| `hashes` | `array` | Yes | List of file content hashes to register in the asset store. |
| `messages` | `array` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->UpsertHash()->create([
  "errors" => null, // array
  "hashes" => null, // array
  "messages" => null, // array
  "success" => null, // bool
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): UpsertHashEntity`

Create a new `UpsertHashEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```php
$client = new CloudflarePagesSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
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

