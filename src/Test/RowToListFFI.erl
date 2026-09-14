-module(test_rowToListFFI@foreign).
-export([runRowToListFFI/1]).
nil_dict() -> #{keys => fun(_Proxy) -> 0 end}.
cons_dict(Tail) -> #{keys => fun(_Proxy) -> 1 + (maps:get(keys, Tail))(proxy) end}.
runRowToListFFI(_Ignored) ->
  Dict = cons_dict(cons_dict(cons_dict(cons_dict(cons_dict(nil_dict()))))),
  (maps:get(keys, Dict))(proxy).
