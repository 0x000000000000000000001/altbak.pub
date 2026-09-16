-module(test_primesFFI@foreign).
-export([runPrimesFFI/1]).
range(Start, Curr, Acc) when Curr < Start -> Acc;
range(Start, Curr, Acc) -> range(Start, Curr - 1, {cons, Curr, Acc}).
reverse({nil}, Acc) -> Acc;
reverse({cons, X, Xs}, Acc) -> reverse(Xs, {cons, X, Acc}).
filter(_P, {nil}, Acc) -> reverse(Acc, {nil});
filter(P, {cons, X, Xs}, Acc) ->
  case P(X) of
    true -> filter(P, Xs, {cons, X, Acc});
    false -> filter(P, Xs, Acc)
  end.
sieve({nil}) -> {nil};
sieve({cons, P, Xs}) ->
  {cons, P, sieve(filter(fun(X) -> X rem P =/= 0 end, Xs, {nil}))}.
sum_list({nil}, Acc) -> Acc;
sum_list({cons, X, Xs}, Acc) -> sum_list(Xs, Acc + X).
runPrimesFFI(N) -> sum_list(sieve(range(2, N, {nil})), 0).
