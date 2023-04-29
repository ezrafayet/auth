-- auth_handler.lua

local _M = {}

local function createCookie(name, value, expiration)
    return string.format("%s=%s; Expires=%s; Path=/; HttpOnly; Secure", name, value, expiration)
end

function login_service(time, convertTimestampToDate)
    local access_token = "access_token_here"
    local refresh_token = "refresh_token_here"

    local cookies = {}

    table.insert(cookies, createCookie("access_token", access_token, convertTimestampToDate(time + 900)))
    table.insert(cookies, createCookie("refresh_token", refresh_token, convertTimestampToDate(time + 259200)))

    return cookies, 200, nil
end

function _M.login()
    -- Login logic here
    local time = ngx.time();
    local cookies, status, err = login_service(time, ngx.cookie_time);
    -- todo: handle error
    ngx.header["Set-Cookie"] = cookies;
    ngx.header.content_type = "application/json; charset=utf-8";
    ngx.say("{\"status\":\"success\"}");
end

function _M.logout()
    local expired = ngx.time() - 1
    local refresh_cookie = string.format("refresh_token=; Expires=%s; Path=/; HttpOnly; Secure", ngx.cookie_time(expired))
    local access_cookie = string.format("access_token=; Expires=%s; Path=/; HttpOnly; Secure", ngx.cookie_time(expired))
    ngx.header["Set-Cookie"] = {refresh_cookie, access_cookie}
    ngx.header.content_type = "application/json; charset=utf-8"
    ngx.say('{"status":"success"}')
end

return _M