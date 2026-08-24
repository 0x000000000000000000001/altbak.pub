-module(test_tCOFFI@foreign).
-export([runTCOFFI/1]).
runTCOFFI(N) -> sum(N, 0).
sum(0, Acc) -> Acc;
sum(N, Acc) -> sum(N - 1, Acc + 1).
