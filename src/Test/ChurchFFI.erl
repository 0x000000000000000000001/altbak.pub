-module(test_churchFFI@foreign).
-export([runChurchFFI/1]).
church(0) -> fun(_F) -> fun(X) -> X end end;
church(N) -> 
  C = church(N - 1),
  fun(F) -> fun(X) -> F((C(F))(X)) end end.
unchurch(C) -> (C(fun(X) -> X + 1 end))(0).
mul(M, N) -> fun(F) -> fun(X) -> (M(N(F)))(X) end end.
runChurchFFI(N) -> 
  C10 = church(N),
  C100 = mul(C10, C10),
  C10k = mul(C100, C100),
  C100k = mul(C10k, C10),
  unchurch(C100k).
