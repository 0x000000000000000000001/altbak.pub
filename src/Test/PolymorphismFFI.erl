-module(test_polymorphismFFI@foreign).
-export([runPolymorphismFFI/1]).
runPolymorphismFFI(N) -> trunc(poly_loop(N, #{ <<"area">> => fun(R) -> math:pi() * maps:get(<<"radius">>, R) * maps:get(<<"radius">>, R) end }, #{ <<"radius">> => 10.0 }, 0.0)).
poly_loop(0, _Dict, _Shape, Acc) -> Acc;
poly_loop(N, Dict, Shape, Acc) -> poly_loop(N - 1, Dict, Shape, Acc + (maps:get(<<"area">>, Dict))(Shape)).
