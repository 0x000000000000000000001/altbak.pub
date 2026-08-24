-module(test_astTreeFFICheatcode@foreign).
-export([runAstTreeFFICheatcode/1]).
runAstTreeFFICheatcode(N) -> eval(buildTree(N)).
buildTree(0) -> 1;
buildTree(N) -> { add, { mul, N, buildTree(N - 1) }, { sub, buildTree(N - 1), 1 } }.
eval(N) when is_integer(N) -> N;
eval({ add, A, B }) -> eval(A) + eval(B);
eval({ mul, A, B }) -> eval(A) * eval(B);
eval({ sub, A, B }) -> eval(A) - eval(B).
