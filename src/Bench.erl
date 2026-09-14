-module(bench@foreign).
-export([benchNow/0, opaque/1, formatNumber/1]).

benchNow() -> fun() ->
  erlang:convert_time_unit(erlang:monotonic_time() - erlang:system_info(start_time),
                          native, microsecond) * 1.0
end.
opaque(A) -> fun() -> A end.
formatNumber(N) -> erlang:float_to_binary(N, [{decimals, 2}]).
