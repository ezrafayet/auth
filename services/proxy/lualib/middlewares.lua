-- access.lua

local _M = {}

-- sets headers that must be passed to the upstream services
function _M.set_headers()
    ngx.req.set_header("X-Forwarded-For", ngx.var.remote_addr)
    ngx.req.set_header("X-Forwarded-Proto", ngx.var.scheme)
    ngx.req.set_header("X-Content-Type-Options", "nosniff")
    ngx.req.set_header("X-Frame-Options", "DENY")
    ngx.req.set_header("X-XSS-Protection", "1; mode=block")
    ngx.req.set_header("X-Request-Id", ngx.var.request_id)
end

-- checks if the request is authorized.
-- it only checks if the authorization header is valid
-- it is not resource-specific
function _M.control_access_token()
    local res = ngx.location.capture("/api/internal/v1/auth/token/authorize", { method = ngx.HTTP_POST })

    if not res then
        ngx.say("{\"status\": \"error\", \"message\": \"Internal server error\"}")
        ngx.exit(ngx.HTTP_INTERNAL_SERVER_ERROR)
        return
    end

    ngx.status = res.status
    ngx.say(res.body)
    ngx.exit(res.status)
end

return _M