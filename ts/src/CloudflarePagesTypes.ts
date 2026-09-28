// Typed models for the CloudflarePages SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface PagesDeployment {
  errors: any[]
  filters?: any[]
  id?: string
  messages: any[]
  result: Record<string, any>
  success: boolean
}

export interface PagesDeploymentLoadMatch {
  account_id: string
  id: string
  project_id: string
}

export interface PagesDeploymentListMatch {
  account_id: string
  project_name?: string
  env?: string
  page?: number
  per_page?: number
  deployment_id?: string
  project_id?: string
}

export interface PagesDeploymentCreateData {
  account_id: string
  deployment_id?: string
  project_id?: string
  project_name?: string
  errors: any[]
  filters?: any[]
  id?: string
  messages: any[]
  result: Record<string, any>
  success: boolean

  // Selects a custom action instead of the plain create:
  //   'retry' | 'rollback'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface PagesDeploymentRemoveMatch {
  account_id: string
  deployment_id?: string
  project_id: string
  tail_id?: string
  id?: string
  force?: boolean
}

export interface PagesDomain {
  errors: any[]
  id?: string
  messages: any[]
  name: string
  result: Record<string, any>
  success: boolean
}

export interface PagesDomainLoadMatch {
  account_id: string
  id: string
  project_id: string
}

export interface PagesDomainListMatch {
  account_id: string
  project_name: string
}

export interface PagesDomainCreateData {
  account_id: string
  project_name: string
  errors: any[]
  id?: string
  messages: any[]
  name: string
  result: Record<string, any>
  success: boolean
}

export interface PagesDomainUpdateData {
  account_id: string
  id: string
  project_id: string
  errors?: any[]
  messages?: any[]
  name?: string
  result?: Record<string, any>
  success?: boolean
}

export interface PagesDomainRemoveMatch {
  account_id: string
  id: string
  project_id: string
}

export interface PagesPagesAssetsCheckMissing {
  hashes: any[]
}

export interface PagesPagesAssetsCheckMissingCreateData {
  hashes: any[]
}

export interface PagesProject {
  build_config?: Record<string, any>
  deployment_configs?: Record<string, any>
  errors: any[]
  id?: string
  messages: any[]
  name: string
  production_branch: string
  result: Record<string, any>
  source: Record<string, any>
  success: boolean
}

export interface PagesProjectLoadMatch {
  account_id: string
  id: string
}

export interface PagesProjectListMatch {
  account_id: string
  page?: number
  per_page?: number
  project_name?: string
}

export interface PagesProjectCreateData {
  account_id: string
  build_config?: Record<string, any>
  deployment_configs?: Record<string, any>
  errors: any[]
  id?: string
  messages: any[]
  name: string
  production_branch: string
  result: Record<string, any>
  source: Record<string, any>
  success: boolean

  // Selects a custom action instead of the plain create:
  //   'purge_build_cache' | 'source'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface PagesProjectUpdateData {
  account_id: string
  id: string
  build_config?: Record<string, any>
  deployment_configs?: Record<string, any>
  errors?: any[]
  messages?: any[]
  name?: string
  production_branch?: string
  result?: Record<string, any>
  source?: Record<string, any>
  success?: boolean
}

export interface PagesProjectRemoveMatch {
  account_id: string
  id: string

  // Selects a custom action instead of the plain remove:
  //   'source'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface Upload {
  errors: any[]
  messages: any[]
  success: boolean
}

export interface UploadCreateData {
  errors: any[]
  messages: any[]
  success: boolean
}

export interface UpsertHash {
  errors: any[]
  hashes: any[]
  messages: any[]
  success: boolean
}

export interface UpsertHashCreateData {
  errors: any[]
  hashes: any[]
  messages: any[]
  success: boolean
}

