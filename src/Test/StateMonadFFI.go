package Test_StateMonadFFI

type stateResult[S, A any] struct {
	val   A
	state S
}

type State[S, A any] func(S) stateResult[S, A]

func runState[S, A any](action State[S, A], initial S) stateResult[S, A] {
	return action(initial)
}

func bindState[S, A, B any](action State[S, A], next func(A) State[S, B]) State[S, B] {
	return func(state S) stateResult[S, B] {
		first := action(state)
		return next(first.val)(first.state)
	}
}

func pureState[S, A any](value A) State[S, A] {
	return func(state S) stateResult[S, A] { return stateResult[S, A]{val: value, state: state} }
}

func get[S any]() State[S, S] {
	return func(state S) stateResult[S, S] { return stateResult[S, S]{val: state, state: state} }
}

func put[S any](state S) State[S, struct{}] {
	return func(_ S) stateResult[S, struct{}] { return stateResult[S, struct{}]{val: struct{}{}, state: state} }
}

func modify[S any](f func(S) S) State[S, struct{}] {
	return bindState(get[S](), func(state S) State[S, struct{}] { return put(f(state)) })
}

func chainModifications(n int) State[int, struct{}] {
	if n == 0 {
		return pureState[int](struct{}{})
	}
	return bindState(modify(func(x int) int { return x + 1 }), func(_ struct{}) State[int, struct{}] {
		return chainModifications(n - 1)
	})
}

func runManyTimes_StateMonad(n int, depth int, acc int) int {
	if n == 0 {
		return acc
	}
	return runManyTimes_StateMonad(n-1, depth, acc+runState(chainModifications(depth), 0).state)
}

func RunStateMonadFFI(limit int) int {
	// The shared wrapper supplies depth 60; PureScript performs 20 repetitions.
	return runManyTimes_StateMonad(20, limit, 0)
}
