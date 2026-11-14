-module(test_stateMonadFFICheatcode@foreign).
-export([runStateMonadFFICheatcode/1]).
runStateMonadFFICheatcode(N) ->
  put(state, 0),
  loop(N * 20),
  erase(state).
loop(0) -> ok;
loop(N) -> 
  put(state, get(state) + 1),
  loop(N - 1).
