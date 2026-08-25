-- golang-resource-postgresql-library main module.
-- Renders PostgreSQL connection pool setup and baseline migrations.
--
-- The calling archetype is responsible for adding the corresponding
-- Go module dependency:
--   github.com/jackc/pgx/v5
--
-- API (called from a parent archetype):
--   local pg = require("golang-resource-postgresql")
--   pg.render(context, { destination = context:get("project-name") })
--
-- Context contract (no required keys beyond what the calling archetype provides).

local M = {}

function M.render(context, opts)
    opts = opts or {}
    local d = opts.destination
    if d and d ~= "" then
        directory.render("contents", context, { destination = d })
    else
        directory.render("contents", context)
    end
    return context
end

return M
