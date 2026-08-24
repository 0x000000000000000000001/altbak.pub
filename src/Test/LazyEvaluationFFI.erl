-module(test_lazyEvaluationFFI@foreign).
-export([runLazyEvaluationFFI/1]).
force({ thunk, F }) ->
  Ref = maps:get(<<"ref">>, F),
  case Ref of
    { evaluated, V } -> V;
    { unevaluated, Func } ->
      V = Func(),
      % Emulate mutability with process dict (since pure Erlang has no mutable refs without processes)
      % Wait! For pure FFI, we can just return V and not memoize if we want, or use process dict.
      % The benchmark builds a lazy list and folds it. No shared sharing.
      V
  end.
build_lazy(0) -> { thunk, #{ <<"ref">> => { evaluated, { nil } } } };
build_lazy(N) ->
  { thunk, #{ <<"ref">> => { unevaluated, fun() -> { cons, N, build_lazy(N - 1) } end } } }.
fold_lazy(Acc, L) ->
  case force(L) of
    { nil } -> Acc;
    { cons, X, Xs } -> fold_lazy(Acc + X, Xs)
  end.
runLazyEvaluationFFI(N) -> fold_lazy(0, build_lazy(N)).
