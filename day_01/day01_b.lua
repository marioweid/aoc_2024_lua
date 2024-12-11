local file = 'input.txt'

-- list
local a_list = {}
local b_list = {}

for line in io.lines(file) do
    local a, b = string.match(line, '(%d+)%s+(%d+)')
    table.insert(a_list, a)
    table.insert(b_list, b)
end

local count_table = {}
for _, entry in pairs(b_list) do
    local index = entry
    if count_table[index] == nil then
        count_table[index] = 1
    else
        count_table[index] = count_table[index] + 1
    end
end

print("Counted 📋: ")
for k, v in pairs(count_table) do
    print(k .. " --- ".. v)
end

local result = 0
for _, entry in pairs(a_list) do
    if count_table[entry] == nil then
        
    else
        result = result + entry * count_table[entry]
    end
end
print(result)