-module(test_fibFFICheatcode@foreign).
-export([runFibFFICheatcode/1]).
runFibFFICheatcode(N) when N < 2 -> N;
runFibFFICheatcode(N) -> runFibFFICheatcode(N - 1) + runFibFFICheatcode(N - 2).
