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

func TestPagesDeploymentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.PagesDeployment(nil)
		if ent == nil {
			t.Fatal("expected non-nil PagesDeploymentEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"pages_deployment": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.PagesDeployment(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.SharedConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.PagesDeployment(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := pages_deploymentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "pages_deployment." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		pagesDeploymentRef01Ent := client.PagesDeployment(nil)
		pagesDeploymentRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "pages_deployment"}), "pages_deployment_ref01"))
		pagesDeploymentRef01Data["account_id"] = setup.idmap["account01"]
		pagesDeploymentRef01Data["deployment_id"] = setup.idmap["deployment01"]
		pagesDeploymentRef01Data["project_id"] = setup.idmap["project01"]
		pagesDeploymentRef01Data["project_name"] = setup.idmap["project_name01"]

		pagesDeploymentRef01DataResult, err := pagesDeploymentRef01Ent.Create(pagesDeploymentRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		pagesDeploymentRef01Data = core.ToMapAny(entityData(pagesDeploymentRef01DataResult))
		if pagesDeploymentRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if pagesDeploymentRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		pagesDeploymentRef01Match := map[string]any{
			"account_id": setup.idmap["account01"],
			"deployment_id": setup.idmap["deployment01"],
			"project_id": setup.idmap["project01"],
		}

		pagesDeploymentRef01ListResult, err := pagesDeploymentRef01Ent.List(pagesDeploymentRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pagesDeploymentRef01List, pagesDeploymentRef01ListOk := pagesDeploymentRef01ListResult.([]any)
		if !pagesDeploymentRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", pagesDeploymentRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(pagesDeploymentRef01List), map[string]any{"id": pagesDeploymentRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// LOAD
		pagesDeploymentRef01MatchDt0 := map[string]any{
			"id": pagesDeploymentRef01Data["id"],
		}
		pagesDeploymentRef01DataDt0Loaded, err := pagesDeploymentRef01Ent.Load(pagesDeploymentRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		pagesDeploymentRef01DataDt0LoadResult := core.ToMapAny(entityData(pagesDeploymentRef01DataDt0Loaded))
		if pagesDeploymentRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if pagesDeploymentRef01DataDt0LoadResult["id"] != pagesDeploymentRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		pagesDeploymentRef01MatchRm0 := map[string]any{
			"id": pagesDeploymentRef01Data["id"],
		}
		_, err = pagesDeploymentRef01Ent.Remove(pagesDeploymentRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		pagesDeploymentRef01MatchRt0 := map[string]any{
			"account_id": setup.idmap["account01"],
			"deployment_id": setup.idmap["deployment01"],
			"project_id": setup.idmap["project01"],
		}

		pagesDeploymentRef01ListRt0Result, err := pagesDeploymentRef01Ent.List(pagesDeploymentRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pagesDeploymentRef01ListRt0, pagesDeploymentRef01ListRt0Ok := pagesDeploymentRef01ListRt0Result.([]any)
		if !pagesDeploymentRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", pagesDeploymentRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(pagesDeploymentRef01ListRt0), map[string]any{"id": pagesDeploymentRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func pages_deploymentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "pages_deployment", "PagesDeploymentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read pages_deployment test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse pages_deployment test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"pages_deployment01", "pages_deployment02", "pages_deployment03", "account01", "deployment01", "project01", "project_name01"},
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
	entidEnvRaw := os.Getenv("CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID": idmap,
		"CLOUDFLARE_PAGES_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_PAGES_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_PAGES_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID"])
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
