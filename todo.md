# Optimization techniques by backend

Audit dated **September 9, 2026**. This matrix describes the techniques actually integrated into the compilers used by altbak, their coverage, and their limitations. It serves as a work list; **it is not a performance ranking**. The reference measurements remain those in the [README](README.md).

All eleven backends referenced by the runners are included. **Our five backends occupy the first columns: sharpurs, purust, phpurs, gopurs, and javapurs.** The other six follow on the right: psgo, official JS, ES, purescm, purerl, and Wasm, which is experimental across the 14 core cases. The README's “Native FP-style” and “Native hand-optimized” columns are handwritten implementations: their shortcuts are not attributed to the compilers.

## Rating legend

- 🟢 **Complete within the defined scope**: the mechanism is integrated and used by the relevant code generation path. This does not certify every possible program.
- 🟡 **Partial**: a significant restriction on types, expression shapes, local/intermodule scope, ABI, or code generation path. [Audit notes](optimization-audit.md#exact-scope-of-the-rows) explain the restriction.
- 🔴 **Missing**: no corresponding mechanism is integrated into the audited pipeline; a primitive, prototype, or unused module is insufficient.
- ⚪ **Gray — unnecessary / not applicable in this form**: the target's representation, value management, or native tail calls already cover the need. Gray cases are explained in the [audit notes](optimization-audit.md#exact-scope-of-the-rows).

**🟢 (PBO) / 🟡 (PBO)**: the shared PBO infrastructure provides a transformation or analysis actually used for this capability, sometimes complemented by the target code generator. The label means neither that PBO implements the entire technique on its own nor that all backends use it with the same coverage. It distinguishes this contribution from passes belonging to the backend, the PureScript frontend, or the runtime. For ES and purescm, it refers to their version of PBO, not our fork's TAST extensions.

A conservative proof of purity, arity, or non-escape is necessary for correctness; it is not inherently a shortcoming. “Partial” identifies a coverage limit worth knowing about. These cells should not be added up into a maturity percentage: some rows depend on the same passes and carry different weight.

## Matrix

The headers link to each backend's evidence and limitations. **JS** refers to the official code generator; **ES** refers to Arista's `purs-backend-es`. **Sharpurs emits F#**, with C# interoperability, rather than translating PureScript function bodies into C#.

| Technique | [sharpurs](optimization-audit.md#sharpurs) | [purust](optimization-audit.md#purust) | [phpurs](optimization-audit.md#phpurs) | [gopurs](optimization-audit.md#gopurs) | [javapurs](optimization-audit.md#javapurs) | [psgo](optimization-audit.md#psgo) | [JS](optimization-audit.md#official-js) | [ES](optimization-audit.md#es) | [purescm](optimization-audit.md#purescm) | [purerl](optimization-audit.md#purerl) | [Wasm](optimization-audit.md#wasm) |
| --- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| Partial evaluation, β-reduction, simplification | 🟡 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 |
| Automatic inlining and dead code elimination | 🟡 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 |
| Elimination of known type class dictionaries | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🟡 | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🟡 |
| Using TAST to select representations | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| Effective `TypeApp` type instantiation | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| Monomorphization by cloning polymorphic functions | 🔴 | 🔴 | 🔴 | 🟡 (PBO) | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| Automatic uncurrying / wrapper-worker | 🟡 | 🟡 | 🟢 | 🟡 | 🟡 | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🟡 | 🟢 |
| Direct emission of primitive operations | 🟡 (PBO) | 🟢 (PBO) | 🟡 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 |
| Keeping numeric parameters and locals unboxed | 🟡 | 🟡 | ⚪ | 🟡 | 🟡 | 🔴 | ⚪ | ⚪ | ⚪ | ⚪ | 🟡 |
| ADT fields specialized by payload type | 🟡 | 🟡 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 |
| Shared nullary constructors / allocation-free enums | 🟡 | 🟡 | 🟡 | 🟢 | 🟢 | 🟡 | 🟢 | 🟢 | 🟢 | 🟢 | 🟡 |
| Closed records: specialized layout and fixed fields | 🔴 | 🟡 | 🔴 | 🟡 | 🟡 | 🔴 | ⚪ | ⚪ | 🔴 | 🔴 | 🔴 |
| Records: unboxed numeric fields in that layout | 🔴 | 🔴 | ⚪ | 🟡 | 🟡 | 🔴 | ⚪ | ⚪ | ⚪ | ⚪ | 🔴 |
| Record updates without generic copying by key | 🔴 | 🟡 | 🔴 | 🟡 | 🟡 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🟡 |
| Self-tail recursion with bounded stack usage | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 | 🟢 (PBO) | ⚪ | ⚪ | 🟡 |
| Mutual tail recursion with bounded stack usage | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🟢 (PBO) | ⚪ | ⚪ | 🟡 |
| Canonicalizing counted recursion into an induction loop | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| Flattening `let` bindings and `Effect` glue | 🟡 (PBO) | 🟡 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟡 | 🟢 (PBO) | 🟢 (PBO) | 🟡 | 🟢 |
| Limited common subexpression sharing — CSE | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 | 🟡 |
| Loop invariants: hoisting / caching per invocation | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| Dedicated fusion of a thunk producer with its consumer | 🔴 | 🔴 | 🟡 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| Specialization of higher-order functions / closures | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟡 |
| Arrays with primitive storage selected by the backend | 🔴 | 🔴 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 |
| Memory reuse after proving ownership / uniqueness | 🔴 | 🟡 | 🔴 | 🟡 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 |
| Erasure of `newtype` wrappers | 🟡 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 |
| Saturated ADT construction without closure chains | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🟢 (PBO) | 🔴 | 🟢 | 🟢 (PBO) | 🟢 (PBO) | 🟢 | 🟢 |
| Elimination of non-escaping Ref / ST cells | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🔴 | 🟡 | 🟡 (PBO) | 🔴 | 🔴 | 🔴 |
| Borrowing to avoid explicit clones / reference counting | ⚪ | 🟡 | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ | ⚪ |
| Scrutinee fusion / case-of-case | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🔴 | 🟡 (PBO) | 🟡 (PBO) | 🔴 | 🟡 |

[Audit notes: technique scope, evidence, limitations, and source revisions](optimization-audit.md).
