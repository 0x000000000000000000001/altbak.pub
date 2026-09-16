-module(test_fibFFI@foreign).
-export([runFibFFI/1]).
runFibFFI(N) when N < 2 -> N;
runFibFFI(N) -> runFibFFI(N - 1) + runFibFFI(N - 2).
