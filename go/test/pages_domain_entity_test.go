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

func TestPagesDomainEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.PagesDomain(nil)
		if ent == nil {
			t.Fatal("expected non-nil PagesDomainEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"pages_domain": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.PagesDomain(nil).Stream("list", nil, nil) {
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
			for item := range streamSdk.PagesDomain(nil).Stream("list", nil, nil) {
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
		setup := pages_domainBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "pages_domain." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		pagesDomainRef01Ent := client.PagesDomain(nil)
		pagesDomainRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "pages_domain"}), "pages_domain_ref01"))
		pagesDomainRef01Data["account_id"] = setup.idmap["account01"]
		pagesDomainRef01Data["project_id"] = setup.idmap["project01"]
		pagesDomainRef01Data["project_name"] = setup.idmap["project_name01"]

		pagesDomainRef01DataResult, err := pagesDomainRef01Ent.Create(pagesDomainRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		pagesDomainRef01Data = core.ToMapAny(entityData(pagesDomainRef01DataResult))
		if pagesDomainRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if pagesDomainRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		pagesDomainRef01Match := map[string]any{
			"account_id": setup.idmap["account01"],
			"project_name": setup.idmap["project_name01"],
		}

		pagesDomainRef01ListResult, err := pagesDomainRef01Ent.List(pagesDomainRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pagesDomainRef01List, pagesDomainRef01ListOk := pagesDomainRef01ListResult.([]any)
		if !pagesDomainRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", pagesDomainRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(pagesDomainRef01List), map[string]any{"id": pagesDomainRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		pagesDomainRef01DataUp0Up := map[string]any{
			"id": pagesDomainRef01Data["id"],
			"account_id": setup.idmap["account_id"],
			"project_id": setup.idmap["project_id"],
		}

		pagesDomainRef01MarkdefUp0Name := "name"
		pagesDomainRef01MarkdefUp0Value := fmt.Sprintf("Mark01-pages_domain_ref01_%d", setup.now)
		pagesDomainRef01DataUp0Up[pagesDomainRef01MarkdefUp0Name] = pagesDomainRef01MarkdefUp0Value

		pagesDomainRef01ResdataUp0Result, err := pagesDomainRef01Ent.Update(pagesDomainRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		pagesDomainRef01ResdataUp0 := core.ToMapAny(entityData(pagesDomainRef01ResdataUp0Result))
		if pagesDomainRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if pagesDomainRef01ResdataUp0["id"] != pagesDomainRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if pagesDomainRef01ResdataUp0[pagesDomainRef01MarkdefUp0Name] != pagesDomainRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", pagesDomainRef01MarkdefUp0Name, pagesDomainRef01ResdataUp0[pagesDomainRef01MarkdefUp0Name])
		}

		// LOAD
		pagesDomainRef01MatchDt0 := map[string]any{
			"id": pagesDomainRef01Data["id"],
		}
		pagesDomainRef01DataDt0Loaded, err := pagesDomainRef01Ent.Load(pagesDomainRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		pagesDomainRef01DataDt0LoadResult := core.ToMapAny(entityData(pagesDomainRef01DataDt0Loaded))
		if pagesDomainRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if pagesDomainRef01DataDt0LoadResult["id"] != pagesDomainRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		pagesDomainRef01MatchRm0 := map[string]any{
			"id": pagesDomainRef01Data["id"],
		}
		_, err = pagesDomainRef01Ent.Remove(pagesDomainRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		pagesDomainRef01MatchRt0 := map[string]any{
			"account_id": setup.idmap["account01"],
			"project_name": setup.idmap["project_name01"],
		}

		pagesDomainRef01ListRt0Result, err := pagesDomainRef01Ent.List(pagesDomainRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pagesDomainRef01ListRt0, pagesDomainRef01ListRt0Ok := pagesDomainRef01ListRt0Result.([]any)
		if !pagesDomainRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", pagesDomainRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(pagesDomainRef01ListRt0), map[string]any{"id": pagesDomainRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func pages_domainBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "pages_domain", "PagesDomainTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read pages_domain test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse pages_domain test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"pages_domain01", "pages_domain02", "pages_domain03", "account01", "project01", "project_name01"},
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
	entidEnvRaw := os.Getenv("CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID": idmap,
		"CLOUDFLARE_PAGES_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_PAGES_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_PAGES_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}
	// Add account_id alias for update test.
	if idmapResolved["account_id"] == nil {
		idmapResolved["account_id"] = idmapResolved["account01"]
	}
	// Add project_id alias for update test.
	if idmapResolved["project_id"] == nil {
		idmapResolved["project_id"] = idmapResolved["project01"]
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
