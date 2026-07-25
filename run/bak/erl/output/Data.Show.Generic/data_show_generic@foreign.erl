-module(data_show_generic@foreign).
-export([intercalate/2]).

intercalate(S, Xs) ->
  XS1 = lists:map(fun unicode:characters_to_list/1, array:to_list(Xs)),
  Res = string:join(XS1, unicode:characters_to_list(S)),
  unicode:characters_to_binary(Res).
