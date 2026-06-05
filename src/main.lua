local source = debug.getinfo(1, "S").source

if source:sub(1, 1) == "@" then
	source = source:sub(2)
end

local base = source:match("(.*/)") or "./"

package.path = base .. "lib/?.lua;" .. package.path
local Encoder = require("encoder")
local Decoder = require("decoder")
local Util = require("util")

local action = arg[1]
local str = arg[2]
local hidden_text = arg[3]

if action == "length" then
	local total_bits, max_characters = Util.get_available_space(str)

	print(
		"String has a total of "
			.. tostring(total_bits)
			.. " bit capacity it can contain "
			.. tostring(max_characters)
			.. " characters."
	)
elseif action == "crypt" then
	print("CRYPTED TEXT:\n" .. Encoder.encode(str, hidden_text))
elseif action == "decrypt" then
	print('DECRYPTED TEXT:\n"' .. Decoder.decode(str) .. '"')
else
	print("Available commands:\n- length [str]\n- crypt [str] [hidden_text]\n- decrypt [str]")
end
