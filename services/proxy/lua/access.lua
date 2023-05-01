-- access.lua

local _M = {}

-- accessNeeded = user | admin
function _M.has_access(access_needed)
    ngx.say("access")
end

return _M