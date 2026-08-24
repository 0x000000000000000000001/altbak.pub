-module(test_stateMonadFFI@foreign).
-export([runStateMonadFFI/1]).
bind(M, F) -> fun(S) ->
  { A, S1 } = M(S),
  (F(A))(S1)
end.
pure(A) -> fun(S) -> { A, S } end.
get() -> fun(S) -> { S, S } end.
put(S) -> fun(_) -> { 0, S } end.
modify(F) -> bind(get(), fun(S) -> put(F(S)) end).
loop(0) -> pure(0);
loop(N) -> bind(modify(fun(S) -> S + 1 end), fun(_) -> loop(N - 1) end).
runStateMonadFFI(N) ->
  { _, FinalState } = (loop(N * 20))(0),
  FinalState.
