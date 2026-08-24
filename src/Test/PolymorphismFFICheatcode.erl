-module(test_polymorphismFFICheatcode@foreign).
-export([runPolymorphismFFICheatcode/1]).
runPolymorphismFFICheatcode(N) -> trunc(poly_loop(N, 0.0)).
poly_loop(0, Acc) -> Acc;
poly_loop(N, Acc) -> poly_loop(N - 1, Acc + (math:pi() * 10.0 * 10.0)).
