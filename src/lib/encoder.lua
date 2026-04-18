local util = require("lib/util")
local encoder = {}

local function get_encoded_letter(charmap, letter, bit)
	for k, v in pairs(charmap) do
		if k==letter or v==letter then
			return bit==0 and k or v
		end
	end

	return nil
end

function encoder.encode(str, hidden_text)
	local return_value = ""
	local hidden_text_binary = {}
	local letter, found

	-- Turn characters into numbers
	for i = 1, hidden_text:len() do
		letter = hidden_text:sub(i, i)
		found = util.table_find(util.ENCODABLE_CHARACTERS, letter:lower())

		if found then
			table.insert(hidden_text_binary, util.decimal_to_binary(found, util.BITS_PER_CHARACTER))
		else
			print("Illegal character in hidden string:", '"'..letter..'"')
		end
	end

	-- Check if lengths match
	local hidden_text_len = hidden_text:len()
	local _, max_characters = util.get_available_space(str)

	if hidden_text_len > max_characters then
		print("WARNING: Hidden text is too long to be hidden in this string, it will be cut!")
	elseif hidden_text_len < max_characters then -- Insert stop signal, so the decoder knows when the text ends
		table.insert(hidden_text_binary, util.decimal_to_binary(util.EMPTY_CHAR_BIT, util.BITS_PER_CHARACTER))
	end

	-- Turn raw binary to FunCodex binary
	local parsing = {}
	local letter_charmap, bit

	for i = 1, str:len() do
		letter = str:sub(i, i)

		_, letter_charmap = util.get_bit(letter)

		if #hidden_text_binary==0 or letter_charmap==nil then
			return_value = return_value..letter
			goto continue
		elseif #parsing==0 then -- Binary chunk is empty, get another one
			parsing = table.remove(hidden_text_binary, 1)
		end

		bit = table.remove(parsing, 1)

		return_value = return_value..get_encoded_letter(letter_charmap, letter, bit)

		::continue::
	end

	return return_value
end

return encoder
