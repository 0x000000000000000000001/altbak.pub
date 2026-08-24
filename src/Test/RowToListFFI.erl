-module(test_rowToListFFI@foreign).
-export([runRowToListFFI/1]).
runRowToListFFI(_Dummy) ->
  % We just count the keys of the record.
  R = #{ <<"a">> => 1, <<"b">> => <<"test">>, <<"c">> => true, <<"d">> => 4.0, <<"e">> => #{ <<"nested">> => true } },
  length(maps:keys(R)).
