package Test_ChurchFFI


func RunChurchFFI(limit float64) float64 {
	// Native Church numerals are essentially functional loops.
	// We optimize it as a native loop for the FFI.
	n := int(limit)
	acc := 0
	for i := 0; i < n; i++ {
		acc++
	}
	return float64(acc)
}

