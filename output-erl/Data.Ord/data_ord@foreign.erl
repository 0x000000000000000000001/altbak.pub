-module(data_ord@foreign).
-export([ordBooleanImpl/5, ordIntImpl/5, ordNumberImpl/5, ordStringImpl/5, ordCharImpl/5, ordArrayImpl/1]).

unsafeCompareImpl(LT, EQ, GT, X, Y) ->
  if
    X < Y -> LT;
    X > Y -> GT;
    true  -> EQ
  end.

ordBooleanImpl(LT, EQ, GT, X, Y) -> unsafeCompareImpl(LT, EQ, GT, X, Y).
ordIntImpl(LT, EQ, GT, X, Y) -> unsafeCompareImpl(LT, EQ, GT, X, Y).
ordNumberImpl(LT, EQ, GT, X, Y) -> unsafeCompareImpl(LT, EQ, GT, X, Y).
ordStringImpl(LT, EQ, GT, X, Y) -> unsafeCompareImpl(LT, EQ, GT, X, Y).
ordCharImpl(LT, EQ, GT, X, Y) -> unsafeCompareImpl(LT, EQ, GT, X, Y).


ordArrayImpl(F, [X|XS], [Y|YS]) ->
  case (F(X))(Y) of
    0 -> ordArrayImpl(F, XS, YS);
    N -> N
  end;
ordArrayImpl(_, [], []) -> 0;
ordArrayImpl(_, _, []) -> -1;
ordArrayImpl(_, [], _) -> 1.

ordArrayImpl(F) ->
  fun (XS) ->
    fun (YS) ->
      ordArrayImpl(F, array:to_list(XS), array:to_list(YS))
    end
  end.
