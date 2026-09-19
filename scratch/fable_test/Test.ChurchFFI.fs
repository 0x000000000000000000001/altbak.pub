[<AutoOpen>]
module PureScript_Test_ChurchFFI

open System
open System.Collections.Generic

module Test_ChurchFFI_FFI =
    type Church = (int -> int) -> int -> int
    let rec fromInt n : Church =
        if n = 0 then fun _ value -> value
        else
            let previous = fromInt (n - 1)
            fun f value -> f (previous f value)
    let multiply (m: Church) (n: Church) : Church = fun f value -> m (n f) value
    let square n = multiply (fromInt n) (fromInt n)
    let runChurchFFI (input: obj) =
        let n = unbox<int> input
        let fourthPower = multiply (square n) (square n)
        let fifthPower = multiply fourthPower (fromInt n)
        fifthPower ((+) 1) 0 :> obj
    

let Test_ChurchFFI_runChurchFFI = box (fun (arg0: obj) -> box (Test_ChurchFFI_FFI.``runChurchFFI`` (unbox arg0)))


let Test_ChurchFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Church Numerals FFI (100k Closure Applications):"))))

let Test_ChurchFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_ChurchFFI_runChurchFFI))) (box ((box dummy)))))))))))))))
