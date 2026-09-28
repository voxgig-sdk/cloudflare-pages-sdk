

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { CloudflarePagesSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('PagesPagesAssetsCheckMissingEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_PAGES_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_PAGES_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflarePagesSDK.test()
    const ent = testsdk.PagesPagesAssetsCheckMissing()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_PAGES_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'pages_pages_assets_check_missing.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"hashes":{"a":true,"h":"Hashes","n":"hashes","r":true,"sh":"List of file content hashes to check for existence in the asset store.","t":"`$ARRAY`","key$":"hashes","index$":0}},"name":"pages_pages_assets_check_missing","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /pages/assets/check-missing","source":"openapi3","version":2},"g":{},"k":"http","m":"POST","o":"/pages/assets/check-missing","q":{},"r":{},"s":[{"lit":"pages"},{"lit":"assets"},{"lit":"check-missing"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"pages_pages_assets_check_missing","name__orig":"pages_pages_assets_check_missing","Name":"PagesPagesAssetsCheckMissing","name_":"pages_pages_assets_check_missing","name-":"pages-pages-assets-check-missing","NAME":"PAGES_PAGES_ASSETS_CHECK_MISSING","index$":2}, {"active":true,"entity":"pages_pages_assets_check_missing","key$":"BasicPagesPagesAssetsCheckMissingFlow","kind":"basic","name":"BasicPagesPagesAssetsCheckMissingFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"pages_pages_assets_check_missing_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'PagesPagesAssetsCheckMissing', {"POST /pages/assets/check-missing":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"example":{"hashes":["a948904f2f0f479b8f936b8a0c5d9882","b026324c6904b2a9cb4b88d6d61c81d1"]},"schema":{"type":"object","properties":{"hashes":{"description":"List of file content hashes to check for existence in the asset store.","type":"array","items":{"type":"string"},"example":["a948904f2f0f479b8f936b8a0c5d9882","b026324c6904b2a9cb4b88d6d61c81d1"],"key$":"hashes"}},"required":["hashes"],"x-ref":"#/components/schemas/pages_pages_assets_check_missing_request","index$":1}}}},"parameters":[]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const pages_pages_assets_check_missing_ref01_ent = client.PagesPagesAssetsCheckMissing()
    let pages_pages_assets_check_missing_ref01_data = setup.data.new.pages_pages_assets_check_missing['pages_pages_assets_check_missing_ref01']

    pages_pages_assets_check_missing_ref01_data = (await pages_pages_assets_check_missing_ref01_ent.create(pages_pages_assets_check_missing_ref01_data)).data()
    assert(null != pages_pages_assets_check_missing_ref01_data)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/pages_pages_assets_check_missing/PagesPagesAssetsCheckMissingTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = CloudflarePagesSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['pages_pages_assets_check_missing01','pages_pages_assets_check_missing02','pages_pages_assets_check_missing03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID': idmap,
    'CLOUDFLARE_PAGES_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_PAGES_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_PAGES_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_PAGES_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_PAGES_TEST_PAGES_PAGES_ASSETS_CHECK_MISSING_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new CloudflarePagesSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
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
    ]))
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
  }

  return setup
}
  
