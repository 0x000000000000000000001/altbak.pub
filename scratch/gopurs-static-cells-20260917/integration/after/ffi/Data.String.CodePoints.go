package CodePoints

import "unicode/utf8"

func decodeWTF8(s string) []rune {
	var res []rune
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size == 1 {
			if i+2 < len(s) && s[i] == 0xED && s[i+1] >= 0xA0 && s[i+1] <= 0xBF && (s[i+2]&0xC0) == 0x80 {
				cp := rune(int64(s[i]&0x0F)<<12 | int64(s[i+1]&0x3F)<<6 | int64(s[i+2]&0x3F))
				res = append(res, cp)
				i += 3
				continue
			}
		}
		res = append(res, r)
		i += size
	}
	return res
}

func encodeWTF8(runes []rune) string {
	var bytes []byte
	buf := make([]byte, 4)
	for _, r := range runes {
		if r >= 0xD800 && r <= 0xDFFF {
			bytes = append(bytes, 0xED, byte(0xA0|((r>>6)&0x3F)), byte(0x80|(r&0x3F)))
		} else {
			n := utf8.EncodeRune(buf, r)
			bytes = append(bytes, buf[:n]...)
		}
	}
	return string(bytes)
}

func _UnsafeCodePointAt0(fallback interface{}, str string) int64 {
	runes := decodeWTF8(str)
	if len(runes) > 0 {
		return int64(runes[0])
	}
	return 0
}

func _CodePointAt(fallback interface{}, just func(interface{}) interface{}, nothing interface{}, unsafeCodePointAt0 interface{}, index int64, str string) interface{} {
	runes := decodeWTF8(str)
	if index < 0 || index >= int64(len(runes)) {
		return nothing
	}
	return just(int64(runes[index]))
}

func _CountPrefix(fallback interface{}, unsafeCodePointAt0 interface{}, pred func(int64) bool, str string) int64 {
	runes := decodeWTF8(str)
	count := int64(0)
	for _, r := range runes {
		if !pred(int64(r)) {
			break
		}
		count++
	}
	return count
}

func _FromCodePointArray(singleton interface{}, cps []interface{}) string {
	runes := make([]rune, len(cps))
	for i, cp := range cps {
		switch v := cp.(type) {
		case int64:
			runes[i] = rune(v)
		case float64:
			runes[i] = rune(v)
		case int:
			runes[i] = rune(v)
		}
	}
	return encodeWTF8(runes)
}

func _Singleton(fallback interface{}, cp int64) string {
	return encodeWTF8([]rune{rune(cp)})
}

func _Take(fallback interface{}, n int64, str string) string {
	runes := decodeWTF8(str)
	if n < 0 {
		n = 0
	}
	if n >= int64(len(runes)) {
		return str
	}
	return encodeWTF8(runes[:n])
}

func _ToCodePointArray(fallback interface{}, unsafeCodePointAt0 interface{}, str string) []interface{} {
	runes := decodeWTF8(str)
	res := make([]interface{}, len(runes))
	for i, r := range runes {
		res[i] = int64(r)
	}
	return res
}
