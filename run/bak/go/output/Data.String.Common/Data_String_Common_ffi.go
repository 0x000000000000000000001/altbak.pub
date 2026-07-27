package Data_String_Common

import (
	"strings"
)

func _LocaleCompare(lt interface{}, eq interface{}, gt interface{}, s1 string, s2 string) interface{} {
	cmp := strings.Compare(s1, s2)
	if cmp < 0 {
		return lt
	} else if cmp > 0 {
		return gt
	}
	return eq
}

func Replace(s1 string, s2 string, s3 string) string {
	return strings.Replace(s3, s1, s2, 1)
}

func ReplaceAll(s1 string, s2 string, s3 string) string {
	return strings.ReplaceAll(s3, s1, s2)
}

func Split(sep string, s string) []string {
	return strings.Split(s, sep)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func JoinWith(s string, xs []string) string {
	return strings.Join(xs, s)
}
