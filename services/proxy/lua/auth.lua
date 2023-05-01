local _M = {}

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

function _Handlers.register ()
    ngx.header.content_type = "application/json; charset=utf-8";
    ngx.say("{\"status\":\"success register\"}");
end

function _Handlers.ask_email_verification (capture)
    local user_id = capture[1]
    ngx.header.content_type = "application/json; charset=utf-8";
    ngx.say("{\"status\":\"success ask email verification\"}");
end

function _Handlers.submit_email_verification (capture)
    local user_id = capture[1]
    local token = capture[2]
    ngx.log(ngx.ERR, "submit_email_verification: " .. capture[1] .. " " .. capture[2])
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
-- NB: This router is not optimized for performance.
-- It should be replaced by a proper router.

local routes = {
    {pattern = "/api/internal/v1/auth/register", method = "POST", handler = _Handlers.register},
    {pattern = "/api/internal/v1/auth/users/*/email-verification", method = "POST", handler = _Handlers.ask_email_verification},
    {pattern = "/api/internal/v1/auth/users/*/email-verification/*", method = "PATCH", handler = _Handlers.submit_email_verification},
    {pattern = "/api/internal/v1/auth/login", method = "POST", handler = _Handlers.login},
    {pattern = "/api/internal/v1/auth/logout", method = "POST", handler = _Handlers.logout},
    {pattern = "/api/internal/v1/auth/refresh", method = "POST", handler = _Handlers.refresh_token},
}

local function match_route(_routes, uri, method)
    local uri_parts = {}
    for part in string.gmatch(uri, "[^/]+") do
        table.insert(uri_parts, part)
    end
    for _, route in ipairs(_routes) do
        local route_parts = {}
        for part in string.gmatch(route.pattern, "[^/]+") do
            table.insert(route_parts, part)
        end
        if #route_parts == #uri_parts then
            local captures = {}
            local route_matched = true
            for i, route_part in ipairs(route_parts) do
                if route_part == "*" then
                    table.insert(captures, uri_parts[i])
                elseif route_part ~= uri_parts[i] then
                    route_matched = false
                    break
                end
            end
            if route_matched and route.method == method then
                return route.handler, captures
            end
        end
    end
    return nil, nil
end

function _M.router ()
    local uri = ngx.var.uri
    if string.sub(uri, -1) == "/" then
        uri = string.sub(uri, 1, -2)
    end
    local method = ngx.req.get_method()
    local route_handler, urlParams = match_route(routes, uri, method)

    if route_handler then
        route_handler(urlParams)
    else
        ngx.exec("/proxy_auth")
    end
end

return _M