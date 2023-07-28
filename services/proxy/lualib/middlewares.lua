-- access.lua

local _M = {}

-- This function sets several HTTP headers for the current request.
--
-- The 'X-Forwarded-For' header is set to the IP address of the client.
-- The 'X-Forwarded-Proto' header is set to the scheme of the request (http or https).
-- The 'X-Content-Type-Options' header is set to 'nosniff' to prevent the browser from doing MIME-type sniffing.
-- The 'X-Frame-Options' header is set to 'DENY' to prevent clickjacking attacks.
-- The 'X-XSS-Protection' header is set to '1; mode=block' to enable XSS filtering.
-- The 'X-Request-Id' header is set to the unique ID of the request from NGINX.
function _M.set_headers()
    ngx.req.set_header("X-Forwarded-For", ngx.var.remote_addr)
    ngx.req.set_header("X-Forwarded-Proto", ngx.var.scheme)
    ngx.req.set_header("X-Content-Type-Options", "nosniff")
    ngx.req.set_header("X-Frame-Options", "DENY")
    ngx.req.set_header("X-XSS-Protection", "1; mode=block")
    ngx.req.set_header("X-Request-Id", ngx.var.request_id)
end

-- This function controls the existence and validity of the token provided in the authorization header.
--
-- It does not verify the token's permissions against a specific resource.
-- The endpoint it calls must not expose sensitive data as the response is passed to the client.
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