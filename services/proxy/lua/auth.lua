local _M = {}

-- Utils

local function createCookie(name, value, expiration)
    return string.format("%s=%s; Expires=%s; Path=/; HttpOnly; Secure", name, value, expiration)
end

-- Services

function login_service(time, convertTimestampToDate)
    local access_token = "access_token_here"
    local refresh_token = "refresh_token_here"

    local cookies = {}

    table.insert(cookies, createCookie("access_token", access_token, convertTimestampToDate(time + 900)))
    table.insert(cookies, createCookie("refresh_token", refresh_token, convertTimestampToDate(time + 259200)))

    return cookies, 200, nil
end

-- Handlers

local _Handlers = {}

function _Handlers.register ()
    ngx.header.content_type = "application/json; charset=utf-8";
    ngx.say("{\"status\":\"success register\"}");
end

function _Handlers.ask_email_verification (urlParams)
    ngx.header.content_type = "application/json; charset=utf-8";
    ngx.say("{\"status\":\"success ask email verification\"}");
end

function _Handlers.submit_email_verification (urlParams)
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

function _Handlers.refresh_token ()
    ngx.header.content_type = "application/json; charset=utf-8"
    ngx.say('{"status":"success refresh"}')
end

-- Router

local routes = {
    {pattern = "^/api/internal/v1/auth/register$", method = "POST", handler = _Handlers.register},
    {pattern = "^/api/internal/v1/auth/users/([%w-]+)/email-verification$", method = "POST", handler = _Handlers.ask_email_verification},
    {pattern = "^/api/internal/v1/auth/users/([%w-]+)/email-verification/([%w-]+)$", method = "PATCH", handler = _Handlers.submit_email_verification},
    {pattern = "^/api/internal/v1/auth/login$", method = "POST", handler = _Handlers.login},
    {pattern = "^/api/internal/v1/auth/logout$", method = "POST", handler = _Handlers.logout},
    {pattern = "^/api/internal/v1/auth/refresh$", method = "POST", handler = _Handlers.refresh_token},
}

local function match_route(route_patterns, uri, method)
    for _, route in ipairs(route_patterns) do
        local captures = {string.match(uri, route.pattern)}
        if #captures > 0 and route.method == method then
            return route.handler, captures
        end
    end
    return nil, nil
end

function _M.router ()
    ngx.log(ngx.ERR, "===== YO", ngx.var.uri, " ", ngx.req.get_method())
    local uri = ngx.var.uri
    local method = ngx.req.get_method()
    local route_handler, urlParams = match_route(routes, uri, method)

    if route_handler then
        route_handler(urlParams)
    else
        ngx.exec("/proxy_auth")
    end
end

return _M