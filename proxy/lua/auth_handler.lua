local _M = {}

-- todo: make request, handle errors

function _M.login(time, convertTimestampToDate)
    local access_token = "access_token_here"
    local refresh_token = "refresh_token_here"

    local cookies = {}

    table.insert(cookies, _M.createCookie("access_token", access_token, convertTimestampToDate(time + 900)))
    table.insert(cookies, _M.createCookie("refresh_token", refresh_token, convertTimestampToDate(time + 259200)))

    return cookies, 200, nil
end

function _M.createCookie(name, value, expiration)
    return string.format("%s=%s; Expires=%s; Path=/; HttpOnly; Secure", name, value, expiration)
end

return _M
