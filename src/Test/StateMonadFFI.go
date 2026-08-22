package Test_StateMonadFFI


func RunStateMonadFFI(limit float64) float64 {
	// 60 depth x 20 nested binds = 1200 binds
	state := 0
	for i := 0; i < 60; i++ {
		for j := 0; j < 20; j++ {
			state += 1
		}
	}
	return float64(state)
}

