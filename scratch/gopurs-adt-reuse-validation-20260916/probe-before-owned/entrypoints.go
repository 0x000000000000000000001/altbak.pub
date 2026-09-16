package main

import ps "gopurs/output/purescript"

// The build root is always empty; the same entry is selected by benchmark act.
func buildFresh(n int64) *ps.Constructor_Test_RBTree_T { return ps.Call_Test_RBTree_buildTree(n, nil) }

// The caller must supply exclusive nodes and retain no root/subtree snapshots.
func insertFresh(key int64, root *ps.Constructor_Test_RBTree_T) *ps.Constructor_Test_RBTree_T { return ps.Call_Test_RBTree_insert(key, root) }
