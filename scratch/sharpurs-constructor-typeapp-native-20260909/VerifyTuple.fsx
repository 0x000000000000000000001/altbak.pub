#if BASELINE
#r "before/bin/Release/net8.0/Program.dll"
#else
#r "after/bin/Release/net8.0/Program.dll"
#endif

open System
open Sharpurs_Prelude
open PureScript_Data_Tuple

let mutable checks = 0
let check label condition =
    if not condition then failwith label
    checks <- checks + 1

let apply fn value = sharpurs_apply fn value
let payloads: obj list =
    [ box Int32.MinValue; box Int32.MaxValue; box "λ and text"
      box [|box 1; box 2|]; box (Map.ofList ["field", box 17]) ]

for left in payloads do
    // A retained first constructor stage is reused with distinct payloads.
    let partial = apply Data_Tuple_Tuple left
    for right in payloads do
        let original = apply partial right
        check "constructor keeps first payload identity" (Object.ReferenceEquals(apply Data_Tuple_fst original, left))
        check "constructor keeps second payload identity" (Object.ReferenceEquals(apply Data_Tuple_snd original, right))
        let swapped = apply Data_Tuple_swap original
        check "swap preserves second payload as first" (Object.ReferenceEquals(apply Data_Tuple_fst swapped, right))
        check "swap preserves first payload as second" (Object.ReferenceEquals(apply Data_Tuple_snd swapped, left))
        let restored = apply Data_Tuple_swap swapped
        check "swap twice restores first payload" (Object.ReferenceEquals(apply Data_Tuple_fst restored, left))
        check "swap twice restores second payload" (Object.ReferenceEquals(apply Data_Tuple_snd restored, right))

printfn "Generated Tuple constructor/swap: %d value, projection, partial reuse and payload-identity checks passed" checks
