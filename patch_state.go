package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
)

type StateTuple struct {
	state int64
	val   int64
}

func Call_Test_StateMonad_chainModifications_native(v_0_loop int64) func(int64) StateTuple {
	if v_0_loop == 0 {
		return func(s_1 int64) StateTuple {
			return StateTuple{state: s_1, val: 0}
		}
	}

	return func(s_2 int64) StateTuple {
		r1 := StateTuple{state: s_2 + 1, val: 0}
		g_prime := Call_Test_StateMonad_chainModifications_native(v_0_loop - 1)
		return g_prime(r1.state)
	}
}
