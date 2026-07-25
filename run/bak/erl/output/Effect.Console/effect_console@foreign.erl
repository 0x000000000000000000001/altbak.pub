-module(effect_console@foreign).
-export([log/1,warn/1,error/1,info/1,clear/0, time/1, timeLog/1, timeEnd/1]).

log(S) -> fun () -> io:put_chars(S), io:nl() end.
warn(S) -> fun () -> io:put_chars(S), io:nl() end.
error(S) -> fun () -> io:put_chars(S), io:nl() end.
info(S) -> fun () -> io:put_chars(S), io:nl() end.

time(S) -> fun () -> erlang:put(S, erlang:system_time(microsecond)) end.

timeLog(S) -> fun () -> 
  case erlang:get(S) of
    undefined -> throw(io:format("No timer ~s", [S]));
    Start ->
      Diff = erlang:system_time(microsecond) - Start,
      io:format("~s: ~pms~n", [S, Diff/1000])
  end
end.

timeEnd(S) -> fun () -> 
  case erlang:get(S) of
    undefined -> throw(io:format("No timer ~s", [S]));
    Start ->
      erlang:erase(S),
      Diff = erlang:system_time(microsecond) - Start,
      io:format("~s: ~pms~n", [S, Diff/1000])
  end
end.

clear() -> fun () -> io:format("\ec") end.