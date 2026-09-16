-module(test_churchFFICheatcode@foreign).
-export([runChurchFFICheatcode/1]).
% c100k = 10^5, so N * N * N * N * N
runChurchFFICheatcode(N) -> (N * N * N * N * N).
