-module(test_rowToListFFI@foreign).
-export([runRowToListFFI/1]).
nil_dict() -> #{keys => fun(_Proxy) -> 0 end}.
cons_dict(Tail) -> #{keys => fun(_Proxy) -> 1 + (maps:get(keys, Tail))(proxy) end}.
% Erlang has no static row constraints. Keep the recursive dictionary together
% with the heterogeneous record through each row constructor.
nil_row() -> {row, [], nil_dict()}.
cons_row(Name, Value, {row, Tail, Dictionary}) ->
  {row, [{Name, Value} | Tail], cons_dict(Dictionary)}.
keys({row, _Values, Dictionary}) -> (maps:get(keys, Dictionary))(proxy).
runRowToListFFI(_Ignored) ->
  Record = cons_row(a, 1, cons_row(b, <<"two">>, cons_row(c, true,
             cons_row(d, 4.0, cons_row(e, <<"five">>, nil_row()))))),
  keys(Record).
