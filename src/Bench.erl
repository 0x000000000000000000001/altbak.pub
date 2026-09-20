-module(bench@foreign).
-export([benchNow/0, opaque/1, formatNumber/1, measureBatch/3]).

benchNow() -> fun() ->
  erlang:convert_time_unit(erlang:monotonic_time() - erlang:system_info(start_time),
                          native, nanosecond) / 1000.0
end.
opaque(A) -> fun() -> A end.
formatNumber(N) -> erlang:float_to_binary(N, [{decimals, 6}]).
measureBatch(Iterations, Expected, Act) -> fun() ->
  Start = erlang:monotonic_time(),
  Result = measure_loop(Iterations, Act, 0),
  Elapsed = erlang:convert_time_unit(erlang:monotonic_time() - Start, native, nanosecond) / 1000.0,
  case Result =:= Expected of
    true -> Elapsed / Iterations;
    false -> erlang:error({unstable_benchmark_result, Result, Expected})
  end
end.
measure_loop(0, _Act, Result) -> Result;
measure_loop(N, Act, _Previous) ->
  Result = Act(),
  put(benchmark_result, Result),
  measure_loop(N - 1, Act, Result).
