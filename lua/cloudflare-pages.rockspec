package = "voxgig-sdk-cloudflare-pages-sdk"
version = "0.0.1-1"
source = {
  -- git+https (GitHub dropped git:// in 2022); pin the install to the release
  -- tag pushed by `make publish`, and point at the lua/ subdir of the monorepo.
  url = "git+https://github.com/voxgig-sdk/cloudflare-pages-sdk.git",
  tag = "lua/v0.0.1",
  dir = "cloudflare-pages-sdk/lua"
}
description = {
  summary = "Unofficial generated Lua SDK for the Cloudflare Pages public API. Not affiliated with or endorsed by the upstream API provider.",
  homepage = "https://github.com/voxgig-sdk/cloudflare-pages-sdk",
  issues_url = "https://github.com/voxgig-sdk/cloudflare-pages-sdk/issues",
  license = "MIT",
  labels = { "voxgig", "sdk", "generated-sdk", "openapi", "api-client", "cloudflare-pages" }
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["cloudflare-pages_sdk"] = "cloudflare-pages_sdk.lua",
    ["config"] = "config.lua",
    ["config_shared"] = "config_shared.lua",
    ["config_plugins"] = "config_plugins.lua",
    ["schema"] = "schema.lua",
    ["features"] = "features.lua",
    ["feature.base_feature"] = "feature/base_feature.lua",
    ["feature.test_feature"] = "feature/test_feature.lua",
  }
}
