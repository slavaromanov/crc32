#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <inttypes.h>

uint32_t crc32(const char *s) {
	static uint32_t c = 0, crc = 0xffffffff, table[256];

	for (int i = 0; i < 256; i++) {
		c = i;
		for (int j = 0; j < 8; j++) c = c & 1 ? 0xEDB88320 ^ (c >> 1) : (c >> 1);
		table[i] = c;
	}
	for (const char *p = s; p < s + strlen(s); p++)
		crc = (crc >> 8) ^ table[crc & 0xff ^ *p];
	return ~crc;
}


int main(int argc, char *argv[]) {
	if (argc == 2)
		return printf("%lx\n", crc32(argv[1]));
	else
		return 1;
}
