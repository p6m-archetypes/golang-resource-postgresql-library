-- golang-resource-postgresql-library standalone entry point.
--
-- Parents consuming this library should depend on it with `library: true`
-- and call `require("golang-resource-postgresql").render(context, { destination = "project-name" })`.
-- This script runs when the archetype is invoked directly to retrofit an
-- existing project:
--
--   archetect render .../golang-resource-postgresql-library <project-dir>

local context = Context.new()
require("lib").render(context)
return context
