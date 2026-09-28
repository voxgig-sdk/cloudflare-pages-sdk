

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


describe('UploadEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_PAGES_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_PAGES_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflarePagesSDK.test()
    const ent = testsdk.Upload()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_PAGES_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'upload.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"errors":{"a":true,"h":"Errors","n":"errors","r":true,"t":"`$ARRAY`","key$":"errors","index$":0},"messages":{"a":true,"h":"Messages","n":"messages","r":true,"t":"`$ARRAY`","key$":"messages","index$":1},"success":{"a":true,"h":"Success","n":"success","r":true,"sh":"Whether the API call was successful.","t":"`$BOOLEAN`","key$":"success","index$":2}},"name":"upload","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /pages/assets/upload","source":"openapi3","version":2},"g":{},"k":"http","m":"POST","o":"/pages/assets/upload","q":{},"r":{},"s":[{"lit":"pages"},{"lit":"assets"},{"lit":"upload"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"upload","name__orig":"upload","Name":"Upload","name_":"upload","name-":"upload","NAME":"UPLOAD","index$":4}, {"active":true,"entity":"upload","key$":"BasicUploadFlow","kind":"basic","name":"BasicUploadFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"upload_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'Upload', {"POST /pages/assets/upload":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"example":[{"base64":true,"key":"b026324c6904b2a9cb4b88d6d61c81d1","metadata":{"contentType":"text/plain"},"value":"SGVsbG8sIFdvcmxkIQ=="}],"schema":{"type":"array","items":{"properties":{"base64":{"description":"Whether value is base64 encoded.","type":"boolean","example":true},"key":{"description":"File content hash used as the object key in the Pages asset store.","type":"string","example":"b026324c6904b2a9cb4b88d6d61c81d1"},"metadata":{"type":"object","properties":{"contentType":{"description":"MIME type for the uploaded file.","type":"string","example":"text/plain"}},"required":["contentType"]},"value":{"description":"File content. When base64 is true, this value is base64 encoded.","type":"string","example":"SGVsbG8sIFdvcmxkIQ=="}},"required":["key","value","metadata","base64"],"type":"object"},"x-ref":"#/components/schemas/pages_pages_assets_upload_request","index$":1}}}},"parameters":[]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const upload_ref01_ent = client.Upload()
    let upload_ref01_data = setup.data.new.upload['upload_ref01']

    upload_ref01_data = (await upload_ref01_ent.create(upload_ref01_data)).data()
    assert(null != upload_ref01_data)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/upload/UploadTestData.json')

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
    ['upload01','upload02','upload03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_PAGES_TEST_UPLOAD_ENTID': idmap,
    'CLOUDFLARE_PAGES_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_PAGES_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_PAGES_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_PAGES_TEST_UPLOAD_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_PAGES_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_PAGES_TEST_UPLOAD_ENTID']
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
  
