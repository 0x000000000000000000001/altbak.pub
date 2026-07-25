-module(data_unfoldable1@foreign).
-export([unfoldr1ArrayImpl/6]).

unfoldr1ArrayImpl(IsNothing, FromJust, Fst, Snd, F, B) ->
  begin
    Loop1 = fun Loop(Result, BB) ->
        begin
          Tuple = F(BB),
          Result1 = [Fst(Tuple)|Result],
          Maybe = Snd(Tuple),
          case IsNothing(Maybe) of
            true -> lists:reverse(Result1);
            false -> Loop(Result1, FromJust(Maybe))
          end
        end
    end,
    array:from_list(Loop1([], B))
  end.
