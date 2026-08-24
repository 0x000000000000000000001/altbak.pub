-module(test_tCOFFICheatcode@foreign).
-export([runTCOFFICheatcode/1]).
runTCOFFICheatcode(N) -> sum(N, 0).
sum(0, Acc) -> Acc;
sum(N, Acc) -> sum(N - 1, Acc + 1).
