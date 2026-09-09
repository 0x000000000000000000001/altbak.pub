# Optimization techniques by backend

Audit dated **September 9, 2026**: **46 techniques across 11 backends**, cross-checked against their invoked passes and code generation paths. This matrix describes integration, coverage, and limitations. It serves as a work list; **it is not a ranking of backend speed**. The reference measurements remain those in the [README](README.md).

All eleven backends referenced by the runners are included. **Our five backends occupy the first columns: sharpurs, purust, phpurs, gopurs, and javapurs.** The other six follow on the right: psgo, official JS, ES, purescm, purerl, and Wasm, which is experimental across the 14 core cases. The README's “Native FP-style” and “Native hand-optimized” columns are handwritten implementations: their shortcuts are not attributed to the compilers.

## Rating legend

- 🟢 **Complete within the defined scope**: the mechanism is integrated and used by the relevant code generation path. This does not certify every possible program.
- 🟡 **Partial**: a significant restriction on types, expression shapes, local/intermodule scope, ABI, or code generation path. [Audit notes](optimization-audit.md#exact-scope-of-the-rows) explain the restriction.
- 🔴 **Missing**: no corresponding mechanism is integrated into the audited pipeline; a primitive, prototype, or unused module is insufficient.
- ⚪ **Gray — unnecessary / not applicable in this form**: the target's representation, value management, or native tail calls already cover the need. Gray cases are explained in the [audit notes](optimization-audit.md#exact-scope-of-the-rows).

**🟢 (PBO) / 🟡 (PBO)**: the shared PBO infrastructure provides a transformation or analysis actually used for this capability, sometimes complemented by the target code generator. The label means neither that PBO implements the entire technique on its own nor that all backends use it with the same coverage. It distinguishes this contribution from passes belonging to the backend, the PureScript frontend, or the runtime. For ES and purescm, it refers to their version of PBO, not our fork's TAST extensions.

A conservative proof of purity, arity, or non-escape is necessary for correctness; it is not inherently a shortcoming. “Partial” identifies a coverage limit worth knowing about. These cells should not be added up into a maturity percentage: some rows depend on the same passes and carry different weight.

## Matrix

**Rows are sorted by estimated performance importance, highest first**, prioritizing our five backends and the core benchmarks. This is a qualitative estimate; neighboring positions can vary by workload. [Ordering rationale](optimization-audit.md#estimated-impact-order).

**Cell references:** backend columns are **A–K**, and technique rows are **1–46**. For example, **E12** means **javapurs × ADT fields specialized by payload type**. Coordinates follow the current table order; the audit uses separate stable technique IDs.

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
| 11 | Self-tail recursion with bounded stack usage | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 | 🟢 (PBO) | ⚪ | ⚪ | 🟡 |
| 12 | ADT fields specialized by payload type | 🟡 | 🟡 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 |
| 13 | Memory reuse after proving ownership / uniqueness (FBIP) | 🔴 | 🟡 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 14 | Loop invariants: hoisting / caching per invocation | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 15 | Dedicated fusion of a thunk producer with its consumer | 🔴 | 🟡 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 16 | Scrutinee fusion / case-of-case | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟡 |
| 17 | Scalar replacement of known / non-escaping records and arrays | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟡 |
| 18 | Last-use moves / borrowing to avoid explicit reference counting | ⚪ | 🟡 | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ |
| 19 | Closed records: specialized layout and fixed fields | 🔴 | 🟡 | 🔴 | 🟡 | 🟡 | 🔴 | ⚪ | ⚪ | 🔴 | 🔴 | 🔴 |
| 20 | Record updates without generic copying by key | 🔴 | 🟢 | 🔴 | 🟡 | 🟡 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🟡 |
| 21 | Records: unboxed numeric fields in that layout | 🔴 | 🔴 | ⚪ | 🟡 | 🟡 | 🔴 | ⚪ | ⚪ | ⚪ | ⚪ | 🔴 |
| 22 | Record-update coalescing / read-after-write forwarding | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟡 |
| 23 | Arrays with primitive storage selected by the backend | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 |
| 24 | Saturated ADT construction without closure chains | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🔴 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 |
| 25 | Erasure of `newtype` wrappers | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 |
| 26 | Nullable / niche representations for ordinary payload ADTs | 🔴 | 🔴 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 27 | Unboxed value results for selected payload ADTs | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 28 | Erasure of ordinary single-field data wrappers | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 29 | Factoring pattern tests into decision trees | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟢 |
| 30 | Flattening `let` bindings and `Effect` glue | 🟡 (PBO) | 🟡 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟡 (PBO) | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 |
| 31 | Native N-ary calls for explicit FnN / EffectFnN (1–10) | 🔴 | 🟡 (PBO) | 🟢 (PBO) | 🟡 (PBO) | 🔴 | 🟢 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🔴 |
| 32 | Eta reduction / forwarding-wrapper elimination | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🟡 | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🔴 |
| 33 | Elimination of non-escaping Ref / ST cells | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🟡 (PBO) | 🔴 | 🔴 | 🔴 |
| 34 | Library loop / collection traversal intrinsics | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 |
| 35 | Collection construction / concatenation fusion | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🟡 | 🔴 |
| 36 | Static representation of capture-free local lambdas | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 37 | Class-specific native layouts for retained dictionaries | 🔴 | 🟡 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 38 | Recursive local workers with explicit capture parameters | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 |
| 39 | Shared nullary constructors / allocation-free enums | 🟢 | 🟡 | 🟡 | 🟢 | 🟢 | 🟡 | 🟢 | 🟢 | 🟢 | 🟢 | 🟢 |
| 40 | Bypassing unchanged constructor reconstruction | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 41 | Native linked-list representation / intrinsics | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟢 | 🟢 | 🔴 |
| 42 | Mutual tail recursion with bounded stack usage | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🟢 (PBO) | ⚪ | ⚪ | 🟡 |
| 43 | Canonicalizing counted recursion into an induction loop | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| 44 | Limited common subexpression sharing — CSE | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 |
| 45 | Shared continuations for non-tail cases | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | 🟢 |
| 46 | Factoring identical branch tails | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 |

[Audit notes: technique scope, evidence, limitations, and source revisions](optimization-audit.md).

## Perceus, FBIP and sticky sharing

These names describe related but distinct mechanisms. **Perceus** combines precise reference counting, reuse analysis and specialization; **FBIP** means *functional but in-place*, where functional code can reuse uniquely owned storage while preserving shared values. A runtime pointer named `PerceusPtr` alone does not establish the complete compiler algorithm. See the [Perceus paper, sections 2.2–2.7](https://www.microsoft.com/en-us/research/wp-content/uploads/2020/11/perceus-tr-v4.pdf).

| Mechanism | Relation to the matrix | Current Purust scope |
| --- | --- | --- |
| **FBIP / reuse analysis** | **B13**, with fixed-field record updates in **B20** | Partial for ADTs: recognized consumed nodes and rotations reuse unique `Rc` cells, with a copying fallback when required. Record setters use `PerceusPtr::make_mut`. Neither cell claims universal in-place execution. |
| **Precise reference counting, drop specialization and dup/drop fusion** | **B18** covers last-use moves and borrowing that reduce reference-count traffic | These optimizations are integrated for selected paths. A systematic Perceus pass that places releases at the earliest safe point and specializes/fuses reference-count operations across all branches is not established. |
| **Reuse specialization: leave unchanged fields in place** | Extends **B13**; fixed-field record setters already contribute through **B20** | Record setters modify the selected field when unique. ADT reuse often extracts and rebuilds the whole payload; specializing those writes remains planned work. **B40** is a different optimization: returning the original constructor when the replacement is already equal. |
| **Propagating uniqueness / extending inferred borrows** | Further work within **B13/B18** | Some generated reconstruction paths retest uniqueness; read-only traversals can still clone child pointers. Removing these costs requires usage proofs and measurements, not just native TAST types. |
| **Sticky sharing / saturating reference counts** | Runtime capability not separately rated among the 46 rows | `PerceusPtr` saturates at `u32::MAX`; subsequent clones/drops leave the count unchanged and the object stays alive. Native ADTs such as RBTree use `std::rc::Rc`, so this mechanism does not apply to that path. No benchmark gain is attributed to saturation. |

Evidence: [ownership and reuse](optimization-audit.md#purust), [record setters and last-use generation](../purust/purust/src/Purust/CodeGen.purs), and the [saturating pointer runtime](../purust/purust/tests/runtime/perceus_ptr/src/lib.rs). The [Purust plan](../purust/purust/todo.md) breaks the remaining work into measured baby steps. Shared nullary constructors in **B39** are another distinct mechanism: sharing a value does not imply an immortal or saturated reference count.

This terminology cross-reference adds no blanket Perceus rating for another backend and does not change the existing cell ratings or coordinates.
