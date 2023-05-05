-- auth.lua
-- todo: must be unit-tested

-- Utils

local function create_cookie(name, value, expiration)
    return string.format("%s=%s; Expires=%s; Path=/; HttpOnly; Secure", name, value, expiration)
end

-- Services

function get_login_cookie(current_timestamp, convert_timestamp_to_date)
    local access_token = "access_token_here"
    local refresh_token = "refresh_token_here"

    local cookies = {}

    table.insert(cookies, create_cookie("access_token", access_token, convert_timestamp_to_date(current_timestamp + 900)))
    table.insert(cookies, create_cookie("refresh_token", refresh_token, convert_timestamp_to_date(current_timestamp + 259200)))

    return cookies, 200, nil
end

-- Handlers

local _Handlers = {}

function _Handlers.email_verification_code ()
    local res = ngx.location.capture("/api/proxy-iam-public", { method = ngx.HTTP_PATCH })
    if res.status >= 200 and res.status < 400 then
        -- todo
        ngx.status = res.status
        ngx.say("{\"status\":\"success\"}");
        return
    end
    ngx.status = res.status
    ngx.say(res.body);
end

function _Handlers.login ()
    local res = ngx.location.capture("/api/proxy-iam-public", { method = ngx.HTTP_POST })
    if res.status >= 200 and res.status < 400 then
        -- todo
        local cookies = get_login_cookie(ngx.time(), ngx.cookie_time);
        ngx.header["Set-Cookie"] = cookies;
        ngx.status = res.status
        ngx.say("{\"status\":\"success\"}");
        return
    end
    ngx.status = res.status
    ngx.say(res.body);
end

-- logout sets cookies to expire in the past
-- it should also blacklist the refresh token
function _Handlers.logout ()
    local expired = ngx.time() - 1
    local refresh_cookie = string.format("refresh_token=; Expires=%s; Path=/; HttpOnly; Secure", ngx.cookie_time(expired))
    local access_cookie = string.format("access_token=; Expires=%s; Path=/; HttpOnly; Secure", ngx.cookie_time(expired))
    ngx.header["Set-Cookie"] = {refresh_cookie, access_cookie}
    ngx.header.content_type = "application/json; charset=utf-8"
    ngx.say('{"status":"success logout"}')
end

function _Handlers.refresh ()
    local res = ngx.location.capture("/api/proxy-iam-public", { method = ngx.HTTP_POST })
    if res.status >= 200 and res.status < 300 then
        -- todo
        ngx.status = res.status
        ngx.say("{\"status\":\"success\"}");
        return
    end
    ngx.status = res.status
    ngx.say(res.body);
end

return _Handlers