package sdktest

import (
	"encoding/json"
	"fmt"
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

func TestPagesProjectEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.PagesProject(nil)
		if ent == nil {
			t.Fatal("expected non-nil PagesProjectEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"pages_project": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.PagesProject(nil).Stream("list", nil, nil) {
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
			for item := range streamSdk.PagesProject(nil).Stream("list", nil, nil) {
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
		setup := pages_projectBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "pages_project." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_PAGES_TEST_PAGES_PROJECT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		pagesProjectRef01Ent := client.PagesProject(nil)
		pagesProjectRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "pages_project"}), "pages_project_ref01"))
		pagesProjectRef01Data["account_id"] = setup.idmap["account01"]
		pagesProjectRef01Data["project_name"] = setup.idmap["project_name01"]

		pagesProjectRef01DataResult, err := pagesProjectRef01Ent.Create(pagesProjectRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		pagesProjectRef01Data = core.ToMapAny(entityData(pagesProjectRef01DataResult))
		if pagesProjectRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if pagesProjectRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		pagesProjectRef01Match := map[string]any{
			"account_id": setup.idmap["account01"],
			"project_name": setup.idmap["project_name01"],
		}

		pagesProjectRef01ListResult, err := pagesProjectRef01Ent.List(pagesProjectRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pagesProjectRef01List, pagesProjectRef01ListOk := pagesProjectRef01ListResult.([]any)
		if !pagesProjectRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", pagesProjectRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(pagesProjectRef01List), map[string]any{"id": pagesProjectRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		pagesProjectRef01DataUp0Up := map[string]any{
			"id": pagesProjectRef01Data["id"],
			"account_id": setup.idmap["account_id"],
		}

		pagesProjectRef01MarkdefUp0Name := "name"
		pagesProjectRef01MarkdefUp0Value := fmt.Sprintf("Mark01-pages_project_ref01_%d", setup.now)
		pagesProjectRef01DataUp0Up[pagesProjectRef01MarkdefUp0Name] = pagesProjectRef01MarkdefUp0Value

		pagesProjectRef01ResdataUp0Result, err := pagesProjectRef01Ent.Update(pagesProjectRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		pagesProjectRef01ResdataUp0 := core.ToMapAny(entityData(pagesProjectRef01ResdataUp0Result))
		if pagesProjectRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if pagesProjectRef01ResdataUp0["id"] != pagesProjectRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if pagesProjectRef01ResdataUp0[pagesProjectRef01MarkdefUp0Name] != pagesProjectRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", pagesProjectRef01MarkdefUp0Name, pagesProjectRef01ResdataUp0[pagesProjectRef01MarkdefUp0Name])
		}

		// LOAD
		pagesProjectRef01MatchDt0 := map[string]any{
			"id": pagesProjectRef01Data["id"],
		}
		pagesProjectRef01DataDt0Loaded, err := pagesProjectRef01Ent.Load(pagesProjectRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		pagesProjectRef01DataDt0LoadResult := core.ToMapAny(entityData(pagesProjectRef01DataDt0Loaded))
		if pagesProjectRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if pagesProjectRef01DataDt0LoadResult["id"] != pagesProjectRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		pagesProjectRef01MatchRm0 := map[string]any{
			"id": pagesProjectRef01Data["id"],
		}
		_, err = pagesProjectRef01Ent.Remove(pagesProjectRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		pagesProjectRef01MatchRt0 := map[string]any{
			"account_id": setup.idmap["account01"],
			"project_name": setup.idmap["project_name01"],
		}

		pagesProjectRef01ListRt0Result, err := pagesProjectRef01Ent.List(pagesProjectRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pagesProjectRef01ListRt0, pagesProjectRef01ListRt0Ok := pagesProjectRef01ListRt0Result.([]any)
		if !pagesProjectRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", pagesProjectRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(pagesProjectRef01ListRt0), map[string]any{"id": pagesProjectRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func pages_projectBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "pages_project", "PagesProjectTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read pages_project test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse pages_project test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"pages_project01", "pages_project02", "pages_project03", "account01", "project_name01"},
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
	entidEnvRaw := os.Getenv("CLOUDFLARE_PAGES_TEST_PAGES_PROJECT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_PAGES_TEST_PAGES_PROJECT_ENTID": idmap,
		"CLOUDFLARE_PAGES_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_PAGES_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_PAGES_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_PAGES_TEST_PAGES_PROJECT_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}
	// Add account_id alias for update test.
	if idmapResolved["account_id"] == nil {
		idmapResolved["account_id"] = idmapResolved["account01"]
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
