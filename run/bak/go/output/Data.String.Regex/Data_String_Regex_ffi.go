package Data_String_Regex

import "regexp"
func _Match(just func(interface{}) interface{}, nothing interface{}, r *regexp.Regexp, s string) interface{} { return nothing }
func _ReplaceBy(r *regexp.Regexp, f func(string, []interface{}) string, s string) string { return s }
func _Search(just func(interface{}) interface{}, nothing interface{}, r *regexp.Regexp, s string) interface{} { return nothing }

func FlagsImpl(r interface{}) string { return "" }
func RegexImpl(left interface{}, right interface{}, s1 string, s2 string) interface{} { return left }
func Replace(r interface{}, s1 string, s2 string) string { return s2 }
func ShowRegexImpl(r interface{}) string { return "" }
func Source(r interface{}) string { return "" }
func Split(r interface{}, s string) []string { return []string{s} }
func Test(r interface{}, s string) bool { return false }
