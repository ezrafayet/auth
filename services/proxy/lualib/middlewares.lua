-- access.lua
-- todo: must be unit-tested

local _M = {}

function extract_user_id_from_jwt(jwt)
    local segments = {}
    for segment in string.gmatch(jwt, "[^%.]+") do
        table.insert(segments, segment)
    end

    if #segments ~= 3 then
        return nil, "invalid JWT format"
    end

    local payload_b64 = segments[2]
    local payload_json, err = ngx.decode_base64(payload_b64)
    if not payload_json then
        return nil, "error decoding JWT payload: " .. err
    end

    local payload = cjson.decode(payload_json)
    return payload.userId
end

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
    ngx.req.set_header("X-Request-Id", "rid-" .. ngx.var.request_id)
end

function set_headers_initiator ()
    local cookie_access_token = ngx.var.cookie_access_token
    local api_access_token = ngx.req.get_headers()["Authorization"]
    local cli_request = ngx.req.get_headers()["X_Is-Cli"]
    local user_agent = ngx.req.get_headers()["User-Agent"]
    local is_web_user_agent = user_agent and (string.find(user_agent, "Mozilla") ~= nil)

    if (cookie_access_token ~= nil and #cookie_access_token ~= 0) then
        local user_id, _ = extract_user_id_from_jwt(cookie_access_token)
        ngx.req.set_header("X-Initiator-Id", user_id)
        ngx.req.set_header("X-Initiator-Type", "web")
    elseif (api_access_token ~= nil and #api_access_token ~= 0) then
        ngx.req.set_header("X-Initiator-Id", api_access_token)
        if (cli_request == "true") then
            ngx.req.set_header("X-Initiator-Type", "cli")
        else
            ngx.req.set_header("X-Initiator-Type", "api")
        end
    else
        ngx.req.set_header("X-Initiator-Id", "")
        if (is_web_user_agent) then
            ngx.req.set_header("X-Initiator-Type", "web")
        else
            ngx.req.set_header("X-Initiator-Type", "api")
        end
    end
end

-- sets headers that must be passed to the upstream services
function _M.set_headers()
    set_headers_forwarded_for()
    set_headers_forwarded_proto()
    set_headers_content_type_options()
    set_headers_frame_options()
    set_headers_xss_protection()
    set_headers_request_id()
    set_headers_initiator()
end

function _M.control_access_token()
    -- todo
end

return _M