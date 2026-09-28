# CloudflarePages SDK exists test

require "minitest/autorun"
require_relative "../CloudflarePages_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = CloudflarePagesSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
