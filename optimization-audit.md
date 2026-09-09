# Optimization audit notes

Supporting evidence for the [backend optimization matrix](todo.md#matrix), audited on **September 9, 2026**. The numbers used in these notes refer to the technique index below; the main matrix has no numbered cells.

## Technique index

1. Partial evaluation, β-reduction, simplification
2. Automatic inlining and dead code elimination
3. Elimination of known type class dictionaries
4. Using TAST to select representations
5. Effective `TypeApp` type instantiation
6. Monomorphization by cloning polymorphic functions
7. Automatic uncurrying / wrapper-worker
8. Direct emission of primitive operations
9. Keeping numeric parameters and locals unboxed
10. ADT fields specialized by payload type
11. Shared nullary constructors / allocation-free enums
12. Closed records: specialized layout and fixed fields
13. Records: unboxed numeric fields in that layout
14. Record updates without generic copying by key
15. Self-tail recursion with bounded stack usage
16. Mutual tail recursion with bounded stack usage
17. Canonicalizing counted recursion into an induction loop
18. Flattening `let` bindings and `Effect` glue
19. Limited common subexpression sharing — CSE
20. Loop invariants: hoisting / caching per invocation
21. Dedicated fusion of a thunk producer with its consumer
22. Specialization of higher-order functions / closures
23. Arrays with primitive storage selected by the backend
24. Memory reuse after proving ownership / uniqueness
25. Erasure of `newtype` wrappers
26. Saturated ADT construction without closure chains
27. Elimination of non-escaping Ref / ST cells
28. Borrowing to avoid explicit clones / reference counting
29. Scrutinee fusion / case-of-case

## Exact scope of the rows

- **01–03**: optimize known values and implementations. Green in 02 covers the bindings/imports handled by the optimizer, not a promise to remove every unused export in the program. Yellow in 03 distinguishes these reductions from systematically eliminating all user-defined dictionaries.
- **04–06**: three separate capabilities. Reading TAST does not mean using every type; substituting `forall a` at a call does not mean creating a standalone `f_Int` function. Row 06 requires this **function cloning**, beyond regions that become monomorphic after inlining. A `Monomorphize.purs` module present in a PBO that is not invoked receives no credit.
- **07**: **automatic** uncurrying of ordinary functions. Explicit `FnN`/`EffectFnN` support alone is insufficient; constructors are covered separately in 26. A worker may still use `Object`, `obj`, or `Value`: see 09.
- **08–10**: a native operator, a primitive local, and a primitive ADT field are three different outcomes. Gray in 09 applies to ABIs that already pass their language's native scalars without adding an `Integer`/`Double` box. Their physical numeric storage is managed by the runtime. Red in 10 does not deny the existence of ADT classes, tuples, or records: it indicates that payloads are not specialized by type.
- **11**: a singleton, shared literal, or immediate tag. Coverage limited to entirely nullary enums leaves the cell yellow for mixed ADTs.
- **12–14**: record layout, field types, and update strategy. A layout with fixed fields does not guarantee unboxed numeric payloads. Native JS objects already have named properties; a pass replacing an artificial Map with a record is unnecessary there, hence gray in 12. V8 decides their machine layout. Gray in 13 applies the same distinction as 09 to targets with native/tagged scalar values; it does not mean their records or access paths are optimal. PHP `stdClass`, Scheme association lists, and Erlang maps still have alternative layouts to explore in 12.
- **15–17**: self TCO, mutual TCO, and counter recognition are separate. Scheme and Erlang provide native tail calls, so an additional backend trampoline is unnecessary, hence gray in 15/16. This does not remove the cost of their closures. Yellow for F# depends on the emitted tail-call forms and then F#/.NET; yellow for Wasm depends on `return_call` and compatible representations. Row 17 requires transforming the recursive pattern, not merely emitting a `for` in an FFI or collection intrinsic.
- **19–21**: frontend CSE mainly shares synthetic dictionary applications; an invariant cache avoids repeating a computation during an invocation; thunk fusion avoids building/forcing the chain. These three techniques are not interchangeable. Neither an internal compiler cache nor a global constant initialized once satisfies row 20.
- **22**: targeted coverage of higher-order functions (HOFs)/closures through known values and inlining, specialized representations, or private workers. Mechanisms differ by column and are detailed below. No yellow cell means general defunctionalization of all closures. CPS combinator fusion achieved through inlining is counted here, not as the dedicated thunk pass in row 21.
- **23**: `[]int64`, `int[]`, `Vec<i64>`, or an equivalent selected by the compiler. A speculative V8 optimization, a PHP array of zvals, or handwritten FFI is insufficient. Wasm is yellow for its explicit primitive arrays; ordinary `Array Int` is not generally specialized.
- **24 / 27 / 28**: reusing a unique object, removing a local cell, and borrowing a field instead of cloning a reference are distinct. Gray in 28 applies to targets where the backend does not emit such explicit clones/reference counts. PHP's zval management belongs to the engine; in Go, the presence of a historical `Rc` field does not prove it is used.
- **29**: scrutinee fusion here means eliminating an ADT that is constructed and then immediately destructured: reducing a `case` over a known constructor, or distributing the consumer into the producer's branches (*case-of-case*). Yellow covers these local forms, including after inlining; it does not promise fusion across a producer call that remains opaque. This row is separate from thunk fusion in 21.

## Provenance and methodology

The `bin/*/run` scripts, dependencies actually selected, entry points, invoked passes, generic fallback paths, existing tests, and previously generated outputs were cross-checked. No benchmark or full build was run for this documentation change. The cited tests document the cases the project checks; they were not rerun during this audit. Sources and outputs do not establish a measured performance gain.

| Column | Inspected version / revision | Identified pipeline |
|---|---|---|
| sharpurs | `7ebdd3f` | TAST purs fork → selected PBO kernels + generic generation from TAST → F# |
| purust | `bf888346` | TAST purs fork → dedicated PBO → Purust → Rust |
| phpurs | `be46bdfc` | TAST purs fork → dedicated PBO → Phpurs → PHP |
| gopurs | `a2263644`; subsequent internal rename `a100d62` | TAST purs fork → dedicated PBO → Gopurs → Go |
| javapurs | `2164c48` | TAST purs fork → dedicated PBO → Javapurs → Java |
| psgo | distribution `0.1.0`; declared sources `355eb9b0`, identical to local `purescript-native/src` at `9829dddc` | purs from PATH, here `0.15.15` → psgo → Go |
| Official JS | purs from PATH, here `0.15.15` | CoreFn → CoreImp optimizer → JS → Node |
| ES | npm `purs-backend-es@1.4.3`, commit `5e764325` | CoreFn → PBO **from this release** → JS → Node |
| purescm | npm `1.12.0` and corresponding local sources; PBO pinned to `54b82ac` | vanilla purs `0.15.15` → purescm's PBO → Scheme → Chez |
| purerl | binary `0.0.24`, local sources `abe9d99b`, and output carrying that version | vanilla purs `0.15.15` → Erlang optimizer → BEAM |
| Wasm | `b98778f0` | standard purs `0.15.16` + externs → its own MIR → WasmGC/Binaryen → Node |

The five PBO copies dedicated to our local backends were inspected at commit `6e256c6a`; the frontend fork at commit `65010ea`. **This extended PBO must not be attributed to the ES 1.4.3 npm package.** The package's exact source was identified through `gitHead` in its [npm metadata](https://registry.npmjs.org/purs-backend-es/1.4.3), then read at the [published commit](https://github.com/aristanetworks/purescript-backend-optimizer/tree/5e7643253ebc9db16f3c9ab702fa756bd3e41b80).

The [psgo distribution](https://github.com/rowtype-yoga/psgo/tree/v0.1.0) points to `andyarvanitis/purescript-native`, branch `golang`, but does not pin its build commit: binary identity was not reproduced. For purescm/purerl, local sources, tool versions, and outputs agree on the mechanisms described; no bit-for-bit rebuild establishing provenance is claimed.

Local links below are relative to this repository in the `htdocs` workspace. Line numbers may move; function and pass names serve as reference points. The revisions above identify the audited state despite concurrent work. A capability receives no credit merely because it appears in an old scratch file, a TODO, or an experiment that is not integrated.

## Shared frontend and PBO evidence

- **Limited CSE, row 19**: [Make.hs](../purescript/src/Language/PureScript/Make.hs), `optimizeCoreFn`, then serialization in [Make/Actions.hs](../purescript/src/Language/PureScript/Make/Actions.hs). The [official 0.15.15 selection](https://github.com/purescript/purescript/blob/v0.15.15/src/Language/PureScript/CoreFn/CSE.hs#L369) targets simple `IsSyntheticApp` expressions: dictionaries, superclasses, and `IsSymbol`. This is not general CSE across all arithmetic computations. Wasm additionally has `SShared` sharing and CAFs.
- **PBO, rows 01–03, 18, 22, 25**: [Builder](../purescript-backend-optimizer-javapurs/src/PureScript/Backend/Optimizer/Builder.purs), `buildModules` → `toBackendModule`; [Convert](../purescript-backend-optimizer-javapurs/src/PureScript/Backend/Optimizer/Convert.purs), selection of used bindings and `IsNewtype`; [Semantics](../purescript-backend-optimizer-javapurs/src/PureScript/Backend/Optimizer/Semantics.purs), evaluation, inlining, `shouldUnpackCtor`, `shouldUnpackRecord`, and let floating. Known aggregates are genuinely eliminated, but this does not constitute general uniqueness analysis.
- **TypeApp, row 05**: in the same `Semantics`, `evalTypeApp` and `instantiateNeutralType` substitute types under `ForAll`, and `evalExternFromImpl` actually uses the instantiation before the value arguments. [TypeApp tests](../purescript-backend-optimizer-javapurs/test/typeapp.mjs). Ignoring a **residual** TypeApp in the final code generator therefore does not mean the entire pipeline ignored it.
- **Operations and effects, rows 08 and 18**: [Semantics/Foreign](../purescript-backend-optimizer-javapurs/src/PureScript/Backend/Optimizer/Semantics/Foreign.purs) recognizes functions including `intAdd`, `intMul`, `bindE`, and `pureE`, then builds `PrimOp` / `EffectBind` / `EffectPure`. The target code generator then emits the operators and statements: the contribution is shared, hence `(PBO)`.
- **TCO, row 15 for Phpurs/Gopurs/Javapurs, 15–16 for ES**: [Codegen/Tco](../purescript-backend-optimizer-javapurs/src/PureScript/Backend/Optimizer/Codegen/Tco.purs) analyzes tail calls and loop roles; the cited code generators use that analysis. Emitting loops or the dispatcher remains specific to the backend. Purust and Sharpurs use their own paths for this capability.
- **Saturated construction, row 26**: `evalApp` in `Semantics` transforms a fully applied `NeutCtorDef` into `NeutData`, then reification produces `CtorSaturated`. The code generators consume this form; Sharpurs does so in its kernels, with a separate path for generic declarations.
- **Attribution boundaries**: row 04 describes the code generators' representation choices based on the TAST produced by the frontend; CSE in row 19 also comes from the frontend. Merely carrying this information does not earn a `(PBO)` label. In row 07, `shouldUncurryAbs` is disabled in the audited local PBO (`Nothing`): our five backends implement their own workers. Gopurs monomorphization in 06 does, however, actually call the PBO monomorphizer, as detailed below.
- **Scrutinee fusion, row 29**: in `Semantics`, `shouldUnpackCtor` removes a constructor used only for field reads or tag tests. `shouldDistributeBranches` pushes the consumer into the branches under the guards `size <= 128`, `result == KnownNeutral`, and uses limited to accesses/tests; `RewriteDistBranchesLet`, `evalAccessor`, and `evalPrimOp` then simplify known fields and tags. These mechanisms exist in our five dedicated PBO copies, in ES 1.4.3, and in the PBO pinned by purescm. Sharpurs benefits through the selected paths described below. The [KnownConstructors03 source case](../purescript-backend-optimizer-gopurs/backend-es/test-snapshots/src/Snapshot.KnownConstructors03.purs) constructs a `Just`/`Nothing` and then pattern-matches it; the [reference output](../purescript-backend-optimizer-gopurs/backend-es/test-snapshots/snapshots-out/Snapshot.KnownConstructors03.js) contains only an `if` returning the final strings, without that intermediate ADT. This snapshot illustrates the shared PBO infrastructure, not a new measurement of the five backends.
- The entry points cited for each backend determine whether this result is actually emitted. Sharpurs consumes only part of it; the historical ES/purescm packages use their own PBO versions.

## sharpurs

- **01–05, 18**: [Main](../sharpurs/sharpurs/src/Main.purs) and [CodeGen](../sharpurs/sharpurs/src/Sharpurs/CodeGen.purs) select ADT/Thunk/Int kernels and [Optimized](../sharpurs/sharpurs/src/Sharpurs/Optimized.purs) wrappers. The general path still traverses **the original TAST declarations**, rather than all the optimized PBO declarations. The mere presence of `buildModules` therefore does not justify green ratings everywhere.
- **07–10**: [DirectCall](../sharpurs/sharpurs/src/Sharpurs/DirectCall.purs) supports monomorphic NonRec bindings and exactly saturated qualified calls; `obj` workers/wrappers remain. [IntKernel](../sharpurs/sharpurs/src/Sharpurs/IntKernel.purs), [IntArithmetic](../sharpurs/sharpurs/src/Sharpurs/IntArithmetic.purs) and [AdtLayout](../sharpurs/sharpurs/src/Sharpurs/AdtLayout.purs) create native paths for supported Int/Bool/closed ADTs. External or polymorphic fields can cause the kernel to be rejected; the fallback is a DU with `obj` fields.
- **11, 26**: DU cases without payloads and shared native values; saturated construction directly as a DU or through a registered native factory. The machine representation and final sharing of nullary constructors also depend on F#.
- **12–14, 23–24**: records use `Map<string,obj>`, updates use `Map.add`, and generic arrays are boxed. The structural sharing provided by `Map` is neither a closed record layout nor a Sharpurs ownership analysis.
- **15–17**: `let rec … and …` groups, `_tco` entry points and saturated direct calls prepare tail calls for F#/.NET. Sharpurs does not emit a general trampoline/dispatcher that itself guarantees bounded stack usage, or a counted loop. See the [TCO output](../altbak.pub-sharpurs/run/bak/sharp/output/Main/Test.TCO.fs).
- **20–22, 25**: [ThunkKernel](../sharpurs/sharpurs/src/Sharpurs/ThunkKernel.purs) is integrated, including in the bundle and the [LazyEvaluation output](../altbak.pub-sharpurs/run/bak/sharp/output/Main/Test.LazyEvaluation.fs). It specializes private `unit -> int` regions and removes recognized `defer`/`force` helpers. **The computation's closure chain is still built, and the worker is called again on every iteration**: yellow 22, red 20/21. The newtype does not add a wrapper class, but its identity constructor and boxed calls remain on the generic path: yellow 25.
- **27**: `Main` deliberately filters out foreign Effect/ST semantics. No general elimination of local cells in the kernels.
- The experimental `AdtInterop` module is not imported by `Main`: its test capabilities are not attributed to the normal pipeline.

## purust

- **01–05, 18, 25**: [Main](../purust/purust/src/Main.purs) actually passes `backendMod` to the generator. [TypeApp tests](../purust/purust/tests/tast/type-instantiation.mjs) verify the emission of `i64`/`f64` loops after instantiation and the `UnknownType` fallback. There is no general pass that creates clones per instantiation: yellow 05 does not imply yellow 06.
- **07–10, 22, 26**: [CodeGen](../purust/purust/src/Purust/CodeGen.purs), direct calls to globals with known arity, partial applications and overapplications, primitive parameters/returns, fields derived from `dataDecls`, type classes represented as structs. Recursive ADTs use `Rc`, and type variables remain generic. `FuncN::Static` / `Shared` distinguishes functions without captures from shared closures. Neither `FuncN` nor the later monomorphization of Rust generics constitutes the PureScript cloning described in row 06.
- **11**: [DataLayout](../purust/purust/src/Purust/DataLayout.purs) makes entirely nullary enums `Copy`. [ShareNullaries](../purust/purust/src/Purust/ShareNullaries.purs) shares nullary constructors in local strict constructions and stops at scope/call boundaries; there is no universal global singleton for nullary constructors in mixed types.
- **12–14**: record structs are based on label sets, but their fields are **`Option<UnknownType>`**. Getters/setters dispatch on shape. Updates use `PerceusPtr::make_mut`: mutate when unique, copy the struct otherwise. The layout is partially structured, but numeric record fields are not yet unboxed.
- **15–17, 20–21**: a single `mbLoop`, calls to the same name → temporaries, `continue`, `loop`. No mutual dispatcher, counter canonicalization, invariant cache or dedicated thunk fusion pass. Partial evaluation can nevertheless eliminate known combinators, row 22.
- **23**: `LitArray` converts elements to `Any`, then uses `mk_array(vec![…])`. This is not a `Vec<i64>` inferred from `Array Int`.
- **24**: [OwnedFields](../purust/purust/src/Purust/OwnedFields.purs), [ReturnCells](../purust/purust/src/Purust/ReturnCells.purs), return workers, `Rc::get_mut` and the shared fallback are actually wired in. The [PerceusPtr runtime](../purust/purust/tests/runtime/perceus_ptr/src/lib.rs), although located under `tests/`, is a dependency of the generated code. The [RBTree output](../altbak.pub-purust/run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs) calls `balance__purust_reuse` on the unique path. Coverage is limited to recognized extractions/returns.
- **27–28**: no general scalarization of Ref cells. However, the [function borrow tests](../purust/purust/tests/codegen/function-borrows.mjs) and the field/tag borrow paths cover an active mechanism that avoids some clones/reference count increments. Borrowing does not permit mutation of a shared value.

## phpurs

- **01–05, 18, 25**: [Main](../phpurs/phpurs/src/Main.purs) does call the dedicated PBO. TAST annotations become PHP scalar signatures/fields and proofs for private regions; they do not turn zvals into raw machine fields.
- **07**: [Printer](../phpurs/phpurs/src/Phpurs/Printer.purs), through the `PhpCall` branch using `allArities`, emits real PHP function calls, including namespace-qualified calls, with overapplication handling. The AST name `PhpDirectCall` alone was not sufficient evidence of this.
- **08, 10–11, 26**: [CodeGen](../phpurs/phpurs/src/Phpurs/CodeGen.purs), scalar mapping, classes derived from `dataDecls`, `CtorSaturated` emitted as `new`. Nullary bindings are shared, but a generic `CtorSaturated` can still allocate. [EnumRegions](../phpurs/phpurs/src/Phpurs/EnumRegions.purs) and [NullableConstructors](../phpurs/phpurs/src/Phpurs/NullableConstructors.purs) reduce enums/nullary leaves in proven private regions while preserving the public representation. Yellow 08 reflects the recognized mappings, rather than certifying all Int32 arithmetic.
- **12, 14, 23**: residual records remain `stdClass` instances with generic cloning, and arrays remain PHP arrays of zvals. The [Records output](../altbak.pub-phpurs/run/bak/php/output/Test.Records/index.php) retains nested clones. No closed record layout or compact numeric array storage is selected.
- **15–17**: `CodeGen` requires a single recursive binding, both locally and at the top level; reassignments followed by `goto`. No mutual TCO or counted loop pass.
- **21–22**: [PartialBindings](../phpurs/phpurs/src/Phpurs/PartialBindings.purs), [ThunkFusion](../phpurs/phpurs/src/Phpurs/ThunkFusion.purs) and [CompactLoops](../phpurs/phpurs/src/Phpurs/CompactLoops.purs) are wired in. Immediately consumed chains and local scalar closures can become compact private workers/callable objects; public closures and cases of unknown provenance remain unchanged. This is not general defunctionalization.
- **02, 24**: [TailInline](../phpurs/phpurs/src/Phpurs/TailInline.purs) and [CopyCleanup](../phpurs/phpurs/src/Phpurs/CopyCleanup.purs) are also wired in. Cleaning up copies/temporaries or preserving simultaneous TCO assignment is not a uniqueness analysis that permits ADT mutation.
- **20, 27**: the [LazyEvaluation output](../altbak.pub-phpurs/run/bak/php/output/Test.LazyEvaluation/index.php) recomputes its fused worker on every iteration. The Church cache tried in scratch is not integrated. ST primitives recognized by PBO remain `TODO_PrimEffect` if they reach the generator: no cell scalarization is wired in.

## gopurs

- **01–06, 22**: [Main](../gopurs/gopurs/src/Main.purs), `monomorphizeModules` before `buildModules`; [Monomorphization](../gopurs/gopurs/src/Gopurs/Monomorphization.purs) and the [PBO monomorphizer](../purescript-backend-optimizer-gopurs/src/PureScript/Backend/Optimizer/Monomorphize.purs). Transitive collection, cloning and call rewriting are indeed present, with types/dictionaries/static arguments. Foreigns, insufficiently concrete instantiations and some recursive locals remain generic. [ClassMetadata](../gopurs/gopurs/src/Gopurs/ClassMetadata.purs) also consumes `classDecls`.
- **07–14, 23, 26**: [CodeGen](../gopurs/gopurs/src/Gopurs/CodeGen.purs), `Call_Module_name` workers, arities 1–10, `exprTypeToGoType`, structured fields and updates, `CtorSaturated`, map/filter/foldl intrinsics; [Printer](../gopurs/gopurs/src/Gopurs/Printer.purs), struct copies by value and slice conversions. The `Value` fallback and `[]Value ↔ []int64` bridges prevent any promise of native representation throughout. The [Value](../gopurs/gopurs/src/Gopurs/Runtime.purs) runtime already stores scalar bits in `IntVal`, without systematically allocating a separate box for each value.
- **10–11, 25**: [AdtMetadata](../gopurs/gopurs/src/Gopurs/AdtMetadata.purs) selects entirely nullary enums and single-constructor ADTs with payloads for pointer representation; [ConstructorMetadata](../gopurs/gopurs/src/Gopurs/ConstructorMetadata.purs) identifies transparent wrappers. More general sums and types without a selected representation retain `Value` paths.
- **12–14**: records have real native fields in recognized cases. The `isClosedRowTail` policy also accepts `Just Any`; it is not equivalent to the strict proof of a closed row used by Javapurs.
- **15–17**: `isSelfRecursiveLoop` requires a top-level group of size 1. Locally, the passed `currentLoopCtx` contains only the current function, despite the names `mutRecBinds` and `combinedLoopCtx`. The code emits `for {}`/`continue`, without a mutual dispatcher or a counter canonicalization pass.
- **21–24**: [ThunkFusion](../gopurs/gopurs/src/Gopurs/ThunkFusion.purs), `optimizeThunkProducers`, replaces an immediately consumed recursive Int thunk producer with a strict worker, under proofs that the operations are total. `normalizeFreshIntArrayRoundtrip` reuses a freshly created filter buffer; this yellow 24 **does not mean mutation of a unique ADT**. `constructorReuse` returns an already equal object without modifying it. No `Rc` check currently permits generated ADT mutation.
- **20, 27**: no program invariant cache in the pipeline. `PrimEffect` retains an indirect cell for RefNew/Read/Write; any subsequent escape optimization by Go does not constitute a Gopurs pass.

## javapurs

- **01–05, 18, 22, 25**: [Main](../javapurs/javapurs/src/Main.purs) consumes the dedicated PBO. Reductions based on known values, dictionaries and inlinable HOFs come from this foundation. Yellow 22 describes this structural specialization; there is no general representation of numeric closures or Java defunctionalization. No call to the clone monomorphizer, hence red 06.
- **07**: [DirectCalls](../javapurs/javapurs/src/Javapurs/DirectCalls.purs) retains the curried wrapper and creates a static worker used for consecutive lambdas of arity 2–32, for nonrecursive globals in the same module declared earlier. Workers still take `Object` parameters. Partial applications, local/unknown functions and initialization boundaries retain their compatible path. [Tests](../javapurs/javapurs/test/direct-calls.mjs).
- **08–10, 15**: [CodeGen](../javapurs/javapurs/src/Javapurs/CodeGen.purs), primitives and TCO; [IntLoops](../javapurs/javapurs/src/Javapurs/IntLoops.purs), counters with a proven Int representation → `int`. Ordinary ADT signatures and fields remain `Object`. [Printer](../javapurs/javapurs/src/Javapurs/Printer.purs) emits direct `continue` and keeps `TcoLoop` for method/closure boundaries that require this fallback; this is not general mutual TCO.
- **11, 26**: constructor classes, saturated construction with `new Ctor(args)` and nullary sharing through a holder. [Nullary constructor tests](../javapurs/javapurs/test/nullary-constructors.mjs). These are Java classes, not a benefit of Valhalla or ADTs automatically emitted as sealed Java records.
- **12–14**: [RecordShapes](../javapurs/javapurs/src/Javapurs/RecordShapes.purs), [RecordTypes](../javapurs/javapurs/src/Javapurs/RecordTypes.purs), [RecordPrinter](../javapurs/javapurs/src/Javapurs/RecordPrinter.purs). Proven closed rows → an `AbstractMap` class with fixed fields; `Int` → `int`, nested record → Map reference, other contents → `Object`. Specialized access/copying, Map fallbacks for unknown values/FFI and overly wide shapes. No native typing of all fields or general intermodule monomorphization. [Record tests](../javapurs/javapurs/test/typed-records.mjs).
- **17**: [CountedLoops](../javapurs/javapurs/src/Javapurs/CountedLoops.purs) recognizes an Int countdown to zero in steps of one whose body uses only supported primitives; the negative domain retains the fallback path. This does not transform every tail recursion into a `for` loop. [Tests](../javapurs/javapurs/test/counted-loops.mjs).
- **20**: [LoopInvariants](../javapurs/javapurs/src/Javapurs/LoopInvariants.purs) and [PureInvariants](../javapurs/javapurs/src/Javapurs/PureInvariants.purs) cache closed, deeply pure calls with Int results on their first successful use during an invocation. No sharing across branches/occurrences, no global cache, no anticipation of effects/exceptions. [Tests](../javapurs/javapurs/test/loop-invariants.mjs).
- **21, 23–24, 27**: no dedicated thunk producer fusion pass; arrays use `Object[]`, no ownership mechanism for recycling ADT nodes, no backend elimination of Ref/ST cells. Any allocation elimination by HotSpot is a separate layer.

## psgo

- The [runner](bin/psgo/run) invokes the binary from PATH. The relevant code is [purescript-native/CodeGen/IL](../purescript-native/src/CodeGen/IL.hs); this is not Gopurs under another name.
- **01–03, 08, 18**: [Optimizer](../purescript-native/src/CodeGen/IL/Optimizer.hs) and [Inliner](../purescript-native/src/CodeGen/IL/Optimizer/Inliner.hs): local reductions, recognized standard operators/dictionaries, IIFEs, composition, and dead code. [MagicDo](../purescript-native/src/CodeGen/IL/Optimizer/MagicDo.hs) handles recognized bind/pure/discard forms; this does not validate all its other branches.
- **27**: an old `inlineST` pass exists and is invoked, but `isSTFunc` expects `Indexer(StringLiteral, Var)`, whereas `qualifiedToIL` emits `Indexer(Var, Var)` for qualified functions. No corresponding intermediate conversion was found. **It is therefore not credited to the current pipeline**. This case illustrates why the presence of a pass, even one that is invoked, is insufficient without checking the forms it receives.
- **04–07, 09–14, 23**: parameters and variables remain `Any`, ADTs/records are `Dict`, and arrays are generic: [Printer](../purescript-native/src/CodeGen/IL/Printer.hs). Explicit `Fn0…Fn10`/EffFn functions are uncurried, but ordinary functions and even saturated constructors still go through `Apply`; hence the red ratings for 07 and 26.
- **11**: global constants use `Once`, which also shares some nullary objects; there are no primitive enums.
- **15–17**: [TCO](../purescript-native/src/CodeGen/IL/Optimizer/TCO.hs) transforms only calls to the function's own name. The loop retains a step function and `Any` values. No mutual groups or counter canonicalization.
- **20–24**: composition inlining does not constitute Lazy producer fusion or callback specialization with a worker. No invariant caching, type-selected primitive arrays, or proven destructive reuse. **25**: newtype applications and patterns are erased in `IL.hs`.
- **29**: the IL generator retains the scrutinee followed by its tag/field tests. The audited local passes do not fuse ADT construction with the consuming `case`; composition inlining does not provide this transformation.

## Official JS

- [JS runner](bin/js/run), standard purs `0.15.15` found in PATH. [CoreImp optimizer](../purescript/src/Language/PureScript/CoreImp/Optimizer.hs): local inlining, eta reduction, IIFEs, thunk elimination, standard operators, local dead code, and MagicDo. The local fork's files retain these mechanisms; their TAST additions are not credited to the standard JS generator.
- **03, 07–08, 22**: known operators/dictionaries are optimized, but ordinary functions retain curried unary calls. No general automatic worker generation or HOF specialization comparable to the kernels in the other columns. FnN API wrappers are a separate explicit case.
- **10–14, 25–26**: [CodeGen/JS 0.15.15](https://github.com/purescript/purescript/blob/v0.15.15/src/Language/PureScript/CodeGen/JS.hs), native properties, named constructors, shared `.value` for nullary constructors, direct `new` when saturated, erased newtypes. `ObjectUpdate` with a known field list becomes a new, explicitly constructed object; otherwise, it uses a key-based copy loop. No primitive array or ADT payload layout selected from the TAST.
- **15–16**: [TCO 0.15.15](https://github.com/purescript/purescript/blob/v0.15.15/src/Language/PureScript/CoreImp/Optimizer/TCO.hs), self-calls and certain nested-function dependencies in tail position. The mutual case is yellow for this local scope; there is no general dispatcher for distinct top-level groups like ES's.
- **18, 27**: [MagicDo](../purescript/src/Language/PureScript/CoreImp/Optimizer/MagicDo.hs), recognized Effect sequences and STRefs confined to a `run` are lowered to local variables. Escaping refs remain objects.
- No dedicated pass for recursive counters, invariant caching, Lazy producer fusion, or ADT uniqueness. V8 may subsequently optimize the machine code, but that does not provide these passes in the backend.
- **29**: the generator retains the scrutinee followed by constructor tests (`instanceof`) and field reads. The audited CoreImp simplifications do not perform ADT construction/deconstruction fusion or case-of-case; any subsequent elimination by V8 belongs to the runtime.

## ES

- Sources for **the exact 1.4.3 release**: [Semantics](https://github.com/aristanetworks/purescript-backend-optimizer/blob/5e7643253ebc9db16f3c9ab702fa756bd3e41b80/src/PureScript/Backend/Optimizer/Semantics.purs), [Convert ES](https://github.com/aristanetworks/purescript-backend-optimizer/blob/5e7643253ebc9db16f3c9ab702fa756bd3e41b80/backend-es/src/PureScript/Backend/Optimizer/Codegen/EcmaScript/Convert.purs). This version does not contain our TAST/TypeApp/monomorphizer.
- **07, 22**: `shouldUncurryAbs` automatically uncurries **local** lambdas used at a single full arity. [Snapshot without FnN](https://github.com/aristanetworks/purescript-backend-optimizer/blob/5e7643253ebc9db16f3c9ab702fa756bd3e41b80/backend-es/test/snapshots-out/Snapshot.UncurriedLocalAbs01.js). Ordinary globals remain curried. The [CPS fusion](https://github.com/aristanetworks/purescript-backend-optimizer/blob/5e7643253ebc9db16f3c9ab702fa756bd3e41b80/backend-es/test/snapshots-out/Snapshot.Fusion01.js) achieved through directives/inlining is real, with intermediates still present; this is not a general pass for Lazy fusion or for every map/filter/fold chain.
- **11–14, 26**: constructors/products and native JS properties; shared nullary values. Residual updates use a generic spread: this release does not retain `copyFields` in its `ExprUpdate`. Eliminating an update when the entire record is already known belongs to 01, not to a dedicated closed-record copy path in 14. The JS engine determines machine layouts; there is no general typed-array selection from `Array Int`.
- **15–16**: `codegenTcoMutualLoopBindings` emits a common dispatcher, a branch number, shared arguments, and entry wrappers. [Mutual recursion snapshot](https://github.com/aristanetworks/purescript-backend-optimizer/blob/5e7643253ebc9db16f3c9ab702fa756bd3e41b80/backend-es/test/snapshots-out/Snapshot.Tco04.js). This capability is actually present, whereas Gopurs and Javapurs do not currently emit it.
- **18, 27**: PBO's shared `Codegen/Tco` analysis provides the `total` and `readWrite` usage counts used by the ES generator. `canUnboxRef` requires all uses to be read/write; local cell → mutable variable, with an object fallback when the ref escapes. [ST snapshot](https://github.com/aristanetworks/purescript-backend-optimizer/blob/5e7643253ebc9db16f3c9ab702fa756bd3e41b80/backend-es/test/snapshots-out/Snapshot.STRun01.js). Unboxing a cell does not prove a statically typed machine Int local.
- **17, 20, 24**: Effect/ST intrinsics emit loops but do not canonicalize counted recursion; `floatLet` is not LICM; no ownership of immutable values is emitted.

## purescm

- [Builder](../purescm/src/PureScript/Backend/Chez/Builder.purs) invokes its PBO and reads standard CoreFn. [Convert](../purescm/src/PureScript/Backend/Chez/Convert.purs) and [Syntax](../purescm/src/PureScript/Backend/Chez/Syntax.purs): curried ordinary global applications, explicit uncurried FnN functions, let/letrec, effect chains, and native `fx`/`fl` operations. Inlinable HOFs benefit from PBO; no TAST-guided clones or general global worker generation.
- **07**: the [PBO pinned](../purescm/spago.yaml) to commit `54b82ac23143dd47bfafaf52dec81dd7ccf2b490` enables `shouldUncurryAbs` for local lambdas used at a single full arity. `RewriteUncurry` produces the `UncurriedAbs` / `UncurriedApp` consumed by `Convert`; [Syntax](../purescm/src/PureScript/Backend/Chez/Syntax.purs), `mkUncurriedFn` / `runUncurriedFn`, actually emits a multi-argument lambda and call. The cell is therefore **🟡 (PBO)**; ordinary global functions remain curried.
- **10–11, 26**: ADTs use `define-record-type`, dedicated accessors, nullary constructors as quoted symbols, and `List` as Scheme pairs/lists; direct saturated construction. The [RBTree output](run/bak/scm/output/Test.RBTree/lib.ss) confirms this representation. Scheme record fields remain generic language values, not layouts per payload instantiation.
- **12–14, 23**: PureScript records use association lists; [runtime](../purescm/lib/purescm/runtime.ss), `record-ref`/`record-remove`/`record-set`. Arrays are generic flexvectors. Do not confuse the `define-record-type` forms **for ADTs** with the representation of **PureScript records**.
- **15–17, 20–21, 24, 27**: Scheme tail calls ensure bounded stack usage; no counter canonicalization, invariant caching, dedicated thunk fusion, destructive reuse proven through uniqueness, or emitted ref scalar replacement. `EffectRefNew` remains a Scheme box.

## purerl

- [CodeGen](../purerl/src/Language/PureScript/Erl/CodeGen.hs) and [Optimizer](../purerl/src/Language/PureScript/Erl/CodeGen/Optimizer.hs): fixed-point simplifications, recognized standard dictionaries, MagicDo, unused-function elimination. **07**: `generateFunctionOverloads`, arity computation, saturated global calls, and wrappers; some overloads nevertheless just call the curried version. The [Polymorphism output](run/bak/erl/output/Test.Polymorphism/test_polymorphism@ps.erl) shows this fallback.
- Erlang `-spec` declarations provide type information, but not a native layout per TAST instantiation. ADTs → tagged tuples, nullary constructors → literals, erased newtypes, direct saturated constructors. Records → maps, `maps:get` access, `EMapUpdate` updates; neither specialized closed records nor type-selected numeric arrays.
- **15–16, 18**: native BEAM tail calls; [TCO output](run/bak/erl/output/Test.TCO/test_tCO@ps.erl) with the direct function `deepTailRec/2`. MagicDo flattening is limited to recognized forms.
- **19–20**: the `?MEMOIZE` macro in the output **does not prove active memoization**: the `erlc` runner does not define `PURERL_MEMOIZE`, so the macro is the identity. [Memoize](../purerl/src/Language/PureScript/Erl/CodeGen/Optimizer/Memoize.hs) adds the annotation; this does not provide invariant caching. The limited upstream CSE, however, is real.
- No dedicated pass for counted loops, thunk producer fusion, HOF specialization beyond the standard inliners, primitive arrays, ADT ownership, or elimination of non-escaping ST cells was identified in this pipeline.
- **29**: the generated code retains ADT tuples and `ECaseOf` expressions. The Maybe/Either inliners construct cases without fusing them with an ADT producer. The MagicDo rewrite that pushes a thunk call into the branches of a case handles a different form; it is insufficient to credit this row.

## Wasm

- [Runner](bin/wasm/run): 14 core cases, Node ≥22, standard frontend **0.15.16**, import checks that prohibit replacing algorithms with JS. This backend uses **its own MIR**, not our PBO or our TAST. It nevertheless uses externs and representation inference: the red ratings for 04/05 do not mean “no types available”.
- **01–03, 18, 22, 25**: [MiddleEnd](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/MiddleEnd.purs), [Simplify](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/MiddleEnd/Optimize/Simplify.purs), [DictElim](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/MiddleEnd/Optimize/DictElim.purs), [Impurify](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/MiddleEnd/Optimize/Impurify.purs), reachability. [Specialize](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/MiddleEnd/Optimize/Specialize.purs) creates workers for a known function argument that remains invariant during recursion, with explicit captures; this is not monomorphization of ForAll types.
- **07–10, 26**: [Lower](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/Lower.purs), `RCallKnown` and direct constructors; [Unbox](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/Lower/Unbox.purs), fixed-point signatures, i32/f64 locals, an intermodule/generic ABI that retains boxes. [Collect](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/Lower/Collect.purs) uses externs for ADT fields, with a boxed fallback.
- **11–14, 23**: [Codegen](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/Codegen.purs), entirely nullary enums as i31ref, ADT GC structs, records as parallel label/value arrays with key-based lookup, generic eqref arrays. Known `copyFields` allows explicit record reconstruction, without an unboxed numeric payload. [Intrinsics](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/Intrinsics.purs) exposes explicit i32/i64 arrays, without general `Array Int` specialization.
- **15–17**: `return_call` for direct tail calls with a compatible return representation, including between mutually recursive top-level functions. Tail calls through `RApply` remain excluded. No dedicated counter canonicalization pass was identified.
- **19–21, 24, 27**: `SShared` sharing and [CAFs](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/Codegen/Caf.purs) lifted to global scope outside cycles. Neither constitutes Javapurs's per-invocation cache. No dedicated thunk producer fusion, ownership of unique values, or elimination of Ref cells: [Prim](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/Codegen/Prim.purs) retains the GC RefNew/Read/Write helpers.
- **29**: [DictElim](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/MiddleEnd/Optimize/DictElim.purs) selects `Semantics.normalize` with `useNbE = true`. [Semantics](../purescript-backend-wasm/compiler/src/PureScript/Backend/Wasm/MiddleEnd/Optimize/Semantics.purs), `evalCase` / `selectAlt` / `matchSem`, reduces a case whose constructors are known and substitutes their fields. This reduction is partial and specific to the Wasm MIR; it does not constitute general fusion through an opaque producer.

## Host compiler and runtime — separate from the matrix's passes

| Backend | Configuration observed in runners / generated projects | Implications for interpretation |
|---|---|---|
| sharpurs | F# project targeting `net8.0`, Release | Tail calls, final DU representation, JIT, and GC also depend on F#/.NET. |
| purust | Cargo Release, but generated profile uses `opt-level=1`, `debug=true` | Do not assume O3/LTO/PGO. LLVM optimizations are separate from Purust's code generation. |
| phpurs | OPcache/JIT enabled by the runner, mode `1255`, 128 MB buffer | The JIT and zvals are not Phpurs passes. |
| gopurs | Go; `GOGC=800`; PGO only with `--pgo` | Go compiler choices, escape analysis, and GC contribute to the result; PGO is not enabled by default. |
| javapurs | Java on the JVM; harness includes warm-up | The JIT may inline, devirtualize, and eliminate some allocations; this audit does not prove it does so at every site. |
| psgo | Go; `GOGC=1000` | The Go compiler and its GC contribute to the result; these host mechanisms are not psgo passes. |
| JS / ES | Node/V8 | Numeric representations, hidden classes, and speculative optimizations belong to the engine. |
| purescm | Chez, `--optimize-level 3` | The Scheme compiler and its native tail calls provide additional optimizations. |
| purerl | `erlc`, then BEAM | Tail calls and term management are provided by Erlang/BEAM. |
| Wasm | MIR optimizations enabled, Binaryen including O3 after assembly, Node/WasmGC | The optimization matrix does not attempt to inventory every internal Binaryen/V8 pass. |

Verifiable references: [Sharpurs runner](bin/sharp/run), [Purust runner](bin/rust/run), [Phpurs runner](bin/php/run), [Gopurs runner](bin/go/run), [Javapurs runner](bin/java/run), [generated Cargo profile](../purust/purust/src/Main.purs), [F# project and application runtime](../sharpurs/sharpurs/src/Main.purs), [other runners](bin), [Wasm/Binaryen pipeline](../purescript-backend-wasm/purs-wasm/src/PursWasm/CLI/Build.purs).

## Work items identified by the audit

| Backend | Concrete gap or extension to investigate | Distinction to preserve |
|---|---|---|
| sharpurs | Extend consumption of optimized PBO output and native layouts beyond the kernels | The thunk kernel removes wrappers but still builds the chain and repeatedly calls the worker. |
| purust | Primitive record fields and arrays; extend typed paths | Memory reuse already exists; an `Option<UnknownType>` field is not a native Int field. |
| phpurs | Closed record layouts; reduce remaining closure costs; proven invariants | Fusion and enum/nullable regions are already integrated; the scratch Church cache is not. |
| gopurs | Mutual TCO dispatcher; extend native representations across slice/ADT boundaries | Monomorphization already exists; `Rc` does not prove that ADT reuse is active. |
| javapurs | Extend workers, type ADT payloads and arrays | Invariant caching and Int records already exist; current workers still use `Object`. |
| psgo | Align ST patterns with the generated IR; typed representations and automatic saturated calls | FnN is explicit; the historical ST pass does not recognize the qualified references being emitted. |
| JS / ES | Evaluate broader global uncurrying and remaining intermediate values | ES already has mutual TCO, local uncurrying, and ref scalarization. |
| purescm | PureScript record representation and ordinary curried global calls | PBO can already uncurry local lambdas; PureScript records remain association lists. |
| purerl | Cost of curried overloads and residual dictionaries | Tail calls are native; the runner's MEMOIZE macro is not an active cache. |
| Wasm | Record layouts with fixed fields; extend unboxed signatures and collections | Externs and inference already exist; TAST and type-based monomorphization remain to be integrated. |

These are **coverage gaps**, not promised numerical gains or a measured priority order. Before implementing one: choose a concrete case, inspect its current output, then compare a short experiment against the corresponding README baseline.
