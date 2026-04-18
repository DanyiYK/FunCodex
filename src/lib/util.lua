local util = {}

function util.table_find(tb, value)
	for k, v in ipairs(tb) do
		if value == v then
			return value
		end
	end

	return nil
end

-- Takes a number and returns a table of 0 and 1s
function util.decimal_to_binary(number, size)
	local return_value = {}
	local x = number

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

function util.binary_to_decimal(binary_table)
	local return_value = 0

	for i = #binary_table, 1, -1 do
		if binary_table[i] == 1 then
			return_value = return_value + 2 ^ (#binary_table - i)
		end
	end

	return return_value
end

return util
