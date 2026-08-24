package Test_StateMonadFFI


type stateResult struct {
	val   interface{}
	state int
}

type State func(int) stateResult

func runState(s State, init int) stateResult {
	return s(init)
}

func bindState(s State, g func(interface{}) State) State {
	return func(state int) stateResult {
		r1 := s(state)
		gPrime := g(r1.val)
		return gPrime(r1.state)
	}
}

func pureState(a interface{}) State {
	return func(s int) stateResult {
		return stateResult{val: a, state: s}
	}
}

func get() State {
	return func(s int) stateResult {
		return stateResult{val: s, state: s}
	}
}

func put(s int) State {
	return func(_ int) stateResult {
		return stateResult{val: nil, state: s}
	}
}

func modify(f func(int) int) State {
	return bindState(get(), func(s interface{}) State {
		return put(f(s.(int)))
	})
}

func chainModifications(n int) State {
	if n == 0 {
		return pureState(nil)
	}
	return bindState(modify(func(x int) int { return x + 1 }), func(_ interface{}) State {
		return chainModifications(n - 1)
	})
}

func runManyTimes_StateMonad(n int, acc int) int {
	if n == 0 {
		return acc
	}
	return runManyTimes_StateMonad(n-1, acc+runState(chainModifications(60), 0).state)
}

func RunStateMonadFFI(limit int) int {
	dummy := limit
	return (runManyTimes_StateMonad(dummy, 0))
}
