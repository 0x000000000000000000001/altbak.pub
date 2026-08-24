package Test_PolymorphismFFICheatcode


type Showable_cheatcode interface {
	Show() string
}
type MyInt_cheatcode int
func (m MyInt_cheatcode) Show() string { return "Int" }

func RunPolymorphismFFICheatcode(limit float64) float64 {
	n := int(limit)
	count := 0
	var s Showable_cheatcode = MyInt_cheatcode(0)
	for i := 0; i < n; i++ {
		if s.Show() == "Int" {
			count++
		}
	}
	return float64(count)
}

