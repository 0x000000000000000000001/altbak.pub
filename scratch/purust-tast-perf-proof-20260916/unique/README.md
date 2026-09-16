# Experiment: eliminate proved-redundant Rc uniqueness checks

Production repositories were only read. `prepare.py` extracts the current generated RBTree kernel and replaces 103 syntactic occurrences of `std::rc::Rc::get_mut` with `assumed_unique_mut`. It retains all insertion, balancing, rebuilding, allocation, clone, and depth bodies. The two include-compatible kernels have the same `Tree` layout and standard `Rc` representation. Both remove the same three unrelated functions (`Test_RBTree_T`, `describe`, `act`) to avoid linking the generated application dependencies; the hot kernel uses direct enum construction, not the removed constructor adapter.

`kernel_unique.rs` changes one decision: nonempty cells are assumed exclusively owned and immediately yield a mutable payload; empty cells still use native `Rc::get_mut`. The empty guard matters: the generated algorithm shares `E` leaves and also temporarily replaces unique consumed cells with `E`. This experiment leaves those two cases to the existing dynamic test.

The specialized helper contains unsafe Rust, confined to this scratch binary. Its assumptions are **stronger than current `usageCount + escapes`**:

1. The tree is locally built from `E`; no old tree version, subtree alias, or Weak pointer survives into a mutating call.
2. Every nonempty subtree has exclusive ownership, including disjoint siblings.
3. All transfers, rotations, and generated read temporaries preserve that invariant at each mutable call; no immutable borrow remains live across mutation.
4. FFI or unknown calls cannot retain references to the tree.

This is an experimental worker with a restricted input domain, **not an implementation for arbitrary persistent trees**. The general callable function would need to retain its dynamic fallback or dispatch to a separately proven worker. It must not use this helper for persistent snapshots. Timing it estimates the value of such a proof; it does not establish that the compiler already produces the proof.

## Reproduce correctness

```sh
python3 prepare.py
rustc --edition=2021 -O --cfg verify_unique check.rs -o check
./check > correctness.tsv
```

The validation build checks `strong_count == 1` and `weak_count == 0` at **every** specialized nonempty mutation before the unsafe conversion. It also checks exact sorted keys, BST order, equal black heights, no red-red edge, black root, recursive child uniqueness, and the generated depth result. Four three-key rotations, 100,000 descending insertions, 100,000 ascending insertions, a 100,000-key deterministic permutation, 100,000 insertions with 4,096 distinct keys, and the original `buildTree` entry point passed.

| Insertion order | Mutation sites executed | Empty cells retaining dynamic check | Final depth |
| --- | ---: | ---: | ---: |
| Descending 100k | 2,483,932 | 0 | 22 |
| Ascending 100k | 2,783,866 | 299,934 | 22 |
| Shuffled 100k | 1,956,041 | 146,697 | 24 |
| Duplicates 100k | 1,349,338 | 104,307 | 17 |

Assertions demonstrate the assumptions on these executions, not on all possible programs. Compile timing variants **without** `--cfg verify_unique`; the correctness build includes counters and cannot produce representative timings. The coordinating agent runs timing binaries sequentially using an identical harness.

## Additional site-provenance variant

`kernel_unique_sites.rs` also removes the empty-payload guard. Its stronger contract is that **every cell reaching any mutable site is exclusive**, including the transient `E` installed by `__purust_take`. The actually shared empty leaves must never reach those sites. An implementation would require provenance of the consumed/reusable cell, not just a check that the payload constructor is nonempty.

`check_sites.rs` validates `strong_count == 1` and `weak_count == 0` on **every** mutable access, including `E`. Compiling with `rustc --edition=2021 -O --cfg verify_unique check_sites.rs -o check_sites` and running `./check_sites > correctness_sites.tsv` passed the same full fixture collection. This evidence supports the stronger precondition for these executions; the compiler proof remains unimplemented. It permits the final timing comparison to remove the extra tag-test cost introduced by the initial experimental helper.

## Addable Haskell information required

The current `Ann` is a five-tuple with optional `(Int, Bool)` usage data (`CoreFn/Ann.hs:42`). `CoreFn/Usage.hs:112` treats both the function and argument of any `App` as escaping. It does not compute return aliasing, parameter ownership modes, or a heap separation relation. Consequently a `False` flag cannot be reinterpreted as a deep-ownership certificate.

The existing Haskell structures can support a conservative additional analysis:

- `Expr` exposes application, lambdas, cases, lets, constructors, and typed applications (`CoreFn/Expr.hs:19`). `DataDecl` provides constructors and field types (`CoreFn/Module.hs:14`). These are sufficient to identify tree-valued fields and track how case-bound fields flow into a return.
- Add function summaries such as `readsOnly(parameter)`, `retains(parameter)`, `resultAliases(parameter/fieldPath)`, `returnsFresh`, and **conditional** `preservesDisjointOwnership(parameters)`. The first useful targeted proof would say: `insert` preserves exclusive ownership of nonempty nodes **provided its tree input was already exclusive**, while `depth` only reads it.
- A case on an exclusively owned `T` may split its ownership token into disjoint tokens for its two children and a token for the reusable parent cell. The constructors in `balance` use each child token once. Reassembling these tokens preserves exclusive ownership. This is a relational/field-sensitive analysis, not a single `unique` boolean deduced from a usage count.
- Recursive `ins`/`buildTree` require conservative fixed-point summaries; unknown and foreign functions default to unknown/retaining. For this benchmark the complete recursive cluster is in one module, so no cross-module optimization is required for the first proof.
- Immutable/shared nullary `E` is an explicit exception. Exclusive nonempty nodes must not be confused with an entirely unshared heap.
- `Make/Actions.hs:256` already invokes `Usage.computeUsage`; an ownership-summary pass can run alongside it. `CoreFn/ToJSON.hs:135` serializes annotations; summaries can be emitted separately at module/function level to avoid wrapping every AST node.
- `CoreFn/Desugar.hs:308` converts type structure; it should not infer dynamic uniqueness there. Ownership is a flow analysis over expressions and calls after lowering.
- PBO and Purust must invalidate or recompute proofs after inlining, duplication, lambda lifting, and newly inserted clones. A Haskell source-level proof alone does not authorize unchecked mutation in arbitrary transformed Rust.

Candidate payload, illustrative and not implemented: `parameterEffects`, `resultOrigins`, `ownershipPreconditions`, `preservesDisjointness`, plus stable binder/occurrence identities. Each certificate needs a defined proof phase and conservative unknown state. Start with borrowing summaries and single-module ownership specializations before general higher-order proofs.
