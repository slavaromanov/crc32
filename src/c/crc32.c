#include <stdio.h>
#include <string.h>
#include <inttypes.h>

#define FF            UINT8_MAX
#define CRC           UINT32_MAX
#define POLY          0xEDB88320

static  uint32_t      table[256];
static  uint8_t       table_is_empty = 1;

static void make_table() {
    uint32_t c = 0;
    for (int i = 0; i < 256; i++, c = i) {
      for (int j = 0; j < 8; j++)
        c = c & 1 ? POLY ^ (c >> 1) : (c >> 1);
      table[i] = c;
    }
    table_is_empty = 0;
}    

uint32_t crc32(const char *buf) {
    if (table_is_empty)
      make_table();
    uint32_t crc = CRC;
    const size_t len = strlen(buf);
    const char *p, *q;
    
    q = buf + len;
    for (p = buf; p < q; p++) {
      crc = (crc >> 8) ^ table[crc & FF ^ *p];
    }
    return ~crc;
}

int main(int argc, char *argv[]) {
  const char *s = "The quick brown fox jumps over the lazy dog";
  printf("%lx\n", crc32(s));
  return 0;
}
