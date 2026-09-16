package Test_PolymorphismFFICheatcode

type Monoidish_Cheatcode interface {
	mempty_() int
	mappend_(int, int) int
}
type IntMonoidish_Cheatcode struct{}
func (IntMonoidish_Cheatcode) mempty_() int { return 1 }
func (IntMonoidish_Cheatcode) mappend_(x, y int) int { return x + y }

func RunPolymorphismFFICheatcode(limit int) int {
	n := int(limit)
	acc := 0
	var m Monoidish_Cheatcode = IntMonoidish_Cheatcode{}
	for i := 0; i < n; i++ {
		acc = m.mappend_(acc, m.mempty_())
	}
	return acc
}
