[<AutoOpen>]
module PureScript_Test_RBTreeFFICheatcode

open System
open System.Collections.Generic

module Test_RBTreeFFICheatcode_FFI =
    let runRBTreeFFICheatcode (n: obj) =
        let n' = unbox<int> n
        let set = System.Collections.Generic.SortedSet<int>()
        for i in 1 .. n' do set.Add(i) |> ignore
        set.Count :> obj
    

let Test_RBTreeFFICheatcode_runRBTreeFFICheatcode = box (fun (arg0: obj) -> box (Test_RBTreeFFICheatcode_FFI.``runRBTreeFFICheatcode`` (unbox arg0)))


let Test_RBTreeFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Red-Black Tree FFICheatcode (100k Worst-Case Insertions):"))))

let Test_RBTreeFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 100000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RBTreeFFICheatcode_runRBTreeFFICheatcode))) (box ((box dummy)))))))))))))))
