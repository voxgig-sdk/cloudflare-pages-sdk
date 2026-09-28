-- Typed models for the CloudflarePages SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
-- params (op.<name>.points[].g.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class PagesDeployment
---@field errors table
---@field filters? table
---@field id? string
---@field messages table
---@field result table
---@field success boolean

---@class PagesDeploymentLoadMatch
---@field account_id string
---@field id string
---@field project_id string

---@class PagesDeploymentListMatch
---@field account_id string
---@field project_name? string
---@field env? string
---@field page? number
---@field per_page? number
---@field deployment_id? string
---@field project_id? string

---@class PagesDeploymentCreateData
---@field account_id string
---@field deployment_id? string
---@field project_id? string
---@field project_name? string
---@field errors table
---@field filters? table
---@field id? string
---@field messages table
---@field result table
---@field success boolean

---@class PagesDeploymentRemoveMatch
---@field account_id string
---@field deployment_id? string
---@field project_id string
---@field tail_id? string
---@field id? string
---@field force? boolean

---@class PagesDomain
---@field errors table
---@field id? string
---@field messages table
---@field name string
---@field result table
---@field success boolean

---@class PagesDomainLoadMatch
---@field account_id string
---@field id string
---@field project_id string

---@class PagesDomainListMatch
---@field account_id string
---@field project_name string

---@class PagesDomainCreateData
---@field account_id string
---@field project_name string
---@field errors table
---@field id? string
---@field messages table
---@field name string
---@field result table
---@field success boolean

---@class PagesDomainUpdateData
---@field account_id string
---@field id string
---@field project_id string
---@field errors? table
---@field messages? table
---@field name? string
---@field result? table
---@field success? boolean

---@class PagesDomainRemoveMatch
---@field account_id string
---@field id string
---@field project_id string

---@class PagesPagesAssetsCheckMissing
---@field hashes table

---@class PagesPagesAssetsCheckMissingCreateData
---@field hashes table

---@class PagesProject
---@field build_config? table
---@field deployment_configs? table
---@field errors table
---@field id? string
---@field messages table
---@field name string
---@field production_branch string
---@field result table
---@field source table
---@field success boolean

---@class PagesProjectLoadMatch
---@field account_id string
---@field id string

---@class PagesProjectListMatch
---@field account_id string
---@field page? number
---@field per_page? number
---@field project_name? string

---@class PagesProjectCreateData
---@field account_id string
---@field build_config? table
---@field deployment_configs? table
---@field errors table
---@field id? string
---@field messages table
---@field name string
---@field production_branch string
---@field result table
---@field source table
---@field success boolean

---@class PagesProjectUpdateData
---@field account_id string
---@field id string
---@field build_config? table
---@field deployment_configs? table
---@field errors? table
---@field messages? table
---@field name? string
---@field production_branch? string
---@field result? table
---@field source? table
---@field success? boolean

---@class PagesProjectRemoveMatch
---@field account_id string
---@field id string

---@class Upload
---@field errors table
---@field messages table
---@field success boolean

---@class UploadCreateData
---@field errors table
---@field messages table
---@field success boolean

---@class UpsertHash
---@field errors table
---@field hashes table
---@field messages table
---@field success boolean

---@class UpsertHashCreateData
---@field errors table
---@field hashes table
---@field messages table
---@field success boolean

local M = {}

return M
