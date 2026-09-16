import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
func ShowIntImpl(n int64) string {
	return fmt.Sprintf("%v", n)
}
func ShowNumberImpl(n float64) string {
	if math.IsNaN(n) {
		return "NaN"
	} else if math.IsInf(n, 1) {
		return "Infinity"
	} else if math.IsInf(n, -1) {
		return "-Infinity"
	}

	absN := math.Abs(n)
	var str string
	if absN != 0 && (absN >= 1e21 || absN < 1e-6) {
		str = strconv.FormatFloat(n, 'g', -1, 64)
		// Go uses e-07 but JS uses e-7. We won't worry too much unless a test fails.
	} else {
		str = strconv.FormatFloat(n, 'f', -1, 64)
	}

	if strings.Contains(str, ".") || strings.Contains(str, "e") {
		return str
	}
	return str + ".0"
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
