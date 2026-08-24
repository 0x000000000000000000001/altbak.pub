package Test_AckermannFFICheatcode


func ack_cheatcode(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if m > 0 && n == 0 {
		return ack_cheatcode(m-1, 1)
	}
	return ack_cheatcode(m-1, ack_cheatcode(m, n-1))
}

func RunAckermannFFICheatcode(limit int) int {
	return (ack_cheatcode(3, 4))
}

