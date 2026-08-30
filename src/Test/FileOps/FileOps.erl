-module(test_fileOps@foreign).
-export([writeFileSync/1, readFileSync/1, loopE/1]).

writeFileSync(Path) ->
    fun(Content) ->
        fun() ->
            ok = file:write_file(Path, Content),
            {}
        end
    end.

readFileSync(Path) ->
    fun() ->
        {ok, Binary} = file:read_file(Path),
        Binary
    end.

loopE(0) ->
    fun(_Action) -> fun() -> {} end end;
loopE(N) ->
    fun(Action) ->
        fun() ->
            loop_erl(N, Action),
            {}
        end
    end.

loop_erl(0, _) -> ok;
loop_erl(N, Action) ->
    Action(),
    loop_erl(N - 1, Action).
