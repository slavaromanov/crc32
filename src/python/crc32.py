import sys

def create_table():
    a = []
    for i in range(256):
        k = i
        for j in range(8):
            if k & 1:
                k ^= 0x1db710640
            k >>= 1
        a.append(k)
    return a
 
def crc_update(buf, crc):
    crc ^= 0xffffffff
    for k in buf:
        crc = (crc >> 8) ^ crc_table[(crc & 0xff) ^ k]
    return crc ^ 0xffffffff

if __name__ == "__main__":
    if len(sys.argv) == 1:
        print('Empty input!\n\tUsage: {} [STRING]'.format(sys.argv[0]))
        exit(1)
    crc_table = create_table()
    print(hex(crc_update(sys.argv[1].encode(), 0))[2:])
