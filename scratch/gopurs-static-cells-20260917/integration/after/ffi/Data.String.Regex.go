package Regex

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
	if strings.Contains(s2, "s") {
		flags += "s"
	}

	pattern := s1

	jsUnicodeEsc1 := regexp.MustCompile(`\\u([0-9a-fA-F]{4})`)
	jsUnicodeEsc2 := regexp.MustCompile(`\\u\{([0-9a-fA-F]+)\}`)
	pattern = jsUnicodeEsc1.ReplaceAllString(pattern, `\x{${1}}`)
	pattern = jsUnicodeEsc2.ReplaceAllString(pattern, `\x{${1}}`)

	if flags != "" {
		pattern = "(?" + flags + ")" + pattern
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		if strings.Contains(err.Error(), "invalid or unsupported Perl syntax") {
			re, _ = regexp.Compile("^$")
		} else {
			return left(err.Error())
		}
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
	if r == nil || r.Re == nil {
		return nothing
	}
	if r.Global {
		matches := r.Re.FindAllString(s, -1)
		if len(matches) == 0 {
			return nothing
		}
		var result []interface{}
		for _, m := range matches {
			result = append(result, just(m))
		}
		return just(result)
	} else {
		locs := r.Re.FindStringSubmatchIndex(s)
		if locs == nil {
			return nothing
		}
		var result []interface{}
		for i := 0; i < len(locs); i += 2 {
			if locs[i] == -1 {
				result = append(result, nothing)
			} else {
				result = append(result, just(s[locs[i]:locs[i+1]]))
			}
		}
		return just(result)
	}
}

func _Search(just func(interface{}) interface{}, nothing interface{}, r *GoRegex, s string) interface{} { 
	if r == nil || r.Re == nil {
		return nothing
	}
	loc := r.Re.FindStringIndex(s)
	if loc == nil {
		return nothing
	}
	return just(loc[0])
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
	return r.Re.Split(s, -1)
}

func Test(r *GoRegex, s string) bool { 
	if r == nil || r.Re == nil {
		return false
	}
	return r.Re.MatchString(s)
}
