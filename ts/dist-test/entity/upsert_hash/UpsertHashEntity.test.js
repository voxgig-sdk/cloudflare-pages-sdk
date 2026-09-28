"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('UpsertHashEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when CLOUDFLARE_PAGES_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('CLOUDFLARE_PAGES_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.CloudflarePagesSDK.test();
        const ent = testsdk.UpsertHash();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.CLOUDFLARE_PAGES_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'upsert_hash.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "errors": { "a": true, "h": "Errors", "n": "errors", "r": true, "t": "`$ARRAY`", "key$": "errors", "index$": 0 }, "hashes": { "a": true, "h": "Hashes", "n": "hashes", "r": true, "sh": "List of file content hashes to register in the asset store.", "t": "`$ARRAY`", "key$": "hashes", "index$": 1 }, "messages": { "a": true, "h": "Messages", "n": "messages", "r": true, "t": "`$ARRAY`", "key$": "messages", "index$": 2 }, "success": { "a": true, "h": "Success", "n": "success", "r": true, "sh": "Whether the API call was successful.", "t": "`$BOOLEAN`", "key$": "success", "index$": 3 } }, "name": "upsert_hash", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /pages/assets/upsert-hashes", "source": "openapi3", "version": 2 }, "g": {}, "k": "http", "m": "POST", "o": "/pages/assets/upsert-hashes", "q": {}, "r": {}, "s": [{ "lit": "pages" }, { "lit": "assets" }, { "lit": "upsert-hashes" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "upsert_hash", "name__orig": "upsert_hash", "Name": "UpsertHash", "name_": "upsert_hash", "name-": "upsert-hash", "NAME": "UPSERT_HASH", "index$": 5 }, { "active": true, "entity": "upsert_hash", "key$": "BasicUpsertHashFlow", "kind": "basic", "name": "BasicUpsertHashFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "upsert_hash_ref01" }, "m": {}, "o": "create", "s": [], "v": [], "index$": 0 }] }, 'UpsertHash', { "POST /pages/assets/upsert-hashes": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "example": { "hashes": ["a948904f2f0f479b8f936b8a0c5d9882", "b026324c6904b2a9cb4b88d6d61c81d1"] }, "schema": { "type": "object", "properties": { "hashes": { "description": "List of file content hashes to register in the asset store.", "type": "array", "items": { "type": "string" }, "example": ["a948904f2f0f479b8f936b8a0c5d9882", "b026324c6904b2a9cb4b88d6d61c81d1"], "key$": "hashes" } }, "required": ["hashes"], "x-ref": "#/components/schemas/pages_pages_assets_upsert_hashes_request", "index$": 1 } } } }, "parameters": [] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const upsert_hash_ref01_ent = client.UpsertHash();
        let upsert_hash_ref01_data = setup.data.new.upsert_hash['upsert_hash_ref01'];
        upsert_hash_ref01_data = (await upsert_hash_ref01_ent.create(upsert_hash_ref01_data)).data();
        (0, node_assert_1.default)(null != upsert_hash_ref01_data);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/upsert_hash/UpsertHashTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.CloudflarePagesSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['upsert_hash01', 'upsert_hash02', 'upsert_hash03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'CLOUDFLARE_PAGES_TEST_UPSERT_HASH_ENTID': idmap,
        'CLOUDFLARE_PAGES_TEST_LIVE': 'FALSE',
        'CLOUDFLARE_PAGES_TEST_EXPLAIN': 'FALSE',
        'CLOUDFLARE_PAGES_APIKEY': '',
    });
    idmap = env['CLOUDFLARE_PAGES_TEST_UPSERT_HASH_ENTID'];
    const live = 'TRUE' === env.CLOUDFLARE_PAGES_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['CLOUDFLARE_PAGES_TEST_UPSERT_HASH_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.CloudflarePagesSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.CLOUDFLARE_PAGES_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.CLOUDFLARE_PAGES_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=UpsertHashEntity.test.js.map