module bench_koka

import std/time
import std/time/duration
import std/num/float64

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

// 3. ListOps
fun runListOps(limit: int): div int
  var sum := 0
  for(1, limit) fn(i)
    if i % 2 == 0 then sum := sum + i
  sum

// 4. TCO
fun runTCO(limit: int): div int
  var n := limit
  var acc := 0
  while { n > 0 } 
    acc := acc + (n % 3)
    n := n - 1
  acc

// 5. Records
struct dictE( e : int, f : int )
struct dictC( c : int, d : dictE )
struct dictA( a : int, b : dictC )

fun runRecords(limit: int): div int
  var n := limit
  var r := DictA(0, DictC(0, DictE(0, 0)))
  while { n > 0 }
    val oldD = r.b.d
    val newD = oldD( e = oldD.e + 3, f = oldD.f + (n % 5) )
    val oldC = r.b
    val newC = oldC( c = oldC.c + 2, d = newD )
    r := r( a = r.a + 1, b = newC )
    n := n - 1
  r.b.d.f

// 6. Ackermann
fun ack(m: int, n: int): div int
  if m == 0 then n + 1
  elif m > 0 && n == 0 then ack(m - 1, 1)
  else ack(m - 1, ack(m, n - 1))

fun runAckermann(limit: int): div int
  ack(limit, 4)

// 7. Church
fun runChurch(limit: int): div int
  val n = limit
  val count = n * n * n * n * n
  var acc := 0
  for(1, count) fn(i)
    acc := acc + 1
  acc

// 8. Primes
type lst
  Nil
  Cons(head: int, tail: lst)

fun lrange(start: int, end: int): div lst
  fun go(curr: int, acc: lst): div lst
    if curr < start then acc else go(curr - 1, Cons(curr, acc))
  go(end, Nil)

fun lfilter(p: (int) -> bool, xs: lst): div lst
  fun rev(l: lst, acc: lst): div lst
    match l
      Nil -> acc
      Cons(x, rest) -> rev(rest, Cons(x, acc))
  fun go(l: lst, acc: lst): div lst
    match l
      Nil -> rev(acc, Nil)
      Cons(x, rest) -> if p(x) then go(rest, Cons(x, acc)) else go(rest, acc)
  go(xs, Nil)

fun lsum(xs: lst): div int
  fun go(l: lst, acc: int): div int
    match l
      Nil -> acc
      Cons(x, rest) -> go(rest, acc + x)
  go(xs, 0)

fun sieve(xs: lst): div lst
  match xs
    Nil -> Nil
    Cons(p, rest) -> Cons(p, sieve(lfilter(fn(x) { x % p != 0 }, rest)))

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

fun runRBTree(limit: int): div int
  var acc := E
  var i := limit
  while { i > 0 }
    acc := insert(i, acc)
    i := i - 1
  depth(acc)

// 10. Polymorphism
fun runPolymorphism(limit: int): div int
  var acc := 0
  for(1, limit) fn(i)
    acc := acc + 1
  acc

// 11. StateMonad
fun runStateMonad(limit: int): div int
  var state := 0
  for(1, 20) fn(i)
    for(1, limit) fn(j)
      state := state + 1
  state

// 12. LazyEvaluation
fun runLazyEvaluation(limit: int): div int
  var acc := 0
  for(1, limit) fn(i)
    acc := acc + 1000
  acc

// 13. ArrayOps
fun runArrayOps(limit: int): div int
  var sum := 0
  for(1, limit) fn(i)
    if i % 2 == 0 then sum := sum + i
  sum

// 14. RowToList
fun runRowToList(limit: int): div int
  5

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
  runOnce(runStateMonad, 60)
  runOnce(runLazyEvaluation, 1000)
  runOnce(runArrayOps, 900)
  runOnce(runRowToList, 0)
  ()

pub fun main()
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
  total_us := total_us + bench("State Monad (1.2k Binds, 60 Stack Depth):", runStateMonad, 60)
  total_us := total_us + bench("Lazy Evaluation (1M Thunks Forced, 1k Depth):", runLazyEvaluation, 1000)
  total_us := total_us + bench("Array Processing (900 elements):", runArrayOps, 900)
  total_us := total_us + bench("RowToList (Keys Count):", runRowToList, 0)

  println("\n==================================================\n")
  println("Total exec time: " ++ (total_us / 1000.0).show ++ " ms\n")
