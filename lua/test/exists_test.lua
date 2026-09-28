-- CloudflarePages SDK exists test

local sdk = require("cloudflare-pages_sdk")

describe("CloudflarePagesSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
