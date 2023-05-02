-- access.lua

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

-- set_header_variables sets variables that must be
-- passed as headers to the upstream services
function _M.set_header_variables()
    -- Set operation ID
    local oid = ngx.req.get_headers()["X-Operation-Id"]

    if (oid ~= nil and #oid ~= 0) then
        ngx.var["x_operation_id"] = ngx.req.get_headers()["X-Operation-Id"]
    else
        ngx.var["x_operation_id"] = string.format("oid-%x", math.random(1000000, 9999999))
    end

    -- Set initiator ID and type
    local cookie_access_token = ngx.var.cookie_access_token
    local api_access_token = ngx.req.get_headers()["Authorization"]
    local initiator_type = ngx.req.get_headers()["X_Initiator_Type"]
    local user_agent = ngx.req.get_headers()["User-Agent"]

    if (cookie_access_token ~= nil and #cookie_access_token ~= 0) then
        local user_id, err = extract_user_id_from_jwt(cookie_access_token)
        -- todo: handle error
        ngx.var["x_initiator_id"] = user_id;
        ngx.var["x_initiator_type"] = "web";
    elif (api_access_token ~= nil and #api_access_token ~= 0 and initiator_type ~= "cli")
        ngx.var["x_initiator_id"] = api_access_token;
        ngx.var["x_initiator_type"] = "api";
    elif (initiator_type == "cli")
        ngx.var["x_initiator_id"] = api_access_token;
        ngx.var["x_initiator_type"] = "cli";
    elif (user_agent ~= nil and #user_agent ~= 0)
        ngx.var["x_initiator_id"] = "";
        ngx.var["x_initiator_type"] = "web";
    else
        -- consider returning an error
        ngx.var["x_initiator_id"] = "";
        ngx.var["x_initiator_type"] = "api";
    end
end

function _M.control_access_token()
    -- todo
end

return _M