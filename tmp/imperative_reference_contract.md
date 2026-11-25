# Optimized imperative C reference

The main C column measures `bench_native.c`: an optimized imperative
implementation of the core workloads. Mutable data, native arrays, efficient
allocation and fused loops are allowed. It is not required to reproduce
PureScript's closures, dictionaries or persistent intermediate structures.
The separate functional C translation remains an archived experiment.

## Opaque inputs

The workload parameter must be obtained at run time on every timed invocation.
The compiler must not replace the input read with the literal used by the
benchmark driver, nor reuse an entire invocation's result across a batch.
The result is consumed, and expected values are verification data only.
This rule also applies to the other language columns; a source-level identity
function alone is not evidence that an optimizing compiler respects it.

Opacity does not forbid algebraic optimization of a function of its runtime
argument. For example, a loop may become a multiplication by the value just
read. It forbids replacing that value with the driver's known `1000` and
returning the precomputed answer. Multiple runtime inputs and inspection of
optimized output must distinguish these two situations.

Inputs and results follow the PureScript programs:

- AST: depth; build and evaluate the expression tree.
- Fibonacci: index; the recursive Fibonacci computation.
- List and Array: range endpoint; sum the even elements of the specified range.
  The array range descends when its endpoint is below its start; the custom
  list range is empty in that case.
- TCO: countdown; sum the remainders modulo three.
- Records: update count; compute the final nested field after the updates.
- Ackermann: first argument, with the second argument fixed at four.
- Church: base numeral; compute the composition's numeric result from it.
- Primes: upper bound; compute the prime sum.
- RBTree: descending insertion count; construct and balance the tree, then
  compute its depth. In-place Okasaki rotations are allowed; a formula returning
  the expected depth without constructing the tree is not a reference for this
  workload.
- Polymorphism: iteration count; compute the repeated integer operation.
- State: repetition count, with sixty state updates inside each repetition.
- Lazy: repetition count, with a chain depth of one thousand. An imperative
  computation may avoid materializing closures while computing the same value.
- RowToList: read the opaque value `10000`, then count the five-field record's
  keys. The PureScript source deliberately ignores that opaque value. A result
  independent of it is therefore expected; the count must come from the record
  structure, not a hard-coded oracle.

Integer operations preserve the Int32 results of the workloads without C
signed-overflow undefined behavior. Validation includes additional feasible
inputs, boundary cases and red-black-tree invariants, not just the fourteen
published answers.

## Timing and allocation

Construction, allocation and release of each invocation's working storage
belong to that timed invocation. There is no arena preallocated outside the
measurement or cached answer shared between calls. Stack storage and an arena
created within the invocation are legitimate imperative implementations.
This does not make C allocation and Go/JS garbage collection identical.

Build before measuring. Use a monotonic clock, three global warmups, three
local warmups, calibrated batches, the minimum of ten batches per process and
the median of three independent processes. Archive exact timings, inputs,
oracles, optimized-output checks, commands and source/executable hashes.
