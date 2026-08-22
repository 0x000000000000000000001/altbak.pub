package Test_PolymorphismFFI


type Showable interface {
	Show() string
}
type MyInt int
func (m MyInt) Show() string { return "Int" }

func RunPolymorphismFFI(limit float64) float64 {
	n := int(limit)
	count := 0
	var s Showable = MyInt(0)
	for i := 0; i < n; i++ {
		if s.Show() == "Int" {
			count++
		}
	}
	return float64(count)
}

