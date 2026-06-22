-module(data_eq@foreign).
-export([eqBooleanImpl/2, eqIntImpl/2, eqNumberImpl/2, eqCharImpl/2, eqStringImpl/2, eqArrayImpl/3]).

eqBooleanImpl(A, B) -> A =:= B.
eqIntImpl(A, B) -> A =:= B.
eqNumberImpl(A, B) -> A =:= B.
eqCharImpl(A, B) -> A =:= B.
eqStringImpl(A, B) -> A =:= B.

eqArrayImpl@1(F, [X|Xs], [Y|Ys]) ->
  case (F(X))(Y) of
    true -> eqArrayImpl@1(F,Xs,Ys);
    false -> false
  end;
eqArrayImpl@1(_, [], []) -> true;
eqArrayImpl@1(_, _, _) -> false.

eqArrayImpl(F, Xs, Ys) -> eqArrayImpl@1(F, array:to_list(Xs), array:to_list(Ys)).
