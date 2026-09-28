-- CloudflarePages SDK error

local CloudflarePagesError = {}
CloudflarePagesError.__index = CloudflarePagesError


function CloudflarePagesError.new(code, msg, ctx)
  local self = setmetatable({}, CloudflarePagesError)
  self.is_sdk_error = true
  self.sdk = "CloudflarePages"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function CloudflarePagesError:error()
  return self.msg
end


function CloudflarePagesError:__tostring()
  return self.msg
end


return CloudflarePagesError
