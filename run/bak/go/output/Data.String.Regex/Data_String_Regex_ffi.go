package Data_String_Regex

import "gopurs/output/gopurs_runtime"

import (
	"regexp"
	"strings"
)

type GoRegex struct {
	Re     *regexp.Regexp
	Global bool
	Flags  string
	Source string
}

func RegexImpl(left func(string) interface{}, right func(interface{}) interface{}, s1 string, s2 string) interface{} {
	flags := ""
	if strings.Contains(s2, "i") {
		flags += "i"
	}
	if strings.Contains(s2, "m") {
		flags += "m"
	}

	pattern := s1
	if flags != "" {
		pattern = "(?" + flags + ")" + s1
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return left(err.Error())
	}

	goRegex := &GoRegex{
		Re:     re,
		Global: strings.Contains(s2, "g"),
		Flags:  s2,
		Source: s1,
	}
	return right(goRegex)
}

func _ReplaceBy(just func(interface{}) interface{}, nothing interface{}, regex *GoRegex, f func(string) func([]interface{}) string, s string) string {
	if regex != nil {
	}
	if regex == nil || regex.Re == nil {
		return s
	}
	matches := regex.Re.FindAllStringSubmatchIndex(s, -1)
	if !regex.Global && len(matches) > 1 {
		matches = matches[:1]
	}
	if len(matches) == 0 {
		return s
	}

	var sb strings.Builder
	lastMatchEnd := 0
	for _, matchIdxs := range matches {
		fullMatch := s[matchIdxs[0]:matchIdxs[1]]
		groups := make([]interface{}, 0)
		for i := 2; i < len(matchIdxs); i += 2 {
			if matchIdxs[i] == -1 {
				groups = append(groups, nothing)
			} else {
				groups = append(groups, just(s[matchIdxs[i]:matchIdxs[i+1]]))
			}
		}

		replacement := f(fullMatch)(groups)
		sb.WriteString(s[lastMatchEnd:matchIdxs[0]])
		sb.WriteString(replacement)
		lastMatchEnd = matchIdxs[1]
	}
	sb.WriteString(s[lastMatchEnd:])
	return sb.String()
}

func Replace(regex *GoRegex, s1 string, s2 string) string {
	if regex == nil || regex.Re == nil {
		return s2
	}
	if regex.Global {
		return regex.Re.ReplaceAllString(s2, s1)
	}
	
	loc := regex.Re.FindStringSubmatchIndex(s2)
	if loc == nil {
		return s2
	}
	
	var res []byte
	res = append(res, s2[:loc[0]]...)
	res = regex.Re.ExpandString(res, s1, s2, loc)
	res = append(res, s2[loc[1]:]...)
	return string(res)
}

func _Match(just func(interface{}) interface{}, nothing interface{}, r *GoRegex, s string) interface{} { 
	panic("Not implemented: Regex._Match")
}

func _Search(just func(interface{}) interface{}, nothing interface{}, r *GoRegex, s string) interface{} { 
	panic("Not implemented: Regex._Search")
}

func FlagsImpl(r *GoRegex) map[string]interface{} { 
	if r == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"global": r.Global,
		"ignoreCase": strings.Contains(r.Flags, "i"),
		"multiline": strings.Contains(r.Flags, "m"),
		"dotAll": strings.Contains(r.Flags, "s"),
		"sticky": false,
		"unicode": true,
	}
}

func ShowRegexImpl(r *GoRegex) string { 
	if r == nil {
		return "//"
	}
	return "/" + r.Source + "/" + r.Flags
}

func Source(r *GoRegex) string { 
	if r == nil {
		return ""
	}
	return r.Source 
}

func Split(r *GoRegex, s string) []string { 
	if r == nil || r.Re == nil {
		return []string{s}
	}
	if r.Global {
		return r.Re.Split(s, -1)
	}
	return r.Re.Split(s, 2)
}

func Test(r *GoRegex, s string) bool { 
	if r == nil || r.Re == nil {
		return false
	}
	return r.Re.MatchString(s)
}


// --- Auto-generated FFI wrappers ---
func Call_regexImpl(arg0 func(string) interface{}, arg1 func(interface{}) interface{}, arg2 string, arg3 string) interface{} {
	return RegexImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_RegexImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := RegexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call__ReplaceBy(arg0 func(interface{}) interface{}, arg1 interface{}, arg2 *GoRegex, arg3 func(string) func([]interface{}) string, arg4 string) string {
	return _ReplaceBy(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs__ReplaceBy = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*GoRegex](arg2)
	go_arg3 := func(p0_0 string) func([]interface{}) string {
			inner_res0 := gopurs_runtime.Apply(arg3, gopurs_runtime.Box(p0_0))
			return func(p1_0 []interface{}) string {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return gopurs_runtime.Unbox[string](inner_res1)
		}
		}
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := _ReplaceBy(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call_replace(arg0 *GoRegex, arg1 string, arg2 string) string {
	return Replace(arg0, arg1, arg2)
}
var _Gopurs_Replace = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*GoRegex](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Replace(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call__Match(arg0 func(interface{}) interface{}, arg1 interface{}, arg2 *GoRegex, arg3 string) interface{} {
	return _Match(arg0, arg1, arg2, arg3)
}
var _Gopurs__Match = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*GoRegex](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := _Match(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call__Search(arg0 func(interface{}) interface{}, arg1 interface{}, arg2 *GoRegex, arg3 string) interface{} {
	return _Search(arg0, arg1, arg2, arg3)
}
var _Gopurs__Search = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*GoRegex](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := _Search(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_flagsImpl(arg0 *GoRegex) map[string]interface{} {
	return FlagsImpl(arg0)
}
var _Gopurs_FlagsImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*GoRegex](arg0)
	go_res := FlagsImpl(go_arg0)
	return func() gopurs_runtime.Value {
			res_map := make(map[string]gopurs_runtime.Value)
			for k, v := range go_res { res_map[k] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Record(res_map)
		}()
})
func Call_showRegexImpl(arg0 *GoRegex) string {
	return ShowRegexImpl(arg0)
}
var _Gopurs_ShowRegexImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*GoRegex](arg0)
	go_res := ShowRegexImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_source(arg0 *GoRegex) string {
	return Source(arg0)
}
var _Gopurs_Source = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*GoRegex](arg0)
	go_res := Source(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_split(arg0 *GoRegex, arg1 string) []string {
	return Split(arg0, arg1)
}
var _Gopurs_Split = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*GoRegex](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Split(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_test(arg0 *GoRegex, arg1 string) bool {
	return Test(arg0, arg1)
}
var _Gopurs_Test = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*GoRegex](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Test(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
