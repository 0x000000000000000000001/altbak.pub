-module(test_churchFFI@foreign).
-export([runChurchFFI/1]).
church(0) -> fun(_F) -> fun(X) -> X end end;
church(N) -> 
  C = church(N - 1),
  fun(F) -> fun(X) -> F((C(F))(X)) end end.
unchurch(C) -> (C(fun(X) -> X + 1 end))(0).
mul(M, N) -> fun(F) -> fun(X) -> (M(N(F)))(X) end end.
c100(N) -> mul(church(N), church(N)).
c10k(N) -> mul(c100(N), c100(N)).
runChurchFFI(N) -> unchurch(mul(c10k(N), church(N))).
