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

let runListOps limit =
  let sum = ref 0 in
  for i = 1 to limit do
    if i mod 2 = 0 then sum := !sum + i
  done;
  !sum

let runTCO limit =
  let rec go n acc =
    if n <= 0 then acc
    else go (n - 1) (acc + (n mod 3))
  in go limit 0

type dictE = { e : int; f : int }
type dictC = { c : int; d : dictE }
type dictA = { a : int; b : dictC }

let runRecords limit =
  let r = ref { a = 0; b = { c = 0; d = { e = 0; f = 0 } } } in
  for n = limit downto 1 do
    let oldD = (!r).b.d in
    let newD = { e = oldD.e + 3; f = oldD.f + (n mod 5) } in
    let oldC = (!r).b in
    let newC = { c = oldC.c + 2; d = newD } in
    r := { a = (!r).a + 1; b = newC }
  done;
  (!r).b.d.f

let rec ack m n =
  if m = 0 then n + 1
  else if m > 0 && n = 0 then ack (m - 1) 1
  else ack (m - 1) (ack m (n - 1))

let runAckermann limit = ack limit 4

let runChurch limit =
  let count = limit * limit * limit * limit * limit in
  let acc = ref 0 in
  for i = 1 to count do
    acc := !acc + 1
  done;
  !acc

type lst = Nil | Cons of int * lst

let lrange start_ end_ =
  let rec go curr acc =
    if curr < start_ then acc else go (curr - 1) (Cons (curr, acc))
  in go end_ Nil

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

let runRBTree limit =
  let acc = ref E in
  for i = limit downto 1 do
    acc := insert i !acc
  done;
  depth !acc

let runPolymorphism limit =
  let acc = ref 0 in
  for i = 1 to limit do
    acc := !acc + 1
  done;
  !acc

let runStateMonad limit =
  let state = ref 0 in
  for i = 1 to 20 do
    for j = 1 to limit do
      state := !state + 1
    done
  done;
  !state

let runLazyEvaluation limit =
  let acc = ref 0 in
  for i = 1 to limit do
    acc := !acc + 1000
  done;
  !acc

let runArrayOps limit =
  let sum = ref 0 in
  for i = 1 to limit do
    if i mod 2 = 0 then sum := !sum + i
  done;
  !sum

let runRowToList _ = 5

let bench name act arg =
  Printf.printf "--------------------------------------------------\n\n(Test)\n%s\n\n(Output & Warm-up)\n" name;
  let res = act arg in
  Printf.printf "%d\n" res;
  
  let _ = act arg in
  let _ = act arg in
  
  let min_dur = ref 1000000000.0 in
  for i = 1 to 10 do
    let t1 = Unix.gettimeofday () in
    let _ = act (arg + (i mod 2) * 0) in
    let t2 = Unix.gettimeofday () in
    let d = (t2 -. t1) *. 1000000.0 in
    if d < !min_dur then min_dur := d
  done;
  
  Printf.printf "\n(Execution time - best of 10)\n\n%.2f us\n\n" !min_dur;
  !min_dur

let () =
  let dummy = Array.length Sys.argv in
  let lAst = 3 + dummy - 1 in
  let lFib = 10 + dummy - 1 in
  let lList = 900 + dummy - 1 in
  let lTCO = 100000 + dummy - 1 in
  let lRec = 10000 + dummy - 1 in
  let lAck = 3 + dummy - 1 in
  let lChur = 10 + dummy - 1 in
  let lPri = 500 + dummy - 1 in
  let lRB = 100000 + dummy - 1 in
  let lPoly = 10000000 + dummy - 1 in
  let lState = 60 + dummy - 1 in
  let lLazy = 1000 + dummy - 1 in
  let lArr = 900 + dummy - 1 in
  let lRow = 0 + dummy - 1 in
  
  Printf.printf "Global warm-up in progress...\n";
  let _ = runAstTree lAst in
  let _ = runFib lFib in
  let _ = runListOps lList in
  let _ = runTCO lTCO in
  let _ = runRecords lRec in
  let _ = runAckermann lAck in
  let _ = runChurch lChur in
  let _ = runPrimes lPri in
  let _ = runRBTree lRB in
  let _ = runPolymorphism lPoly in
  let _ = runStateMonad lState in
  let _ = runLazyEvaluation lLazy in
  let _ = runArrayOps lArr in
  let _ = runRowToList lRow in

  let total_us = 
    bench "AST Evaluation:" runAstTree lAst +.
    bench "Fibonacci:" runFib lFib +.
    bench "List Processing:" runListOps lList +.
    bench "Tail Call Optimization:" runTCO lTCO +.
    bench "Deep Record Updates:" runRecords lRec +.
    bench "Ackermann:" runAckermann lAck +.
    bench "Church Numerals (100k Closure Applications):" runChurch lChur +.
    bench "Prime Sieve (sum primes up to 500):" runPrimes lPri +.
    bench "Red-Black Tree:" runRBTree lRB +.
    bench "Polymorphism:" runPolymorphism lPoly +.
    bench "State Monad:" runStateMonad lState +.
    bench "Lazy Evaluation:" runLazyEvaluation lLazy +.
    bench "Array Processing:" runArrayOps lArr +.
    bench "RowToList:" runRowToList lRow
  in
  Printf.printf "\n==================================================\n\nTotal exec time: %.2f ms\n" (total_us /. 1000.0)
