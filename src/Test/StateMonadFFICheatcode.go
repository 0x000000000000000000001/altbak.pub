package Test_StateMonadFFICheatcode


func RunStateMonadFFICheatcode(limit int) int {
	// The shared FFI wrapper supplies the depth; evaluate 20 repetitions without closures.
	state := 0
	for i := 0; i < 20; i++ {
		for j := 0; j < limit; j++ {
			state += 1
		}
	}
	return (state)
}
