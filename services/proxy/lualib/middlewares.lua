-- access.lua

local _M = {}

function set_headers_forwarded_for()
    ngx.req.set_header("X-Forwarded-For", ngx.var.remote_addr)
end

function set_headers_forwarded_proto()
    ngx.req.set_header("X-Forwarded-Proto", ngx.var.scheme)
end

function set_headers_content_type_options()
    ngx.req.set_header("X-Content-Type-Options", "nosniff")
end

function set_headers_frame_options()
    ngx.req.set_header("X-Frame-Options", "DENY")
end

function set_headers_xss_protection()
    ngx.req.set_header("X-XSS-Protection", "1; mode=block")
end

function set_headers_request_id ()
    ngx.req.set_header("X-Request-Id", ngx.var.request_id)
end

-- sets headers that must be passed to the upstream services
function _M.set_headers()
    set_headers_forwarded_for()
    set_headers_forwarded_proto()
    set_headers_content_type_options()
    set_headers_frame_options()
    set_headers_xss_protection()
    set_headers_request_id()
end

-- todo: must be testable and tested
function _M.control_access_token()
    local res = ngx.location.capture("/api/internal/v1/auth/token/authorize", { method = ngx.HTTP_POST })

    if not res then
        ngx.say("{\"status\": \"error\", \"message\": \"Internal server error\"}")
        ngx.exit(ngx.HTTP_INTERNAL_SERVER_ERROR)
        return
    end

    if res.status < 200 or res.status >= 400 then
        ngx.status = res.status
        ngx.say(res.body)
        ngx.exit(res.status)
        return
    end

    return ngx.OK
end

return _M