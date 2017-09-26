local uint8 = 0xff
local uint32 = 0xffffffff
local crc_table = {}

function hex(str)
  return ("%x"):format(str)
end

function xor(x, y) 
  return (~x & y) | (x & ~y) 
end

function even(c)
  if c&1 == 1 then
    c = xor(c, 0x1db710640)
  end
  return c
end

for i=1,256 do
  local c = i - 1
  for j=1,8 do
    c = even(c) >> 1
  end
  crc_table[i] = c
end

function crc32(s)
  local crc = uint32
  for c in s:gmatch('.') do
    crc = xor((crc >> 8), crc_table[xor(crc & uint8, string.byte(c)) + 1])
  end
  return xor(crc, uint32)
end

if #arg == 0 then
  print(string.format("Usage: %s [STRING]", arg[0]))
  os.exit(1)
end
print(hex(crc32(arg[1])))
