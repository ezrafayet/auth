-- auth_router.lua
local _M = {}

local handlerV1 = require("auth_handler")

local routes = {
    ["/api/internal/v1/auth/login"] = { ["POST"] = handlerV1.login },
    ["/api/internal/v1/auth/logout"] = { ["POST"] = handlerV1.logout },
}

function _M.run()
    local route_handler = routes[ngx.var.uri][ngx.req.get_method()]

    if route_handler then
        route_handler()
    else
        ngx.exec("/proxy_auth")
    end
end

return _M

--To implement:

--# Creates a new user with a username, an email and a password
--# POST /api/internal/v1/auth/register (PUBLIC)
--# For the password flow only
--
--# Sends verification email to user
--# POST /api/internal/v1/auth/users/{userId}/email-verification (PUBLIC)
--# For both password and magic link flows
--
--# Verifies user email
--# PATCH /api/v1/auth/users/{userId}/email-verification/{verificationToken} (PUBLIC)
--# For both password and magic link flows
--
--# User requests to login
--# POST /api/internal/v1/auth/login (PUBLIC)
--# For the password flow only
--#
--# # User requests to logout
--# # POST /api/internal/v1/auth/logout (PUBLIC)
--# # For both password and magic link flows
--#
--# # User requests a new access token
--# # POST /api/v1/internal/auth/refresh (PUBLIC)
