<?php
declare(strict_types=1);

// Typed models for the CloudflarePages SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** PagesDeployment entity data model. */
class PagesDeployment
{
    public array $errors;
    public ?array $filters = null;
    public ?string $id = null;
    public array $messages;
    public array $result;
    public bool $success;
}

/** Request payload for PagesDeployment#load. */
class PagesDeploymentLoadMatch
{
    public string $account_id;
    public string $id;
    public string $project_id;
}

/** Request payload for PagesDeployment#list. */
class PagesDeploymentListMatch
{
    public string $account_id;
    public ?string $project_name = null;
    public ?string $env = null;
    public ?int $page = null;
    public ?int $per_page = null;
    public ?string $deployment_id = null;
    public ?string $project_id = null;
}

/** Request payload for PagesDeployment#create. */
class PagesDeploymentCreateData
{
    public string $account_id;
    public ?string $deployment_id = null;
    public ?string $project_id = null;
    public ?string $project_name = null;
    public array $errors;
    public ?array $filters = null;
    public ?string $id = null;
    public array $messages;
    public array $result;
    public bool $success;
}

/** Request payload for PagesDeployment#remove. */
class PagesDeploymentRemoveMatch
{
    public string $account_id;
    public ?string $deployment_id = null;
    public string $project_id;
    public ?string $tail_id = null;
    public ?string $id = null;
    public ?bool $force = null;
}

/** PagesDomain entity data model. */
class PagesDomain
{
    public array $errors;
    public ?string $id = null;
    public array $messages;
    public string $name;
    public array $result;
    public bool $success;
}

/** Request payload for PagesDomain#load. */
class PagesDomainLoadMatch
{
    public string $account_id;
    public string $id;
    public string $project_id;
}

/** Request payload for PagesDomain#list. */
class PagesDomainListMatch
{
    public string $account_id;
    public string $project_name;
}

/** Request payload for PagesDomain#create. */
class PagesDomainCreateData
{
    public string $account_id;
    public string $project_name;
    public array $errors;
    public ?string $id = null;
    public array $messages;
    public string $name;
    public array $result;
    public bool $success;
}

/** Request payload for PagesDomain#update. */
class PagesDomainUpdateData
{
    public string $account_id;
    public string $id;
    public string $project_id;
    public ?array $errors = null;
    public ?array $messages = null;
    public ?string $name = null;
    public ?array $result = null;
    public ?bool $success = null;
}

/** Request payload for PagesDomain#remove. */
class PagesDomainRemoveMatch
{
    public string $account_id;
    public string $id;
    public string $project_id;
}

/** PagesPagesAssetsCheckMissing entity data model. */
class PagesPagesAssetsCheckMissing
{
    public array $hashes;
}

/** Request payload for PagesPagesAssetsCheckMissing#create. */
class PagesPagesAssetsCheckMissingCreateData
{
    public array $hashes;
}

/** PagesProject entity data model. */
class PagesProject
{
    public ?array $build_config = null;
    public ?array $deployment_configs = null;
    public array $errors;
    public ?string $id = null;
    public array $messages;
    public string $name;
    public string $production_branch;
    public array $result;
    public array $source;
    public bool $success;
}

/** Request payload for PagesProject#load. */
class PagesProjectLoadMatch
{
    public string $account_id;
    public string $id;
}

/** Request payload for PagesProject#list. */
class PagesProjectListMatch
{
    public string $account_id;
    public ?int $page = null;
    public ?int $per_page = null;
    public ?string $project_name = null;
}

/** Request payload for PagesProject#create. */
class PagesProjectCreateData
{
    public string $account_id;
    public ?array $build_config = null;
    public ?array $deployment_configs = null;
    public array $errors;
    public ?string $id = null;
    public array $messages;
    public string $name;
    public string $production_branch;
    public array $result;
    public array $source;
    public bool $success;
}

/** Request payload for PagesProject#update. */
class PagesProjectUpdateData
{
    public string $account_id;
    public string $id;
    public ?array $build_config = null;
    public ?array $deployment_configs = null;
    public ?array $errors = null;
    public ?array $messages = null;
    public ?string $name = null;
    public ?string $production_branch = null;
    public ?array $result = null;
    public ?array $source = null;
    public ?bool $success = null;
}

/** Request payload for PagesProject#remove. */
class PagesProjectRemoveMatch
{
    public string $account_id;
    public string $id;
}

/** Upload entity data model. */
class Upload
{
    public array $errors;
    public array $messages;
    public bool $success;
}

/** Request payload for Upload#create. */
class UploadCreateData
{
    public array $errors;
    public array $messages;
    public bool $success;
}

/** UpsertHash entity data model. */
class UpsertHash
{
    public array $errors;
    public array $hashes;
    public array $messages;
    public bool $success;
}

/** Request payload for UpsertHash#create. */
class UpsertHashCreateData
{
    public array $errors;
    public array $hashes;
    public array $messages;
    public bool $success;
}

