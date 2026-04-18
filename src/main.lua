local Encoder = require("./lib/encoder")
local Decoder = require("./lib/decoder")
local Util = require("./lib/util")

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
