-module(test_lazyEvaluationFFICheatcode@foreign).
-export([runLazyEvaluationFFICheatcode/1]).
runLazyEvaluationFFICheatcode(N) -> sum(N, 0).
sum(0, Acc) -> Acc;
sum(N, Acc) -> sum(N - 1, Acc + N).
