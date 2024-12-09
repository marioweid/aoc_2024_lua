local file = 'input.txt'

local function sort_func(left, right)
    return left < right
end

-- list
local a_list = {}
local b_list = {}

for line in io.lines(file) do
    local a, b = string.match(line, '(%d+)%s+(%d+)')
    table.insert(a_list, a)
    table.insert(b_list, b)
end

for _, v in pairs(a_list) do
    print(v)
end

table.sort(a_list, sort_func)
table.sort(b_list, sort_func)

print("Sorted 📋: ")

local abs_values = {}
for i = 1, #a_list do
    local abs_val = math.abs(a_list[i] - b_list[i])
    print("abs(" .. a_list[i] .. "-" .. - b_list[i] .. "=" .. abs_val .. ")")
    table.insert(abs_values, abs_val)
end

local result = 0
for _, value in pairs(abs_values) do
    result = result + value
end

print("Result 🏁:" .. result)