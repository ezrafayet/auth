-- auth.lua

-- Utils

local function create_cookie(name, value, expiration)
    return string.format("%s=%s; Expires=%s; Path=/; HttpOnly; Secure", name, value, expiration)
end

-- Services

function login_service(current_timestamp, convert_timestamp_to_date)
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
    -- Process the response here, e.g., modify, log, etc.
    --ngx.say(res.body)

    ngx.header.content_type = "application/json; charset=utf-8";
    ngx.say("{\"status\":\"success submit email verification\"}");
end

function _Handlers.login ()
    local cookies, status, err = login_service(ngx.time(), ngx.cookie_time);
    -- todo: handle error
    ngx.header["Set-Cookie"] = cookies;
    ngx.header.content_type = "application/json; charset=utf-8";
    ngx.say("{\"status\":\"success login\"}");
end

function _Handlers.logout ()
    local expired = ngx.time() - 1
    local refresh_cookie = string.format("refresh_token=; Expires=%s; Path=/; HttpOnly; Secure", ngx.cookie_time(expired))
    local access_cookie = string.format("access_token=; Expires=%s; Path=/; HttpOnly; Secure", ngx.cookie_time(expired))
    ngx.header["Set-Cookie"] = {refresh_cookie, access_cookie}
    ngx.header.content_type = "application/json; charset=utf-8"
    ngx.say('{"status":"success logout"}')
end

function _Handlers.refresh ()
    ngx.header.content_type = "application/json; charset=utf-8"
    ngx.say('{"status":"success refresh"}')
end

return _Handlers