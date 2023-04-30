-- access.lua

local _M = {}

-- accessNeeded = user | admin
function _M.hasAccess(accessNeeded)
    ngx.say("access")
end

return _M