-module(test_primesFFICheatcode@foreign).
-export([runPrimesFFICheatcode/1]).
runPrimesFFICheatcode(N) -> primes(N, 2, []).
primes(Limit, Curr, Acc) when Curr > Limit -> sum(Acc, 0);
primes(Limit, Curr, Acc) ->
  case any_divides(Curr, Acc) of
    true -> primes(Limit, Curr + 1, Acc);
    false -> primes(Limit, Curr + 1, [Curr | Acc])
  end.
any_divides(_N, []) -> false;
any_divides(N, [X | Xs]) ->
  if N rem X == 0 -> true;
     true -> any_divides(N, Xs)
  end.
sum([], Acc) -> Acc;
sum([X | Xs], Acc) -> sum(Xs, Acc + X).
