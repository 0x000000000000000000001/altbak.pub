#r "after/bin/Release/net8.0/Program.dll"

open PureScript_Test_Polymorphism
open Sharpurs_Prelude

let runGeneric dictionary count initial =
    sharpurs_apply (sharpurs_apply (sharpurs_apply Test_Polymorphism_polyLoop dictionary) (box count)) initial

let strings : obj = box (Map.ofList [
    "mempty_", box "x"
    "mappend_", box (fun (left: obj) -> box (fun (right: obj) -> box (unbox<string> left + unbox<string> right)))
])

let mutable checkedCount = 0
let check label expected actual =
    if actual <> expected then failwithf "%s: expected %A, got %A" label expected actual
    checkedCount <- checkedCount + 1

check "generic Int" 17 (unbox<int> (runGeneric Test_Polymorphism_intMonoidish 10 (box 7)))
check "generic String" "prefixxxxxx" (unbox<string> (runGeneric strings 5 (box "prefix")))
check "generic String empty loop" "prefix" (unbox<string> (runGeneric strings 0 (box "prefix")))
check "generic Int after String" 9 (unbox<int> (runGeneric Test_Polymorphism_intMonoidish 10 (box -1)))
for _ in 1 .. 2 do
    check "native Int action" "10000000" (unbox<string> ((unbox<obj -> obj> Test_Polymorphism_act) null))
printfn "Real generated program: %d generic Int/String and repeated native action checks passed" checkedCount
