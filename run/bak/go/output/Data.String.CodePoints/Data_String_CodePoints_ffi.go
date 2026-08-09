package Data_String_CodePoints

import "gopurs/output/gopurs_runtime"


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


// --- Auto-generated FFI wrappers ---
var _Gopurs__CodePointAt = // TAST: (Func [(Func [Int, String] (ADT ["Data","Maybe","Maybe"] [Int])), (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), (Func [String] Int), Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := gopurs_runtime.Unbox[int64](arg4)
	go_arg5 := gopurs_runtime.Unbox[string](arg5)
	go_res := _CodePointAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__CountPrefix = // TAST: (Func [(Func [(Func [Int] Boolean), String] Int), (Func [String] Int), (Func [Int] Boolean), String] Int)
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := func(p0_0 int64) bool {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := _CountPrefix(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__FromCodePointArray = // TAST: (Func [(Func [Int] String), (Array Int)] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := _FromCodePointArray(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__Singleton = // TAST: (Func [(Func [Int] String), Int] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := _Singleton(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__Take = // TAST: (Func [(Func [Int, String] String), Int, String] String)
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := _Take(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ToCodePointArray = // TAST: (Func [(Func [String] (Array Int)), (Func [String] Int), String] (Array Int))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := _ToCodePointArray(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs__UnsafeCodePointAt0 = // TAST: (Func [(Func [String] Int), String] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := _UnsafeCodePointAt0(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})