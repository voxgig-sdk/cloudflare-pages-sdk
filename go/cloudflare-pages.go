package voxgigcloudflarepagessdk

import (
	"github.com/voxgig-sdk/cloudflare-pages-sdk/go/core"
	"github.com/voxgig-sdk/cloudflare-pages-sdk/go/entity"
	"github.com/voxgig-sdk/cloudflare-pages-sdk/go/feature"
	_ "github.com/voxgig-sdk/cloudflare-pages-sdk/go/utility"
)

// Type aliases preserve external API.
type CloudflarePagesSDK = core.CloudflarePagesSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type CloudflarePagesEntity = core.CloudflarePagesEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type CloudflarePagesError = core.CloudflarePagesError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewPagesDeploymentEntityFunc = func(client *core.CloudflarePagesSDK, entopts map[string]any) core.CloudflarePagesEntity {
		return entity.NewPagesDeploymentEntity(client, entopts)
	}
	core.NewPagesDomainEntityFunc = func(client *core.CloudflarePagesSDK, entopts map[string]any) core.CloudflarePagesEntity {
		return entity.NewPagesDomainEntity(client, entopts)
	}
	core.NewPagesPagesAssetsCheckMissingEntityFunc = func(client *core.CloudflarePagesSDK, entopts map[string]any) core.CloudflarePagesEntity {
		return entity.NewPagesPagesAssetsCheckMissingEntity(client, entopts)
	}
	core.NewPagesProjectEntityFunc = func(client *core.CloudflarePagesSDK, entopts map[string]any) core.CloudflarePagesEntity {
		return entity.NewPagesProjectEntity(client, entopts)
	}
	core.NewUploadEntityFunc = func(client *core.CloudflarePagesSDK, entopts map[string]any) core.CloudflarePagesEntity {
		return entity.NewUploadEntity(client, entopts)
	}
	core.NewUpsertHashEntityFunc = func(client *core.CloudflarePagesSDK, entopts map[string]any) core.CloudflarePagesEntity {
		return entity.NewUpsertHashEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewCloudflarePagesSDK = core.NewCloudflarePagesSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewCloudflarePagesSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *CloudflarePagesSDK  { return NewCloudflarePagesSDK(nil) }
func Test() *CloudflarePagesSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
