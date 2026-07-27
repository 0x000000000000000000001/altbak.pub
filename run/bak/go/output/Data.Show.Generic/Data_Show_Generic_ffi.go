package Data_Show_Generic

import "strings"
func Intercalate(separator string, arr []string) string {
	return strings.Join(arr, separator)
}
