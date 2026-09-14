-module(test_ackermannFFICheatcode@foreign).
-export([runAckermannFFICheatcode/1]).
runAckermannFFICheatcode(M) -> ack(M, 4).
ack(0, N) -> N + 1;
ack(M, 0) -> ack(M - 1, 1);
ack(M, N) -> ack(M - 1, ack(M, N - 1)).
