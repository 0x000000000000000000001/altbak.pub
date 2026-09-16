-module(test_arrayOpsFFI@foreign).
-export([runArrayOpsFFI/1]).
runArrayOpsFFI(N) ->
  Step = case N >= 1 of true -> 1; false -> -1 end,
  Input = array:from_list(lists:seq(1, N, Step)),
  Filtered = array:foldl(fun(_I, X, Acc) ->
    case X rem 2 of
      0 -> array:set(array:size(Acc), X, Acc);
      _ -> Acc
    end
  end, array:new(), Input),
  array:foldl(fun(_I, X, Acc) -> Acc + X end, 0, Filtered).
