"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('PagesProjectEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when CLOUDFLARE_PAGES_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('CLOUDFLARE_PAGES_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.CloudflarePagesSDK.test();
        const ent = testsdk.PagesProject();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.CLOUDFLARE_PAGES_TEST_LIVE;
        for (const op of ['create', 'list', 'update', 'load', 'remove']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'pages_project.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "build_config": { "a": true, "h": "Build Config", "n": "build_config", "r": false, "sh": "Configs for the project build process.", "t": "`$OBJECT`", "key$": "build_config", "index$": 0 }, "deployment_configs": { "a": true, "h": "Deployment Configs", "n": "deployment_configs", "r": false, "sh": "Configs for deployments in a project.", "t": "`$OBJECT`", "key$": "deployment_configs", "index$": 1 }, "errors": { "a": true, "h": "Errors", "n": "errors", "r": true, "t": "`$ARRAY`", "key$": "errors", "index$": 2 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "t": "`$STRING`", "key$": "id", "index$": 3 }, "messages": { "a": true, "h": "Messages", "n": "messages", "r": true, "t": "`$ARRAY`", "key$": "messages", "index$": 4 }, "name": { "a": true, "h": "Name", "n": "name", "op": { "update": { "req": false, "type": "`$STRING`" } }, "r": true, "sh": "Name for the Pages project.", "t": "`$STRING`", "key$": "name", "index$": 5 }, "production_branch": { "a": true, "h": "Production Branch", "n": "production_branch", "op": { "update": { "req": false, "type": "`$STRING`" } }, "r": true, "sh": "Production branch of the project.", "t": "`$STRING`", "key$": "production_branch", "index$": 6 }, "result": { "a": true, "h": "Result", "n": "result", "r": true, "t": "`$OBJECT`", "key$": "result", "index$": 7 }, "source": { "a": true, "h": "Source", "n": "source", "r": true, "sh": "Configs for the project source control.", "t": "`$OBJECT`", "key$": "source", "index$": 8 }, "success": { "a": true, "h": "Success", "n": "success", "r": true, "sh": "Whether the API call was successful.", "t": "`$BOOLEAN`", "key$": "success", "index$": 9 } }, "id": { "field": "id", "name": "id" }, "name": "pages_project", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /accounts/{account_id}/pages/projects/{project_name}/purge_build_cache", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "this-is-my-project-01", "k": "param", "n": "project_name", "or": "project_name", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "POST", "o": "/accounts/{account_id}/pages/projects/{project_name}/purge_build_cache", "q": { "$action": "purge_build_cache", "exist": ["account_id", "project_name"] }, "r": {}, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "pages" }, { "lit": "projects" }, { "var": "project_name" }, { "lit": "purge_build_cache" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "POST /accounts/{account_id}/pages/projects/{project_name}/source", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "this-is-my-project-01", "k": "param", "n": "project_name", "or": "project_name", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "POST", "o": "/accounts/{account_id}/pages/projects/{project_name}/source", "q": { "$action": "source", "exist": ["account_id", "project_name"] }, "r": {}, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "pages" }, { "lit": "projects" }, { "var": "project_name" }, { "lit": "source" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }, { "a": true, "co": { "id": "POST /accounts/{account_id}/pages/projects", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/accounts/{account_id}/pages/projects", "q": { "exist": ["account_id"] }, "r": {}, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "pages" }, { "lit": "projects" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 2 }], "key$": "create" }, "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /accounts/{account_id}/pages/projects", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }], "query": [{ "a": true, "ex": 1, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "ex": 10, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/accounts/{account_id}/pages/projects", "q": { "exist": ["account_id", "page", "per_page"] }, "r": {}, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "pages" }, { "lit": "projects" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "GET /accounts/{account_id}/pages/projects/{project_name}/upload-token", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "this-is-my-project-01", "k": "param", "n": "project_name", "or": "project_name", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/accounts/{account_id}/pages/projects/{project_name}/upload-token", "q": { "exist": ["account_id", "project_name"] }, "r": {}, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "pages" }, { "lit": "projects" }, { "var": "project_name" }, { "lit": "upload-token" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /accounts/{account_id}/pages/projects/{project_name}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "this-is-my-project-01", "k": "param", "n": "id", "or": "project_name", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/accounts/{account_id}/pages/projects/{project_name}", "q": { "exist": ["account_id", "id"] }, "r": { "param": { "project_name": "id" } }, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "pages" }, { "lit": "projects" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" }, "remove": { "input": "data", "name": "remove", "points": [{ "a": true, "co": { "id": "DELETE /accounts/{account_id}/pages/projects/{project_name}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "this-is-my-project-01", "k": "param", "n": "id", "or": "project_name", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/accounts/{account_id}/pages/projects/{project_name}", "q": { "exist": ["account_id", "id"] }, "r": { "param": { "project_name": "id" } }, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "pages" }, { "lit": "projects" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "DELETE /accounts/{account_id}/pages/projects/{project_name}/source", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "this-is-my-project-01", "k": "param", "n": "project_name", "or": "project_name", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/accounts/{account_id}/pages/projects/{project_name}/source", "q": { "$action": "source", "exist": ["account_id", "project_name"] }, "r": {}, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "pages" }, { "lit": "projects" }, { "var": "project_name" }, { "lit": "source" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }], "key$": "remove" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PATCH /accounts/{account_id}/pages/projects/{project_name}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "this-is-my-project-01", "k": "param", "n": "id", "or": "project_name", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PATCH", "o": "/accounts/{account_id}/pages/projects/{project_name}", "q": { "exist": ["account_id", "id"] }, "r": { "param": { "project_name": "id" } }, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "pages" }, { "lit": "projects" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [] }, "key$": "pages_project", "name__orig": "pages_project", "Name": "PagesProject", "name_": "pages_project", "name-": "pages-project", "NAME": "PAGES_PROJECT", "index$": 3 }, { "active": true, "entity": "pages_project", "key$": "BasicPagesProjectFlow", "kind": "basic", "name": "BasicPagesProjectFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "pages_project_ref01" }, "m": { "account_id": "account01", "project_name": "project_name01" }, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": {}, "m": { "account_id": "account01", "project_name": "project_name01" }, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "pages_project_ref01" } }], "index$": 1 }, { "a": true, "d": { "account_id": "account01" }, "i": { "ref": "pages_project_ref01", "srcdatavar": "pages_project_ref01_data", "suffix": "_up0", "textfield": "name" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-pages_project_ref01" } }], "v": [], "index$": 2 }, { "a": true, "d": {}, "i": { "ref": "pages_project_ref01", "srcdatavar": "pages_project_ref01_data", "suffix": "_dt0" }, "m": { "account_id": "account01", "id": "pages_project01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-pages_project_ref01" } }], "index$": 3 }, { "a": true, "d": {}, "i": { "ref": "pages_project_ref01", "suffix": "_rm0" }, "m": { "account_id": "account01", "id": "pages_project01" }, "o": "remove", "s": [], "v": [], "index$": 4 }, { "a": true, "d": {}, "i": { "suffix": "_rt0" }, "m": { "account_id": "account01", "project_name": "project_name01" }, "o": "list", "s": [], "v": [{ "apply": "ItemNotExists", "def": { "ref": "pages_project_ref01" } }], "index$": 5 }] }, 'PagesProject', { "POST /accounts/{account_id}/pages/projects/{project_name}/purge_build_cache": { "protocol": "http", "parameters": [{ "name": "project_name", "in": "path", "required": true, "schema": { "description": "Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.", "type": "string", "example": "this-is-my-project-01", "pattern": "^[a-z0-9][a-z0-9-]*$", "x-auditable": true, "x-ref": "#/components/schemas/pages_project_name" }, "index$": 0 }, { "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/pages_identifier" }, "index$": 1 }] }, "POST /accounts/{account_id}/pages/projects/{project_name}/source": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "description": "Repository and deployment settings for connecting a Git repository to a Pages project.", "type": "object", "properties": { "config": { "allOf": [{ "type": "object", "properties": { "deployments_enabled": {}, "owner": {}, "owner_id": {}, "path_excludes": {}, "path_includes": {}, "pr_comments_enabled": {}, "preview_branch_excludes": {}, "preview_branch_includes": {}, "preview_deployment_setting": {}, "production_branch": {}, "production_deployments_enabled": {}, "repo_id": {}, "repo_name": {} }, "required": ["deployments_enabled", "owner", "owner_id", "path_excludes", "path_includes", "pr_comments_enabled", "preview_branch_excludes", "preview_branch_includes", "preview_deployment_setting", "production_branch", "production_deployments_enabled", "repo_id", "repo_name"], "x-ref": "#/components/schemas/pages_source_config" }], "x-fern-property-name": "source-config" }, "type": { "description": "The source control management provider.", "type": "string", "example": "github", "enum": ["github", "gitlab"], "x-auditable": true, "x-ref": "#/components/schemas/pages_source_type" } }, "required": ["type", "config"], "x-ref": "#/components/schemas/pages_source_connect_request" } } } }, "parameters": [{ "name": "project_name", "in": "path", "required": true, "schema": { "description": "Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.", "type": "string", "example": "this-is-my-project-01", "pattern": "^[a-z0-9][a-z0-9-]*$", "x-auditable": true, "x-ref": "#/components/schemas/pages_project_name" }, "index$": 0 }, { "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/pages_identifier" }, "index$": 1 }] }, "POST /accounts/{account_id}/pages/projects": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "type": "object", "properties": { "build_config": { "description": "Configs for the project build process.", "type": "object", "properties": { "build_caching": { "description": "Enable build caching for the project.", "type": "boolean", "example": true, "x-auditable": true }, "build_command": { "description": "Command used to build project.", "type": "string", "example": "npm run build", "x-auditable": true }, "destination_dir": { "description": "Output directory of the build.", "type": "string", "example": "build", "x-auditable": true }, "root_dir": { "description": "Directory to run the command.", "type": "string", "example": "/", "x-auditable": true }, "web_analytics_tag": { "description": "The classifying tag for analytics.", "type": "string", "example": "cee1c73f6e4743d0b5e6bb1a0bcaabcc", "nullable": true, "x-auditable": true }, "web_analytics_token": { "description": "The auth token for analytics.", "type": "string", "example": "021e1057c18547eca7b79f2516f06o7x", "nullable": true, "x-sensitive": true } }, "key$": "build_config" }, "deployment_configs": { "description": "Configs for deployments in a project.", "type": "object", "properties": { "preview": { "description": "Configs for preview deploys.", "allOf": [{}] }, "production": { "description": "Configs for production deploys.", "allOf": [{}] } }, "key$": "deployment_configs" }, "name": { "description": "Name for the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.", "type": "string", "example": "my-pages-app", "x-auditable": true, "key$": "name" }, "production_branch": { "description": "Production branch of the project. Used to identify production deployments.", "type": "string", "example": "main", "x-auditable": true, "key$": "production_branch" }, "source": { "description": "Configs for the project source control.", "type": "object", "properties": { "config": { "type": "object", "properties": { "deployments_enabled": {}, "owner": {}, "owner_id": {}, "path_excludes": {}, "path_includes": {}, "pr_comments_enabled": {}, "preview_branch_excludes": {}, "preview_branch_includes": {}, "preview_deployment_setting": {}, "production_branch": {}, "production_deployments_enabled": {}, "repo_id": {}, "repo_name": {} } }, "type": { "description": "The source control management provider.", "type": "string", "example": "github", "enum": ["github", "gitlab"], "x-auditable": true } }, "required": ["type", "config"], "key$": "source" } }, "required": ["name", "production_branch"], "index$": 1 } } } }, "parameters": [{ "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/pages_identifier" }, "index$": 0 }] }, "GET /accounts/{account_id}/pages/projects": { "protocol": "http", "parameters": [{ "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/pages_identifier" }, "index$": 0 }, { "name": "page", "in": "query", "schema": { "description": "Page number of results to return.", "type": "integer", "example": 1 }, "index$": 1 }, { "name": "per_page", "in": "query", "schema": { "description": "Number of results to return per page.", "type": "integer", "example": 10 }, "index$": 2 }] }, "GET /accounts/{account_id}/pages/projects/{project_name}/upload-token": { "protocol": "http", "parameters": [{ "name": "project_name", "in": "path", "required": true, "schema": { "description": "Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.", "type": "string", "example": "this-is-my-project-01", "pattern": "^[a-z0-9][a-z0-9-]*$", "x-auditable": true, "x-ref": "#/components/schemas/pages_project_name" }, "index$": 0 }, { "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/pages_identifier" }, "index$": 1 }] }, "GET /accounts/{account_id}/pages/projects/{project_name}": { "protocol": "http", "parameters": [{ "name": "project_name", "in": "path", "required": true, "schema": { "description": "Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.", "type": "string", "example": "this-is-my-project-01", "pattern": "^[a-z0-9][a-z0-9-]*$", "x-auditable": true, "x-ref": "#/components/schemas/pages_project_name" }, "index$": 0 }, { "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/pages_identifier" }, "index$": 1 }] }, "DELETE /accounts/{account_id}/pages/projects/{project_name}": { "protocol": "http", "parameters": [{ "name": "project_name", "in": "path", "required": true, "schema": { "description": "Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.", "type": "string", "example": "this-is-my-project-01", "pattern": "^[a-z0-9][a-z0-9-]*$", "x-auditable": true, "x-ref": "#/components/schemas/pages_project_name" }, "index$": 0 }, { "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/pages_identifier" }, "index$": 1 }] }, "DELETE /accounts/{account_id}/pages/projects/{project_name}/source": { "protocol": "http", "parameters": [{ "name": "project_name", "in": "path", "required": true, "schema": { "description": "Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.", "type": "string", "example": "this-is-my-project-01", "pattern": "^[a-z0-9][a-z0-9-]*$", "x-auditable": true, "x-ref": "#/components/schemas/pages_project_name" }, "index$": 0 }, { "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/pages_identifier" }, "index$": 1 }] }, "PATCH /accounts/{account_id}/pages/projects/{project_name}": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "type": "object", "properties": { "build_config": { "description": "Configs for the project build process.", "type": "object", "properties": { "build_caching": { "description": "Enable build caching for the project.", "type": "boolean", "example": true, "x-auditable": true }, "build_command": { "description": "Command used to build project.", "type": "string", "example": "npm run build", "x-auditable": true }, "destination_dir": { "description": "Output directory of the build.", "type": "string", "example": "build", "x-auditable": true }, "root_dir": { "description": "Directory to run the command.", "type": "string", "example": "/", "x-auditable": true }, "web_analytics_tag": { "description": "The classifying tag for analytics.", "type": "string", "example": "cee1c73f6e4743d0b5e6bb1a0bcaabcc", "nullable": true, "x-auditable": true }, "web_analytics_token": { "description": "The auth token for analytics.", "type": "string", "example": "021e1057c18547eca7b79f2516f06o7x", "nullable": true, "x-sensitive": true } }, "key$": "build_config" }, "deployment_configs": { "description": "Configs for deployments in a project.", "type": "object", "properties": { "preview": { "description": "Configs for preview deploys.", "allOf": [{}] }, "production": { "description": "Configs for production deploys.", "allOf": [{}] } }, "key$": "deployment_configs" }, "name": { "description": "Name for the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.", "type": "string", "example": "my-pages-app", "x-auditable": true, "key$": "name" }, "production_branch": { "description": "Production branch of the project. Used to identify production deployments.", "type": "string", "example": "main", "x-auditable": true, "key$": "production_branch" }, "source": { "description": "Configs for the project source control.", "type": "object", "properties": { "config": { "type": "object", "properties": { "deployments_enabled": {}, "owner": {}, "owner_id": {}, "path_excludes": {}, "path_includes": {}, "pr_comments_enabled": {}, "preview_branch_excludes": {}, "preview_branch_includes": {}, "preview_deployment_setting": {}, "production_branch": {}, "production_deployments_enabled": {}, "repo_id": {}, "repo_name": {} } }, "type": { "description": "The source control management provider.", "type": "string", "example": "github", "enum": ["github", "gitlab"], "x-auditable": true } }, "required": ["type", "config"], "key$": "source" } }, "example": { "deployment_configs": { "production": { "compatibility_date": "2022-01-01T00:00:00Z", "compatibility_flags": ["url_standard"], "env_vars": { "NODE_VERSION": { "value": "22" }, "delete_this_env_var": null, "secret_var": { "type": "secret_text", "value": "A_CMS_API_TOKEN" } } } } }, "index$": 1 } } } }, "parameters": [{ "name": "project_name", "in": "path", "required": true, "schema": { "description": "Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.", "type": "string", "example": "this-is-my-project-01", "pattern": "^[a-z0-9][a-z0-9-]*$", "x-auditable": true, "x-ref": "#/components/schemas/pages_project_name" }, "index$": 0 }, { "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/pages_identifier" }, "index$": 1 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const pages_project_ref01_ent = client.PagesProject();
        let pages_project_ref01_data = setup.data.new.pages_project['pages_project_ref01'];
        pages_project_ref01_data['account_id'] = setup.idmap['account01'];
        pages_project_ref01_data['project_name'] = setup.idmap['project_name01'];
        pages_project_ref01_data = (await pages_project_ref01_ent.create(pages_project_ref01_data)).data();
        (0, node_assert_1.default)(null != pages_project_ref01_data.id);
        // LIST
        const pages_project_ref01_match = {};
        pages_project_ref01_match['account_id'] = setup.idmap['account01'];
        pages_project_ref01_match['project_name'] = setup.idmap['project_name01'];
        const pages_project_ref01_list = (await pages_project_ref01_ent.list(pages_project_ref01_match)).map((e) => e.data());
        (0, node_assert_1.default)(!isempty(select(pages_project_ref01_list, { id: pages_project_ref01_data.id })));
        // UPDATE
        const pages_project_ref01_data_up0 = {};
        pages_project_ref01_data_up0.id = pages_project_ref01_data.id;
        pages_project_ref01_data_up0['account_id'] = setup.idmap['account_id'];
        const pages_project_ref01_markdef_up0 = { name: 'name', value: 'Mark01-pages_project_ref01_' + setup.now };
        pages_project_ref01_data_up0[pages_project_ref01_markdef_up0.name] = pages_project_ref01_markdef_up0.value;
        const pages_project_ref01_resdata_up0 = (await pages_project_ref01_ent.update(pages_project_ref01_data_up0)).data();
        (0, node_assert_1.default)(pages_project_ref01_resdata_up0.id === pages_project_ref01_data_up0.id);
        (0, node_assert_1.default)(pages_project_ref01_resdata_up0[pages_project_ref01_markdef_up0.name] === pages_project_ref01_markdef_up0.value);
        // LOAD
        const pages_project_ref01_match_dt0 = {};
        pages_project_ref01_match_dt0.id = pages_project_ref01_data.id;
        const pages_project_ref01_data_dt0 = (await pages_project_ref01_ent.load(pages_project_ref01_match_dt0)).data();
        (0, node_assert_1.default)(pages_project_ref01_data_dt0.id === pages_project_ref01_data.id);
        // REMOVE
        const pages_project_ref01_match_rm0 = { id: pages_project_ref01_data.id };
        await pages_project_ref01_ent.remove(pages_project_ref01_match_rm0);
        // LIST
        const pages_project_ref01_match_rt0 = {};
        pages_project_ref01_match_rt0['account_id'] = setup.idmap['account01'];
        pages_project_ref01_match_rt0['project_name'] = setup.idmap['project_name01'];
        const pages_project_ref01_list_rt0 = (await pages_project_ref01_ent.list(pages_project_ref01_match_rt0)).map((e) => e.data());
        (0, node_assert_1.default)(isempty(select(pages_project_ref01_list_rt0, { id: pages_project_ref01_data.id })));
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/pages_project/PagesProjectTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.CloudflarePagesSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['pages_project01', 'pages_project02', 'pages_project03', 'account01', 'project_name01'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'CLOUDFLARE_PAGES_TEST_PAGES_PROJECT_ENTID': idmap,
        'CLOUDFLARE_PAGES_TEST_LIVE': 'FALSE',
        'CLOUDFLARE_PAGES_TEST_EXPLAIN': 'FALSE',
        'CLOUDFLARE_PAGES_APIKEY': '',
    });
    idmap = env['CLOUDFLARE_PAGES_TEST_PAGES_PROJECT_ENTID'];
    const live = 'TRUE' === env.CLOUDFLARE_PAGES_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['CLOUDFLARE_PAGES_TEST_PAGES_PROJECT_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.CloudflarePagesSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.CLOUDFLARE_PAGES_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.CLOUDFLARE_PAGES_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=PagesProjectEntity.test.js.map