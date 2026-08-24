-module(test_recordsFFICheatcode@foreign).
-export([runRecordsFFICheatcode/1]).
runRecordsFFICheatcode(N) -> updateRec(N, 0, 0, 0, 0, 0).
updateRec(0, _A, _C, _E, F, _N) -> F;
updateRec(N, A, C, E, F, _OldN) ->
  updateRec(N - 1, A + 1, C + 2, E + 3, F + (N rem 5), N).
