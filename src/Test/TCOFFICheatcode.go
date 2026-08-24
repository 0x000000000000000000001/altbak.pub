package Test_TCOFFICheatcode

func RunTCOFFICheatcode(limit int) int {
	n := int(limit)
	acc := 0
	for n > 0 {
		acc += (n % 3)
		n--
	}
	return (acc)
}
