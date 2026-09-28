package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/cloudflare-pages-sdk/go"
	"github.com/voxgig-sdk/cloudflare-pages-sdk/go/core"

	vs "github.com/voxgig-sdk/cloudflare-pages-sdk/go/utility/struct"
)

func TestPagesPagesAssetsCheckMissingEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.PagesPagesAssetsCheckMissing(nil)
		if ent == nil {
			t.Fatal("expected non-nil PagesPagesAssetsCheckMissingEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := pages_pages_assets_check_missingBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "pages_pages_assets_check_missing." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		pagesPagesAssetsCheckMissingRef01Ent := client.PagesPagesAssetsCheckMissing(nil)
		pagesPagesAssetsCheckMissingRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "pages_pages_assets_check_missing"}), "pages_pages_assets_check_missing_ref01"))

		pagesPagesAssetsCheckMissingRef01DataResult, err := pagesPagesAssetsCheckMissingRef01Ent.Create(pagesPagesAssetsCheckMissingRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		pagesPagesAssetsCheckMissingRef01Data = core.ToMapAny(entityData(pagesPagesAssetsCheckMissingRef01DataResult))
		if pagesPagesAssetsCheckMissingRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func pages_pages_assets_check_missingBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "pages_pages_assets_check_missing", "PagesPagesAssetsCheckMissingTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read pages_pages_assets_check_missing test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse pages_pages_assets_check_missing test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"pages_pages_assets_check_missing01", "pages_pages_assets_check_missing02", "pages_pages_assets_check_missing03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID": idmap,
		"CLOUDFLARE_PAGES_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_PAGES_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_PAGES_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["CLOUDFLARE_PAGES_TEST_LIVE"] == "TRUE" {
		// An empty map, not a nil one: Merge returns nil when its last entry
		// is nil, and BasicSetup is normally called with no extras - so a
		// bare nil silently discarded the apikey and server values below.
		extraOpts := extra
		if extraOpts == nil {
			extraOpts = map[string]any{}
		}

		mergedOpts := vs.Merge([]any{
			// liveClientOptions() FIRST, so the generated fields below win:
			// sdk-test-control.json's test.client.options adds to the live
			// client, it does not redirect it.
			liveClientOptions(),
			map[string]any{
				"apikey": env["CLOUDFLARE_PAGES_APIKEY"],
			},
			extraOpts,
		})
		client = sdk.NewCloudflarePagesSDK(core.ToMapAny(mergedOpts))
	}

	live := env["CLOUDFLARE_PAGES_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["CLOUDFLARE_PAGES_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
