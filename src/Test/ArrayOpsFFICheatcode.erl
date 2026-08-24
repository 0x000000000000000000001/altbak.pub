-module(test_arrayOpsFFICheatcode@foreign).
-export([runArrayOpsFFICheatcode/1]).
runArrayOpsFFICheatcode(N) -> sum_evens(N, 0).
sum_evens(0, Acc) -> Acc;
sum_evens(N, Acc) ->
  if N rem 2 == 0 -> sum_evens(N - 1, Acc + N);
     true -> sum_evens(N - 1, Acc)
  end.
