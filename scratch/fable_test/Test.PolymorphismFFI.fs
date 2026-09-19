[<AutoOpen>]
module PureScript_Test_PolymorphismFFI

open System
open System.Collections.Generic

module Test_PolymorphismFFI_FFI =
    type Monoidish = { mempty: int; mappend: int -> int -> int }
    let rec polyLoop dictionary count acc =
        if count = 0 then acc
        else polyLoop dictionary (count - 1) (dictionary.mappend acc dictionary.mempty)
    let runPolymorphismFFI (n: obj) =
        polyLoop { mempty = 1; mappend = (+) } (unbox<int> n) 0 :> obj
    

let Test_PolymorphismFFI_runPolymorphismFFI = box (fun (arg0: obj) -> box (Test_PolymorphismFFI_FFI.``runPolymorphismFFI`` (unbox arg0)))


let Test_PolymorphismFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Polymorphism FFI (10M Type Class Dict Lookups):"))))

let Test_PolymorphismFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10000000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_PolymorphismFFI_runPolymorphismFFI))) (box ((box dummy)))))))))))))))
