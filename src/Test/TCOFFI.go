package Test_TCOFFI


func RunTCOFFI(limit float64) float64 {
	n := int(limit)
	acc := 0
	for i := n; i > 0; i-- {
		acc += 1
	}
	return float64(acc)
}

