package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "CloudflarePages",
			"slug": "cloudflare-pages",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://api.cloudflare.com/client/v4",
			"auth": map[string]any{
				"prefix": "",
				"name": "X-Auth-Email",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"pages_deployment": map[string]any{},
				"pages_domain": map[string]any{},
				"pages_pages_assets_check_missing": map[string]any{},
				"pages_project": map[string]any{},
				"upload": map[string]any{},
				"upsert_hash": map[string]any{},
			},
		},
		"entity": map[string]any{
			"pages_deployment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "errors",
						"title": "Errors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "filters",
						"title": "Filters",
						"type": "`$ARRAY`",
						"short": "Filters to apply to the tail session.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "messages",
						"title": "Messages",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "success",
						"title": "Success",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the API call was successful.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "pages_deployment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "deployments",
									},
									map[string]any{
										"var": "deployment_id",
									},
									map[string]any{
										"lit": "retry",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"deployments",
									"{deployment_id}",
									"retry",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "deployment_id",
											"orig": "deployment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"$action": "retry",
									"exist": []any{
										"account_id",
										"deployment_id",
										"project_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/rollback",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "deployments",
									},
									map[string]any{
										"var": "deployment_id",
									},
									map[string]any{
										"lit": "rollback",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"deployments",
									"{deployment_id}",
									"rollback",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "deployment_id",
											"orig": "deployment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"$action": "rollback",
									"exist": []any{
										"account_id",
										"deployment_id",
										"project_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/tails",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "deployments",
									},
									map[string]any{
										"var": "deployment_id",
									},
									map[string]any{
										"lit": "tails",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"deployments",
									"{deployment_id}",
									"tails",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "deployment_id",
											"orig": "deployment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"deployment_id",
										"project_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/deployments",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_name",
									},
									map[string]any{
										"lit": "deployments",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_name}",
									"deployments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "project_name",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"project_name",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/deployments",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_name",
									},
									map[string]any{
										"lit": "deployments",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_name}",
									"deployments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "project_name",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
									"query": []any{
										map[string]any{
											"name": "env",
											"orig": "env",
											"type": "`$STRING`",
											"kind": "query",
											"example": "preview",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 10,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"env",
										"page",
										"per_page",
										"project_name",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/history/logs",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "deployments",
									},
									map[string]any{
										"var": "deployment_id",
									},
									map[string]any{
										"lit": "history",
									},
									map[string]any{
										"lit": "logs",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"deployments",
									"{deployment_id}",
									"history",
									"logs",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "deployment_id",
											"orig": "deployment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"deployment_id",
										"project_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "deployments",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"deployments",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"deployment_id": "id",
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "id",
											"orig": "deployment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
										"project_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/tails/{tail_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "deployments",
									},
									map[string]any{
										"var": "deployment_id",
									},
									map[string]any{
										"lit": "tails",
									},
									map[string]any{
										"var": "tail_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"deployments",
									"{deployment_id}",
									"tails",
									"{tail_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "deployment_id",
											"orig": "deployment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
										map[string]any{
											"name": "tail_id",
											"orig": "tail_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"deployment_id",
										"project_id",
										"tail_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "deployments",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"deployments",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"deployment_id": "id",
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "id",
											"orig": "deployment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
									"query": []any{
										map[string]any{
											"name": "force",
											"orig": "force",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"force",
										"id",
										"project_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"pages_domain": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "errors",
						"title": "Errors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "messages",
						"title": "Messages",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "Fully qualified domain name for the Pages project, such as `example.com`.",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "success",
						"title": "Success",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the API call was successful.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "pages_domain",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/domains",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_name",
									},
									map[string]any{
										"lit": "domains",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_name}",
									"domains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "project_name",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"project_name",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/domains",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_name",
									},
									map[string]any{
										"lit": "domains",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_name}",
									"domains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "project_name",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"project_name",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"domains",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"domain_name": "id",
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "id",
											"orig": "domain_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "example.com",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
										"project_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"domains",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"domain_name": "id",
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "id",
											"orig": "domain_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "example.com",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
										"project_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_id",
									},
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_id}",
									"domains",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"domain_name": "id",
										"project_name": "project_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "id",
											"orig": "domain_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "example.com",
										},
										map[string]any{
											"name": "project_id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
										"project_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"pages_pages_assets_check_missing": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "hashes",
						"title": "Hashes",
						"type": "`$ARRAY`",
						"req": true,
						"short": "List of file content hashes to check for existence in the asset store.",
					},
				},
				"name": "pages_pages_assets_check_missing",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/assets/check-missing",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "assets",
									},
									map[string]any{
										"lit": "check-missing",
									},
								},
								"parts": []any{
									"pages",
									"assets",
									"check-missing",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"pages_project": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "build_config",
						"title": "Build Config",
						"type": "`$OBJECT`",
						"short": "Configs for the project build process.",
					},
					map[string]any{
						"name": "deployment_configs",
						"title": "Deployment Configs",
						"type": "`$OBJECT`",
						"short": "Configs for deployments in a project.",
					},
					map[string]any{
						"name": "errors",
						"title": "Errors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "messages",
						"title": "Messages",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "Name for the Pages project.",
					},
					map[string]any{
						"name": "production_branch",
						"title": "Production Branch",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "Production branch of the project.",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "source",
						"title": "Source",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Configs for the project source control.",
					},
					map[string]any{
						"name": "success",
						"title": "Success",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the API call was successful.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "pages_project",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/purge_build_cache",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_name",
									},
									map[string]any{
										"lit": "purge_build_cache",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_name}",
									"purge_build_cache",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "project_name",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"$action": "purge_build_cache",
									"exist": []any{
										"account_id",
										"project_name",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/source",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_name",
									},
									map[string]any{
										"lit": "source",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_name}",
									"source",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "project_name",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"$action": "source",
									"exist": []any{
										"account_id",
										"project_name",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/pages/projects",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/pages/projects",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 10,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"page",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/upload-token",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_name",
									},
									map[string]any{
										"lit": "upload-token",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_name}",
									"upload-token",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "project_name",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"project_name",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"project_name": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"project_name": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}/source",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "project_name",
									},
									map[string]any{
										"lit": "source",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{project_name}",
									"source",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "project_name",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"$action": "source",
									"exist": []any{
										"account_id",
										"project_name",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/pages/projects/{project_name}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "projects",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"pages",
									"projects",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"project_name": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "id",
											"orig": "project_name",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "this-is-my-project-01",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"upload": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "errors",
						"title": "Errors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "messages",
						"title": "Messages",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "success",
						"title": "Success",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the API call was successful.",
					},
				},
				"name": "upload",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/assets/upload",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "assets",
									},
									map[string]any{
										"lit": "upload",
									},
								},
								"parts": []any{
									"pages",
									"assets",
									"upload",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"upsert_hash": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "errors",
						"title": "Errors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "hashes",
						"title": "Hashes",
						"type": "`$ARRAY`",
						"req": true,
						"short": "List of file content hashes to register in the asset store.",
					},
					map[string]any{
						"name": "messages",
						"title": "Messages",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "success",
						"title": "Success",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the API call was successful.",
					},
				},
				"name": "upsert_hash",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/assets/upsert-hashes",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"lit": "assets",
									},
									map[string]any{
										"lit": "upsert-hashes",
									},
								},
								"parts": []any{
									"pages",
									"assets",
									"upsert-hashes",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
