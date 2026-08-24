-module(test_ackermannFFI@foreign).
-export([runAckermannFFI/1]).
runAckermannFFI(_Dummy) -> ack(3, 4).
ack(0, N) -> N + 1;
ack(M, 0) when M > 0 -> ack(M - 1, 1);
ack(M, N) -> ack(M - 1, ack(M, N - 1)).
