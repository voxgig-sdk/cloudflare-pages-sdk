# CloudflarePages Python SDK Reference

Complete API reference for the CloudflarePages Python SDK.


## CloudflarePagesSDK

### Constructor

```python
from cloudflarepages_sdk import CloudflarePagesSDK

client = CloudflarePagesSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CloudflarePagesSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = CloudflarePagesSDK.test()
```


### Instance Methods

#### `PagesDeployment(data=None)`

Create a new `PagesDeploymentEntity` instance. Pass `None` for no initial data.

#### `PagesDomain(data=None)`

Create a new `PagesDomainEntity` instance. Pass `None` for no initial data.

#### `PagesPagesAssetsCheckMissing(data=None)`

Create a new `PagesPagesAssetsCheckMissingEntity` instance. Pass `None` for no initial data.

#### `PagesProject(data=None)`

Create a new `PagesProjectEntity` instance. Pass `None` for no initial data.

#### `Upload(data=None)`

Create a new `UploadEntity` instance. Pass `None` for no initial data.

#### `UpsertHash(data=None)`

Create a new `UpsertHashEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## PagesDeploymentEntity

```python
pages_deployment = client.PagesDeployment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `list` | Yes |  |
| `filters` | `list` | No | Filters to apply to the tail session. |
| `id` | `str` | No |  |
| `messages` | `list` | Yes |  |
| `result` | `dict` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PagesDeployment().create({
    "account_id": "example_account_id",  # str
    "errors": [],  # list
    "messages": [],  # list
    "result": {},  # dict
    "success": True,  # bool
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PagesDeployment().list({"account_id": "example"})
for pages_deployment in results:
    print(pages_deployment)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PagesDeployment().load({"id": "pages_deployment_id", "account_id": "account_id", "project_id": "project_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.PagesDeployment().remove({"id": "pages_deployment_id", "account_id": "account_id", "project_id": "project_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PagesDeploymentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PagesDomainEntity

```python
pages_domain = client.PagesDomain()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `list` | Yes |  |
| `id` | `str` | No |  |
| `messages` | `list` | Yes |  |
| `name` | `str` | Yes | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `dict` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PagesDomain().create({
    "account_id": "example_account_id",  # str
    "project_name": "example_project_name",  # str
    "errors": [],  # list
    "messages": [],  # list
    "name": "example_name",  # str
    "result": {},  # dict
    "success": True,  # bool
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PagesDomain().list({"account_id": "example", "project_name": "example"})
for pages_domain in results:
    print(pages_domain)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PagesDomain().load({"id": "pages_domain_id", "account_id": "account_id", "project_id": "project_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.PagesDomain().remove({"id": "pages_domain_id", "account_id": "account_id", "project_id": "project_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.PagesDomain().update({
    "id": "pages_domain_id",
    "account_id": "account_id",
    "project_id": "project_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PagesDomainEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PagesPagesAssetsCheckMissingEntity

```python
pages_pages_assets_check_missing = client.PagesPagesAssetsCheckMissing()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `hashes` | `list` | Yes | List of file content hashes to check for existence in the asset store. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PagesPagesAssetsCheckMissing().create({
    "hashes": [],  # list
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PagesPagesAssetsCheckMissingEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PagesProjectEntity

```python
pages_project = client.PagesProject()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_config` | `dict` | No | Configs for the project build process. |
| `deployment_configs` | `dict` | No | Configs for deployments in a project. |
| `errors` | `list` | Yes |  |
| `id` | `str` | No |  |
| `messages` | `list` | Yes |  |
| `name` | `str` | Yes | Name for the Pages project. |
| `production_branch` | `str` | Yes | Production branch of the project. |
| `result` | `dict` | Yes |  |
| `source` | `dict` | Yes | Configs for the project source control. |
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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PagesProject().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PagesProject().list({"account_id": "example"})
for pages_project in results:
    print(pages_project)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PagesProject().load({"id": "pages_project_id", "account_id": "account_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.PagesProject().remove({"id": "pages_project_id", "account_id": "account_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.PagesProject().update({
    "id": "pages_project_id",
    "account_id": "account_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PagesProjectEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UploadEntity

```python
upload = client.Upload()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `list` | Yes |  |
| `messages` | `list` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Upload().create({
    "errors": [],  # list
    "messages": [],  # list
    "success": True,  # bool
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UploadEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpsertHashEntity

```python
upsert_hash = client.UpsertHash()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `list` | Yes |  |
| `hashes` | `list` | Yes | List of file content hashes to register in the asset store. |
| `messages` | `list` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.UpsertHash().create({
    "errors": [],  # list
    "hashes": [],  # list
    "messages": [],  # list
    "success": True,  # bool
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpsertHashEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```python
client = CloudflarePagesSDK({
    "feature": {
        "test": {"active": True},
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

