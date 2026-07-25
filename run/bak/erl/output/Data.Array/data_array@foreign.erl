-module(data_array@foreign).
-export([range/2,replicate/2,fromFoldableImpl/2,length/1,snoc/2,unconsImpl/3,indexImpl/4,findMapImpl/4,findIndexImpl/4,findLastIndexImpl/4,'_insertAt'/5,'_deleteAt'/4,'_updateAt'/5,reverse/1,concat/1,filter/2,partition/2,scanl/3,scanr/3,intersperse/2,sortByImpl/3,slice/3,zipWith/3,unsafeIndexImpl/2,any/2,all/2]).

% ------------------------------------------------------------------------------
%  Array creation --------------------------------------------------------------
% ------------------------------------------------------------------------------

range(Start, End) -> array:from_list(lists:seq(Start, End, case End < Start of true -> -1; false -> 1 end)).

replicate(N, X) -> array:new(max(N,0), {default, X}).

fromFoldableImpl(Foldr,Xs) -> array:from_list( ((Foldr(fun (H) -> fun (T) -> [H|T] end end))([]))(Xs) ).


% ------------------------------------------------------------------------------
%  Array size ------------------------------------------------------------------
% ------------------------------------------------------------------------------

length(Xs) -> array:size(Xs).

% ------------------------------------------------------------------------------
%  Extending arrays ------------------------------------------------------------
% ------------------------------------------------------------------------------

snoc(L,E) -> array:set(array:size(L), E, L).

% ------------------------------------------------------------------------------
%  Non-indexed reads -----------------------------------------------------------
% ------------------------------------------------------------------------------

unconsImpl(Empty, Next, Xs) ->
  case array:size(Xs) of
    0 -> Empty(unit);
    _ -> (Next(array:get(0, Xs)))(array:from_list(tl(array:to_list(Xs))))
  end.

% ------------------------------------------------------------------------------
%  Indexed operations ----------------------------------------------------------
% ------------------------------------------------------------------------------

indexImpl(Just,Nothing,Xs,I) ->
  case I < 0 orelse I >= array:size(Xs) of
    true -> Nothing;
    false -> Just(array:get(I,Xs))
  end.

findMapImpl(Nothing,IsJust,F,Xs) ->
  begin
    N = array:size(Xs) - 1,
    Find = fun
      Find(I) when I>N -> Nothing;
      Find(I) -> 
        Z = F(array:get(I,Xs)),
        case IsJust(Z) of
          true -> Z;
          false -> Find(I+1)
      end
    end,
    Find(0)
  end.

findIndexImpl(Just,Nothing,F,Xs) ->
  begin
    N = array:size(Xs) - 1,
    Find = fun
      Find(I) when I>N -> Nothing;
      Find(I) -> case F(array:get(I,Xs)) of
        true -> Just(I);
        false -> Find(I+1)
      end
    end,
    Find(0)
  end.

findLastIndexImpl(Just,Nothing,F,Xs) ->
  begin
    Find = fun
      Find(I) when I < 0 -> Nothing;
      Find(I) -> case F(array:get(I,Xs)) of
        true -> Just(I);
        false -> Find(I-1)
      end
    end,
    Find(array:size(Xs) - 1)
  end.

'_insertAt'(Just,Nothing,I,A,L) ->
  Size = array:size(L),
  InsertF = fun
    InsertF(LL, J) when J > Size -> LL;
    InsertF(LL, J) when J =:= I -> InsertF(array:set(J, A, LL), J+1);
    InsertF(LL, J) when J > I -> InsertF(array:set(J, array:get(J-1, L), LL), J+1)
  end,
  case I < 0 orelse I > Size of
    true -> Nothing;
    false -> Just(InsertF(L, I))
  end.

'_deleteAt'(Just,Nothing,I,L) ->
  Size = array:size(L),
  DeleteF = fun
    DeleteF(LL, J) when J-1 > Size -> LL;
    DeleteF(LL, J) when J >= I -> DeleteF(array:set(J, array:get(J+1, L), LL), J+1)
  end,
  case I < 0 orelse I > Size - 1 of
    true -> Nothing;
    false -> Just(array:resize(Size-1, DeleteF(L, I)))
  end.

'_updateAt'(Just,Nothing,I,A,L) ->
  case I < 0 orelse I >= array:size(L) of
    true -> Nothing;
    false -> Just(array:set(I, A, L))
  end.

%%------------------------------------------------------------------------------
%% Transformations -------------------------------------------------------------
%%------------------------------------------------------------------------------

reverse(L) -> array:from_list(lists:reverse(array:to_list(L))).

concat(Ls) -> array:from_list(lists:append(array:to_list(array:map(fun (_, X) -> array:to_list(X) end, Ls)))).

filter(F,L) -> array:from_list([X || X <- array:to_list(L), F(X)]).

partition(F,A) -> begin
  {Yes, No} = lists:partition(F, array:to_list(A)),
  #{yes => array:from_list(Yes), no => array:from_list(No)}
end.

scanl(F,B,Xs) -> 
  { Res, _ } = array:foldl(fun (I,V,{Arr, Acc}) -> 
                                  Acc1 = (F(Acc))(V),
                                  { array:set(I, Acc1, Arr), Acc1 }
                           end,
                           { array:new(array:size(Xs)), B},
                           Xs),
  Res.

scanr(F,B,Xs) ->
  { Res, _ } = array:foldr(fun (I,V,{Arr, Acc}) ->
                                  Acc1 = (F(V))(Acc),
                                  { array:set(I, Acc1, Arr), Acc1 }
                           end,
                           { array:new(array:size(Xs)), B},
                           Xs),
  Res.

intersperse(X,A) -> array:from_list(lists:join(X,array:to_list(A))).

% ------------------------------------------------------------------------------
%  Sorting ---------------------------------------------------------------------
% ------------------------------------------------------------------------------



% exports.sortByImpl = (function () {
%   function mergeFromTo(compare, fromOrdering, xs1, xs2, from, to) {
%     var mid;
%     var i;
%     var j;
%     var k;
%     var x;
%     var y;
%     var c;

%     mid = from + ((to - from) >> 1);
%     if (mid - from > 1) mergeFromTo(compare, fromOrdering, xs2, xs1, from, mid);
%     if (to - mid > 1) mergeFromTo(compare, fromOrdering, xs2, xs1, mid, to);

%     i = from;
%     j = mid;
%     k = from;
%     while (i < mid && j < to) {
%       x = xs2[i];
%       y = xs2[j];
%       c = fromOrdering(compare(x)(y));
%       if (c > 0) {
%         xs1[k++] = y;
%         ++j;
%       }
%       else {
%         xs1[k++] = x;
%         ++i;
%       }
%     }
%     while (i < mid) {
%       xs1[k++] = xs2[i++];
%     }
%     while (j < to) {
%       xs1[k++] = xs2[j++];
%     }
%   }

%   return function (compare) {
%     return function (fromOrdering) {
%       return function (xs) {
%         var out;

%         if (xs.length < 2) return xs;

%         out = xs.slice(0);
%         mergeFromTo(compare, fromOrdering, out, xs.slice(0), 0, xs.length);

%         return out;
%       };
%     };
%   };
% })();

sortByImpl(Compare, FromOrdering, Xs) -> 
  array:from_list(lists:sort(fun (X, Y) -> case FromOrdering((Compare(X))(Y)) of 1 -> false; _ -> true end end, array:to_list(Xs))).

% ------------------------------------------------------------------------------
%  Subarrays -------------------------------------------------------------------
% ------------------------------------------------------------------------------

slice_index_(I,L) -> 
  case I < 0 of 
    true -> array:size(L) + I;
    false -> I
  end.

slice(S,E,L) ->
  S1 = slice_index_(S,L),
  E1 = slice_index_(E,L),
  case E1 < S1 of
    true -> array:new(0);
    false -> array:from_list(lists:sublist(array:to_list(L), S+1, E-S))
  end.

% ------------------------------------------------------------------------------
%  Zipping ---------------------------------------------------------------------
% ------------------------------------------------------------------------------

zipWith(F,Xs,Ys) -> array:from_list(lists:zipwith(fun (X, Y) -> (F(X))(Y) end, array:to_list(Xs), array:to_list(Ys))).

unsafeIndexImpl(A,N) -> array:get(N,A).


%%------------------------------------------------------------------------------
%% Folding ---------------------------------------------------------------------
%%------------------------------------------------------------------------------

any(P,Xs) -> 
  N = array:size(Xs) - 1,
  Any = fun
    Any(I) when I>N -> false;
    Any(I) -> 
      case P(array:get(I,Xs)) of
        true -> true;
        false -> Any(I+1)
    end
  end,
  Any(0).

all(P,Xs) -> 
  N = array:size(Xs) - 1,
  All = fun
    All(I) when I>N -> true;
    All(I) -> 
      case P(array:get(I,Xs)) of
        true -> All(I+1);
        false -> false
    end
  end,
  All(0).
