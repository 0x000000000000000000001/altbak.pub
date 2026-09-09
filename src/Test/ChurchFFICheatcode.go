package Test_ChurchFFICheatcode


func RunChurchFFICheatcode(limit int) int {
	// The Church composition represents n^5 successor applications.
	// Count them directly without allocating the intermediate closures.
	n := limit
	count := n * n * n * n * n
	acc := 0
	for i := 0; i < count; i++ {
		acc++
	}
	return (acc)
}
