-module(test_polymorphismFFI@foreign).
-export([runPolymorphismFFI/1]).
runPolymorphismFFI(N) ->
  Dict = #{mempty => 1, mappend => fun(X) -> fun(Y) -> X + Y end end},
  poly_loop(N, Dict, 0).
poly_loop(0, _Dict, Acc) -> Acc;
poly_loop(N, Dict, Acc) ->
  Append = maps:get(mappend, Dict),
  poly_loop(N - 1, Dict, (Append(Acc))(maps:get(mempty, Dict))).
