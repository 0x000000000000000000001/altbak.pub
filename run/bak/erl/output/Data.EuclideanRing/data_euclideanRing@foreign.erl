-module(data_euclideanRing@foreign).
-export([intDegree/1, intDiv/2, intMod/2, numDiv/2]).

intDegree(X) -> abs(X).

intDiv(_, 0) -> 0;
intDiv(X, Y) when Y > 0 -> floor_(X / Y);
intDiv(X, Y) -> -floor_(X / -Y).

% Not supported natively until erlang 20
floor_(X) ->
    T = erlang:trunc(X),
    case (X - T) of
        Neg when Neg < 0 -> T - 1;
        Pos when Pos > 0 -> T;
        _ -> T
    end.

intMod(_, 0) -> 0;
intMod(X, Y) -> 
    YY = erlang:abs(Y),
    (X rem YY + YY) rem YY.

numDiv(X, Y) -> X / Y.
