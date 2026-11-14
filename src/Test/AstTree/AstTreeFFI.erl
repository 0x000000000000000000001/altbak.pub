-module(test_astTreeFFI@foreign).
-export([runAstTreeFFI/1]).
buildTree(0) -> { val, 1 };
buildTree(N) -> { add, { mul, { val, N }, buildTree(N - 1) }, { sub, buildTree(N - 1), { val, 1 } } }.
eval({ val, N }) -> N;
eval({ add, A, B }) -> eval(A) + eval(B);
eval({ mul, A, B }) -> eval(A) * eval(B);
eval({ sub, A, B }) -> eval(A) - eval(B).
runAstTreeFFI(N) -> eval(buildTree(N)).
