local util = {}

util.BITS_PER_CHARACTER = 5

--[[ 
--Char maps
--NOTE: apparently Lua does not support utf8 natively
--this means that a cyrillic character cannot be converted to lowercase or uppercase
--normally. (Please Lua, add support for UTF-8 ;()
--]]
util.UPPER_SIMILAR_CHARACTERS = { -- List of characters that are only similar in uppercase
	["У"] = "Y",
	["К"] = "K",
	["Е"] = "E",
	["Х"] = "X",
	["А"] = "A",
	["Р"] = "P",
	["О"] = "O",
	["С"] = "C",

	-- Exclusive characters
	["М"] = "M",
	["Т"] = "T",
	["Н"] = "H",
	["В"] = "B",
	["З"] = "3",
}
util.LOWER_SIMILAR_CHARACTERS = {
	["у"] = "y",
	["к"] = "k",
	["е"] = "e",
	["х"] = "x",
	["а"] = "a",
	["р"] = "p",
	["о"] = "o",
	["с"] = "c",
}
util.SPECIAL_SIMILAR_CHARACTERS = {
	[" "] = " ", -- TODO change this to an empty character
}

util.REGISTERED_CHARMAPS = {
	util.UPPER_SIMILAR_CHARACTERS,
	util.LOWER_SIMILAR_CHARACTERS,
	util.SPECIAL_SIMILAR_CHARACTERS,
}

-- List of characters that can be encoded
util.ENCODABLE_CHARACTERS = (function()
	local str = "abcdefghijklmnopqrstuvwxyz "
	local tb = {}
	for i = 1, str:len() do
		tb[i] = str:sub(i, i)
	end
	return tb
end)()

util.EMPTY_CHAR_BIT = 0 -- Equal to ""

function util.table_find(tb, value)
	for k, v in ipairs(tb) do
		if value == v then
			return k
		end
	end

	return nil
end

-- Takes a number and returns a table of 0 and 1s
function util.decimal_to_binary(number, size)
	local return_value = {}

	while number > 0 do
		table.insert(return_value, number % 2)

		number = math.floor(number / 2)
	end

	-- Check for empty spaces
	if #return_value < size then
		for i = #return_value + 1, size do
			return_value[i] = 0
		end
	end

	-- Reverse the result
	local old = return_value
	return_value = {}

	for k, v in ipairs(old) do
		return_value[size + 1 - k] = v
	end

	return return_value
end

-- Takes a table of bits and turns it into a number
function util.binary_to_decimal(binary_table)
	local return_value = 0

	for i = #binary_table, 1, -1 do
		if binary_table[i] == 1 then
			return_value = return_value + 2 ^ (#binary_table - i)
		end
	end

	return return_value
end

-- Checks a table and returns 0 if the value is a key, 1 if the value is a value
-- Returns nil if it's neither of those
function util.key_or_value(tb, value)
	for k, v in pairs(tb) do
		if k == value then
			return 0
		elseif v == value then
			return 1
		end
	end

	return nil
end

-- Returns a tuple:
-- int: value of the bit
-- table: the dict where the char was found
--
-- It can be used both to check if char is encodable and to check its hidden value.
-- It's fun to think but this logic can be applied to us humans too, we all have a hidden value
-- we just need an util.get_bit in our life to reveal it.
function util.get_bit(char)
	local bit_value

	for _, dict in pairs(util.REGISTERED_CHARMAPS) do
		bit_value = util.key_or_value(dict, char)

		if bit_value then
			return bit_value, dict
		end
	end

	return nil
end

-- Returns a tuple:
-- int: the amount of bits available
-- int: the amount of characters that can be stored
function util.get_available_space(str)
	local count = 0
	local letter

	for i = 1, str:len() do
		letter = str:sub(i, i)

		count = count + (util.get_bit(letter) ~= nil and 1 or 0)
	end

	return count, math.floor(count / util.BITS_PER_CHARACTER)
end

return util
