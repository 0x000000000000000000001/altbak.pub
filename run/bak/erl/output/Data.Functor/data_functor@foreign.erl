-module(data_functor@foreign).
-export([arrayMap/2]).

arrayMap(F, Arr) ->
  array:map(fun (_,X) -> F(X) end, Arr).
