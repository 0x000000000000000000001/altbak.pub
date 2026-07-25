-module(data_traversable@foreign).
-export([traverseArrayImpl/5]).

traverseArrayImpl(Apply, Map, Pure, F, Array) ->
  Cons = fun (X) -> fun (XS) -> [X|XS] end end,
  BuildFrom = fun (_, X, YS) ->
    (Apply( (Map(Cons))(F(X)) ))(YS)
  end,
  ResultList = array:foldr(BuildFrom, Pure([]), Array),
  (Map(fun array:from_list/1))(ResultList).
