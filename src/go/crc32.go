package main

import	"fmt"
import	"os"

	const sep            = " "
	const poly    uint32 = 0xEDB88320

	var code   int
	var table  = make([]uint32, 256)

func join(a []string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		return a[0] + sep + a[1]
	case 3:
		return a[0] + sep + a[1] + sep + a[2]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

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
        if len(os.Args) != 2 {
	  fmt.Printf("%s [STRING or WORD] ...\n", os.Args[0])
	  code = 1
        } else {
	  fmt.Printf("%0x\n", crc32(os.Args[1]))
        }
	os.Exit(code)
}
