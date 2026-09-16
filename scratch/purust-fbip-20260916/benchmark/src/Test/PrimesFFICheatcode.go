package Test_PrimesFFICheatcode


func RunPrimesFFICheatcode(limit int) int {
	n := limit
	if n < 2 {
		return 0
	}
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}
	
	for p := 2; p*p <= n; p++ {
		if sieve[p] {
			for i := p * p; i <= n; i += p {
				sieve[i] = false
			}
		}
	}
	
	sum := 0
	for p := 2; p <= n; p++ {
		if sieve[p] {
			sum += p
		}
	}
	return sum
}
