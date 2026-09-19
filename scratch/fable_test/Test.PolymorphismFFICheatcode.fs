[<AutoOpen>]
module PureScript_Test_PolymorphismFFICheatcode

open System
open System.Collections.Generic

module Test_PolymorphismFFICheatcode_FFI =
    let runPolymorphismFFICheatcode (n: obj) =
        let n' = unbox<int> n
        let mutable sum = 0
        for i in 1 .. n' do sum <- sum + 1
        sum :> obj
    

let Test_PolymorphismFFICheatcode_runPolymorphismFFICheatcode = box (fun (arg0: obj) -> box (Test_PolymorphismFFICheatcode_FFI.``runPolymorphismFFICheatcode`` (unbox arg0)))


let Test_PolymorphismFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Polymorphism FFICheatcode (10M Type Class Dict Lookups):"))))

let Test_PolymorphismFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10000000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_PolymorphismFFICheatcode_runPolymorphismFFICheatcode))) (box ((box dummy)))))))))))))))
