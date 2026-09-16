# Safe ownership representation experiment

`kernel.rs` is a handwritten lowering of the same four-case Okasaki RBTree insertion algorithm. It uses `Link = Option<Owner<Node>>`, a mutable borrow during insertion, a shared borrow during depth traversal, and rotations which move existing cells and child links. It contains no unsafe code. All four balancing cases preserve the exact shape and colors of the generated kernel in the tested cases.

Two compile configurations share every algorithm body:

- `--cfg rc_owner`: `Owner<T> = Rc<T>`, exclusive access through `Rc::make_mut`.
- No cfg: `Owner<T> = Box<T>`, exclusive access through `Box::as_mut`.

The only differences are the owner alias and the `own()` helper. This isolates the effect of Rc versus Box **within this alternative lowering**. The difference from Purust's generated kernel combines several transformations: field borrowing, a different empty-leaf representation, reusable rotations, owned workers, and representation changes. Its total gain must not be attributed to a single TAST annotation or to already implemented compiler behavior.

For a uniquely built tree, each distinct key allocates one nonempty node. Empty leaves are `None` and allocate nothing. The generated baseline has one allocated shared `E` in addition to those nodes. The Box form also avoids the Rc allocation header; this is part of the representation experiment rather than an incidental omitted workload.

The timing interface is `pub fn run(n: i64) -> i64`: insert `n, n-1, ..., 1`, compute depth, drop the tree, return depth. It includes destruction. The parent agent supplies one identical timing harness and executes binaries sequentially.

Validation command:

```sh
rustc --edition=2021 -O check.rs -o check_box
./check_box > correctness_box.tsv
rustc --edition=2021 -O --cfg rc_owner check.rs -o check_rc
./check_rc > correctness_rc.tsv
```

`generated.rs` is a copy of the unchanged extracted baseline with only `crate::` qualified as `crate::generated::` so it can coexist in a module for comparison.

Both versions passed four small rotation cases, 100k descending insertions, 100k ascending insertions, a shuffled permutation of 100k distinct keys, and 100k insertions with 4096 distinct keys. Checks cover exact keys, BST ordering, root color, red-red exclusion, black height, depth, and the full pre-order shape with every node color/key against the generated baseline. The Rc version additionally preserves 200 snapshots while subsequent versions are built. The Box version deliberately has no persistent snapshot API: a compiler must prove the surrounding region requires no sharing, or retain a general Rc worker for shared inputs.
