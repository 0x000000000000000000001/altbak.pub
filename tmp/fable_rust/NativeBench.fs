module NativeBench

// Direct F# translations of src/Test/*.purs. Keep the source abstractions:
// immutable lists/arrays/records, higher-order Church/State/Lazy functions,
// generic dictionaries and a type-indexed row dictionary. The compiler may
// optimize them; the source must not replace them with numeric shortcuts.

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

type FpList<'a> = Nil | Cons of 'a * FpList<'a>

let private rangeList start finish =
    let rec go current acc =
        if current < start then acc
        else go (current - 1) (Cons (current, acc))
    go finish Nil

let private filterEvens values =
    let rec go remaining acc =
        match remaining with
        | Nil -> acc
        | Cons (value, rest) ->
            if value % 2 = 0 then go rest (Cons (value, acc))
            else go rest acc
    go values Nil

let rec private foldList f acc values =
    match values with
    | Nil -> acc
    | Cons (value, rest) -> foldList f (f acc value) rest

let runListOps limit = foldList (+) 0 (filterEvens (rangeList 1 limit))

let runTCO limit =
    let rec go n acc =
        if n = 0 then acc
        else go (n - 1) (acc + n % 3)
    go limit 0

type DictE = { e: int; f: int }
type DictC = { c: int; d: DictE }
type DictA = { a: int; b: DictC }

let runRecords limit =
    let rec update n record =
        if n = 0 then record
        else
            update (n - 1)
                { record with a = record.a + 1
                              b = { record.b with c = record.b.c + 2
                                                  d = { record.b.d with e = record.b.d.e + 3
                                                                        f = record.b.d.f + n % 5 } } }
    let initial = { a = 0; b = { c = 0; d = { e = 0; f = 0 } } }
    (update limit initial).b.d.f

let rec private ack m n =
    if m = 0 then n + 1
    elif n = 0 then ack (m - 1) 1
    else ack (m - 1) (ack m (n - 1))

let runAckermann limit = ack limit 4

type Church<'a> = ('a -> 'a) -> 'a -> 'a

let private zeroC : Church<'a> = fun _ x -> x
// Passing the function as an argument avoids stock Fable's Rust borrow/move
// conflict when the same function is both the callee and an inner argument.
let private applyChurchStep (f: 'a -> 'a) (x: 'a) = f x
let private succC (n: Church<'a>) : Church<'a> = fun f x -> applyChurchStep f (n f x)
let private mulC (m: Church<'a>) (n: Church<'a>) : Church<'a> = fun f x -> m (n f) x
let rec private fromInt n : Church<int> =
    if n = 0 then zeroC else succC (fromInt (n - 1))
let private c100 n = mulC (fromInt n) (fromInt n)
let private c10k n = mulC (c100 n) (c100 n)
let private c100k n = mulC (c10k n) (fromInt n)
let runChurch limit = c100k limit (fun x -> x + 1) 0

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
    let rec build n acc =
        if n = 0 then acc else build (n - 1) (insert n acc)
    depth (build limit E)

type Monoidish<'a> = { Mempty: 'a; Mappend: 'a -> 'a -> 'a }

let private polyLoop (dictionary: Monoidish<'a>) limit initial =
    let rec go n acc =
        if n = 0 then acc
        else go (n - 1) (dictionary.Mappend acc dictionary.Mempty)
    go limit initial

let private intMonoidish = { Mempty = 1; Mappend = fun x y -> x + y }
let runPolymorphism limit = polyLoop intMonoidish limit 0

type StateResult<'s, 'a> = { Value: 'a; State: 's }
type State<'s, 'a> = State of ('s -> StateResult<'s, 'a>)

let private runState (State f) s = f s
let private bindState (State f) g =
    State (fun s ->
        let r = f s
        let (State next) = g r.Value
        next r.State)
let private pureState a = State (fun s -> { Value = a; State = s })
let private getState<'s> : State<'s, 's> = State (fun s -> { Value = s; State = s })
let private putState s = State (fun _ -> { Value = (); State = s })
let private modifyState f = bindState getState (fun s -> putState (f s))
let rec private chainModifications n =
    if n = 0 then pureState ()
    else bindState (modifyState (fun x -> x + 1)) (fun _ -> chainModifications (n - 1))

let runStateMonad limit =
    let rec go n acc =
        if n = 0 then acc
        else go (n - 1) (acc + (runState (chainModifications 60) 0).State)
    go limit 0

// Explicit functions, not System.Lazy: PureScript's custom Lazy does not memoize.
type Thunk<'a> = Thunk of (unit -> 'a)
let private defer f = Thunk f
let private force (Thunk f) = f ()
let rec private buildThunks n acc =
    if n = 0 then acc
    else buildThunks (n - 1) (defer (fun () -> force acc + 1))

let runLazyEvaluation limit =
    let rec go n acc =
        if n = 0 then acc
        else go (n - 1) (acc + force (buildThunks 1000 (defer (fun () -> 0))))
    go limit 0

let runArrayOps limit =
    let values =
        if limit >= 1 then Array.init limit (fun i -> i + 1)
        else Array.init (2 - limit) (fun i -> 1 - i)
    values |> Array.filter (fun x -> x % 2 = 0) |> Array.fold (+) 0

// F# has no RowToList constraint. A heterogeneous row and its type-indexed
// recursive dictionary preserve keysNil/keysCons without hardcoding the count.
type RowNil = RowNil
type RowCons<'head, 'tail> = { Head: 'head; Tail: 'tail }
type RecordKeys<'row> = { KeysImpl: unit -> int }
// Construct the phantom-typed record in a generic scope: stock Fable's Rust
// emitter otherwise leaves its phantom type unbound at concrete constructors.
let private recordKeys<'row> implementation : RecordKeys<'row> = { KeysImpl = implementation }
let private keysNil : RecordKeys<RowNil> = recordKeys (fun () -> 0)
let private keysCons (tail: RecordKeys<'tail>) : RecordKeys<RowCons<'head, 'tail>> =
    recordKeys (fun () -> 1 + tail.KeysImpl ())
let private keys (dictionary: RecordKeys<'row>) (_record: 'row) = dictionary.KeysImpl ()

let runRowToList (_: int) =
    let record = { Head = 1; Tail = { Head = "two"; Tail = { Head = true; Tail = { Head = 4.0; Tail = { Head = "five"; Tail = RowNil } } } } }
    keys (keysCons (keysCons (keysCons (keysCons (keysCons keysNil))))) record
