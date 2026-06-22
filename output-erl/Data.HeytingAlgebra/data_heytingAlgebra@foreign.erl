-module(data_heytingAlgebra@foreign).
-export([boolConj/2, boolDisj/2, boolNot/1]).

boolConj(A, B) -> A and B.
boolDisj(A, B) -> A or B.
boolNot(A) -> not A.
