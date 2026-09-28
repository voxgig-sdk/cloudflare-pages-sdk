# CloudflarePages SDK utility: make_error

from __future__ import annotations
from cloudflarepages_sdk.core.operation import CloudflarePagesOperation
from cloudflarepages_sdk.core.result import CloudflarePagesResult
from cloudflarepages_sdk.core.control import CloudflarePagesControl
from cloudflarepages_sdk.core.error import CloudflarePagesError


def make_error_util(ctx, err):
    if ctx is None:
        from cloudflarepages_sdk.core.context import CloudflarePagesContext
        ctx = CloudflarePagesContext({}, None)

    op = ctx.op
    if op is None:
        op = CloudflarePagesOperation({})
    opname = op.name
    if opname == "" or opname == "_":
        opname = "unknown operation"

    result = ctx.result
    if result is None:
        result = CloudflarePagesResult({})
    result.ok = False

    if err is None:
        err = result.err
    if err is None:
        err = ctx.make_error("unknown", "unknown error")

    errmsg = ""
    if isinstance(err, CloudflarePagesError):
        errmsg = err.msg
    elif hasattr(err, "msg") and err.msg is not None:
        errmsg = err.msg
    elif isinstance(err, str):
        errmsg = err
    else:
        errmsg = str(err)

    msg = "CloudflarePagesSDK: " + opname + ": " + errmsg
    msg = ctx.utility.clean(ctx, msg)

    result.err = None

    spec = ctx.spec

    if ctx.ctrl.explain is not None:
        ctx.ctrl.explain["err"] = {"message": msg}

    sdk_err = CloudflarePagesError("", msg, ctx)
    sdk_err.result = ctx.utility.clean(ctx, result)
    sdk_err.spec = ctx.utility.clean(ctx, spec)

    # Promote the HTTP status to the top level, so a consumer can branch on
    # `err.status` / `err.not_found` instead of reaching into `err.result`.
    sdk_err.status = -1 if result.status is None else result.status

    if isinstance(err, CloudflarePagesError):
        sdk_err.code = err.code

    ctx.ctrl.err = sdk_err

    # Fire PreUnexpected so observability features (metrics, telemetry, audit,
    # debug) close/record error paths that never reach PreDone (e.g. a PrePoint
    # rbac short-circuit). Fires after ctx.ctrl.err is set so hooks can read the
    # error; features guard against double-recording when PreDone already fired.
    if getattr(ctx, "utility", None) is not None and \
            callable(getattr(ctx.utility, "feature_hook", None)):
        ctx.utility.feature_hook(ctx, "PreUnexpected")

    if ctx.ctrl.throw_err is False:
        return result.resdata

    raise sdk_err
