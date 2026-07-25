-module(data_unfoldable@foreign).
-export([unfoldrArrayImpl/6]).

unfoldrArrayImpl(IsNothing, FromJust, Fst, Snd, F, B) ->
  begin
    Loop1 = fun Loop(Result, BB) ->
        begin
          Maybe = F(BB),
          case IsNothing(Maybe) of
            true -> lists:reverse(Result);
            false -> begin
              Tuple = FromJust(Maybe),
              Loop([Fst(Tuple)|Result], Snd(Tuple))
            end
          end
        end
    end,
    array:from_list(Loop1([], B))
  end.
