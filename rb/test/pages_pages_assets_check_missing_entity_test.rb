# PagesPagesAssetsCheckMissing entity test

require "minitest/autorun"
require "json"
require_relative "../CloudflarePages_sdk"
require_relative "runner"

class PagesPagesAssetsCheckMissingEntityTest < Minitest::Test
  def test_create_instance
    testsdk = CloudflarePagesSDK.test(nil, nil)
    ent = testsdk.PagesPagesAssetsCheckMissing(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = pages_pages_assets_check_missing_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "pages_pages_assets_check_missing." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    pages_pages_assets_check_missing_ref01_ent = client.PagesPagesAssetsCheckMissing(nil)
    pages_pages_assets_check_missing_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.pages_pages_assets_check_missing"), "pages_pages_assets_check_missing_ref01"))

    pages_pages_assets_check_missing_ref01_data_result = pages_pages_assets_check_missing_ref01_ent.create(pages_pages_assets_check_missing_ref01_data, nil)
    pages_pages_assets_check_missing_ref01_data = Helpers.to_map(pages_pages_assets_check_missing_ref01_data_result.respond_to?(:data_get) ? pages_pages_assets_check_missing_ref01_data_result.data_get : pages_pages_assets_check_missing_ref01_data_result)
    assert !pages_pages_assets_check_missing_ref01_data.nil?

  end
end

def pages_pages_assets_check_missing_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "pages_pages_assets_check_missing", "PagesPagesAssetsCheckMissingTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = CloudflarePagesSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["pages_pages_assets_check_missing01", "pages_pages_assets_check_missing02", "pages_pages_assets_check_missing03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID" => idmap,
    "CLOUDFLARE_PAGES_TEST_LIVE" => "FALSE",
    "CLOUDFLARE_PAGES_TEST_EXPLAIN" => "FALSE",
    "CLOUDFLARE_PAGES_APIKEY" => "",
  })

  idmap_resolved = Helpers.to_map(
    env["CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["CLOUDFLARE_PAGES_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      # FIRST, so the generated fields below win: sdk-test-control.json's
      # test.client.options adds to the live client, it does not redirect it.
      Runner.live_client_options,
      {
        "apikey" => env["CLOUDFLARE_PAGES_APIKEY"],
      },
      extra || {},
    ])
    client = CloudflarePagesSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["CLOUDFLARE_PAGES_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["CLOUDFLARE_PAGES_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
