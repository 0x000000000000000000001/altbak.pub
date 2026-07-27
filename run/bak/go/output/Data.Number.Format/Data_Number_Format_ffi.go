package Data_Number_Format

import "fmt"
func ToExponentialNative(fractionDigits int, n float64) string { return fmt.Sprintf("%.*e", fractionDigits, n) }
func ToFixedNative(fractionDigits int, n float64) string { return fmt.Sprintf("%.*f", fractionDigits, n) }
func ToPrecisionNative(precision int, n float64) string { return fmt.Sprintf("%.*g", precision, n) }
func ToString(n float64) string { return fmt.Sprintf("%v", n) }
