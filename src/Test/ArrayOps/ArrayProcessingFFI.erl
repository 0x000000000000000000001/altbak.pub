-module(test_arrayProcessingFFI@foreign).
-export([runArrayOpsFFI/1]).
% Compatibility for the old filename; the canonical foreign module is ArrayOpsFFI.
runArrayOpsFFI(N) -> test_arrayOpsFFI@foreign:runArrayOpsFFI(N).
