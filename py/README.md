# CloudflarePages Python SDK



The Python SDK for the CloudflarePages API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.PagesDeployment()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/cloudflare-pages-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from cloudflarepages_sdk import CloudflarePagesSDK

client = CloudflarePagesSDK({
    "apikey": os.environ.get("CLOUDFLARE_PAGES_APIKEY"),
})
```

### 2. List pagesdeployment records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    pagesdeployments = client.PagesDeployment().list({"account_id": "example"})
    for pagesdeployment in pagesdeployments:
        print(pagesdeployment)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load a pagesdeployment

`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    pagesdeployment = client.PagesDeployment().load({"id": "example_id", "account_id": "example_account_id", "project_id": "example_project_id"})
    print(pagesdeployment)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.PagesDeployment().create({"account_id": "example_account_id", "errors": [], "messages": [], "result": {}, "success": True})

# Remove
client.PagesDeployment().remove({"id": created.data_get()["id"], "account_id": "example_account_id", "project_id": "example_project_id"})
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    pagesdomains = client.PagesDomain().list()
    print(pagesdomains)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = CloudflarePagesSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
pagesdomain = client.PagesDomain().list()
# pagesdomain contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = CloudflarePagesSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### CloudflarePagesSDK

```python
from cloudflarepages_sdk import CloudflarePagesSDK

client = CloudflarePagesSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = CloudflarePagesSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### CloudflarePagesSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `pages_deployment = client.PagesDeployment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `list` |  |
| `filters` | `list` | Filters to apply to the tail session. |
| `id` | `str` |  |
| `messages` | `list` |  |
| `result` | `dict` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```python
pages_deployment = client.PagesDeployment().load({"id": "pages_deployment_id", "account_id": "account_id", "project_id": "project_id"})
```

#### Example: List

```python
pages_deployments = client.PagesDeployment().list({"account_id": "example"})
```

#### Example: Create

```python
pages_deployment = client.PagesDeployment().create({
    "account_id": "example_account_id",  # str
    "errors": [],  # list
    "messages": [],  # list
    "result": {},  # dict
    "success": True,  # bool
})
```


### PagesDomain

Create an instance: `pages_domain = client.PagesDomain()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `list` |  |
| `id` | `str` |  |
| `messages` | `list` |  |
| `name` | `str` | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `dict` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```python
pages_domain = client.PagesDomain().load({"id": "pages_domain_id", "account_id": "account_id", "project_id": "project_id"})
```

#### Example: List

```python
pages_domains = client.PagesDomain().list({"account_id": "example", "project_name": "example"})
```

#### Example: Create

```python
pages_domain = client.PagesDomain().create({
    "account_id": "example_account_id",  # str
    "project_name": "example_project_name",  # str
    "errors": [],  # list
    "messages": [],  # list
    "name": "example_name",  # str
    "result": {},  # dict
    "success": True,  # bool
})
```


### PagesPagesAssetsCheckMissing

Create an instance: `pages_pages_assets_check_missing = client.PagesPagesAssetsCheckMissing()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `hashes` | `list` | List of file content hashes to check for existence in the asset store. |

#### Example: Create

```python
pages_pages_assets_check_missing = client.PagesPagesAssetsCheckMissing().create({
    "hashes": [],  # list
})
```


### PagesProject

Create an instance: `pages_project = client.PagesProject()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `build_config` | `dict` | Configs for the project build process. |
| `deployment_configs` | `dict` | Configs for deployments in a project. |
| `errors` | `list` |  |
| `id` | `str` |  |
| `messages` | `list` |  |
| `name` | `str` | Name for the Pages project. |
| `production_branch` | `str` | Production branch of the project. |
| `result` | `dict` |  |
| `source` | `dict` | Configs for the project source control. |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```python
pages_project = client.PagesProject().load({"id": "pages_project_id", "account_id": "account_id"})
```

#### Example: List

```python
pages_projects = client.PagesProject().list({"account_id": "example"})
```

#### Example: Create

```python
pages_project = client.PagesProject().create({
    "account_id": "example_account_id",  # str
    "errors": [],  # list
    "messages": [],  # list
    "name": "example_name",  # str
    "production_branch": "example_production_branch",  # str
    "result": {},  # dict
    "source": {},  # dict
    "success": True,  # bool
})
```


### Upload

Create an instance: `upload = client.Upload()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `list` |  |
| `messages` | `list` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Create

```python
upload = client.Upload().create({
    "errors": [],  # list
    "messages": [],  # list
    "success": True,  # bool
})
```


### UpsertHash

Create an instance: `upsert_hash = client.UpsertHash()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `list` |  |
| `hashes` | `list` | List of file content hashes to register in the asset store. |
| `messages` | `list` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Create

```python
upsert_hash = client.UpsertHash().create({
    "errors": [],  # list
    "hashes": [],  # list
    "messages": [],  # list
    "success": True,  # bool
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── cloudflarepages_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── schema.py                    -- Generated option + entity specs
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`cloudflarepages_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
pagesdomain = client.PagesDomain()
pagesdomain.list()

# pagesdomain.data_get() now returns the pagesdomain data from the last list
# pagesdomain.match_get() returns the last match criteria
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
