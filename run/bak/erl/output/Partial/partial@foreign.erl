% module Partial
-module(partial@foreign).
-export(['_crashWith'/1]).

'_crashWith'(Msg) ->
  erlang:error(Msg).
