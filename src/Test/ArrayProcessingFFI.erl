-module(test_arrayOpsFFI@foreign).
-export([runArrayOpsFFI/1]).
runArrayOpsFFI(N) ->
  Arr = array:from_list(lists:seq(1, N)),
  Filtered = array:foldr(fun(_I, X, Acc) -> 
    if X rem 2 == 0 -> array:set(array:size(Acc), X, Acc);
       true -> Acc
    end
  end, array:new(), Arr),
  array:foldl(fun(_I, X, Acc) -> Acc + X end, 0, Filtered).
