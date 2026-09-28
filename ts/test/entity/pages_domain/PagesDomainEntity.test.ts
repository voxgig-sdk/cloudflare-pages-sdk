

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


describe('PagesDomainEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_PAGES_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_PAGES_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflarePagesSDK.test()
    const ent = testsdk.PagesDomain()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_PAGES_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'pages_domain.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"errors":{"a":true,"h":"Errors","n":"errors","r":true,"t":"`$ARRAY`","key$":"errors","index$":0},"id":{"a":true,"h":"Id","n":"id","r":false,"t":"`$STRING`","key$":"id","index$":1},"messages":{"a":true,"h":"Messages","n":"messages","r":true,"t":"`$ARRAY`","key$":"messages","index$":2},"name":{"a":true,"h":"Name","n":"name","r":true,"sh":"Fully qualified domain name for the Pages project, such as `example.com`.","t":"`$STRING`","key$":"name","index$":3},"result":{"a":true,"h":"Result","n":"result","r":true,"t":"`$OBJECT`","key$":"result","index$":4},"success":{"a":true,"h":"Success","n":"success","r":true,"sh":"Whether the API call was successful.","t":"`$BOOLEAN`","key$":"success","index$":5}},"id":{"field":"id","name":"id"},"name":"pages_domain","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /accounts/{account_id}/pages/projects/{project_name}/domains","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"this-is-my-project-01","k":"param","n":"project_name","or":"project_name","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"POST","o":"/accounts/{account_id}/pages/projects/{project_name}/domains","q":{"exist":["account_id","project_name"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"pages"},{"lit":"projects"},{"var":"project_name"},{"lit":"domains"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /accounts/{account_id}/pages/projects/{project_name}/domains","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"this-is-my-project-01","k":"param","n":"project_name","or":"project_name","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/accounts/{account_id}/pages/projects/{project_name}/domains","q":{"exist":["account_id","project_name"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"pages"},{"lit":"projects"},{"var":"project_name"},{"lit":"domains"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"example.com","k":"param","n":"id","or":"domain_name","r":true,"t":"`$STRING`","index$":1},{"a":true,"ex":"this-is-my-project-01","k":"param","n":"project_id","or":"project_name","r":true,"t":"`$STRING`","index$":2}]},"k":"http","m":"GET","o":"/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}","q":{"exist":["account_id","id","project_id"]},"r":{"param":{"domain_name":"id","project_name":"project_id"}},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"pages"},{"lit":"projects"},{"var":"project_id"},{"lit":"domains"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"example.com","k":"param","n":"id","or":"domain_name","r":true,"t":"`$STRING`","index$":1},{"a":true,"ex":"this-is-my-project-01","k":"param","n":"project_id","or":"project_name","r":true,"t":"`$STRING`","index$":2}]},"k":"http","m":"DELETE","o":"/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}","q":{"exist":["account_id","id","project_id"]},"r":{"param":{"domain_name":"id","project_name":"project_id"}},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"pages"},{"lit":"projects"},{"var":"project_id"},{"lit":"domains"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"a":true,"co":{"id":"PATCH /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"example.com","k":"param","n":"id","or":"domain_name","r":true,"t":"`$STRING`","index$":1},{"a":true,"ex":"this-is-my-project-01","k":"param","n":"project_id","or":"project_name","r":true,"t":"`$STRING`","index$":2}]},"k":"http","m":"PATCH","o":"/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}","q":{"exist":["account_id","id","project_id"]},"r":{"param":{"domain_name":"id","project_name":"project_id"}},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"pages"},{"lit":"projects"},{"var":"project_id"},{"lit":"domains"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[]},"key$":"pages_domain","name__orig":"pages_domain","Name":"PagesDomain","name_":"pages_domain","name-":"pages-domain","NAME":"PAGES_DOMAIN","index$":1}, {"active":true,"entity":"pages_domain","key$":"BasicPagesDomainFlow","kind":"basic","name":"BasicPagesDomainFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"pages_domain_ref01"},"m":{"account_id":"account01","project_id":"project01","project_name":"project_name01"},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{},"m":{"account_id":"account01","project_name":"project_name01"},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"pages_domain_ref01"}}],"index$":1},{"a":true,"d":{"account_id":"account01","project_id":"project01"},"i":{"ref":"pages_domain_ref01","srcdatavar":"pages_domain_ref01_data","suffix":"_up0","textfield":"name"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-pages_domain_ref01"}}],"v":[],"index$":2},{"a":true,"d":{},"i":{"ref":"pages_domain_ref01","srcdatavar":"pages_domain_ref01_data","suffix":"_dt0"},"m":{"account_id":"account01","id":"pages_domain01","project_id":"project01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-pages_domain_ref01"}}],"index$":3},{"a":true,"d":{},"i":{"ref":"pages_domain_ref01","suffix":"_rm0"},"m":{"account_id":"account01","id":"pages_domain01","project_id":"project01"},"o":"remove","s":[],"v":[],"index$":4},{"a":true,"d":{},"i":{"suffix":"_rt0"},"m":{"account_id":"account01","project_name":"project_name01"},"o":"list","s":[],"v":[{"apply":"ItemNotExists","def":{"ref":"pages_domain_ref01"}}],"index$":5}]}, 'PagesDomain', {"POST /accounts/{account_id}/pages/projects/{project_name}/domains":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","properties":{"name":{"description":"Fully qualified domain name for the Pages project, such as `example.com`.","type":"string","example":"example.com","pattern":"^([a-zA-Z0-9][\\-a-zA-Z0-9]*\\.)+[\\-a-zA-Z0-9]{2,20}$","x-auditable":true,"x-ref":"#/components/schemas/pages_domain_name","key$":"name"}},"required":["name"],"index$":1}}}},"parameters":[{"name":"project_name","in":"path","required":true,"schema":{"description":"Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.","type":"string","example":"this-is-my-project-01","pattern":"^[a-z0-9][a-z0-9-]*$","x-auditable":true,"x-ref":"#/components/schemas/pages_project_name"},"index$":0},{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/pages_identifier"},"index$":1}]},"GET /accounts/{account_id}/pages/projects/{project_name}/domains":{"protocol":"http","parameters":[{"name":"project_name","in":"path","required":true,"schema":{"description":"Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.","type":"string","example":"this-is-my-project-01","pattern":"^[a-z0-9][a-z0-9-]*$","x-auditable":true,"x-ref":"#/components/schemas/pages_project_name"},"index$":0},{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/pages_identifier"},"index$":1}]},"GET /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}":{"protocol":"http","parameters":[{"name":"domain_name","in":"path","required":true,"schema":{"description":"Fully qualified domain name for the Pages project, such as `example.com`.","type":"string","example":"example.com","pattern":"^([a-zA-Z0-9][\\-a-zA-Z0-9]*\\.)+[\\-a-zA-Z0-9]{2,20}$","x-auditable":true,"x-ref":"#/components/schemas/pages_domain_name"},"index$":0},{"name":"project_name","in":"path","required":true,"schema":{"description":"Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.","type":"string","example":"this-is-my-project-01","pattern":"^[a-z0-9][a-z0-9-]*$","x-auditable":true,"x-ref":"#/components/schemas/pages_project_name"},"index$":1},{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/pages_identifier"},"index$":2}]},"DELETE /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}":{"protocol":"http","parameters":[{"name":"domain_name","in":"path","required":true,"schema":{"description":"Fully qualified domain name for the Pages project, such as `example.com`.","type":"string","example":"example.com","pattern":"^([a-zA-Z0-9][\\-a-zA-Z0-9]*\\.)+[\\-a-zA-Z0-9]{2,20}$","x-auditable":true,"x-ref":"#/components/schemas/pages_domain_name"},"index$":0},{"name":"project_name","in":"path","required":true,"schema":{"description":"Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.","type":"string","example":"this-is-my-project-01","pattern":"^[a-z0-9][a-z0-9-]*$","x-auditable":true,"x-ref":"#/components/schemas/pages_project_name"},"index$":1},{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/pages_identifier"},"index$":2}]},"PATCH /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}":{"protocol":"http","parameters":[{"name":"domain_name","in":"path","required":true,"schema":{"description":"Fully qualified domain name for the Pages project, such as `example.com`.","type":"string","example":"example.com","pattern":"^([a-zA-Z0-9][\\-a-zA-Z0-9]*\\.)+[\\-a-zA-Z0-9]{2,20}$","x-auditable":true,"x-ref":"#/components/schemas/pages_domain_name"},"index$":0},{"name":"project_name","in":"path","required":true,"schema":{"description":"Name of the Pages project. Must begin with a lowercase letter or digit and contain only lowercase letters, digits, and hyphens.","type":"string","example":"this-is-my-project-01","pattern":"^[a-z0-9][a-z0-9-]*$","x-auditable":true,"x-ref":"#/components/schemas/pages_project_name"},"index$":1},{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/pages_identifier"},"index$":2}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const pages_domain_ref01_ent = client.PagesDomain()
    let pages_domain_ref01_data = setup.data.new.pages_domain['pages_domain_ref01']
    pages_domain_ref01_data['account_id'] = setup.idmap['account01']
    pages_domain_ref01_data['project_id'] = setup.idmap['project01']
    pages_domain_ref01_data['project_name'] = setup.idmap['project_name01']

    pages_domain_ref01_data = (await pages_domain_ref01_ent.create(pages_domain_ref01_data)).data()
    assert(null != pages_domain_ref01_data.id)


    // LIST
    const pages_domain_ref01_match: any = {}
    pages_domain_ref01_match['account_id'] = setup.idmap['account01']
    pages_domain_ref01_match['project_name'] = setup.idmap['project_name01']

    const pages_domain_ref01_list = (await pages_domain_ref01_ent.list(pages_domain_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(pages_domain_ref01_list, { id: pages_domain_ref01_data.id })))


    // UPDATE
    const pages_domain_ref01_data_up0: any = {}
    pages_domain_ref01_data_up0.id = pages_domain_ref01_data.id
    pages_domain_ref01_data_up0 ['account_id'] = setup.idmap['account_id']
    pages_domain_ref01_data_up0 ['project_id'] = setup.idmap['project_id']

    const pages_domain_ref01_markdef_up0 = { name: 'name', value: 'Mark01-pages_domain_ref01_' + setup.now }
    ;(pages_domain_ref01_data_up0 as any)[pages_domain_ref01_markdef_up0.name] = pages_domain_ref01_markdef_up0.value

    const pages_domain_ref01_resdata_up0 = (await pages_domain_ref01_ent.update(pages_domain_ref01_data_up0)).data()
    assert(pages_domain_ref01_resdata_up0.id === pages_domain_ref01_data_up0.id)

    assert((pages_domain_ref01_resdata_up0 as any)[pages_domain_ref01_markdef_up0.name] === pages_domain_ref01_markdef_up0.value)


    // LOAD
    const pages_domain_ref01_match_dt0: any = {}
    pages_domain_ref01_match_dt0.id = pages_domain_ref01_data.id
    const pages_domain_ref01_data_dt0 = (await pages_domain_ref01_ent.load(pages_domain_ref01_match_dt0)).data()
    assert(pages_domain_ref01_data_dt0.id === pages_domain_ref01_data.id)


    // REMOVE
    const pages_domain_ref01_match_rm0: any = { id: pages_domain_ref01_data.id }
    await pages_domain_ref01_ent.remove(pages_domain_ref01_match_rm0)
  

    // LIST
    const pages_domain_ref01_match_rt0: any = {}
    pages_domain_ref01_match_rt0['account_id'] = setup.idmap['account01']
    pages_domain_ref01_match_rt0['project_name'] = setup.idmap['project_name01']

    const pages_domain_ref01_list_rt0 = (await pages_domain_ref01_ent.list(pages_domain_ref01_match_rt0)).map((e: any) => e.data())

    assert(isempty(select(pages_domain_ref01_list_rt0, { id: pages_domain_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/pages_domain/PagesDomainTestData.json')

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
    ['pages_domain01','pages_domain02','pages_domain03','account01','project01','project_name01'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID': idmap,
    'CLOUDFLARE_PAGES_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_PAGES_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_PAGES_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_PAGES_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_PAGES_TEST_PAGES_DOMAIN_ENTID']
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
  
