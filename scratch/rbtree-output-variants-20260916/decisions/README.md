# Decision variants from current output

`prepare.py` reads the exact requested output and extracts the dependency-free tree kernel without changing its function bodies. The source SHA is recorded in `manifest.json`.

- `kernel-simple-guard.rs` changes only the child-rebuild guard. Its large boolean expression becomes a compact classifier for none/LL/LR/RL/RR; the existing permutation guard remains.
- `kernel-fused-guard.rs` also reuses that classification at the two post-child call sites. The existing LL permutation worker no longer repeats the LL color/constructor checks. Its `get_mut` uniqueness checks, field pattern matches and all generic shared/Weak fallbacks remain.

Important: the generated permutation worker is specialized only for LL. The other rotations retain their original generic implementation. These variants do not replace the tree algorithm, pointer representation, allocation strategy, depth traversal or ownership ABI. No unsafe code is introduced.

`check.rs` compares both variants against the original generated functions using the same enum representation. It checks exact shape, colors, keys and depth; all four minimal rotations; 100,000 ascending and descending keys; 10,000 deterministic shuffled keys; duplicates; 200 retained versions; root and child Weak references. It also exhaustively compares the original and simplified guard for all 722 relevant color/constructor combinations, including invalid red-black trees. These are correctness runs, not timing runs.

No new Haskell TAST fact is intrinsically required for these transformations. The constructor tests, their ordering and branch dominance are already present in the generated decision tree. A backend can share the decision or carry a branch discriminator into its specialized mutation worker. If such a discriminator were carried through the TAST/PBO, the required fact would be that the same unmodified fields satisfy the selected constructor/color pattern. It must be invalidated upon mutation. `usageCount + escapes` alone does not express that fact; static uniqueness is not required here because dynamic Rc checks are kept.

Run `python3 prepare.py` to recreate and validate. Parent orchestration performs sequential, uninstrumented timing with a common harness.
