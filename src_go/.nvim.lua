-- /path/to/dir/.nvim.lua
require("overseer").register_template({
	name = "My project task",
	params = {},
	condition = {
		-- This makes the template only available in the current directory
		-- In case you :cd out later
		dir = vim.fn.getcwd(),
	},
	builder = function()
		return {
			cmd = { "echo" },
			args = { "Hello", "world" },
		}
	end,
})

require("overseer").register_template({
	name = "my tasks",
	params = {},
	condition = {
		-- This makes the template only available in the current directory
		-- In case you :cd out later
		dir = vim.fn.getcwd(),
	},
	builder = function()
		return {
			cmd = { "echo" },
			args = { "Hello", "world" },
		}
	end,
})
