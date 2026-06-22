-module(test_async@foreign).
-export([runAsyncTest/1]).

runAsyncTest(N) ->
    fun() ->
        Parent = self(),
        [spawn(fun() -> Parent ! done end) || _ <- lists:seq(1, N)],
        wait_for_workers(N),
        {}
    end.

wait_for_workers(0) -> ok;
wait_for_workers(N) ->
    receive
        done -> wait_for_workers(N - 1)
    end.
