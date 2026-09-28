# PagesDeployment entity test

import json
import os
import time

import pytest

from cloudflarepages_sdk.utility.voxgig_struct import voxgig_struct as vs
from cloudflarepages_sdk import CloudflarePagesSDK
from cloudflarepages_sdk.core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestPagesDeploymentEntity:

    def test_should_create_instance(self):
        testsdk = CloudflarePagesSDK.test(None, None)
        ent = testsdk.PagesDeployment(None)
        assert ent is not None

    def test_should_stream(self):
        # Feature #4: the entity stream(action, ...) method runs the op
        # pipeline and yields result items. With the streaming feature active
        # it yields the feature's incremental output; otherwise it falls back
        # to the materialised list so stream always yields.
        seed = {
            "entity": {
                "pages_deployment": {
                    "s1": {"id": "s1"},
                    "s2": {"id": "s2"},
                    "s3": {"id": "s3"},
                }
            }
        }

        # Fallback: streaming inactive -> yields the materialised list items.
        base = CloudflarePagesSDK.test(seed, None)
        seen = list(base.PagesDeployment(None).stream("list", None, None))
        assert len(seen) == 3

        # Inbound: streaming active -> yields each item from the feature.
        from cloudflarepages_sdk.config import shared_config
        cfg = shared_config()
        if isinstance(cfg.get("feature"), dict) and "streaming" in cfg["feature"]:
            sdk = CloudflarePagesSDK.test(
                seed, {"feature": {"streaming": {"active": True}}})
            got = []
            for item in sdk.PagesDeployment(None).stream("list", None, None):
                if isinstance(item, list):
                    got.extend(item)
                else:
                    got.append(item)
            assert len(got) == 3

    def test_should_run_basic_flow(self):
        setup = _pages_deployment_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "list", "load", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "pages_deployment." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        pages_deployment_ref01_ent = client.PagesDeployment(None)
        pages_deployment_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.pages_deployment"), "pages_deployment_ref01"))
        pages_deployment_ref01_data["account_id"] = setup["idmap"]["account01"]
        pages_deployment_ref01_data["deployment_id"] = setup["idmap"]["deployment01"]
        pages_deployment_ref01_data["project_id"] = setup["idmap"]["project01"]
        pages_deployment_ref01_data["project_name"] = setup["idmap"]["project_name01"]

        pages_deployment_ref01_data = helpers.to_map(runner.entity_data(pages_deployment_ref01_ent.create(pages_deployment_ref01_data, None)))
        assert pages_deployment_ref01_data is not None
        assert pages_deployment_ref01_data["id"] is not None

        # LIST
        pages_deployment_ref01_match = {
            "account_id": setup["idmap"]["account01"],
            "deployment_id": setup["idmap"]["deployment01"],
            "project_id": setup["idmap"]["project01"],
        }

        pages_deployment_ref01_list_result = pages_deployment_ref01_ent.list(pages_deployment_ref01_match, None)
        assert isinstance(pages_deployment_ref01_list_result, list)

        found_item = vs.select(
            runner.entity_list_to_data(pages_deployment_ref01_list_result),
            {"id": pages_deployment_ref01_data["id"]})
        assert not vs.isempty(found_item)

        # LOAD
        pages_deployment_ref01_match_dt0 = {
            "id": pages_deployment_ref01_data["id"],
        }
        pages_deployment_ref01_data_dt0_loaded = pages_deployment_ref01_ent.load(pages_deployment_ref01_match_dt0, None)
        pages_deployment_ref01_data_dt0_load_result = helpers.to_map(runner.entity_data(pages_deployment_ref01_data_dt0_loaded))
        assert pages_deployment_ref01_data_dt0_load_result is not None
        assert pages_deployment_ref01_data_dt0_load_result["id"] == pages_deployment_ref01_data["id"]

        # REMOVE
        pages_deployment_ref01_match_rm0 = {
            "id": pages_deployment_ref01_data["id"],
        }
        pages_deployment_ref01_ent.remove(pages_deployment_ref01_match_rm0, None)

        # LIST
        pages_deployment_ref01_match_rt0 = {
            "account_id": setup["idmap"]["account01"],
            "deployment_id": setup["idmap"]["deployment01"],
            "project_id": setup["idmap"]["project01"],
        }

        pages_deployment_ref01_list_rt0_result = pages_deployment_ref01_ent.list(pages_deployment_ref01_match_rt0, None)
        assert isinstance(pages_deployment_ref01_list_rt0_result, list)

        not_found_item = vs.select(
            runner.entity_list_to_data(pages_deployment_ref01_list_rt0_result),
            {"id": pages_deployment_ref01_data["id"]})
        assert vs.isempty(not_found_item)



def _pages_deployment_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/pages_deployment/PagesDeploymentTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = CloudflarePagesSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["pages_deployment01", "pages_deployment02", "pages_deployment03", "account01", "deployment01", "project01", "project_name01"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID": idmap,
        "CLOUDFLARE_PAGES_TEST_LIVE": "FALSE",
        "CLOUDFLARE_PAGES_TEST_EXPLAIN": "FALSE",
        "CLOUDFLARE_PAGES_APIKEY": "",
    })

    idmap_resolved = helpers.to_map(
        env.get("CLOUDFLARE_PAGES_TEST_PAGES_DEPLOYMENT_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("CLOUDFLARE_PAGES_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            # FIRST, so the generated fields below win: sdk-test-control.json's
            # test.client.options adds to the live client, it does not
            # redirect it.
            runner.live_client_options(),
            {
                "apikey": env.get("CLOUDFLARE_PAGES_APIKEY"),
            },
            extra or {},
        ])
        client = CloudflarePagesSDK(helpers.to_map(merged_opts))

    _live = env.get("CLOUDFLARE_PAGES_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("CLOUDFLARE_PAGES_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
