type expr =
  | Val of int
  | Add of expr * expr
  | Mul of expr * expr
  | Sub of expr * expr

let rec evalAst = function
  | Val v -> v
  | Add (l, r) -> evalAst l + evalAst r
  | Mul (l, r) -> evalAst l * evalAst r
  | Sub (l, r) -> evalAst l - evalAst r

let rec buildTreeAst = function
  | 0 -> Val 1
  | n -> Add (Mul (Val n, buildTreeAst (n - 1)), Sub (buildTreeAst (n - 1), Val 1))

let runAstTree limit = evalAst (buildTreeAst limit)

let rec fib = function
  | 0 -> 0
  | 1 -> 1
  | n -> fib (n - 1) + fib (n - 2)

let runFib limit = fib limit

type 'a lst = Nil | Cons of 'a * 'a lst

let lrange start_ end_ =
  let rec go curr acc =
    if curr < start_ then acc else go (curr - 1) (Cons (curr, acc))
  in go end_ Nil

(* Test.ListOps intentionally keeps the even elements in reverse order. *)
let filterEvens xs =
  let rec go rest acc = match rest with
    | Nil -> acc
    | Cons (x, tail) ->
        if x mod 2 = 0 then go tail (Cons (x, acc)) else go tail acc
  in go xs Nil

let rec lfoldl f acc = function
  | Nil -> acc
  | Cons (x, xs) -> lfoldl f (f acc x) xs

let runListOps limit = lfoldl ( + ) 0 (filterEvens (lrange 1 limit))

let runTCO limit =
  let rec go n acc =
    if n = 0 then acc
    else go (n - 1) (acc + (n mod 3))
  in go limit 0

type dictE = { e : int; f : int }
type dictC = { c : int; d : dictE }
type dictA = { a : int; b : dictC }

let initial = { a = 0; b = { c = 0; d = { e = 0; f = 0 } } }

let rec updateRec n r =
  if n = 0 then r
  else updateRec (n - 1)
    { a = r.a + 1;
      b = { c = r.b.c + 2;
            d = { e = r.b.d.e + 3; f = r.b.d.f + (n mod 5) } } }

let runRecords limit = (updateRec limit initial).b.d.f

let rec ack m n =
  if m = 0 then n + 1
  else if n = 0 then ack (m - 1) 1
  else ack (m - 1) (ack m (n - 1))

let runAckermann limit = ack limit 4

type 'a church = ('a -> 'a) -> 'a -> 'a

let zeroC : 'a church = fun _ x -> x
let succC (n : 'a church) : 'a church = fun f x -> f (n f x)
let addC (m : 'a church) (n : 'a church) : 'a church =
  fun f x -> m f (n f x)
let mulC (m : 'a church) (n : 'a church) : 'a church =
  fun f x -> m (n f) x

let rec fromInt n : int church =
  if n = 0 then zeroC else succC (fromInt (n - 1))

let toInt (n : int church) = n (fun x -> x + 1) 0
let c10 n = fromInt n
let c100 n = mulC (c10 n) (c10 n)
let c10k n = mulC (c100 n) (c100 n)
let c100k n = mulC (c10k n) (c10 n)
let runChurch limit = toInt (c100k limit)

let lfilter p xs =
  let rec rev l acc = match l with
    | Nil -> acc
    | Cons (x, rest) -> rev rest (Cons (x, acc))
  in
  let rec go l acc = match l with
    | Nil -> rev acc Nil
    | Cons (x, rest) -> if p x then go rest (Cons (x, acc)) else go rest acc
  in go xs Nil

let lsum xs =
  let rec go l acc = match l with
    | Nil -> acc
    | Cons (x, rest) -> go rest (acc + x)
  in go xs 0

let rec sieve = function
  | Nil -> Nil
  | Cons (p, rest) -> Cons (p, sieve (lfilter (fun x -> x mod p <> 0) rest))

let runPrimes limit = lsum (sieve (lrange 2 limit))

type color = R | B
type tree = E | T of color * tree * int * tree

let balance c l v r = match (c, l, r) with
  | (B, T (R, T (R, a, x, b), y, cc), _) -> T (R, T (B, a, x, b), y, T (B, cc, v, r))
  | (B, T (R, a, x, T (R, b, y, cc)), _) -> T (R, T (B, a, x, b), y, T (B, cc, v, r))
  | (B, _, T (R, T (R, b, y, cc), z, d)) -> T (R, T (B, l, v, b), y, T (B, cc, z, d))
  | (B, _, T (R, b, y, T (R, cc, z, d))) -> T (R, T (B, l, v, b), y, T (B, cc, z, d))
  | _ -> T (c, l, v, r)

let rec ins x t = match t with
  | E -> T (R, E, x, E)
  | T (c, l, y, r) ->
      if x < y then balance c (ins x l) y r
      else if x > y then balance c l y (ins x r)
      else T (c, l, y, r)

let insert x t = match ins x t with
  | T (_, l, y, r) -> T (B, l, y, r)
  | E -> E

let rec depth = function
  | E -> 0
  | T (_, l, _, r) ->
      let ld = depth l in
      let rd = depth r in
      if ld > rd then 1 + ld else 1 + rd

let rec buildTree n acc =
  if n = 0 then acc else buildTree (n - 1) (insert n acc)

let runRBTree limit = depth (buildTree limit E)

(* An explicit typed dictionary represents Test.Polymorphism.Monoidish. *)
type 'a monoidish = { mempty_ : 'a; mappend_ : 'a -> 'a -> 'a }

let intMonoidish = { mempty_ = 1; mappend_ = ( + ) }

let polyLoop dict n_init acc_init =
  let rec go n acc =
    if n = 0 then acc else go (n - 1) (dict.mappend_ acc dict.mempty_)
  in go n_init acc_init

let runPolymorphism limit = polyLoop intMonoidish limit 0

(* PureScript's State newtype is erased; retain its function and result record. *)
type ('s, 'a) state_result = { value : 'a; state : 's }
type ('s, 'a) state = 's -> ('s, 'a) state_result

let runState (f : ('s, 'a) state) s = f s

let bindState (f : ('s, 'a) state) (g : 'a -> ('s, 'b) state)
    : ('s, 'b) state =
  fun s ->
    let r1 = f s in
    let next = g r1.value in
    next r1.state

let pureState (a : 'a) : ('s, 'a) state = fun s -> { value = a; state = s }
let get : ('s, 's) state = fun s -> { value = s; state = s }
let put (s : 's) : ('s, unit) state = fun _ -> { value = (); state = s }
let modify f = bindState get (fun s -> put (f s))

let rec chainModifications n =
  if n = 0 then pureState ()
  else bindState (modify (fun x -> x + 1))
    (fun () -> chainModifications (n - 1))

let rec runStateManyTimes n acc =
  if n = 0 then acc
  else runStateManyTimes (n - 1)
    (acc + (runState (chainModifications 60) 0).state)

let runStateMonad limit = runStateManyTimes limit 0

(* These thunks are deliberately non-memoizing, like Test.LazyEvaluation. *)
type 'a thunk = unit -> 'a
let defer (f : unit -> 'a) : 'a thunk = f
let force (f : 'a thunk) = f ()

let rec buildThunks n acc =
  if n = 0 then acc
  else buildThunks (n - 1) (defer (fun () -> force acc + 1))

let rec runLazyManyTimes n acc =
  if n = 0 then acc
  else runLazyManyTimes (n - 1)
    (acc + force (buildThunks 1000 (defer (fun () -> 0))))

let runLazyEvaluation limit = runLazyManyTimes limit 0

let arrayRange first last =
  let step = if first <= last then 1 else -1 in
  Array.init (abs (last - first) + 1) (fun i -> first + step * i)

(* OCaml's Array has no filter. Keep range/filter/fold as separate operations;
   this local buffer implements the same fresh-array contract as Data.Array. *)
let arrayFilter p xs =
  if Array.length xs = 0 then [||]
  else
    let buffer = Array.make (Array.length xs) xs.(0) in
    let rec copy i written =
      if i = Array.length xs then Array.sub buffer 0 written
      else
        let x = xs.(i) in
        if p x then begin
          buffer.(written) <- x;
          copy (i + 1) (written + 1)
        end else copy (i + 1) written
    in copy 0 0

let runArrayOps limit =
  Array.fold_left ( + ) 0 (arrayFilter (fun x -> x mod 2 = 0) (arrayRange 1 limit))

(* A heterogeneous row and its type-indexed dictionary replace RowToList.
   The shared row index enforces that the dictionary describes the record.
   As in PureScript, keys counts the dictionary, without reading field values. *)
type a_label = ALabel
type b_label = BLabel
type c_label = CLabel
type d_label = DLabel
type e_label = ELabel

type _ record_row =
  | RowNil : unit record_row
  | RowCons : 'label * 'field * 'tail record_row ->
      ('label * 'field * 'tail) record_row

type _ record_keys =
  | KeysNil : unit record_keys
  | KeysCons : 'tail record_keys -> ('label * 'field * 'tail) record_keys

let rec keysImpl : type row. row record_keys -> unit -> int =
  fun dict () -> match dict with
  | KeysNil -> 0
  | KeysCons tail -> 1 + keysImpl tail ()

let keys (type row) (dict : row record_keys) (_ : row record_row) = keysImpl dict ()

let runRowToList _ =
  let record = RowCons (ALabel, 1,
    RowCons (BLabel, "two",
      RowCons (CLabel, true,
        RowCons (DLabel, 4.0, RowCons (ELabel, "five", RowNil))))) in
  let dict = KeysCons (KeysCons (KeysCons (KeysCons (KeysCons KeysNil)))) in
  keys dict record

external monotonic_ns : unit -> int64 = "altbak_monotonic_ns"

let consume act arg =
  ignore (Sys.opaque_identity (act (Sys.opaque_identity arg)))

let time_batch act arg iterations =
  let start = monotonic_ns () in
  for _ = 1 to iterations do
    consume act arg
  done;
  let finish = monotonic_ns () in
  Int64.to_float (Int64.sub finish start) /. 1000.0

let calibrate act arg =
  let rec go iterations =
    if time_batch act arg iterations >= 10000.0 || iterations >= 16777216 then
      iterations
    else
      go (iterations * 2)
  in
  go 1

let bench name act arg =
  Printf.printf "--------------------------------------------------\n\n(Test)\n%s\n\n(Output & Warm-up)\n" name;
  let res = Sys.opaque_identity (act (Sys.opaque_identity arg)) in
  Printf.printf "%d\n" res;
  consume act arg;
  consume act arg;

  let iterations = calibrate act arg in
  let min_dur = ref infinity in
  for _ = 1 to 10 do
    let d = time_batch act arg iterations /. float_of_int iterations in
    if d < !min_dur then min_dur := d
  done;

  Printf.printf "\n(Execution time - best of 10)\n\n%.6f μs\nBatch iterations: %d\n\n" !min_dur iterations;
  !min_dur

let () =
  let check_only = match Array.to_list Sys.argv with
    | [_] -> false
    | [_; "--check-only"] -> true
    | _ -> invalid_arg "usage: benchmark [--check-only]"
  in
  let lAst = 3 in
  let lFib = 10 in
  let lList = 900 in
  let lTCO = 100000 in
  let lRec = 10000 in
  let lAck = 3 in
  let lChur = 10 in
  let lPri = 500 in
  let lRB = 100000 in
  let lPoly = 10000000 in
  let lState = 20 in
  let lLazy = 1000 in
  let lArr = 900 in
  let lRow = 10000 in

  if check_only then begin
    let check key act arg =
      Printf.printf "%s=%d\n" key (Sys.opaque_identity (act (Sys.opaque_identity arg)))
    in
    check "AstTree" runAstTree lAst;
    check "Fib" runFib lFib;
    check "ListOps" runListOps lList;
    check "TCO" runTCO lTCO;
    check "Records" runRecords lRec;
    check "Ackermann" runAckermann lAck;
    check "Church" runChurch lChur;
    check "Primes" runPrimes lPri;
    check "RBTree" runRBTree lRB;
    check "Polymorphism" runPolymorphism lPoly;
    check "StateMonad" runStateMonad lState;
    check "LazyEvaluation" runLazyEvaluation lLazy;
    check "ArrayOps" runArrayOps lArr;
    check "RowToList" runRowToList lRow
  end else begin
  Printf.printf "Global warm-up in progress...\n";
  for _ = 1 to 3 do
    consume runAstTree lAst;
    consume runFib lFib;
    consume runListOps lList;
    consume runTCO lTCO;
    consume runRecords lRec;
    consume runAckermann lAck;
    consume runChurch lChur;
    consume runPrimes lPri;
    consume runRBTree lRB;
    consume runPolymorphism lPoly;
    consume runStateMonad lState;
    consume runLazyEvaluation lLazy;
    consume runArrayOps lArr;
    consume runRowToList lRow
  done;

  let t1 = bench "AST Evaluation:" runAstTree lAst in
  let t2 = bench "Fibonacci:" runFib lFib in
  let t3 = bench "List Processing (900 elements):" runListOps lList in
  let t4 = bench "Tail Call Optimization (100k calls):" runTCO lTCO in
  let t5 = bench "Deep Record Updates (10k iterations):" runRecords lRec in
  let t6 = bench "Ackermann (3, 4):" runAckermann lAck in
  let t7 = bench "Church Numerals (100k Closure Applications):" runChurch lChur in
  let t8 = bench "Prime Sieve (sum primes up to 500):" runPrimes lPri in
  let t9 = bench "Red-Black Tree (100k Worst-Case Insertions):" runRBTree lRB in
  let t10 = bench "Polymorphism (10M Type Class Dict Lookups):" runPolymorphism lPoly in
  let t11 = bench "State Monad (1.2k Binds, 60 Stack Depth):" runStateMonad lState in
  let t12 = bench "Lazy Evaluation (1M Thunks Forced, 1k Depth):" runLazyEvaluation lLazy in
  let t13 = bench "Array Processing (900 elements):" runArrayOps lArr in
  let t14 = bench "RowToList (Keys Count):" runRowToList lRow in
  let total_us = t1 +. t2 +. t3 +. t4 +. t5 +. t6 +. t7 +. t8 +.
    t9 +. t10 +. t11 +. t12 +. t13 +. t14
  in
  Printf.printf "\n==================================================\n\nTotal exec time: %.6f ms\n" (total_us /. 1000.0)
  end
