-- CloudflarePages SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "CloudflarePages",
      slug = "cloudflare-pages",
      version = "0.0.1",
      target = "lua",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
        ["optspec"] = {
          ["entity"] = "`$MAP`",
          ["net"] = "`$MAP`",
        },
        ["strict"] = false,
        ["transport"] = "base",
      },
    },
    options = {
      base = "https://api.cloudflare.com/client/v4",
      auth = {
        prefix = "",
        name = "X-Auth-Email",
      },
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["pages_deployment"] = {},
        ["pages_domain"] = {},
        ["pages_pages_assets_check_missing"] = {},
        ["pages_project"] = {},
        ["upload"] = {},
        ["upsert_hash"] = {},
      },
    },
    entity = {
      ["pages_deployment"] = {
        ["fields"] = {
          {
            ["name"] = "errors",
            ["title"] = "Errors",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "filters",
            ["title"] = "Filters",
            ["type"] = "`$ARRAY`",
            ["short"] = "Filters to apply to the tail session.",
          },
          {
            ["name"] = "id",
            ["title"] = "Id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "messages",
            ["title"] = "Messages",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "result",
            ["title"] = "Result",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "success",
            ["title"] = "Success",
            ["type"] = "`$BOOLEAN`",
            ["req"] = true,
            ["short"] = "Whether the API call was successful.",
          },
        },
        ["id"] = {
          ["field"] = "id",
          ["name"] = "id",
        },
        ["name"] = "pages_deployment",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "deployments",
                  },
                  {
                    ["var"] = "deployment_id",
                  },
                  {
                    ["lit"] = "retry",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_id}",
                  "deployments",
                  "{deployment_id}",
                  "retry",
                },
                ["rename"] = {
                  ["param"] = {
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "deployment_id",
                      ["orig"] = "deployment_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["$action"] = "retry",
                  ["exist"] = {
                    "account_id",
                    "deployment_id",
                    "project_id",
                  },
                },
              },
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/rollback",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "deployments",
                  },
                  {
                    ["var"] = "deployment_id",
                  },
                  {
                    ["lit"] = "rollback",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_id}",
                  "deployments",
                  "{deployment_id}",
                  "rollback",
                },
                ["rename"] = {
                  ["param"] = {
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "deployment_id",
                      ["orig"] = "deployment_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["$action"] = "rollback",
                  ["exist"] = {
                    "account_id",
                    "deployment_id",
                    "project_id",
                  },
                },
              },
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/tails",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "deployments",
                  },
                  {
                    ["var"] = "deployment_id",
                  },
                  {
                    ["lit"] = "tails",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_id}",
                  "deployments",
                  "{deployment_id}",
                  "tails",
                },
                ["rename"] = {
                  ["param"] = {
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "deployment_id",
                      ["orig"] = "deployment_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "deployment_id",
                    "project_id",
                  },
                },
              },
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/deployments",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_name",
                  },
                  {
                    ["lit"] = "deployments",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_name}",
                  "deployments",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "project_name",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "project_name",
                  },
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/deployments",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_name",
                  },
                  {
                    ["lit"] = "deployments",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_name}",
                  "deployments",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "project_name",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                  ["query"] = {
                    {
                      ["name"] = "env",
                      ["orig"] = "env",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "preview",
                    },
                    {
                      ["name"] = "page",
                      ["orig"] = "page",
                      ["type"] = "`$INTEGER`",
                      ["kind"] = "query",
                      ["example"] = 1,
                    },
                    {
                      ["name"] = "per_page",
                      ["orig"] = "per_page",
                      ["type"] = "`$INTEGER`",
                      ["kind"] = "query",
                      ["example"] = 10,
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "env",
                    "page",
                    "per_page",
                    "project_name",
                  },
                },
              },
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/history/logs",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "deployments",
                  },
                  {
                    ["var"] = "deployment_id",
                  },
                  {
                    ["lit"] = "history",
                  },
                  {
                    ["lit"] = "logs",
                  },
                },
                ["parts"] = {
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
                ["rename"] = {
                  ["param"] = {
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "deployment_id",
                      ["orig"] = "deployment_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "deployment_id",
                    "project_id",
                  },
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "deployments",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_id}",
                  "deployments",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["deployment_id"] = "id",
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "id",
                      ["orig"] = "deployment_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                    "project_id",
                  },
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/tails/{tail_id}",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "deployments",
                  },
                  {
                    ["var"] = "deployment_id",
                  },
                  {
                    ["lit"] = "tails",
                  },
                  {
                    ["var"] = "tail_id",
                  },
                },
                ["parts"] = {
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
                ["rename"] = {
                  ["param"] = {
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "deployment_id",
                      ["orig"] = "deployment_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                    {
                      ["name"] = "tail_id",
                      ["orig"] = "tail_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "deployment_id",
                    "project_id",
                    "tail_id",
                  },
                },
              },
              {
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "deployments",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_id}",
                  "deployments",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["deployment_id"] = "id",
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "id",
                      ["orig"] = "deployment_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "f64788e9-fccd-4d4a-a28a-cb84f88f6e12",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                  ["query"] = {
                    {
                      ["name"] = "force",
                      ["orig"] = "force",
                      ["type"] = "`$BOOLEAN`",
                      ["kind"] = "query",
                      ["example"] = true,
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
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
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["pages_domain"] = {
        ["fields"] = {
          {
            ["name"] = "errors",
            ["title"] = "Errors",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "id",
            ["title"] = "Id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "messages",
            ["title"] = "Messages",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "name",
            ["title"] = "Name",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "Fully qualified domain name for the Pages project, such as `example.com`.",
          },
          {
            ["name"] = "result",
            ["title"] = "Result",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "success",
            ["title"] = "Success",
            ["type"] = "`$BOOLEAN`",
            ["req"] = true,
            ["short"] = "Whether the API call was successful.",
          },
        },
        ["id"] = {
          ["field"] = "id",
          ["name"] = "id",
        },
        ["name"] = "pages_domain",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/domains",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_name",
                  },
                  {
                    ["lit"] = "domains",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_name}",
                  "domains",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "project_name",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "project_name",
                  },
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/domains",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_name",
                  },
                  {
                    ["lit"] = "domains",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_name}",
                  "domains",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "project_name",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "project_name",
                  },
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "domains",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_id}",
                  "domains",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["domain_name"] = "id",
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "id",
                      ["orig"] = "domain_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "example.com",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                    "project_id",
                  },
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "domains",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_id}",
                  "domains",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["domain_name"] = "id",
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "id",
                      ["orig"] = "domain_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "example.com",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                    "project_id",
                  },
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_id",
                  },
                  {
                    ["lit"] = "domains",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_id}",
                  "domains",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["domain_name"] = "id",
                    ["project_name"] = "project_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "id",
                      ["orig"] = "domain_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "example.com",
                    },
                    {
                      ["name"] = "project_id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                    "project_id",
                  },
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["pages_pages_assets_check_missing"] = {
        ["fields"] = {
          {
            ["name"] = "hashes",
            ["title"] = "Hashes",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
            ["short"] = "List of file content hashes to check for existence in the asset store.",
          },
        },
        ["name"] = "pages_pages_assets_check_missing",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/pages/assets/check-missing",
                ["segments"] = {
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "assets",
                  },
                  {
                    ["lit"] = "check-missing",
                  },
                },
                ["parts"] = {
                  "pages",
                  "assets",
                  "check-missing",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {},
                ["select"] = {},
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["pages_project"] = {
        ["fields"] = {
          {
            ["name"] = "build_config",
            ["title"] = "Build Config",
            ["type"] = "`$OBJECT`",
            ["short"] = "Configs for the project build process.",
          },
          {
            ["name"] = "deployment_configs",
            ["title"] = "Deployment Configs",
            ["type"] = "`$OBJECT`",
            ["short"] = "Configs for deployments in a project.",
          },
          {
            ["name"] = "errors",
            ["title"] = "Errors",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "id",
            ["title"] = "Id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "messages",
            ["title"] = "Messages",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "name",
            ["title"] = "Name",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["op"] = {
              ["update"] = {
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Name for the Pages project.",
          },
          {
            ["name"] = "production_branch",
            ["title"] = "Production Branch",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["op"] = {
              ["update"] = {
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Production branch of the project.",
          },
          {
            ["name"] = "result",
            ["title"] = "Result",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "source",
            ["title"] = "Source",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
            ["short"] = "Configs for the project source control.",
          },
          {
            ["name"] = "success",
            ["title"] = "Success",
            ["type"] = "`$BOOLEAN`",
            ["req"] = true,
            ["short"] = "Whether the API call was successful.",
          },
        },
        ["id"] = {
          ["field"] = "id",
          ["name"] = "id",
        },
        ["name"] = "pages_project",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/purge_build_cache",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_name",
                  },
                  {
                    ["lit"] = "purge_build_cache",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_name}",
                  "purge_build_cache",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "project_name",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["$action"] = "purge_build_cache",
                  ["exist"] = {
                    "account_id",
                    "project_name",
                  },
                },
              },
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/source",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_name",
                  },
                  {
                    ["lit"] = "source",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_name}",
                  "source",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "project_name",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["$action"] = "source",
                  ["exist"] = {
                    "account_id",
                    "project_name",
                  },
                },
              },
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/accounts/{account_id}/pages/projects",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                  },
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/accounts/{account_id}/pages/projects",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                  },
                  ["query"] = {
                    {
                      ["name"] = "page",
                      ["orig"] = "page",
                      ["type"] = "`$INTEGER`",
                      ["kind"] = "query",
                      ["example"] = 1,
                    },
                    {
                      ["name"] = "per_page",
                      ["orig"] = "per_page",
                      ["type"] = "`$INTEGER`",
                      ["kind"] = "query",
                      ["example"] = 10,
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "page",
                    "per_page",
                  },
                },
              },
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/upload-token",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_name",
                  },
                  {
                    ["lit"] = "upload-token",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_name}",
                  "upload-token",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "project_name",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "project_name",
                  },
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["project_name"] = "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                  },
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["project_name"] = "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                  },
                },
              },
              {
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}/source",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "project_name",
                  },
                  {
                    ["lit"] = "source",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{project_name}",
                  "source",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "project_name",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["$action"] = "source",
                  ["exist"] = {
                    "account_id",
                    "project_name",
                  },
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/accounts/{account_id}/pages/projects/{project_name}",
                ["segments"] = {
                  {
                    ["lit"] = "accounts",
                  },
                  {
                    ["var"] = "account_id",
                  },
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "projects",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "accounts",
                  "{account_id}",
                  "pages",
                  "projects",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["project_name"] = "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "023e105f4ecef8ad9ca31a8372d0c353",
                    },
                    {
                      ["name"] = "id",
                      ["orig"] = "project_name",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                      ["example"] = "this-is-my-project-01",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                  },
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["upload"] = {
        ["fields"] = {
          {
            ["name"] = "errors",
            ["title"] = "Errors",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "messages",
            ["title"] = "Messages",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "success",
            ["title"] = "Success",
            ["type"] = "`$BOOLEAN`",
            ["req"] = true,
            ["short"] = "Whether the API call was successful.",
          },
        },
        ["name"] = "upload",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/pages/assets/upload",
                ["segments"] = {
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "assets",
                  },
                  {
                    ["lit"] = "upload",
                  },
                },
                ["parts"] = {
                  "pages",
                  "assets",
                  "upload",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {},
                ["select"] = {},
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["upsert_hash"] = {
        ["fields"] = {
          {
            ["name"] = "errors",
            ["title"] = "Errors",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "hashes",
            ["title"] = "Hashes",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
            ["short"] = "List of file content hashes to register in the asset store.",
          },
          {
            ["name"] = "messages",
            ["title"] = "Messages",
            ["type"] = "`$ARRAY`",
            ["req"] = true,
          },
          {
            ["name"] = "success",
            ["title"] = "Success",
            ["type"] = "`$BOOLEAN`",
            ["req"] = true,
            ["short"] = "Whether the API call was successful.",
          },
        },
        ["name"] = "upsert_hash",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/pages/assets/upsert-hashes",
                ["segments"] = {
                  {
                    ["lit"] = "pages",
                  },
                  {
                    ["lit"] = "assets",
                  },
                  {
                    ["lit"] = "upsert-hashes",
                  },
                },
                ["parts"] = {
                  "pages",
                  "assets",
                  "upsert-hashes",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {},
                ["select"] = {},
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
