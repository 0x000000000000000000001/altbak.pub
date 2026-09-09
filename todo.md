# Optimization techniques by backend

Audit dated **September 9, 2026**: **49 techniques across 11 backends**, cross-checked against their invoked passes and code generation paths. This matrix describes integration, coverage, and limitations. It serves as a work list; **it is not a ranking of backend speed**. The reference measurements remain those in the [README](README.md).

All eleven backends referenced by the runners are included. **Our five backends occupy the first columns: sharpurs, purust, phpurs, gopurs, and javapurs.** The other six follow on the right: psgo, official JS, ES, purescm, purerl, and Wasm, which is experimental across the 14 core cases. The README's “Native FP-style” and “Native hand-optimized” columns are handwritten implementations: their shortcuts are not attributed to the compilers.

## Rating legend

- 🟢 **Complete within the defined scope**: the mechanism is integrated and used by the relevant code generation path. This does not certify every possible program.
- 🟡 **Partial**: a significant restriction on types, expression shapes, local/intermodule scope, ABI, or code generation path. [Audit notes](optimization-audit.md#exact-scope-of-the-rows) explain the restriction.
- 🔴 **Missing**: no corresponding mechanism is integrated into the audited pipeline; a primitive, prototype, or unused module is insufficient.
- ⚪ **Gray — unnecessary / not applicable in this form**: the target's representation, value management, or native tail calls already cover the need. Gray cases are explained in the [audit notes](optimization-audit.md#exact-scope-of-the-rows).

**🟢 (PBO) / 🟡 (PBO)**: the shared PBO infrastructure provides a transformation or analysis actually used for this capability, sometimes complemented by the target code generator. The label means neither that PBO implements the entire technique on its own nor that all backends use it with the same coverage. It distinguishes this contribution from passes belonging to the backend, the PureScript frontend, or the runtime. For ES and purescm, it refers to their version of PBO, not our fork's TAST extensions.

A conservative proof of purity, arity, or non-escape is necessary for correctness; it is not inherently a shortcoming. “Partial” identifies a coverage limit worth knowing about. These cells should not be added up into a maturity percentage: some rows depend on the same passes and carry different weight.

## Matrix

**Row order defines performance-investigation priority, highest first**, across our five backends and the core benchmarks. Skip capabilities already complete for the backend being improved. The first ten rows cover broad transformations and their representation prerequisites; rows **11–15** group ADT layouts, FBIP, field-level reuse specialization, borrowing and Perceus-style reference-count optimization. Sticky sharing is last because no current hot-path benefit has been established. This is a qualitative order; each proposed gain still needs an isolated measurement. [Ordering rationale](optimization-audit.md#estimated-impact-order).

**Cell references:** backend columns are **A–K**, and technique rows are **1–49**. For example, **E11** means **javapurs × ADT fields specialized by payload type**. Coordinates follow the current table order; the audit uses separate stable technique IDs.

The headers link to each backend's evidence and limitations. **JS** refers to the official code generator; **ES** refers to Arista's `purs-backend-es`. **Sharpurs emits F#**, with C# interoperability, rather than translating PureScript function bodies into C#.

| # | Technique | **A**<br>[sharpurs](optimization-audit.md#sharpurs) | **B**<br>[purust](optimization-audit.md#purust) | **C**<br>[phpurs](optimization-audit.md#phpurs) | **D**<br>[gopurs](optimization-audit.md#gopurs) | **E**<br>[javapurs](optimization-audit.md#javapurs) | **F**<br>[psgo](optimization-audit.md#psgo) | **G**<br>[JS](optimization-audit.md#official-js) | **H**<br>[ES](optimization-audit.md#es) | **I**<br>[purescm](optimization-audit.md#purescm) | **J**<br>[purerl](optimization-audit.md#purerl) | **K**<br>[Wasm](optimization-audit.md#wasm) |
| :---: | --- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| 1 | Partial evaluation, β-reduction, simplification | 🟡 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 |
| 2 | Automatic inlining and dead code elimination | 🟡 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 |
| 3 | Specialization of higher-order functions / closures | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟡 |
| 4 | Automatic uncurrying / wrapper-worker | 🟡 | 🟡 | 🟢 | 🟡 | 🟡 | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🟢 |
| 5 | Elimination of known type class dictionaries | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🟡 | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🟡 |
| 6 | Monomorphization by cloning polymorphic functions | 🔴 | 🔴 | 🔴 | 🟡 (PBO) | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 7 | Using TAST to select representations | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 8 | Effective `TypeApp` type instantiation | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 9 | Keeping numeric parameters and locals unboxed | 🟡 | 🟡 | ⚪ | 🟡 | 🟡 | 🔴 | ⚪ | ⚪ | ⚪ | ⚪ | 🟡 |
| 10 | Direct emission of primitive operations | 🟡 (PBO) | 🟢 (PBO) | 🟡 (PBO) | 🟢 (PBO) | 🟡 (PBO) | 🟡 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 |
| 11 | ADT fields specialized by payload type | 🟡 | 🟡 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 |
| 12 | Memory reuse after proving ownership / uniqueness (FBIP) | 🔴 | 🟡 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 13 | ADT reuse specialization: write only changed fields | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 14 | Last-use moves / borrowing to avoid explicit reference counting | ⚪ | 🟡 | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ |
| 15 | Perceus: branch-local release specialization / dup-drop fusion | ⚪ | 🔴 | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ |
| 16 | Self-tail recursion with bounded stack usage | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 | 🟢 (PBO) | ⚪ | ⚪ | 🟡 |
| 17 | Loop invariants: hoisting / caching per invocation | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 18 | Dedicated fusion of a thunk producer with its consumer | 🔴 | 🟡 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 19 | Scrutinee fusion / case-of-case | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟡 |
| 20 | Scalar replacement of known / non-escaping records and arrays | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟡 |
| 21 | Closed records: specialized layout and fixed fields | 🔴 | 🟡 | 🔴 | 🟡 | 🟡 | 🔴 | ⚪ | ⚪ | 🔴 | 🔴 | 🔴 |
| 22 | Record updates without generic copying by key | 🔴 | 🟢 | 🔴 | 🟡 | 🟡 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🟡 |
| 23 | Records: unboxed numeric fields in that layout | 🔴 | 🔴 | ⚪ | 🟡 | 🟡 | 🔴 | ⚪ | ⚪ | ⚪ | ⚪ | 🔴 |
| 24 | Record-update coalescing / read-after-write forwarding | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟡 |
| 25 | Arrays with primitive storage selected by the backend | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 |
| 26 | Saturated ADT construction without closure chains | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🔴 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 |
| 27 | Erasure of `newtype` wrappers | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 |
| 28 | Nullable / niche representations for ordinary payload ADTs | 🔴 | 🔴 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 29 | Unboxed value results for selected payload ADTs | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 30 | Erasure of ordinary single-field data wrappers | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 31 | Factoring pattern tests into decision trees | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟢 |
| 32 | Flattening `let` bindings and `Effect` glue | 🟡 (PBO) | 🟡 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟡 (PBO) | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 |
| 33 | Native N-ary calls for explicit FnN / EffectFnN (1–10) | 🔴 | 🟡 (PBO) | 🟢 (PBO) | 🟡 (PBO) | 🔴 | 🟢 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🔴 |
| 34 | Eta reduction / forwarding-wrapper elimination | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🟡 | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🔴 |
| 35 | Elimination of non-escaping Ref / ST cells | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🟡 (PBO) | 🔴 | 🔴 | 🔴 |
| 36 | Library loop / collection traversal intrinsics | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 |
| 37 | Collection construction / concatenation fusion | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🟡 | 🔴 |
| 38 | Static representation of capture-free local lambdas | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 39 | Class-specific native layouts for retained dictionaries | 🔴 | 🟡 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 40 | Recursive local workers with explicit capture parameters | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 |
| 41 | Shared nullary constructors / allocation-free enums | 🟢 | 🟡 | 🟡 | 🟢 | 🟢 | 🟡 | 🟢 | 🟢 | 🟢 | 🟢 | 🟢 |
| 42 | Bypassing unchanged constructor reconstruction | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 43 | Native linked-list representation / intrinsics | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟢 | 🟢 | 🔴 |
| 44 | Mutual tail recursion with bounded stack usage | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🟢 (PBO) | ⚪ | ⚪ | 🟡 |
| 45 | Canonicalizing counted recursion into an induction loop | 🔴 | 🟡 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 46 | Limited common subexpression sharing — CSE | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 |
| 47 | Shared continuations for non-tail cases | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | 🟢 |
| 48 | Factoring identical branch tails | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 |
| 49 | Sticky sharing / saturating reference counts | ⚪ | 🟡 | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ |

[Audit notes: technique scope, evidence, limitations, and source revisions](optimization-audit.md).

## Latest Purust integrations

- **B13: red → yellow.** A uniquely owned ADT can now update one scalar field while leaving the other fields in place. Calls, conversions and changes to multiple fields retain the existing reconstruction. The rule handles RBTree recoloration and four `Data.Time` setters. Five paired full runners measured **RBTree 18.237 → 17.644 ms**, **total 20.977 → 20.354 ms**. [Scope, tests and measurements](scratch/rust-recolor-field-20260909/REPORT.md).
- **B3/B4 remain yellow; B45: red → yellow.** A typed recursive function producer of the form `build 0 = identity; build n = let previous = build (n - 1) in \f x -> f (previous f x)` now becomes a countdown loop in its saturated wrapper for `n >= 0`. This is one restricted transformation recorded under higher-order specialization, wrapper/worker use and counted-recursion recognition; its gain is counted once. It requires `Int -> (Int -> Int) -> Int -> Int`, preserves the negative fallback and callback order, and does not provide general monomorphization. Five paired full runners measured **Church 1.439 → 0.164 ms**, **total 21.053 → 19.691 ms**; separate allocation counts fell **133,641 → 309**. [Scope, tests and measurements](scratch/rust-function-fusion-20260909/REPORT.md).

- **B14 remains yellow; B22 remains green: one nested child now reuses its cell.** Closed TAST rows and the same local root/field identify the update; all RHS values are staged before detaching the child. Shared roots, independently shared children and captured old versions keep copying. The generated Records kernel has **20,003 → 10,003 allocations/liberations**, with **0.861 → 0.767 ms (−10.9%)** across five paired full runners. Total **19.735 → 19.733 ms** remains within variation. The deepest child still copies once per iteration. [Scope, tests and measurements](scratch/rust-record-child-20260909/REPORT.md).
- **B14 remains yellow; B22 remains green.** Closed-record updates with a local base at last use now evaluate replacement values first, then move the root into the existing setters. Retained aliases still trigger copying. The generated Records kernel eliminates **10,000 allocations/liberations: 30,003 → 20,003**. Two series of five paired full runners confirm **−10.5%** and **−10.1%** on Records; all ten pairs give **0.938 → 0.837 ms**. The total **19.112 → 19.077 ms** does not establish an overall improvement within the observed variation. This preceding root-only step retained both child copies; the child extension above removes one. [Scope, tests and measurements](scratch/rust-record-root-move-20260909/REPORT.md).

These are separate experiments with their own baselines, not successive runs whose times can be subtracted. The latest integration passes **22 codegen tests, 13 TAST tests and all 14 runner results**, including a clean build. The README remains the historical reference. The row order is retained: deeper record-path reuse, changed-child ADT reconstruction, read-only borrowing and record representation still need their own integration proofs and measurements. [Purust baby steps](../purust/purust/todo.md).

## Perceus, FBIP and sticky sharing

These names describe related but distinct mechanisms. **Perceus** combines precise reference counting, reuse analysis and specialization; **FBIP** means *functional but in-place*, where functional code can reuse uniquely owned storage while preserving shared values. A runtime pointer named `PerceusPtr` alone does not establish the complete compiler algorithm. See the [Perceus paper, sections 2.2–2.7](https://www.microsoft.com/en-us/research/wp-content/uploads/2020/11/perceus-tr-v4.pdf).

| Mechanism | Priority / matrix cells | Current Purust scope |
| --- | --- | --- |
| **FBIP / reuse analysis** | **12 — B12**, with fixed-field record updates in **B22** | Partial for ADTs: recognized consumed nodes and rotations reuse unique `Rc` cells, with a copying fallback when required. Record setters use `PerceusPtr::make_mut`. Neither cell claims universal in-place execution. |
| **Reuse specialization: leave unchanged ADT fields in place** | **13 — B13, partial** | [ReuseFields](../purust/purust/src/Purust/ReuseFields.purs) updates one scalar field of a uniquely owned ADT, preserving unchanged fields in place; shared/weak cases retain the reconstruction fallback. Calls, conversions and multiple changed fields are outside this first scope. Updating a child through a call is the next extension to prototype. Fixed-field record setters belong to **B22**. **B42** is different: returning the original constructor when the replacement is already equal. |
| **Last-use moves / borrowing; propagation of uniqueness** | **14 — B14, partial**, also **B12** | Selected moves and borrows are integrated, including closed-record roots moved after replacement values at last use and one immediate child detached after its RHS values. Existing setters reuse unique roots/children with copying for shared versions; the deepest child still copies in Records. Some ADT reconstruction paths retest uniqueness; read-only traversals can still clone child pointers. |
| **Perceus: release specialization and dup/drop fusion** | **15 — B15, missing** | No branch-local compiler pass over reference-count operations was established. Existing last-use moves, borrows and ownership transfers retain their credit in **B12/B14**; they do not by themselves establish this additional transformation. |
| **Sticky sharing / saturating reference counts** | **49 — B49, partial** | `PerceusPtr` saturates at `u32::MAX`; subsequent clones/drops leave the count unchanged and the object stays alive. Native ADTs such as RBTree use `std::rc::Rc`, so this mechanism does not apply to that path. No benchmark gain is attributed to saturation. |

Evidence: [ownership and reuse](optimization-audit.md#purust), [record setters and last-use generation](../purust/purust/src/Purust/CodeGen.purs), and the [saturating pointer runtime](../purust/purust/tests/runtime/perceus_ptr/src/lib.rs). The [Purust plan](../purust/purust/todo.md) breaks the remaining work into measured baby steps. Shared nullary constructors in **B41** are another distinct mechanism: sharing a value does not imply an immortal or saturated reference count.

The three added mechanisms have separate stable audit IDs **47–49**. Existing ratings are retained with their techniques; displayed coordinates change with the priority order. The audit explains the scopes and the gray cells for targets without backend-controlled reference counting.
