# CloudflarePages TypeScript SDK Reference

Complete API reference for the CloudflarePages TypeScript SDK.


## CloudflarePagesSDK

### Constructor

```ts
new CloudflarePagesSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CloudflarePagesSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = CloudflarePagesSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `CloudflarePagesSDK` instance in test mode.


### Instance Methods

#### `PagesDeployment(data?: object)`

Create a new `PagesDeployment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PagesDeploymentEntity` instance.

#### `PagesDomain(data?: object)`

Create a new `PagesDomain` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PagesDomainEntity` instance.

#### `PagesPagesAssetsCheckMissing(data?: object)`

Create a new `PagesPagesAssetsCheckMissing` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PagesPagesAssetsCheckMissingEntity` instance.

#### `PagesProject(data?: object)`

Create a new `PagesProject` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PagesProjectEntity` instance.

#### `Upload(data?: object)`

Create a new `Upload` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `UploadEntity` instance.

#### `UpsertHash(data?: object)`

Create a new `UpsertHash` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `UpsertHashEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `CloudflarePagesSDK.test()`.

**Returns:** `CloudflarePagesSDK` instance in test mode.


---

## PagesDeploymentEntity

```ts
const pages_deployment = client.PagesDeployment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `any[]` | Yes |  |
| `filters` | `any[]` | No | Filters to apply to the tail session. |
| `id` | `string` | No |  |
| `messages` | `any[]` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `success` | `boolean` | Yes | Whether the API call was successful. |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `retry` | `/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry` | `client.PagesDeployment().create({ $action: 'retry', ... })` |
| `rollback` | `/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/rollback` | `client.PagesDeployment().create({ $action: 'rollback', ... })` |

An action returns that action's OWN response, which is not necessarily a
PagesDeployment record — check the API definition for its shape.

```ts
const result = await client.PagesDeployment().create({
  $action: 'retry',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PagesDeployment().create({
  account_id: 'example_account_id',
  errors: [],
  messages: [],
  result: {},
  success: true,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PagesDeployment().list({ account_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PagesDeployment().load({ id: 'pages_deployment_id', account_id: 'account_id', project_id: 'project_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.PagesDeployment().remove({ id: 'pages_deployment_id', account_id: 'account_id', project_id: 'project_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PagesDeploymentEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflarePagesSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PagesDomainEntity

```ts
const pages_domain = client.PagesDomain()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `any[]` | Yes |  |
| `id` | `string` | No |  |
| `messages` | `any[]` | Yes |  |
| `name` | `string` | Yes | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `Record<string, any>` | Yes |  |
| `success` | `boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PagesDomain().create({
  account_id: 'example_account_id',
  project_name: 'example_project_name',
  errors: [],
  messages: [],
  name: 'example_name',
  result: {},
  success: true,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PagesDomain().list({ account_id: "example", project_name: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PagesDomain().load({ id: 'pages_domain_id', account_id: 'account_id', project_id: 'project_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.PagesDomain().remove({ id: 'pages_domain_id', account_id: 'account_id', project_id: 'project_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.PagesDomain().update({
  id: 'pages_domain_id',
  account_id: 'account_id',
  project_id: 'project_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PagesDomainEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflarePagesSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PagesPagesAssetsCheckMissingEntity

```ts
const pages_pages_assets_check_missing = client.PagesPagesAssetsCheckMissing()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `hashes` | `any[]` | Yes | List of file content hashes to check for existence in the asset store. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PagesPagesAssetsCheckMissing().create({
  hashes: [],
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PagesPagesAssetsCheckMissingEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflarePagesSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PagesProjectEntity

```ts
const pages_project = client.PagesProject()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_config` | `Record<string, any>` | No | Configs for the project build process. |
| `deployment_configs` | `Record<string, any>` | No | Configs for deployments in a project. |
| `errors` | `any[]` | Yes |  |
| `id` | `string` | No |  |
| `messages` | `any[]` | Yes |  |
| `name` | `string` | Yes | Name for the Pages project. |
| `production_branch` | `string` | Yes | Production branch of the project. |
| `result` | `Record<string, any>` | Yes |  |
| `source` | `Record<string, any>` | Yes | Configs for the project source control. |
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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `purge_build_cache` | `/accounts/{account_id}/pages/projects/{project_name}/purge_build_cache` | `client.PagesProject().create({ $action: 'purge_build_cache', ... })` |
| `source` | `/accounts/{account_id}/pages/projects/{project_name}/source` | `client.PagesProject().create({ $action: 'source', ... })` |
| `source` | `/accounts/{account_id}/pages/projects/{project_name}/source` | `client.PagesProject().remove({ $action: 'source', ... })` |

An action returns that action's OWN response, which is not necessarily a
PagesProject record — check the API definition for its shape.

```ts
const result = await client.PagesProject().create({
  $action: 'purge_build_cache',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PagesProject().create({
  account_id: 'example_account_id',
  errors: [],
  messages: [],
  name: 'example_name',
  production_branch: 'example_production_branch',
  result: {},
  source: {},
  success: true,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PagesProject().list({ account_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PagesProject().load({ id: 'pages_project_id', account_id: 'account_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.PagesProject().remove({ id: 'pages_project_id', account_id: 'account_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.PagesProject().update({
  id: 'pages_project_id',
  account_id: 'account_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PagesProjectEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflarePagesSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## UploadEntity

```ts
const upload = client.Upload()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `any[]` | Yes |  |
| `messages` | `any[]` | Yes |  |
| `success` | `boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Upload().create({
  errors: [],
  messages: [],
  success: true,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `UploadEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflarePagesSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## UpsertHashEntity

```ts
const upsert_hash = client.UpsertHash()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errors` | `any[]` | Yes |  |
| `hashes` | `any[]` | Yes | List of file content hashes to register in the asset store. |
| `messages` | `any[]` | Yes |  |
| `success` | `boolean` | Yes | Whether the API call was successful. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.UpsertHash().create({
  errors: [],
  hashes: [],
  messages: [],
  success: true,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `UpsertHashEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflarePagesSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```ts
const client = new CloudflarePagesSDK({
  feature: {
    test: { active: true },
  }
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

