-module(data_foldable@foreign).
-export([foldrArray/3, foldlArray/3]).

foldrArray(F, Init, Xs) ->
  array:foldr(fun (_, X, A) -> (F(X))(A) end, Init, Xs).

foldlArray(F, Init, Xs) ->
  array:foldl(fun (_, X, A) -> (F(A))(X) end, Init, Xs).
