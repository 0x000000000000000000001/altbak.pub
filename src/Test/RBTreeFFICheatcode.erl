-module(test_rBTreeFFICheatcode@foreign).
-export([runRBTreeFFICheatcode/1]).
runRBTreeFFICheatcode(N) -> 
  Tree = lists:foldl(fun(X, T) -> gb_trees:enter(X, X, T) end, gb_trees:empty(), lists:seq(1, N)),
  gb_trees:size(Tree). % Just a dummy O(1) op to avoid returning the whole tree
