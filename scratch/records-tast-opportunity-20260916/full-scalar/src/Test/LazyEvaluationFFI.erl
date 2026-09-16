-module(test_lazyEvaluationFFI@foreign).
-export([runLazyEvaluationFFI/1]).
% The source Lazy is a thunk, without memoization.
build_thunks(0, Acc) -> Acc;
build_thunks(N, Acc) -> build_thunks(N - 1, fun() -> Acc() + 1 end).
run_many(0, Acc) -> Acc;
run_many(N, Acc) ->
  Thunk = build_thunks(1000, fun() -> 0 end),
  run_many(N - 1, Acc + Thunk()).
runLazyEvaluationFFI(N) -> run_many(N, 0).
