package main

import "fmt"

const poly uint32 = 0xEDB88320
var   table       = make([]uint32, 256)

func init() {
	for i := 0; i < 256; i++ {
		c := uint32(i)
		for j := 0; j < 8; j++ {
			b := c&1 == 1
			c >>= 1
			if b {
				c ^= poly
			}
		}
		table[i] = c
	}
}

func crc32(s string) uint32 {
	crc := ^uint32(0)
	for i := 0; i < len(s); i++ {
		crc = table[byte(crc)^s[i]] ^ (crc >> 8)
	}
	return ^crc
}

func main() {
	fmt.Printf("%0x\n", crc32("The quick brown fox jumps over the lazy dog"))
}
