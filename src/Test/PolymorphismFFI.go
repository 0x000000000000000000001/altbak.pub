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
	n := n_init
	acc := acc_init
	for n > 0 {
		acc = dict.mappend_(acc)(dict.mempty_())
		n--
	}
	return acc
}

func RunPolymorphismFFI(limit int) int {
	dummy := limit
	return (polyLoop(IntMonoidish{}, dummy, 0))
}
