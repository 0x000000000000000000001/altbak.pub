package Test_TCOFFI


func RunTCOFFI(limit float64) float64 {
	n := int(limit)
	acc := 0
	for n > 0 {
		acc += (n % 3)
		n -= 1
	}
	return float64(acc)
}
