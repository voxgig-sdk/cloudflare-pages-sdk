# CloudflarePages SDK utility: make_context

from cloudflarepages_sdk.core.context import CloudflarePagesContext


def make_context_util(ctxmap, basectx):
    return CloudflarePagesContext(ctxmap, basectx)
