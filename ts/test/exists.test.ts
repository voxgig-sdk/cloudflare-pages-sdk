
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { CloudflarePagesSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = CloudflarePagesSDK.test()
    equal(testsdk instanceof CloudflarePagesSDK, true,
      'CloudflarePagesSDK.test() must return a client synchronously')
  })

})
