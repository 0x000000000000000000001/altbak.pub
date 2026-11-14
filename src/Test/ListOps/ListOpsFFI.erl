-module(test_listOpsFFI@foreign).
-export([runListOpsFFI/1]).
range(Start, Curr, Acc) when Curr < Start -> Acc;
range(Start, Curr, Acc) -> range(Start, Curr - 1, { cons, Curr, Acc }).
filterEvens({ nil }, Acc) -> Acc;
filterEvens({ cons, X, Xs }, Acc) ->
  if X rem 2 == 0 -> filterEvens(Xs, { cons, X, Acc });
     true -> filterEvens(Xs, Acc)
  end.
foldl(_F, Acc, { nil }) -> Acc;
foldl(F, Acc, { cons, X, Xs }) -> foldl(F, F(Acc, X), Xs).
runListOpsFFI(N) -> foldl(fun(Acc, X) -> Acc + X end, 0, filterEvens(range(1, N, { nil }), { nil })).
