-module(test_stateMonadFFI@foreign).
-export([runStateMonadFFI/1]).
bind(M, F) -> fun(S) ->
  { A, S1 } = M(S),
  (F(A))(S1)
end.
pure(A) -> fun(S) -> { A, S } end.
get_state() -> fun(S) -> { S, S } end.
put_state(S) -> fun(_) -> { 0, S } end.
modify(F) -> bind(get_state(), fun(S) -> put_state(F(S)) end).
loop(0) -> pure(0);
loop(N) -> bind(modify(fun(S) -> S + 1 end), fun(_) -> loop(N - 1) end).
runStateMonadFFI(Depth) -> run_many(20, Depth, 0).
run_many(0, _Depth, Acc) -> Acc;
run_many(N, Depth, Acc) ->
  {_, State} = (loop(Depth))(0),
  run_many(N - 1, Depth, Acc + State).
