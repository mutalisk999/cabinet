package utils

import "unicode/utf8"

func BytesToRunes(bytesBuf []byte) []rune {
	buf := bytesBuf
	runes := make([]rune, 0)
	for len(buf) > 0 {
		r, size := utf8.DecodeRune(buf)
		buf = buf[size:]
		runes = append(runes, r)
	}
	return runes
}
