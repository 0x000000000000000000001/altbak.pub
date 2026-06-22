-module(effect@foreign).
-export([pureE/1,bindE/1,untilE/1,whileE/1,forE/1, foreachE/1]).

pureE(A) -> fun () -> A end.

bindE(A) ->
  fun(F) ->
    fun() ->
      (F(A()))()
    end
  end.

untilE(F) ->
  fun Loop() ->
    case F() of
      true -> unit;
      _ -> Loop()
    end
  end.

whileE(F) ->
  fun (A) ->
    fun Loop() ->
      case F() of
        true ->
          begin
            A(),
            Loop()
          end;
        _ -> unit
      end
    end
  end.


forE(Lo,Hi,F) when Lo < Hi ->
  begin
    (F(Lo))(),
    forE(Lo+1,Hi,F)
  end;
forE(_,_,_) -> unit.

forE(Lo) ->
  fun (Hi) ->
    fun (F) ->
      fun () ->
        forE(Lo,Hi,F)
      end
    end
  end.

foreachE(As) ->
  fun (F) ->
    fun () ->
      lists:foreach(fun (X) -> (F(X))() end, array:to_list(As))
    end
  end.
