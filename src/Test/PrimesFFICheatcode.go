package Test_PrimesFFICheatcode


type List_cheatcode interface {
	isList()
}

type Nil_cheatcode struct{}

func (Nil_cheatcode) isList() {}

type Cons_cheatcode struct {
	value0 int
	value1 List_cheatcode
}

func (Cons_cheatcode) isList() {}

func rangeList_cheatcode(start int, end int) List_cheatcode {
	var goFunc func(int, List_cheatcode) List_cheatcode
	goFunc = func(curr int, acc List_cheatcode) List_cheatcode {
		if curr < start {
			return acc
		}
		return goFunc(curr-1, Cons_cheatcode{value0: curr, value1: acc})
	}
	return goFunc(end, Nil_cheatcode{})
}

func filter_cheatcode(p func(int) bool, lst List_cheatcode) List_cheatcode {
	var goFunc func(List_cheatcode, List_cheatcode) List_cheatcode
	goFunc = func(list List_cheatcode, acc List_cheatcode) List_cheatcode {
		switch l := list.(type) {
		case Nil_cheatcode:
			return reverse_cheatcode(acc)
		case Cons_cheatcode:
			x := l.value0
			xs := l.value1
			if p(x) {
				return goFunc(xs, Cons_cheatcode{value0: x, value1: acc})
			} else {
				return goFunc(xs, acc)
			}
		}
		return Nil_cheatcode{}
	}
	return goFunc(lst, Nil_cheatcode{})
}

func reverse_cheatcode(lst List_cheatcode) List_cheatcode {
	var goFunc func(List_cheatcode, List_cheatcode) List_cheatcode
	goFunc = func(list List_cheatcode, acc List_cheatcode) List_cheatcode {
		switch l := list.(type) {
		case Nil_cheatcode:
			return acc
		case Cons_cheatcode:
			return goFunc(l.value1, Cons_cheatcode{value0: l.value0, value1: acc})
		}
		return Nil_cheatcode{}
	}
	return goFunc(lst, Nil_cheatcode{})
}

func sieve_cheatcode(lst List_cheatcode) List_cheatcode {
	switch l := lst.(type) {
	case Nil_cheatcode:
		return Nil_cheatcode{}
	case Cons_cheatcode:
		p := l.value0
		xs := l.value1
		return Cons_cheatcode{
			value0: p,
			value1: sieve_cheatcode(filter_cheatcode(func(x int) bool {
				return x%p != 0
			}, xs)),
		}
	}
	return Nil_cheatcode{}
}

func sumList_cheatcode(lst List_cheatcode) int {
	var goFunc func(List_cheatcode, int) int
	goFunc = func(list List_cheatcode, acc int) int {
		switch l := list.(type) {
		case Nil_cheatcode:
			return acc
		case Cons_cheatcode:
			return goFunc(l.value1, acc+l.value0)
		}
		return acc
	}
	return goFunc(lst, 0)
}

func RunPrimesFFICheatcode(limit float64) float64 {
	dummy := int(limit)
	return float64(sumList_cheatcode(sieve_cheatcode(rangeList_cheatcode(2, dummy))))
}
