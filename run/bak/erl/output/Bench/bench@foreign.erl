-module(bench@foreign).
-export([benchNow/0, opaque/1, formatNumber/1]).

benchNow() -> fun() -> 0.0 end.
opaque(A) -> fun() -> A end.
formatNumber(N) -> erlang:float_to_binary(N, [{decimals, 2}]).
