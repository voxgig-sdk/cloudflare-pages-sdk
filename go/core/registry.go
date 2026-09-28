package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewPagesDeploymentEntityFunc func(client *CloudflarePagesSDK, entopts map[string]any) CloudflarePagesEntity

var NewPagesDomainEntityFunc func(client *CloudflarePagesSDK, entopts map[string]any) CloudflarePagesEntity

var NewPagesPagesAssetsCheckMissingEntityFunc func(client *CloudflarePagesSDK, entopts map[string]any) CloudflarePagesEntity

var NewPagesProjectEntityFunc func(client *CloudflarePagesSDK, entopts map[string]any) CloudflarePagesEntity

var NewUploadEntityFunc func(client *CloudflarePagesSDK, entopts map[string]any) CloudflarePagesEntity

var NewUpsertHashEntityFunc func(client *CloudflarePagesSDK, entopts map[string]any) CloudflarePagesEntity

