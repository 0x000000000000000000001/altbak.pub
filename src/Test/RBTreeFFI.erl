-module(test_rBTreeFFI@foreign).
-export([runRBTreeFFI/1]).
insert(X, { leaf }) -> { node, r, { leaf }, X, { leaf } };
insert(X, { node, C, L, V, R }) ->
  if X < V -> balance(C, insert(X, L), V, R);
     X > V -> balance(C, L, V, insert(X, R));
     true  -> { node, C, L, X, R }
  end.
make_black({ node, _, L, V, R }) -> { node, b, L, V, R };
make_black({ leaf }) -> { leaf }.
balance(b, { node, r, { node, r, A, X, B }, Y, C }, Z, D) -> { node, r, { node, b, A, X, B }, Y, { node, b, C, Z, D } };
balance(b, { node, r, A, X, { node, r, B, Y, C } }, Z, D) -> { node, r, { node, b, A, X, B }, Y, { node, b, C, Z, D } };
balance(b, A, X, { node, r, { node, r, B, Y, C }, Z, D }) -> { node, r, { node, b, A, X, B }, Y, { node, b, C, Z, D } };
balance(b, A, X, { node, r, B, Y, { node, r, C, Z, D } }) -> { node, r, { node, b, A, X, B }, Y, { node, b, C, Z, D } };
balance(C, A, X, B) -> { node, C, A, X, B }.
insert_all(0, Acc) -> Acc;
insert_all(N, Acc) -> insert_all(N - 1, make_black(insert(N, Acc))).
find_depth({ leaf }) -> 0;
find_depth({ node, _, L, _, R }) -> 
  DL = find_depth(L),
  DR = find_depth(R),
  if DL > DR -> 1 + DL;
     true -> 1 + DR
  end.
runRBTreeFFI(N) -> find_depth(insert_all(N, { leaf })).
