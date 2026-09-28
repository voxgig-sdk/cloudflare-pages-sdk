# CloudflarePages TypeScript SDK



The TypeScript SDK for the CloudflarePages API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.PagesDeployment()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb` — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/cloudflare-pages-sdk/releases](https://github.com/voxgig-sdk/cloudflare-pages-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { CloudflarePagesSDK } from '@voxgig-sdk/cloudflare-pages-sdk'

const client = new CloudflarePagesSDK({
  apikey: process.env.CLOUDFLARE_PAGES_APIKEY,
})
```

### 2. List pagesdeployment records

`list()` resolves to an array of PagesDeployment ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

```ts
const pagesdeployments = await client.PagesDeployment().list({ account_id: "example" })

for (const pagesdeployment of pagesdeployments) {
  console.log(pagesdeployment)
}
```

### 3. Load a pagesdeployment

`load()` returns the entity directly and throws on failure:

```ts
try {
  const pagesdeployment = await client.PagesDeployment().load({ id: 'example_id', account_id: 'example_account_id', project_id: 'example_project_id' })
  console.log(pagesdeployment)
} catch (err) {
  console.error('load failed:', err)
}
```

### 4. Create, update, and remove

```ts
// Create — returns the created PagesDeployment ENTITY (.data() for the record)
const created = await client.PagesDeployment().create({
  account_id: 'example_account_id',
  errors: [],
  messages: [],
  result: {},
  success: true,
})

// Remove
await client.PagesDeployment().remove({
  id: created.data().id!,
  account_id: 'example_account_id',
  project_id: 'example_project_id',
})
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const pagesdomains = await client.PagesDomain().list()
  console.log(pagesdomains)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = CloudflarePagesSDK.test()

const pagesdomain = await client.PagesDomain().list()
// pagesdomain is the entity, populated with mock response data
// — call pagesdomain.data() for the record itself
console.log(pagesdomain)
```

You can also use the instance method:

```ts
const client = new CloudflarePagesSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.PagesDomain()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data.id)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new CloudflarePagesSDK({
  apikey: '...',
  extend: [logger],
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
cd ts && npm test
```

Live entity tests continue independent operations after errors and attempt
supported cleanup. Their final result reports failures and missing prerequisites
after the remaining work completes. The model and test inputs determine which
API operations the generated scenarios cover.


## Reference

### CloudflarePagesSDK

#### Constructor

```ts
new CloudflarePagesSDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `PagesDeployment(data?)` | `PagesDeploymentEntity` | Create a PagesDeployment entity instance. |
| `PagesDomain(data?)` | `PagesDomainEntity` | Create a PagesDomain entity instance. |
| `PagesPagesAssetsCheckMissing(data?)` | `PagesPagesAssetsCheckMissingEntity` | Create a PagesPagesAssetsCheckMissing entity instance. |
| `PagesProject(data?)` | `PagesProjectEntity` | Create a PagesProject entity instance. |
| `Upload(data?)` | `UploadEntity` | Create an Upload entity instance. |
| `UpsertHash(data?)` | `UpsertHashEntity` | Create an UpsertHash entity instance. |
| `tester(testopts?, sdkopts?)` | `CloudflarePagesSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `CloudflarePagesSDK.test(testopts?, sdkopts?)` | `CloudflarePagesSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): CloudflarePagesSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `void`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

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

Operations: create, list, load, remove.

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

Operations: create, list, load, remove, update.

API path: `/accounts/{account_id}/pages/projects/{project_name}/domains`

#### PagesPagesAssetsCheckMissing

| Field | Description |
| --- | --- |
| `hashes` | List of file content hashes to check for existence in the asset store. |

Operations: create.

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

Operations: create, list, load, remove, update.

API path: `/accounts/{account_id}/pages/projects/{project_name}/purge_build_cache`

#### Upload

| Field | Description |
| --- | --- |
| `errors` |  |
| `messages` |  |
| `success` | Whether the API call was successful. |

Operations: create.

API path: `/pages/assets/upload`

#### UpsertHash

| Field | Description |
| --- | --- |
| `errors` |  |
| `hashes` | List of file content hashes to register in the asset store. |
| `messages` |  |
| `success` | Whether the API call was successful. |

Operations: create.

API path: `/pages/assets/upsert-hashes`



## Entities


### PagesDeployment

Create an instance: `const pages_deployment = client.PagesDeployment()`

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
| `errors` | `any[]` |  |
| `filters` | `any[]` | Filters to apply to the tail session. |
| `id` | `string` |  |
| `messages` | `any[]` |  |
| `result` | `Record<string, any>` |  |
| `success` | `boolean` | Whether the API call was successful. |

#### Example: Load

```ts
const pages_deployment = await client.PagesDeployment().load({ id: 'pages_deployment_id', account_id: 'account_id', project_id: 'project_id' })
```

#### Example: List

```ts
const pages_deployments = await client.PagesDeployment().list({ account_id: "example" })
```

#### Example: Create

```ts
const pages_deployment = await client.PagesDeployment().create({
  account_id: 'example_account_id',
  errors: [],
  messages: [],
  result: {},
  success: true,
})
```


### PagesDomain

Create an instance: `const pages_domain = client.PagesDomain()`

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
| `errors` | `any[]` |  |
| `id` | `string` |  |
| `messages` | `any[]` |  |
| `name` | `string` | Fully qualified domain name for the Pages project, such as `example.com`. |
| `result` | `Record<string, any>` |  |
| `success` | `boolean` | Whether the API call was successful. |

#### Example: Load

```ts
const pages_domain = await client.PagesDomain().load({ id: 'pages_domain_id', account_id: 'account_id', project_id: 'project_id' })
```

#### Example: List

```ts
const pages_domains = await client.PagesDomain().list({ account_id: "example", project_name: "example" })
```

#### Example: Create

```ts
const pages_domain = await client.PagesDomain().create({
  account_id: 'example_account_id',
  project_name: 'example_project_name',
  errors: [],
  messages: [],
  name: 'example_name',
  result: {},
  success: true,
})
```


### PagesPagesAssetsCheckMissing

Create an instance: `const pages_pages_assets_check_missing = client.PagesPagesAssetsCheckMissing()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `hashes` | `any[]` | List of file content hashes to check for existence in the asset store. |

#### Example: Create

```ts
const pages_pages_assets_check_missing = await client.PagesPagesAssetsCheckMissing().create({
  hashes: [],
})
```


### PagesProject

Create an instance: `const pages_project = client.PagesProject()`

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
| `build_config` | `Record<string, any>` | Configs for the project build process. |
| `deployment_configs` | `Record<string, any>` | Configs for deployments in a project. |
| `errors` | `any[]` |  |
| `id` | `string` |  |
| `messages` | `any[]` |  |
| `name` | `string` | Name for the Pages project. |
| `production_branch` | `string` | Production branch of the project. |
| `result` | `Record<string, any>` |  |
| `source` | `Record<string, any>` | Configs for the project source control. |
| `success` | `boolean` | Whether the API call was successful. |

#### Example: Load

```ts
const pages_project = await client.PagesProject().load({ id: 'pages_project_id', account_id: 'account_id' })
```

#### Example: List

```ts
const pages_projects = await client.PagesProject().list({ account_id: "example" })
```

#### Example: Create

```ts
const pages_project = await client.PagesProject().create({
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


### Upload

Create an instance: `const upload = client.Upload()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `any[]` |  |
| `messages` | `any[]` |  |
| `success` | `boolean` | Whether the API call was successful. |

#### Example: Create

```ts
const upload = await client.Upload().create({
  errors: [],
  messages: [],
  success: true,
})
```


### UpsertHash

Create an instance: `const upsert_hash = client.UpsertHash()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errors` | `any[]` |  |
| `hashes` | `any[]` | List of file content hashes to register in the asset store. |
| `messages` | `any[]` |  |
| `success` | `boolean` | Whether the API call was successful. |

#### Example: Create

```ts
const upsert_hash = await client.UpsertHash().create({
  errors: [],
  hashes: [],
  messages: [],
  success: true,
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
cloudflare-pages/
├── src/
│   ├── CloudflarePagesSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { CloudflarePagesSDK } from '@voxgig-sdk/cloudflare-pages-sdk'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const pagesdomain = client.PagesDomain()
await pagesdomain.list()

// pagesdomain.data() now returns the pagesdomain data from the last `list`
// pagesdomain.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
