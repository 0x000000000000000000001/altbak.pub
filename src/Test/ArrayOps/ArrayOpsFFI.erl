-module(test_arrayOpsFFI@foreign).
-export([runArrayOpsFFI/1]).
filter_array(Predicate, Input) ->
  array:foldl(fun(_I, X, Acc) ->
    case Predicate(X) of
      true -> array:set(array:size(Acc), X, Acc);
      false -> Acc
    end
  end, array:new(), Input).
runArrayOpsFFI(N) ->
  Step = case N >= 1 of true -> 1; false -> -1 end,
  Input = array:from_list(lists:seq(1, N, Step)),
  Filtered = filter_array(fun(X) -> X rem 2 =:= 0 end, Input),
  array:foldl(fun(_I, X, Acc) -> Acc + X end, 0, Filtered).
