// Typed models for the CloudflarePages SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/cloudflare-pages-sdk/go/core"
)

// PagesDeployment is the typed data model for the pages_deployment entity.
type PagesDeployment struct {
}

// PagesDeploymentLoadMatch is the typed request payload for PagesDeployment.LoadTyped.
type PagesDeploymentLoadMatch struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
	ProjectId string `json:"project_id"`
}

// PagesDeploymentListMatch is the typed request payload for PagesDeployment.ListTyped.
type PagesDeploymentListMatch struct {
	AccountId string `json:"account_id"`
	ProjectName *string `json:"project_name,omitempty"`
	Env *string `json:"env,omitempty"`
	Page *int `json:"page,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	DeploymentId *string `json:"deployment_id,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
}

// PagesDeploymentCreateData is the typed request payload for PagesDeployment.CreateTyped.
type PagesDeploymentCreateData struct {
	AccountId string `json:"account_id"`
	DeploymentId *string `json:"deployment_id,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
	ProjectName *string `json:"project_name,omitempty"`
	Errors []any `json:"errors"`
	Filters *[]any `json:"filters,omitempty"`
	Id *string `json:"id,omitempty"`
	Messages []any `json:"messages"`
	Result map[string]any `json:"result"`
	Success bool `json:"success"`
}

// PagesDeploymentRemoveMatch is the typed request payload for PagesDeployment.RemoveTyped.
type PagesDeploymentRemoveMatch struct {
	AccountId string `json:"account_id"`
	DeploymentId *string `json:"deployment_id,omitempty"`
	ProjectId string `json:"project_id"`
	TailId *string `json:"tail_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Force *bool `json:"force,omitempty"`
}

// PagesDomain is the typed data model for the pages_domain entity.
type PagesDomain struct {
}

// PagesDomainLoadMatch is the typed request payload for PagesDomain.LoadTyped.
type PagesDomainLoadMatch struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
	ProjectId string `json:"project_id"`
}

// PagesDomainListMatch is the typed request payload for PagesDomain.ListTyped.
type PagesDomainListMatch struct {
	AccountId string `json:"account_id"`
	ProjectName string `json:"project_name"`
}

// PagesDomainCreateData is the typed request payload for PagesDomain.CreateTyped.
type PagesDomainCreateData struct {
	AccountId string `json:"account_id"`
	ProjectName string `json:"project_name"`
	Errors []any `json:"errors"`
	Id *string `json:"id,omitempty"`
	Messages []any `json:"messages"`
	Name string `json:"name"`
	Result map[string]any `json:"result"`
	Success bool `json:"success"`
}

// PagesDomainUpdateData is the typed request payload for PagesDomain.UpdateTyped.
type PagesDomainUpdateData struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
	ProjectId string `json:"project_id"`
	Errors *[]any `json:"errors,omitempty"`
	Messages *[]any `json:"messages,omitempty"`
	Name *string `json:"name,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// PagesDomainRemoveMatch is the typed request payload for PagesDomain.RemoveTyped.
type PagesDomainRemoveMatch struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
	ProjectId string `json:"project_id"`
}

// PagesPagesAssetsCheckMissing is the typed data model for the pages_pages_assets_check_missing entity.
type PagesPagesAssetsCheckMissing struct {
}

// PagesPagesAssetsCheckMissingCreateData is the typed request payload for PagesPagesAssetsCheckMissing.CreateTyped.
type PagesPagesAssetsCheckMissingCreateData struct {
	Hashes []any `json:"hashes"`
}

// PagesProject is the typed data model for the pages_project entity.
type PagesProject struct {
}

// PagesProjectLoadMatch is the typed request payload for PagesProject.LoadTyped.
type PagesProjectLoadMatch struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
}

// PagesProjectListMatch is the typed request payload for PagesProject.ListTyped.
type PagesProjectListMatch struct {
	AccountId string `json:"account_id"`
	Page *int `json:"page,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	ProjectName *string `json:"project_name,omitempty"`
}

// PagesProjectCreateData is the typed request payload for PagesProject.CreateTyped.
type PagesProjectCreateData struct {
	AccountId string `json:"account_id"`
	BuildConfig *map[string]any `json:"build_config,omitempty"`
	DeploymentConfigs *map[string]any `json:"deployment_configs,omitempty"`
	Errors []any `json:"errors"`
	Id *string `json:"id,omitempty"`
	Messages []any `json:"messages"`
	Name string `json:"name"`
	ProductionBranch string `json:"production_branch"`
	Result map[string]any `json:"result"`
	Source map[string]any `json:"source"`
	Success bool `json:"success"`
}

// PagesProjectUpdateData is the typed request payload for PagesProject.UpdateTyped.
type PagesProjectUpdateData struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
	BuildConfig *map[string]any `json:"build_config,omitempty"`
	DeploymentConfigs *map[string]any `json:"deployment_configs,omitempty"`
	Errors *[]any `json:"errors,omitempty"`
	Messages *[]any `json:"messages,omitempty"`
	Name *string `json:"name,omitempty"`
	ProductionBranch *string `json:"production_branch,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Source *map[string]any `json:"source,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// PagesProjectRemoveMatch is the typed request payload for PagesProject.RemoveTyped.
type PagesProjectRemoveMatch struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
}

// Upload is the typed data model for the upload entity.
type Upload struct {
}

// UploadCreateData is the typed request payload for Upload.CreateTyped.
type UploadCreateData struct {
	Errors []any `json:"errors"`
	Messages []any `json:"messages"`
	Success bool `json:"success"`
}

// UpsertHash is the typed data model for the upsert_hash entity.
type UpsertHash struct {
}

// UpsertHashCreateData is the typed request payload for UpsertHash.CreateTyped.
type UpsertHashCreateData struct {
	Errors []any `json:"errors"`
	Hashes []any `json:"hashes"`
	Messages []any `json:"messages"`
	Success bool `json:"success"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
