#r "after/bin/Release/net8.0/Program.dll"

open System
open Sharpurs_Prelude
open PureScript_Data_Maybe
open PureScript_Data_String_CodePoints

let mutable checks = 0
let check label condition =
    if not condition then failwith label
    checks <- checks + 1

let publicUnsurrogate high low =
    unbox<int> (sharpurs_apply (sharpurs_apply Data_String_CodePoints_unsurrogate (box high)) (box low))

let highSamples = [0xD800; 0xD801; 0xD83D; 0xDBFE; 0xDBFF]
let lowSamples = [0xDC00; 0xDC01; 0xDE00; 0xDFFE; 0xDFFF]
let pairs = [for high in highSamples do for low in lowSamples do yield high, low]

for high in highSamples do
    let partial = sharpurs_apply Data_String_CodePoints_unsurrogate (box high)
    for low in lowSamples do
        let expected = 0x10000 + (high - 0xD800) * 0x400 + (low - 0xDC00)
        check "public surrogate conversion" (publicUnsurrogate high low = expected)
        check "raw direct surrogate conversion"
            (unbox<int> (Data_String_CodePoints_unsurrogate_direct (box high) (box low)) = expected)
        check "guarded direct surrogate conversion"
            (unbox<int> (Data_String_CodePoints_unsurrogate_direct_apply (box high) (box low)) = expected)
        check "reused partial surrogate conversion" (unbox<int> (sharpurs_apply partial (box low)) = expected)

let uncons value =
    match unbox<Data_Maybe_Maybe> (sharpurs_apply Data_String_CodePoints_uncons (box value)) with
    | Data_Maybe_Nothingusd_Ctor -> None
    | Data_Maybe_Justusd_Ctor record ->
        let fields = unbox<Map<string, obj>> record
        Some (unbox<int> fields.["head"], unbox<string> fields.["tail"])

check "uncons empty" (uncons "" = None)

let validStrings = pairs |> List.map (fun (high, low) ->
    let text = String([| char high; char low |]) + "tail"
    let expected = 0x10000 + (high - 0xD800) * 0x400 + (low - 0xDC00)
    text, expected, "tail")

let otherStrings = [
    "A", int 'A', ""
    "Atail", int 'A', "tail"
    "é漢", 0xE9, "漢"
    String([| char 0xD800 |]), 0xD800, ""
    String([| char 0xDC00 |]), 0xDC00, ""
    String([| char 0xD800; 'A' |]), 0xD800, "A"
    String([| char 0xDC00; 'A' |]), 0xDC00, "A"
]

for text, expected, tail in validStrings @ otherStrings do
    check "uncons head and UTF-16 tail" (uncons text = Some (expected, tail))
    check "unsafe first code-point fallback"
        (unbox<int> (sharpurs_apply Data_String_CodePoints_unsafeCodePointAt0Fallback (box text)) = expected)

printfn "Generated CodePoints: %d surrogate, partial application and caller checks passed" checks
