<?php
declare(strict_types=1);

// PagesDomain entity test

require_once __DIR__ . '/../cloudflarepages_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class PagesDomainEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = CloudflarePagesSDK::test(null, null);
        $ent = $testsdk->PagesDomain(null);
        $this->assertNotNull($ent);
    }

    // Feature #4: the entity stream(action, ...) method runs the op pipeline
    // and yields result items. With the streaming feature active it yields the
    // feature's incremental output; otherwise it falls back to the materialised
    // list so stream always yields.
    public function test_stream(): void
    {
        $seed = [
            "entity" => [
                "pages_domain" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = CloudflarePagesSDK::test($seed, null);
        $seen = iterator_to_array($base->PagesDomain(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = CloudflarePagesConfig::shared_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = CloudflarePagesSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->PagesDomain(null)->stream("list", null, null) as $item) {
                if (is_array($item) && array_is_list($item)) {
                    foreach ($item as $sub) {
                        $got[] = $sub;
                    }
                } else {
                    $got[] = $item;
                }
            }
            $this->assertCount(3, $got);
        }
    }

    public function test_basic_flow(): void
    {
        $setup = pages_domain_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "update", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "pages_domain." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $pages_domain_ref01_ent = $client->PagesDomain(null);
        $pages_domain_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.pages_domain"), "pages_domain_ref01"));
        $pages_domain_ref01_data["account_id"] = $setup["idmap"]["account01"];
        $pages_domain_ref01_data["project_id"] = $setup["idmap"]["project01"];
        $pages_domain_ref01_data["project_name"] = $setup["idmap"]["project_name01"];

        $pages_domain_ref01_data_result = $pages_domain_ref01_ent->create($pages_domain_ref01_data, null);
        $pages_domain_ref01_data = Helpers::to_map(is_object($pages_domain_ref01_data_result) && method_exists($pages_domain_ref01_data_result, 'data_get') ? $pages_domain_ref01_data_result->data_get() : $pages_domain_ref01_data_result);
        $this->assertNotNull($pages_domain_ref01_data);
        $this->assertNotNull($pages_domain_ref01_data["id"]);

        // LIST
        $pages_domain_ref01_match = [
            "account_id" => $setup["idmap"]["account01"],
            "project_name" => $setup["idmap"]["project_name01"],
        ];

        $pages_domain_ref01_list_result = $pages_domain_ref01_ent->list($pages_domain_ref01_match, null);
        $this->assertIsArray($pages_domain_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($pages_domain_ref01_list_result),
            ["id" => $pages_domain_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // UPDATE
        $pages_domain_ref01_data_up0_up = [
            "id" => $pages_domain_ref01_data["id"],
            "account_id" => $setup["idmap"]["account_id"],
            "project_id" => $setup["idmap"]["project_id"],
        ];

        $pages_domain_ref01_markdef_up0_name = "name";
        $pages_domain_ref01_markdef_up0_value = "Mark01-pages_domain_ref01_" . $setup["now"];
        $pages_domain_ref01_data_up0_up[$pages_domain_ref01_markdef_up0_name] = $pages_domain_ref01_markdef_up0_value;

        $pages_domain_ref01_resdata_up0_result = $pages_domain_ref01_ent->update($pages_domain_ref01_data_up0_up, null);
        $pages_domain_ref01_resdata_up0 = Helpers::to_map(is_object($pages_domain_ref01_resdata_up0_result) && method_exists($pages_domain_ref01_resdata_up0_result, 'data_get') ? $pages_domain_ref01_resdata_up0_result->data_get() : $pages_domain_ref01_resdata_up0_result);
        $this->assertNotNull($pages_domain_ref01_resdata_up0);
        $this->assertEquals($pages_domain_ref01_resdata_up0["id"], $pages_domain_ref01_data_up0_up["id"]);
        $this->assertEquals($pages_domain_ref01_resdata_up0[$pages_domain_ref01_markdef_up0_name], $pages_domain_ref01_markdef_up0_value);

        // LOAD
        $pages_domain_ref01_match_dt0 = [
            "id" => $pages_domain_ref01_data["id"],
        ];
        $pages_domain_ref01_data_dt0_loaded = $pages_domain_ref01_ent->load($pages_domain_ref01_match_dt0, null);
        $pages_domain_ref01_data_dt0_load_result = Helpers::to_map(is_object($pages_domain_ref01_data_dt0_loaded) && method_exists($pages_domain_ref01_data_dt0_loaded, 'data_get') ? $pages_domain_ref01_data_dt0_loaded->data_get() : $pages_domain_ref01_data_dt0_loaded);
        $this->assertNotNull($pages_domain_ref01_data_dt0_load_result);
        $this->assertEquals($pages_domain_ref01_data_dt0_load_result["id"], $pages_domain_ref01_data["id"]);

        // REMOVE
        $pages_domain_ref01_match_rm0 = [
            "id" => $pages_domain_ref01_data["id"],
        ];
        $pages_domain_ref01_ent->remove($pages_domain_ref01_match_rm0, null);

        // LIST
        $pages_domain_ref01_match_rt0 = [
            "account_id" => $setup["idmap"]["account01"],
            "project_name" => $setup["idmap"]["project_name01"],
        ];

        $pages_domain_ref01_list_rt0_result = $pages_domain_ref01_ent->list($pages_domain_ref01_match_rt0, null);
        $this->assertIsArray($pages_domain_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($pages_domain_ref01_list_rt0_result),
            ["id" => $pages_domain_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function pages_domain_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/pages_domain/PagesDomainTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = CloudflarePagesSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["pages_domain01", "pages_domain02", "pages_domain03", "account01", "project01", "project_name01"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID" => $idmap,
        "CLOUDFLARE_PAGES_TEST_LIVE" => "FALSE",
        "CLOUDFLARE_PAGES_TEST_EXPLAIN" => "FALSE",
        "CLOUDFLARE_PAGES_APIKEY" => "",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }
    if (!isset($idmap_resolved["account_id"])) {
        $idmap_resolved["account_id"] = $idmap_resolved["account01"];
    }
    if (!isset($idmap_resolved["project_id"])) {
        $idmap_resolved["project_id"] = $idmap_resolved["project01"];
    }

    if ($env["CLOUDFLARE_PAGES_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            Runner::live_client_options(),
            [
                "apikey" => $env["CLOUDFLARE_PAGES_APIKEY"],
            ],
            // ismap, not a plain "?? []" default: an empty PHP array is a
            // LIST, and a non-map later entry REPLACES the accumulated map in
            // merge - so the no-extras call discarded live_client_options()
            // and the apikey/server map above it.
            Vs::ismap($extra) ? $extra : new \stdClass(),
        ]);
        // "?? []" because merge legitimately answers with a stdClass when every
        // contributing entry is an EMPTY map - an SDK with no apikey and no
        // server variables generates an empty middle entry, so that is the
        // common case, not the edge one. to_map returns null for a non-array by
        // design, and the constructor takes a non-nullable array, so without the
        // fallback every such SDK died on "must be of type array, null given"
        // the moment live mode was switched on. Offline mode never reaches this
        // branch, which is why the offline suite stayed green.
        $client = new CloudflarePagesSDK(Helpers::to_map($merged_opts) ?? []);
    }

    $live = $env["CLOUDFLARE_PAGES_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["CLOUDFLARE_PAGES_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
