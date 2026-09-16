-module(test_arrayOpsFFICheatcode@foreign).
-export([runArrayOpsFFICheatcode/1]).
runArrayOpsFFICheatcode(N) ->
  Step = case N >= 1 of true -> 1; false -> -1 end,
  sum_evens(1, N + Step, Step, 0).
sum_evens(Stop, Stop, _Step, Acc) -> Acc;
sum_evens(I, Stop, Step, Acc) ->
  Next = case I rem 2 of 0 -> Acc + I; _ -> Acc end,
  sum_evens(I + Step, Stop, Step, Next).
