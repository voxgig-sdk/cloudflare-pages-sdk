# CloudflarePages SDK utility: make_context

from projectname_sdk.core.context import CloudflarePagesContext


def make_context_util(ctxmap, basectx):
    return CloudflarePagesContext(ctxmap, basectx)
