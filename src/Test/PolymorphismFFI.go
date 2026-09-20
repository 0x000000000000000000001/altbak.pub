package Test_PolymorphismFFI

type Monoidish[A any] struct {
	Mempty  A
	Mappend func(A) func(A) A
}

func polyLoop[A any](dict Monoidish[A], n int, acc A) A {
	for n != 0 {
		acc = dict.Mappend(acc)(dict.Mempty)
		n--
	}
	return acc
}

var intMonoidish = Monoidish[int]{
	Mempty: 1,
	Mappend: func(x int) func(int) int {
		return func(y int) int { return x + y }
	},
}

func RunPolymorphismFFI(limit int) int {
	return polyLoop(intMonoidish, limit, 0)
}
