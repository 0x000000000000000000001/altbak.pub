module NativeBench

// Handwritten F# counterparts of tmp/bench_ocaml.ml and tmp/bench_haskell.hs.
// Like those native references, ListOps/ArrayOps, Church, Polymorphism, State
// and LazyEvaluation use their numerical loops; RowToList returns the known
// field count. They do not recreate PureScript closures or type dictionaries.
// ASTs, nested immutable records, sieve lists and red-black trees retain the
// reference data structures and algorithms. No sharpurs output is used here.

type Expr =
    | Val of int
    | Add of Expr * Expr
    | Mul of Expr * Expr
    | Sub of Expr * Expr

let rec private evalAst = function
    | Val value -> value
    | Add (left, right) -> evalAst left + evalAst right
    | Mul (left, right) -> evalAst left * evalAst right
    | Sub (left, right) -> evalAst left - evalAst right

let rec private buildAst = function
    | 0 -> Val 1
    | n -> Add (Mul (Val n, buildAst (n - 1)), Sub (buildAst (n - 1), Val 1))

let runAstTree limit = evalAst (buildAst limit)

let rec private fib = function
    | 0 -> 0
    | 1 -> 1
    | n -> fib (n - 1) + fib (n - 2)

let runFib limit = fib limit

let runListOps limit =
    let mutable sum = 0
    for i = 1 to limit do
        if i % 2 = 0 then
            sum <- sum + i
    sum

let runTCO limit =
    let rec go n acc =
        if n <= 0 then acc
        else go (n - 1) (acc + n % 3)
    go limit 0

type DictE = { e: int; f: int }
type DictC = { c: int; d: DictE }
type DictA = { a: int; b: DictC }

let runRecords limit =
    let mutable record = { a = 0; b = { c = 0; d = { e = 0; f = 0 } } }
    for n = limit downto 1 do
        let oldD = record.b.d
        let newD = { e = oldD.e + 3; f = oldD.f + n % 5 }
        let oldC = record.b
        let newC = { c = oldC.c + 2; d = newD }
        record <- { a = record.a + 1; b = newC }
    record.b.d.f

let rec private ack m n =
    if m = 0 then n + 1
    elif m > 0 && n = 0 then ack (m - 1) 1
    else ack (m - 1) (ack m (n - 1))

let runAckermann limit = ack limit 4

let runChurch limit =
    let count = limit * limit * limit * limit * limit
    let mutable acc = 0
    for _ = 1 to count do
        acc <- acc + 1
    acc

type IntList = Nil | Cons of int * IntList

let private listRange start finish =
    let rec go current acc =
        if current < start then acc
        else go (current - 1) (Cons (current, acc))
    go finish Nil

let private listFilter predicate values =
    let rec reverse remaining acc =
        match remaining with
        | Nil -> acc
        | Cons (value, rest) -> reverse rest (Cons (value, acc))
    let rec go remaining acc =
        match remaining with
        | Nil -> reverse acc Nil
        | Cons (value, rest) ->
            if predicate value then go rest (Cons (value, acc))
            else go rest acc
    go values Nil

let private listSum values =
    let rec go remaining acc =
        match remaining with
        | Nil -> acc
        | Cons (value, rest) -> go rest (acc + value)
    go values 0

let rec private sieve = function
    | Nil -> Nil
    | Cons (prime, rest) ->
        Cons (prime, sieve (listFilter (fun value -> value % prime <> 0) rest))

let runPrimes limit = listSum (sieve (listRange 2 limit))

type Color = R | B
type Tree = E | T of Color * Tree * int * Tree

// Named child matches avoid Fable 5.17.2's nested-pattern variable shadowing.
// The four rotations and their tree allocations match the native references.
let private balanceRight left value right =
    match right with
    | T (R, rightLeft, rightValue, rightRight) ->
        match rightLeft with
        | T (R, b, y, c) ->
            T (R, T (B, left, value, b), y, T (B, c, rightValue, rightRight))
        | _ ->
            match rightRight with
            | T (R, c, z, d) ->
                T (R, T (B, left, value, rightLeft), rightValue, T (B, c, z, d))
            | _ -> T (B, left, value, right)
    | _ -> T (B, left, value, right)

let private balance color left value right =
    match color with
    | R -> T (color, left, value, right)
    | B ->
        match left with
        | T (R, leftLeft, leftValue, leftRight) ->
            match leftLeft with
            | T (R, a, x, b) ->
                T (R, T (B, a, x, b), leftValue, T (B, leftRight, value, right))
            | _ ->
                match leftRight with
                | T (R, b, y, c) ->
                    T (R, T (B, leftLeft, leftValue, b), y, T (B, c, value, right))
                | _ -> balanceRight left value right
        | _ -> balanceRight left value right

let rec private ins value tree =
    match tree with
    | E -> T (R, E, value, E)
    | T (color, left, current, right) ->
        if value < current then balance color (ins value left) current right
        elif value > current then balance color left current (ins value right)
        else T (color, left, current, right)

let private insert value tree =
    match ins value tree with
    | T (_, left, current, right) -> T (B, left, current, right)
    | E -> E

let rec private depth = function
    | E -> 0
    | T (_, left, _, right) ->
        let leftDepth = depth left
        let rightDepth = depth right
        if leftDepth > rightDepth then 1 + leftDepth else 1 + rightDepth

let runRBTree limit =
    let mutable tree = E
    for value = limit downto 1 do
        tree <- insert value tree
    depth tree

let runPolymorphism limit =
    let mutable acc = 0
    for _ = 1 to limit do
        acc <- acc + 1
    acc

let runStateMonad limit =
    let mutable state = 0
    for _ = 1 to 20 do
        for _ = 1 to limit do
            state <- state + 1
    state

let runLazyEvaluation limit =
    let mutable acc = 0
    for _ = 1 to limit do
        acc <- acc + 1000
    acc

let runArrayOps limit =
    let mutable sum = 0
    for i = 1 to limit do
        if i % 2 = 0 then
            sum <- sum + i
    sum

let runRowToList (_: int) = 5
