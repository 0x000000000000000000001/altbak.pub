import (
	"fmt"
	"strings"
)

func trimExp(s string) string {
	s = strings.Replace(s, "e+0", "e+", 1)
	s = strings.Replace(s, "e-0", "e-", 1)
	return s
}

func ToExponentialNative(fractionDigits int, n float64) string {
	return trimExp(fmt.Sprintf("%.*e", fractionDigits, n))
}
func ToFixedNative(fractionDigits int, n float64) string {
	return fmt.Sprintf("%.*f", fractionDigits, n)
}
func ToPrecisionNative(precision int, n float64) string {
	return trimExp(fmt.Sprintf("%.*g", precision, n))
}
func ToString(n float64) string {
	return fmt.Sprintf("%v", n)
}
