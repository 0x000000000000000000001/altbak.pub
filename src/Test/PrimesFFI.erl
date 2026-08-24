-module(test_primesFFI@foreign).
-export([runPrimesFFI/1]).
divides(D, N) -> N rem D == 0.
any_divides(_N, { nil }) -> false;
any_divides(N, { cons, X, Xs }) ->
  case divides(X, N) of
    true -> true;
    false -> any_divides(N, Xs)
  end.
primes(Limit, Curr, Acc) when Curr > Limit -> Acc;
primes(Limit, Curr, Acc) ->
  case any_divides(Curr, Acc) of
    true -> primes(Limit, Curr + 1, Acc);
    false -> primes(Limit, Curr + 1, { cons, Curr, Acc })
  end.
sum_list({ nil }) -> 0;
sum_list({ cons, X, Xs }) -> X + sum_list(Xs).
runPrimesFFI(N) -> sum_list(primes(N, 2, { nil })).
