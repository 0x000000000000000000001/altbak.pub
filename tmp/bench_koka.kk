module bench_koka

import std/time
import std/time/duration
import std/num/float64
import std/os/env

// Direct functional ports of src/Test/*.purs. Mutation is confined to the
// timing harness. Koka may optimize/reuse these immutable values itself.

// 1. AstTree
type expr
  Val( v : int )
  Add( l : expr, r : expr )
  Mul( l : expr, r : expr )
  Sub( l : expr, r : expr )

fun evalAst(e: expr): div int
  match e
    Val(v) -> v
    Add(l, r) -> evalAst(l) + evalAst(r)
    Mul(l, r) -> evalAst(l) * evalAst(r)
    Sub(l, r) -> evalAst(l) - evalAst(r)

fun buildTreeAst(n: int): div expr
  if n == 0 then Val(1) else
    Add(Mul(Val(n), buildTreeAst(n - 1)), Sub(buildTreeAst(n - 1), Val(1)))

fun runAstTree(limit: int): div int
  evalAst(buildTreeAst(limit))

// 2. Fib
fun fib(n: int): div int
  if n == 0 then 0
  elif n == 1 then 1
  else fib(n - 1) + fib(n - 2)

fun runFib(limit: int): div int
  fib(limit)

// 3. ListOps: the same strict range, reversed even list and generic fold.
type benchList<a>
  LNil
  LCons(head: a, tail: benchList<a>)

fun lrange(start: int, end: int): div benchList<int>
  fun go(curr: int, acc: benchList<int>): div benchList<int>
    if curr < start then acc else go(curr - 1, LCons(curr, acc))
  go(end, LNil)

fun filterEvens(xs: benchList<int>): div benchList<int>
  fun go(rest: benchList<int>, acc: benchList<int>): div benchList<int>
    match rest
      LNil -> acc
      LCons(x, tail) ->
        if x % 2 == 0 then go(tail, LCons(x, acc)) else go(tail, acc)
  go(xs, LNil)

fun lfoldl(f: (b, a) -> b, acc: b, xs: benchList<a>): div b
  match xs
    LNil -> acc
    LCons(x, rest) -> lfoldl(f, f(acc, x), rest)

fun runListOps(limit: int): div int
  lfoldl(fn(acc, x) acc + x, 0, filterEvens(lrange(1, limit)))

// 4. TCO
fun deepTailRec(n: int, acc: int): div int
  if n == 0 then acc else deepTailRec(n - 1, acc + (n % 3))

fun runTCO(limit: int): div int
  deepTailRec(limit, 0)

// 5. Records: recursive immutable updates at all three nesting levels.
struct dictE(e: int, f: int)
struct dictC(c: int, d: dictE)
struct dictA(a: int, b: dictC)

fun updateRec(n: int, r: dictA): div dictA
  if n == 0 then r else
    val oldD = r.b.d
    val newD = oldD(e = oldD.e + 3, f = oldD.f + (n % 5))
    val oldC = r.b
    val newC = oldC(c = oldC.c + 2, d = newD)
    updateRec(n - 1, r(a = r.a + 1, b = newC))

fun runRecords(limit: int): div int
  updateRec(limit, DictA(0, DictC(0, DictE(0, 0)))).b.d.f

// 6. Ackermann
fun ack(m: int, n: int): div int
  if m == 0 then n + 1
  elif n == 0 then ack(m - 1, 1)
  else ack(m - 1, ack(m, n - 1))

fun runAckermann(limit: int): div int
  ack(limit, 4)

// 7. Church: curried higher-order numerals; no integer multiplication shortcut.
alias church<a> = (a -> div a) -> div (a -> div a)

fun zeroC(): church<a>
  fn(_) fn(x) x

fun succC(n: church<a>): church<a>
  fn(f) fn(x) f(n(f)(x))

fun addC(m: church<a>, n: church<a>): church<a>
  fn(f) fn(x) m(f)(n(f)(x))

fun mulC(m: church<a>, n: church<a>): church<a>
  fn(f) fn(x) m(n(f))(x)

fun fromInt(n: int): div church<int>
  if n == 0 then zeroC() else succC(fromInt(n - 1))

fun toInt(n: church<int>): div int
  n(fn(x) x + 1)(0)

fun c10(n: int): div church<int>
  fromInt(n)

fun c100(n: int): div church<int>
  mulC(c10(n), c10(n))

fun c10k(n: int): div church<int>
  mulC(c100(n), c100(n))

fun c100k(n: int): div church<int>
  mulC(c10k(n), c10(n))

fun runChurch(limit: int): div int
  toInt(c100k(limit))

// 8. Primes: strict allocated lists and a predicate capturing each prime.
fun lreverse(xs: benchList<a>): div benchList<a>
  fun go(rest, acc)
    match rest
      LNil -> acc
      LCons(x, tail) -> go(tail, LCons(x, acc))
  go(xs, LNil)

fun lfilter(p: a -> bool, xs: benchList<a>): div benchList<a>
  fun go(rest, acc)
    match rest
      LNil -> lreverse(acc)
      LCons(x, tail) -> if p(x) then go(tail, LCons(x, acc)) else go(tail, acc)
  go(xs, LNil)

fun lsum(xs: benchList<int>): div int
  fun go(rest: benchList<int>, acc: int): div int
    match rest
      LNil -> acc
      LCons(x, tail) -> go(tail, acc + x)
  go(xs, 0)

fun sieve(xs: benchList<int>): div benchList<int>
  match xs
    LNil -> LNil
    LCons(p, rest) -> LCons(p, sieve(lfilter(fn(x) x % p != 0, rest)))

fun runPrimes(limit: int): div int
  lsum(sieve(lrange(2, limit)))

// 9. RBTree
type color
  R
  B

type tree
  E
  T(c: color, l: tree, v: int, r: tree)

fun balance(c: color, l: tree, v: int, r: tree): tree
  match (c, l, r)
    (B, T(R, T(R, a, x, b), y, cc), _) -> T(R, T(B, a, x, b), y, T(B, cc, v, r))
    (B, T(R, a, x, T(R, b, y, cc)), _) -> T(R, T(B, a, x, b), y, T(B, cc, v, r))
    (B, _, T(R, T(R, b, y, cc), z, d)) -> T(R, T(B, l, v, b), y, T(B, cc, z, d))
    (B, _, T(R, b, y, T(R, cc, z, d))) -> T(R, T(B, l, v, b), y, T(B, cc, z, d))
    _ -> T(c, l, v, r)

fun ins(x: int, t: tree): div tree
  match t
    E -> T(R, E, x, E)
    T(c, l, y, r) ->
      if x < y then balance(c, ins(x, l), y, r)
      elif x > y then balance(c, l, y, ins(x, r))
      else T(c, l, y, r)

fun insert(x: int, t: tree): div tree
  match ins(x, t)
    T(_, l, y, r) -> T(B, l, y, r)
    E -> E

fun depth(t: tree): div int
  match t
    E -> 0
    T(_, l, _, r) -> 
      val ld = depth(l)
      val rd = depth(r)
      if ld > rd then 1 + ld else 1 + rd

fun buildTree(n: int, acc: tree): div tree
  if n == 0 then acc else buildTree(n - 1, insert(n, acc))

fun runRBTree(limit: int): div int
  depth(buildTree(limit, E))

// 10. Polymorphism: Koka has no Haskell-style classes. Pass the same generic
// dictionary explicitly; specialization and inlining remain compiler choices.
struct monoidish<a>(mempty: a, mappend: (a, a) -> a)

fun polyLoop(dict: monoidish<a>, n: int, acc: a): div a
  if n == 0 then acc else
    val append = dict.mappend
    polyLoop(dict, n - 1, append(acc, dict.mempty))

fun runPolymorphism(limit: int): div int
  polyLoop(Monoidish(1, fn(x, y) x + y), limit, 0)

// 11. StateMonad: a polymorphic state/result record and nested closures.
struct stateResult<s,a>(value: a, state: s)
struct state<s,a>(run: s -> div stateResult<s,a>)

fun runState(action: state<s,a>, s: s): div stateResult<s,a>
  val run = action.run
  run(s)

fun bindState(action: state<s,a>, next: a -> div state<s,b>): state<s,b>
  State(fn(s) {
    val r1 = runState(action, s)
    runState(next(r1.value), r1.state)
  })

fun pureState(a: a): state<s,a>
  State(fn(s) StateResult(a, s))

fun getState(): state<s,s>
  State(fn(s) StateResult(s, s))

fun putState(s: s): state<s,()>
  State(fn(_) StateResult((), s))

fun modifyState(f: s -> s): state<s,()>
  bindState(getState(), fn(s) putState(f(s)))

fun chainModifications(n: int): div state<int,()>
  if n == 0 then pureState(()) else
    bindState(modifyState(fn(x) x + 1), fn(_) chainModifications(n - 1))

fun stateManyTimes(n: int, acc: int): div int
  if n == 0 then acc else
    stateManyTimes(n - 1, acc + runState(chainModifications(60), 0).state)

fun runStateMonad(limit: int): div int
  stateManyTimes(limit, 0)

// 12. LazyEvaluation: explicit unit -> a thunks, deliberately not memoized.
struct benchLazy<a>(thunk: () -> div a)

fun deferThunk(f: () -> div a): benchLazy<a>
  BenchLazy(f)

fun forceThunk(lazyValue: benchLazy<a>): div a
  val thunk = lazyValue.thunk
  thunk()

fun buildThunks(n: int, acc: benchLazy<int>): div benchLazy<int>
  if n == 0 then acc else
    buildThunks(n - 1, deferThunk(fn() forceThunk(acc) + 1))

fun lazyManyTimes(n: int, acc: int): div int
  if n == 0 then acc else
    lazyManyTimes(n - 1, acc + forceThunk(buildThunks(1000, deferThunk(fn() 0))))

fun runLazyEvaluation(limit: int): div int
  lazyManyTimes(limit, 0)

// 13. ArrayOps: materialize native vectors for both range and filter.
// Koka's vector library has no filter/foldl. The generic filter stages retained
// elements in a strict list before constructing its result vector.
fun vectorFilter(p: a -> bool, xs: vector<a>): div vector<a>
  fun go(i, acc)
    match xs.at(i)
      Nothing -> acc.reverse.vector
      Just(x) -> if p(x) then go(i + 1, Cons(x, acc)) else go(i + 1, acc)
  go(0, [])

fun vectorFoldl(f: (b, a) -> b, acc: b, xs: vector<a>): div b
  fun go(i, result)
    match xs.at(i)
      Nothing -> result
      Just(x) -> go(i + 1, f(result, x))
  go(0, acc)

fun arrayRange(start: int, end: int): vector<int>
  if start <= end then vector-init(end - start + 1, fn(i) start + i)
  else vector-init(start - end + 1, fn(i) start - i)

fun runArrayOps(limit: int): div int
  vectorFoldl(fn(acc, x) acc + x, 0,
    vectorFilter(fn(x) x % 2 == 0, arrayRange(1, limit)))

// 14. RowToList: Koka has no row kinds. A typed heterogeneous product carries
// the record shape, and recursive keysNil/keysCons dictionaries count it.
// Field wrappers retain names and heterogeneous types, without dynamic boxing.
struct rowNil
struct rowField<a>(name: string, value: a)
struct rowCons<h,t>(head: rowField<h>, tail: t)
struct recordKeys<r>(keysImpl: () -> div int)

fun keysNil(): recordKeys<rowNil>
  RecordKeys(fn() 0)

fun keysCons(tail: recordKeys<t>): recordKeys<rowCons<h,t>>
  val impl = tail.keysImpl
  RecordKeys(fn() 1 + impl())

fun keys(dict: recordKeys<r>, record: r): div int
  val impl = dict.keysImpl
  impl()

fun runRowToList(limit: int): div int
  val record = RowCons(RowField("a", 1), RowCons(RowField("b", "two"),
    RowCons(RowField("c", True), RowCons(RowField("d", 4.0),
    RowCons(RowField("e", "five"), RowNil)))))
  keys(keysCons(keysCons(keysCons(keysCons(keysCons(keysNil()))))), record)

// Benchmarking Framework
// The ndet effect keeps these calls observable to Koka. Volatile C accesses
// keep inputs unknown and results live through the native optimizer as well.
extern opaque(value: int): ndet int
  c inline "({ volatile kk_integer_t bench_value = #1; bench_value; })"

extern consume(value: int): ndet ()
  c inline "({ volatile kk_integer_t bench_value = #1; kk_integer_drop(bench_value, kk_context()); kk_Unit; })"

fun runOnce(act, input)
  val result = act(opaque(input))
  consume(result)
  result

fun runBatch(act, input, count)
  val t1 = ticks()
  for(1, count) fn(_)
    runOnce(act, input)
    ()
  val t2 = ticks()
  (t2 - t1).nano-seconds.float64 / 1000.0

fun bench(name, act, input)
  println("--------------------------------------------------\n\n(Test)\n" ++ name)
  println("\n(Output & Warm-up)")
  println(runOnce(act, input).show)
  runOnce(act, input)
  runOnce(act, input)

  // Calibrate whole batches to at least 10 ms before measuring ten batches.
  var count := 1
  var duration := runBatch(act, input, count)
  while { duration < 10000.0 && count < 16777216 }
    count := count * 2
    duration := runBatch(act, input, count)

  var min_dur := 1.0e30
  for(1, 10) fn(_)
    val d = runBatch(act, input, count) / count.float64
    if d < min_dur then min_dur := d
  val us = min_dur
  println("\n(Execution time - best of 10)\n\n" ++ us.show-fixed(6) ++ " μs\n")
  println("Batch iterations: " ++ count.show ++ "\n")
  us

fun warmup()
  runOnce(runAstTree, 3)
  runOnce(runFib, 10)
  runOnce(runListOps, 900)
  runOnce(runTCO, 100000)
  runOnce(runRecords, 10000)
  runOnce(runAckermann, 3)
  runOnce(runChurch, 10)
  runOnce(runPrimes, 500)
  runOnce(runRBTree, 100000)
  runOnce(runPolymorphism, 10000000)
  runOnce(runStateMonad, 20)
  runOnce(runLazyEvaluation, 1000)
  runOnce(runArrayOps, 900)
  runOnce(runRowToList, 10000)
  ()

fun benchmarkAll()
  println("Global warm-up in progress...")
  for(1, 3) fn(_)
    warmup()

  var total_us := 0.0
  total_us := total_us + bench("AST Evaluation:", runAstTree, 3)
  total_us := total_us + bench("Fibonacci:", runFib, 10)
  total_us := total_us + bench("List Processing (900 elements):", runListOps, 900)
  total_us := total_us + bench("Tail Call Optimization (100k calls):", runTCO, 100000)
  total_us := total_us + bench("Deep Record Updates (10k iterations):", runRecords, 10000)
  total_us := total_us + bench("Ackermann (3, 4):", runAckermann, 3)
  total_us := total_us + bench("Church Numerals (100k Closure Applications):", runChurch, 10)
  total_us := total_us + bench("Prime Sieve (sum primes up to 500):", runPrimes, 500)
  total_us := total_us + bench("Red-Black Tree (100k Worst-Case Insertions):", runRBTree, 100000)
  total_us := total_us + bench("Polymorphism (10M Type Class Dict Lookups):", runPolymorphism, 10000000)
  total_us := total_us + bench("State Monad (1.2k Binds, 60 Stack Depth):", runStateMonad, 20)
  total_us := total_us + bench("Lazy Evaluation (1M Thunks Forced, 1k Depth):", runLazyEvaluation, 1000)
  total_us := total_us + bench("Array Processing (900 elements):", runArrayOps, 900)
  total_us := total_us + bench("RowToList (Keys Count):", runRowToList, 10000)

  println("\n==================================================\n")
  println("Total exec time: " ++ (total_us / 1000.0).show ++ " ms\n")

// Check exactly the same functions and opaque inputs, without warmups/timers.
fun checkOnly()
  println("AstTree=" ++ runOnce(runAstTree, 3).show)
  println("Fib=" ++ runOnce(runFib, 10).show)
  println("ListOps=" ++ runOnce(runListOps, 900).show)
  println("TCO=" ++ runOnce(runTCO, 100000).show)
  println("Records=" ++ runOnce(runRecords, 10000).show)
  println("Ackermann=" ++ runOnce(runAckermann, 3).show)
  println("Church=" ++ runOnce(runChurch, 10).show)
  println("Primes=" ++ runOnce(runPrimes, 500).show)
  println("RBTree=" ++ runOnce(runRBTree, 100000).show)
  println("Polymorphism=" ++ runOnce(runPolymorphism, 10000000).show)
  println("StateMonad=" ++ runOnce(runStateMonad, 20).show)
  println("LazyEvaluation=" ++ runOnce(runLazyEvaluation, 1000).show)
  println("ArrayOps=" ++ runOnce(runArrayOps, 900).show)
  println("RowToList=" ++ runOnce(runRowToList, 10000).show)

pub fun main()
  match get-args()
    [arg] | arg == "--check-only" -> checkOnly()
    _ -> benchmarkAll()
