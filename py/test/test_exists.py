# CloudflarePages SDK exists test

import pytest
from cloudflarepages_sdk import CloudflarePagesSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = CloudflarePagesSDK.test(None, None)
        assert testsdk is not None
