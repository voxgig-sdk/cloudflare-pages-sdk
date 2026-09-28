# frozen_string_literal: true

# Typed models for the CloudflarePages SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# PagesDeployment entity data model.
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] filters
#   @return [Array, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] success
#   @return [Boolean]
PagesDeployment = Struct.new(
  :errors,
  :filters,
  :id,
  :messages,
  :result,
  :success,
  keyword_init: true
)

# Request payload for PagesDeployment#load.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] project_id
#   @return [String]
PagesDeploymentLoadMatch = Struct.new(
  :account_id,
  :id,
  :project_id,
  keyword_init: true
)

# Request payload for PagesDeployment#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] project_name
#   @return [String, nil]
#
# @!attribute [rw] env
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [Integer, nil]
#
# @!attribute [rw] per_page
#   @return [Integer, nil]
#
# @!attribute [rw] deployment_id
#   @return [String, nil]
#
# @!attribute [rw] project_id
#   @return [String, nil]
PagesDeploymentListMatch = Struct.new(
  :account_id,
  :project_name,
  :env,
  :page,
  :per_page,
  :deployment_id,
  :project_id,
  keyword_init: true
)

# Request payload for PagesDeployment#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] deployment_id
#   @return [String, nil]
#
# @!attribute [rw] project_id
#   @return [String, nil]
#
# @!attribute [rw] project_name
#   @return [String, nil]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] filters
#   @return [Array, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] success
#   @return [Boolean]
PagesDeploymentCreateData = Struct.new(
  :account_id,
  :deployment_id,
  :project_id,
  :project_name,
  :errors,
  :filters,
  :id,
  :messages,
  :result,
  :success,
  keyword_init: true
)

# Request payload for PagesDeployment#remove.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] deployment_id
#   @return [String, nil]
#
# @!attribute [rw] project_id
#   @return [String]
#
# @!attribute [rw] tail_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] force
#   @return [Boolean, nil]
PagesDeploymentRemoveMatch = Struct.new(
  :account_id,
  :deployment_id,
  :project_id,
  :tail_id,
  :id,
  :force,
  keyword_init: true
)

# PagesDomain entity data model.
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] success
#   @return [Boolean]
PagesDomain = Struct.new(
  :errors,
  :id,
  :messages,
  :name,
  :result,
  :success,
  keyword_init: true
)

# Request payload for PagesDomain#load.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] project_id
#   @return [String]
PagesDomainLoadMatch = Struct.new(
  :account_id,
  :id,
  :project_id,
  keyword_init: true
)

# Request payload for PagesDomain#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] project_name
#   @return [String]
PagesDomainListMatch = Struct.new(
  :account_id,
  :project_name,
  keyword_init: true
)

# Request payload for PagesDomain#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] project_name
#   @return [String]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] success
#   @return [Boolean]
PagesDomainCreateData = Struct.new(
  :account_id,
  :project_name,
  :errors,
  :id,
  :messages,
  :name,
  :result,
  :success,
  keyword_init: true
)

# Request payload for PagesDomain#update.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] project_id
#   @return [String]
#
# @!attribute [rw] errors
#   @return [Array, nil]
#
# @!attribute [rw] messages
#   @return [Array, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] success
#   @return [Boolean, nil]
PagesDomainUpdateData = Struct.new(
  :account_id,
  :id,
  :project_id,
  :errors,
  :messages,
  :name,
  :result,
  :success,
  keyword_init: true
)

# Request payload for PagesDomain#remove.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] project_id
#   @return [String]
PagesDomainRemoveMatch = Struct.new(
  :account_id,
  :id,
  :project_id,
  keyword_init: true
)

# PagesPagesAssetsCheckMissing entity data model.
#
# @!attribute [rw] hashes
#   @return [Array]
PagesPagesAssetsCheckMissing = Struct.new(
  :hashes,
  keyword_init: true
)

# Request payload for PagesPagesAssetsCheckMissing#create.
#
# @!attribute [rw] hashes
#   @return [Array]
PagesPagesAssetsCheckMissingCreateData = Struct.new(
  :hashes,
  keyword_init: true
)

# PagesProject entity data model.
#
# @!attribute [rw] build_config
#   @return [Hash, nil]
#
# @!attribute [rw] deployment_configs
#   @return [Hash, nil]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] production_branch
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] source
#   @return [Hash]
#
# @!attribute [rw] success
#   @return [Boolean]
PagesProject = Struct.new(
  :build_config,
  :deployment_configs,
  :errors,
  :id,
  :messages,
  :name,
  :production_branch,
  :result,
  :source,
  :success,
  keyword_init: true
)

# Request payload for PagesProject#load.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
PagesProjectLoadMatch = Struct.new(
  :account_id,
  :id,
  keyword_init: true
)

# Request payload for PagesProject#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] page
#   @return [Integer, nil]
#
# @!attribute [rw] per_page
#   @return [Integer, nil]
#
# @!attribute [rw] project_name
#   @return [String, nil]
PagesProjectListMatch = Struct.new(
  :account_id,
  :page,
  :per_page,
  :project_name,
  keyword_init: true
)

# Request payload for PagesProject#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] build_config
#   @return [Hash, nil]
#
# @!attribute [rw] deployment_configs
#   @return [Hash, nil]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] production_branch
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] source
#   @return [Hash]
#
# @!attribute [rw] success
#   @return [Boolean]
PagesProjectCreateData = Struct.new(
  :account_id,
  :build_config,
  :deployment_configs,
  :errors,
  :id,
  :messages,
  :name,
  :production_branch,
  :result,
  :source,
  :success,
  keyword_init: true
)

# Request payload for PagesProject#update.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] build_config
#   @return [Hash, nil]
#
# @!attribute [rw] deployment_configs
#   @return [Hash, nil]
#
# @!attribute [rw] errors
#   @return [Array, nil]
#
# @!attribute [rw] messages
#   @return [Array, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] production_branch
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] source
#   @return [Hash, nil]
#
# @!attribute [rw] success
#   @return [Boolean, nil]
PagesProjectUpdateData = Struct.new(
  :account_id,
  :id,
  :build_config,
  :deployment_configs,
  :errors,
  :messages,
  :name,
  :production_branch,
  :result,
  :source,
  :success,
  keyword_init: true
)

# Request payload for PagesProject#remove.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
PagesProjectRemoveMatch = Struct.new(
  :account_id,
  :id,
  keyword_init: true
)

# Upload entity data model.
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] success
#   @return [Boolean]
Upload = Struct.new(
  :errors,
  :messages,
  :success,
  keyword_init: true
)

# Request payload for Upload#create.
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] success
#   @return [Boolean]
UploadCreateData = Struct.new(
  :errors,
  :messages,
  :success,
  keyword_init: true
)

# UpsertHash entity data model.
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] hashes
#   @return [Array]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] success
#   @return [Boolean]
UpsertHash = Struct.new(
  :errors,
  :hashes,
  :messages,
  :success,
  keyword_init: true
)

# Request payload for UpsertHash#create.
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] hashes
#   @return [Array]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] success
#   @return [Boolean]
UpsertHashCreateData = Struct.new(
  :errors,
  :hashes,
  :messages,
  :success,
  keyword_init: true
)

