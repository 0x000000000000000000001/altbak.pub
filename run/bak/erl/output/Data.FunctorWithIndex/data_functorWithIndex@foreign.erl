-module(data_functorWithIndex@foreign).
-export([mapWithIndexArray/2]).

mapWithIndexArray(F, Xs) ->
    array:map(fun (I, X) -> (F(I))(X) end, Xs).

