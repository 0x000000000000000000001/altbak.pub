-module(control_bind@foreign).
-export([arrayBind/2]).

arrayBind(Arr, F) ->
  array:from_list(lists:append(array:to_list(array:map(fun (_, X) -> array:to_list(F(X)) end, Arr)))).
