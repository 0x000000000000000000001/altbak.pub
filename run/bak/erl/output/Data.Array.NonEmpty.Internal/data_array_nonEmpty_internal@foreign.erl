-module(data_array_nonEmpty_internal@foreign).
-export([foldl1Impl/2, foldr1Impl/2, traverse1Impl/4]).


foldl1Impl(F, A) ->
    Length = array:size(A),
    foldl1Impl_(F, A, array:get(0, A), 1, Length).

foldl1Impl_(F, A, Acc, I, N) when I < N ->
    foldl1Impl_(F, A, (F(Acc))(array:get(I, A)), I+1, N);
foldl1Impl_(_,_,Acc,_,_) -> Acc.

foldr1Impl(F, A) ->
    Length = array:size(A),
    foldr1Impl_(F, A, array:get(Length-1, A), Length-2).

foldr1Impl_(F, A, Acc, I) when I >= 0 ->
    foldr1Impl_(F, A, (F(array:get(I, A)))(Acc), I-1);
foldr1Impl_(_,_,Acc,_) -> Acc.

traverse1Impl(Apply, Map, F, A) ->
    Len = array:size(A),
    Last = array:get(Len-1, A),
    Cons = fun (X) -> fun (XS) -> [X|XS] end end,
    Acc = (Map(fun (X) -> [X] end))(Last),
    BuildFrom = fun (_, X, YS) ->
        (Apply( (Map(Cons))(F(X)) ))(YS)
      end,
    ResultList = array:foldr(BuildFrom, Acc, array:resize(Len-1, A)),
        (Map(fun array:from_list/1))(ResultList).
