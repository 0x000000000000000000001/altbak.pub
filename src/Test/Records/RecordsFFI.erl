-module(test_recordsFFI@foreign).
-export([runRecordsFFI/1]).
updateRec(0, R) -> R;
updateRec(N, R) ->
  B = maps:get(<<"b">>, R),
  D = maps:get(<<"d">>, B),
  NewD = D#{ <<"e">> => maps:get(<<"e">>, D) + 3,
             <<"f">> => maps:get(<<"f">>, D) + (N rem 5) },
  NewB = B#{ <<"c">> => maps:get(<<"c">>, B) + 2,
             <<"d">> => NewD },
  updateRec(N - 1, R#{ <<"a">> => maps:get(<<"a">>, R) + 1,
                       <<"b">> => NewB }).
runRecordsFFI(N) -> 
  R = #{ <<"a">> => 0, <<"b">> => #{ <<"c">> => 0, <<"d">> => #{ <<"e">> => 0, <<"f">> => 0 } } },
  Final = updateRec(N, R),
  maps:get(<<"f">>, maps:get(<<"d">>, maps:get(<<"b">>, Final))).
