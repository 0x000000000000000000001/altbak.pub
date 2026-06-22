-module(data_semiring@foreign).
-export([intAdd/2, intMul/2, numAdd/2, numMul/2]).

intAdd(X, Y) -> X + Y.
intMul(X, Y) -> X * Y.

numAdd(X, Y) -> X + Y.
numMul(X, Y) -> X * Y.
