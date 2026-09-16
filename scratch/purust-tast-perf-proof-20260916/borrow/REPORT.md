# Borrowed recursive depth worker

This scratch experiment changes only `Test_RBTree_depth` in the freshly emitted
Rust. Every constructor, insertion, balancing, recoloring and build function is
byte-for-byte identical between `kernel-baseline.rs` and `kernel-borrow.rs`.

The transformation keeps the emitted depth branch, arithmetic and accessor
structure. Its recursive worker takes `&Rc<Tree>` and its two child accessors
return references instead of cloning. The public entrypoint still accepts and
consumes `Rc<Tree>`; it drops the root after the borrowed traversal returns.
Consequently a full build/depth/drop timing includes destruction in both cases.

`prepare.py` regenerates both kernels from the source whose SHA256 is recorded
in `counts.json`. It runs the original allocation-probe correctness harness on
each through the original one-word native-Rc counter wrapper. It never measures
time. Root orchestration performs all uninstrumented timings separately.

Both kernels pass: 100,000 ordered keys, depth 22, BST and red-black invariants,
all four rotations, exclusion of weak references from in-place mutation, 200
retained persistent snapshots with exact keys and balanced logical lifetimes.

| Logical operations | Generated baseline | Borrowed depth |
| --- | ---: | ---: |
| Build allocation count | 100,001 | 100,001 |
| Build `Rc` clones | 2,283,976 | 2,283,976 |
| Build successful `get_mut` | 2,483,932 | 2,483,932 |
| Depth `Rc` clones | 200,000 | 0 |
| Final payload destructions | 100,001 | 100,001 |

This demonstrates removal of 200,000 clone/release pairs, not a measured speedup
by itself. It requires a transitive read-only, non-retaining parameter summary:
all uses of the argument and its projections are pattern tests, scalar reads or
recursive calls with the same contract; no result retains an input pointer, and
no FFI/effect can observe reference ownership or invalidate the borrow. Current
`usageCount` and `escapes` fields alone do not encode that contract.

Such a summary can be inferred over recursive strongly connected components in
the compiler and serialized in TAST, provided later optimizations preserve or
revalidate it. Unknown/foreign calls must conservatively retain the owned ABI.
Borrowing also works for persistent shared roots: this experiment keeps the
same representations and fallback behavior for all construction paths.

## Mutable-slot insertion worker

`prepare-ins.py` adds a separate prototype, `kernel-borrow-ins.rs`, with original
depth; `kernel-borrow-ins-depth.rs` composes it with borrowed depth. Both retain
the same `Rc<Tree>` layout and public owned ABI. Unlike the first experiment,
this prototype changes the construction path. The added `ins_worker.rs` uses
`&mut Rc<Tree>` for recursive descent under the original dynamic `get_mut` check,
so it needs no temporary sibling clone to move an owned child out of its parent.
The original red/black test, generated reconstruction guard, generated field
permutation and generated generic reuse worker remain in use. The exact
original `ins` body is retained under a new name for shared and weak fallbacks.

This prototype requires a consumed-argument/result-slot contract: the argument
is transferred through the call and its replacement returned to the same slot;
the caller needs no surviving alias to that slot. Exclusivity of the current
cell still comes from `get_mut`, so these measurements do not demonstrate a
static uniqueness proof or removal of reference-count checks.

| Logical operations, unique build | Generated baseline | Mutable slot |
| --- | ---: | ---: |
| Allocations | 100,001 | 100,001 |
| `Rc` clones | 2,283,976 | 200,000 |
| Successful `get_mut` | 2,483,932 | 2,483,932 |
| Extra logical adapter calls | 0 | 100,000 |

Both variants pass the same invariants, exact persistent versions, weak tests
and balanced lifetimes. Their prototype fallback is deliberately conservative:
the 200-version persistence test performs 7,311 clones instead of 5,808, while
allocating the same 1,504 cells. The small checks phase also allocates 23 cells
instead of 21 because the weak fallback temporarily retains the original root.
Thus the experiment targets the unique hot path; it does not establish a gain
for persistent/shared workloads. Each new leaf clones its shared empty pointer
twice instead of once, accounting for the remaining 200,000 build clones.

`check-orders.py` additionally compiles all four uninstrumented kernels and
checks three independent input families without taking timings. All agree on
exact sorted keys, red-black invariants, black height and independently computed
depth: 100,000 ascending keys (black height 17, depth 22), a deterministic shuffle
of 100,000 keys (black height 14, depth 23), and 100,000 insertions containing
duplicates over 997 distinct keys (black height 8, depth 13). Raw results are in
`orders.log`.
