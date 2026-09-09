[<AutoOpen>]
module PureScript_Test_PolymorphismFFI

open System
open System.Collections.Generic

module Test_PolymorphismFFI_FFI =
    type Show = { show: int -> string }
    let showInt : Show = { show = fun x -> string x }
    let print (dict: Show) x = dict.show x
    let runPolymorphismFFI (n: obj) =
        let mutable res = 0
        for i in 1 .. (unbox<int> n) do res <- res + 1
        res :> obj
    

let Test_PolymorphismFFI_runPolymorphismFFI = box (fun (arg0: obj) -> box (Test_PolymorphismFFI_FFI.``runPolymorphismFFI`` (unbox arg0)))


let Test_PolymorphismFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Polymorphism FFI (10M Type Class Dict Lookups):"))))

let Test_PolymorphismFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10000000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_PolymorphismFFI_runPolymorphismFFI))) (box ((box dummy)))))))))))))))
