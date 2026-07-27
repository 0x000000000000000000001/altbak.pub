package Data_Show

import "fmt"
func ShowIntImpl(n int) string {
	return fmt.Sprintf("%v", n)
}
func ShowNumberImpl(n float64) string {
	return fmt.Sprintf("%f", n)
}
func ShowCharImpl(c string) string {
	return fmt.Sprintf("'%s'", c)
}
func ShowStringImpl(s string) string {
	return fmt.Sprintf("%q", s)
}
func ShowArrayImpl(f func(interface{}) string, arr []interface{}) string {
	res := "["
	for i, v := range arr {
		if i > 0 {
			res += ","
		}
		res += f(v)
	}
	res += "]"
	return res
}
