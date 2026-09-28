# CloudflarePages SDK utility: make_context
require_relative '../core/context'
module CloudflarePagesUtilities
  MakeContext = ->(ctxmap, basectx) {
    CloudflarePagesContext.new(ctxmap, basectx)
  }
end
