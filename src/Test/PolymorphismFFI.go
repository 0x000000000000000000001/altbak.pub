package Test_PolymorphismFFI


type Monoidish interface {
	mempty_() int
	mappend_(int) func(int) int
}

type IntMonoidish struct{}

func (IntMonoidish) mempty_() int {
	return 1
}

func (IntMonoidish) mappend_(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func polyLoop(dict Monoidish, n_init int, acc_init int) int {
	var goFunc func(int, int) int
	goFunc = func(n int, acc int) int {
		if n == 0 {
			return acc
		}
		return goFunc(n-1, dict.mappend_(acc)(dict.mempty_()))
	}
	return goFunc(n_init, acc_init)
}

func RunPolymorphismFFI(limit float64) float64 {
	dummy := int(limit)
	return float64(polyLoop(IntMonoidish{}, dummy, 0))
}
