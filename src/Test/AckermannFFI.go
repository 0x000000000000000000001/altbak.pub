package Test_AckermannFFI


func ack(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if m > 0 && n == 0 {
		return ack(m-1, 1)
	}
	return ack(m-1, ack(m, n-1))
}

func RunAckermannFFI(limit float64) float64 {
	return float64(ack(3, 4))
}

