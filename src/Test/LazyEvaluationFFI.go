package Test_LazyEvaluationFFI


func RunLazyEvaluationFFI(limit float64) float64 {
	// A simple loop simulating forced thunks without overhead
	depth := int(limit)
	total := 0
	for i := 0; i < depth*1000; i++ {
		total += 1
	}
	return float64(total)
}

