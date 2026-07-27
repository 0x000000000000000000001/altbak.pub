package Data_String_CodeUnits

import (
	"strings"
)

func FromCharArray(a []string) string {
	var b strings.Builder
	for _, v := range a {
		b.WriteString(v)
	}
	return b.String()
}

func ToCharArray(str string) []string {
	arr := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		arr[i] = string(str[i])
	}
	return arr
}

func Singleton(c string) string {
	return c
}

func _CharAt(just func(string) interface{}, nothing interface{}, idx int, str string) interface{} {
	if idx >= 0 && idx < len(str) {
		return just(string(str[idx]))
	}
	return nothing
}

func _ToChar(just func(string) interface{}, nothing interface{}, str string) interface{} {
	if len(str) == 1 {
		return just(str)
	}
	return nothing
}

func Length(s string) int {
	return len(s)
}

func CountPrefix(p func(string) bool, str string) int {
	i := 0
	for i < len(str) {
		if p(string(str[i])) {
			i++
		} else {
			break
		}
	}
	return i
}

func _IndexOf(just func(int) interface{}, nothing interface{}, x string, s string) interface{} {
	idx := strings.Index(s, x)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func _IndexOfStartingAt(just func(int) interface{}, nothing interface{}, x string, startIdx int, str string) interface{} {
	if startIdx < 0 || startIdx > len(str) {
		return nothing
	}
	idx := strings.Index(str[startIdx:], x)
	if idx == -1 {
		return nothing
	}
	return just(idx + startIdx)
}

func _LastIndexOf(just func(int) interface{}, nothing interface{}, x string, s string) interface{} {
	idx := strings.LastIndex(s, x)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func _LastIndexOfStartingAt(just func(int) interface{}, nothing interface{}, x string, startIdx int, str string) interface{} {
	if startIdx < 0 || startIdx >= len(str) {
		return nothing
	}
	idx := strings.LastIndex(str[:startIdx+len(x)], x)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func Take(idx int, str string) string {
	if idx < 0 {
		idx = 0
	}
	if idx > len(str) {
		idx = len(str)
	}
	return str[:idx]
}

func Drop(idx int, str string) string {
	if idx < 0 {
		idx = 0
	}
	if idx > len(str) {
		idx = len(str)
	}
	return str[idx:]
}

func Slice(start int, end int, str string) string {
	if start < 0 {
		start = len(str) + start
	}
	if end < 0 {
		end = len(str) + end
	}
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > len(str) {
		start = len(str)
	}
	if end > len(str) {
		end = len(str)
	}
	if start > end {
		return ""
	}
	return str[start:end]
}

func SplitAt(idx int, str string) map[string]interface{} {
	if idx < 0 {
		idx = 0
	}
	if idx > len(str) {
		idx = len(str)
	}
	rec := make(map[string]interface{})
	rec["before"] = str[:idx]
	rec["after"] = str[idx:]
	return rec
}
