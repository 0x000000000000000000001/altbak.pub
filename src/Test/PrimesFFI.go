package Test_PrimesFFI


func RunPrimesFFI(limit float64) float64 {
	n := int(limit)
	if n < 2 {
		return 0
	}
	sieve := make([]bool, n+1)
	sum := 0
	for p := 2; p <= n; p++ {
		if !sieve[p] {
			sum += p
			for i := p * p; i <= n; i += p {
				sieve[i] = true
			}
		}
	}
	return float64(sum)
}

