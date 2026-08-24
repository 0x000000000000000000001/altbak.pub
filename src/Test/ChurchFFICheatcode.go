package Test_ChurchFFICheatcode


func RunChurchFFICheatcode(limit int) int {
	// Native Church numerals are essentially functional loops.
	// We optimize it as a native loop for the FFICheatcode.
	n := int(limit)
	acc := 0
	for i := 0; i < n * 10000; i++ {
		acc++
	}
	return (acc)
}

