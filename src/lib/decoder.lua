local util = require("lib/util")
local ustring = require("lib/utf8string")

local decoder = {}

function decoder.decode(str)
	local return_value = ""
	local found_binary = {}

	local utf_str = ustring(str)

	local bit_group = {}

	for i = 1, utf_str:len() do
		letter = utf_str:sub(i, i)

		found, _ = util.get_bit(letter.rawstring)

		if found then
			table.insert(bit_group, found)

			if #bit_group >= util.BITS_PER_CHARACTER then
				table.insert(found_binary, bit_group)

				bit_group = {}
			end
		end
	end

	local decimal_value
	for _, binary_table in ipairs(found_binary) do
		decimal_value = util.binary_to_decimal(binary_table)

		if decimal_value == util.EMPTY_CHAR_BIT then
			break
		end

		return_value = return_value .. (util.ENCODABLE_CHARACTERS[decimal_value] or "[?]")
	end

	return return_value
end

return decoder
