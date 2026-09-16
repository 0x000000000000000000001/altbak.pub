import "math"
func IntDegree(x int64) int64 {
	if x < 0 {
		x = -x
	}
	if x > 2147483647 {
		x = 2147483647
	}
	return x
}
func IntDiv(x int64, y int64) int64 {
	if y == 0 {
		return 0
	}
	if y > 0 {
		return int64(math.Floor(float64(x) / float64(y)))
	}
	return int64(-math.Floor(float64(x) / float64(-y)))
}
func IntMod(x int64, y int64) int64 {
	if y == 0 {
		return 0
	}
	if y < 0 {
		y = -y
	}
	return ((x % y) + y) % y
}
func NumDiv(n1 float64, n2 float64) float64 {
	return n1 / n2
}
