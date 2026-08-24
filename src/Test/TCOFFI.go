package Test_TCOFFI


func RunTCOFFI(limit int) int {
	n := int(limit)
	acc := 0
	for n > 0 {
		acc += (n % 3)
		n -= 1
	}
	return (acc)
}
