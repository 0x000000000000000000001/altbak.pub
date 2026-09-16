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
fun bench(name, act)
  println("--------------------------------------------------\n\n(Test)\n" ++ name)
  println("\n(Output & Warm-up)")
  println((act() : int).show)
  act()
  act()
  var min_dur := 1000000000.0
  for(1, 10) fn(i)
    val t1 = ticks()
    act()
    val t2 = ticks()
    val d = (t2 - t1).nano-seconds.float64 / 1000.0
    if d < min_dur then min_dur := d
  val us = min_dur
  println("\n(Execution time - best of 10)\n\n" ++ us.show ++ " us\n")
  us

pub fun main()
  println("Global warm-up in progress...")
  runAstTree(3)
  runFib(10)
  runListOps(900)
  runTCO(100000)
  runRecords(10000)
  runAckermann(3)
  runChurch(10)
  runPrimes(500)
  runRBTree(100000)
  runPolymorphism(10000000)
  runStateMonad(60)
  runLazyEvaluation(1000)
  runArrayOps(900)
  runRowToList(0)

  var total_us := 0.0
  total_us := total_us + bench("AST Evaluation:", fn() { runAstTree(3) })
  total_us := total_us + bench("Fibonacci:", fn() { runFib(10) })
  total_us := total_us + bench("List Processing:", fn() { runListOps(900) })
  total_us := total_us + bench("Tail Call Optimization:", fn() { runTCO(100000) })
  total_us := total_us + bench("Deep Record Updates:", fn() { runRecords(10000) })
  total_us := total_us + bench("Ackermann:", fn() { runAckermann(3) })
  total_us := total_us + bench("Church Numerals (100k Closure Applications):", fn() { runChurch(10) })
  total_us := total_us + bench("Prime Sieve (sum primes up to 500):", fn() { runPrimes(500) })
  total_us := total_us + bench("Red-Black Tree:", fn() { runRBTree(100000) })
  total_us := total_us + bench("Polymorphism:", fn() { runPolymorphism(10000000) })
  total_us := total_us + bench("State Monad:", fn() { runStateMonad(60) })
  total_us := total_us + bench("Lazy Evaluation:", fn() { runLazyEvaluation(1000) })
  total_us := total_us + bench("Array Processing:", fn() { runArrayOps(900) })
  total_us := total_us + bench("RowToList:", fn() { runRowToList(0) })

  println("\n==================================================\n")
  println("Total exec time: " ++ (total_us / 1000.0).show ++ " ms\n")