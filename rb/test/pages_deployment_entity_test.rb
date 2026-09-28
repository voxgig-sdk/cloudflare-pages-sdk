# PagesDeployment entity test

require "minitest/autorun"
require "json"
require_relative "../CloudflarePages_sdk"
require_relative "runner"

class PagesDeploymentEntityTest < Minitest::Test
  def test_create_instance
    testsdk = CloudflarePagesSDK.test(nil, nil)
    ent = testsdk.PagesDeployment(nil)
    assert !ent.nil?
  end

  # Feature #4: the entity stream(action, ...) method runs the op pipeline and
  # returns an Enumerator over result items. With the streaming feature active
  # it yields the feature's incremental output; otherwise it falls back to the
  # materialised list so stream always yields.
  def test_stream
    seed = {
      "entity" => {
        "pages_deployment" => {
          "s1" => { "id" => "s1" },
          "s2" => { "id" => "s2" },
          "s3" => { "id" => "s3" },
        },
      },
    }

    # Fallback: streaming inactive -> yields the materialised list items.
    base = CloudflarePagesSDK.test(seed, nil)
    seen = base.PagesDeployment(nil).stream("list", nil, nil).to_a
    assert_equal 3, seen.length

    # Inbound: streaming active -> yields each item from the feature.
    cfg = CloudflarePagesConfig.shared_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("streaming")
      sdk = CloudflarePagesSDK.test(seed, { "feature" => { "streaming" => { "active" => true } } })
      got = []
      sdk.PagesDeployment(nil).stream("list", nil, nil).each do |item|
        if item.is_a?(Array)
          got.concat(item)
        else
          got << item
        end
      end
      assert_equal 3, got.length
    end
  end

  def test_basic_flow
    setup = pages_deployment_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "pages_deployment." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    pages_deployment_ref01_ent = client.PagesDeployment(nil)
    pages_deployment_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.pages_deployment"), "pages_deployment_ref01"))
    pages_deployment_ref01_data["account_id"] = setup[:idmap]["account01"]
    pages_deployment_ref01_data["deployment_id"] = setup[:idmap]["deployment01"]
    pages_deployment_ref01_data["project_id"] = setup[:idmap]["project01"]
    pages_deployment_ref01_data["project_name"] = setup[:idmap]["project_name01"]

    pages_deployment_ref01_data_result = pages_deployment_ref01_ent.create(pages_deployment_ref01_data, nil)
    pages_deployment_ref01_data = Helpers.to_map(pages_deployment_ref01_data_result.respond_to?(:data_get) ? pages_deployment_ref01_data_result.data_get : pages_deployment_ref01_data_result)
    assert !pages_deployment_ref01_data.nil?
    assert !pages_deployment_ref01_data["id"].nil?

    # LIST
    pages_deployment_ref01_match = {
      "account_id" => setup[:idmap]["account01"],
      "deployment_id" => setup[:idmap]["deployment01"],
      "project_id" => setup[:idmap]["project01"],
    }

    pages_deployment_ref01_list_result = pages_deployment_ref01_ent.list(pages_deployment_ref01_match, nil)
    assert pages_deployment_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(pages_deployment_ref01_list_result),
      { "id" => pages_deployment_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # LOAD
    pages_deployment_ref01_match_dt0 = {
      "id" => pages_deployment_ref01_data["id"],
    }
    pages_deployment_ref01_data_dt0_loaded = pages_deployment_ref01_ent.load(pages_deployment_ref01_match_dt0, nil)
    pages_deployment_ref01_data_dt0_load_result = Helpers.to_map(pages_deployment_ref01_data_dt0_loaded.respond_to?(:data_get) ? pages_deployment_ref01_data_dt0_loaded.data_get : pages_deployment_ref01_data_dt0_loaded)
    assert !pages_deployment_ref01_data_dt0_load_result.nil?
    assert_equal pages_deployment_ref01_data_dt0_load_result["id"], pages_deployment_ref01_data["id"]

    # REMOVE
    pages_deployment_ref01_match_rm0 = {
      "id" => pages_deployment_ref01_data["id"],
    }
    pages_deployment_ref01_ent.remove(pages_deployment_ref01_match_rm0, nil)

    # LIST
    pages_deployment_ref01_match_rt0 = {
      "account_id" => setup[:idmap]["account01"],
      "deployment_id" => setup[:idmap]["deployment01"],
      "project_id" => setup[:idmap]["project01"],
    }

    pages_deployment_ref01_list_rt0_result = pages_deployment_ref01_ent.list(pages_deployment_ref01_match_rt0, nil)
    assert pages_deployment_ref01_list_rt0_result.is_a?(Array)

    not_found_item = Vs.select(
      Runner.entity_list_to_data(pages_deployment_ref01_list_rt0_result),
      { "id" => pages_deployment_ref01_data["id"] })
    assert Vs.isempty(not_found_item)

  end
end

def pages_deployment_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "pages_deployment", "PagesDeploymentTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = CloudflarePagesSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["pages_deployment01", "pages_deployment02", "pages_deployment03", "account01", "deployment01", "project01", "project_name01"],
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
  entid_env_raw = ENV["CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID" => idmap,
    "CLOUDFLARE_PAGES_TEST_LIVE" => "FALSE",
    "CLOUDFLARE_PAGES_TEST_EXPLAIN" => "FALSE",
    "CLOUDFLARE_PAGES_APIKEY" => "",
  })

  idmap_resolved = Helpers.to_map(
    env["CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID"])
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
