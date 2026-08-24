-module(test_fibFFI@foreign).
-export([runFibFFI/1]).

runFibFFI(N) ->
  if N < 2 -> N;
     true -> runFibFFI(N - 1) + runFibFFI(N - 2)
  end.
