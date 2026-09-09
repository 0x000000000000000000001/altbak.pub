[<AutoOpen>]
module PureScript_Test_StateMonadFFI

open System
open System.Collections.Generic

module Test_StateMonadFFI_FFI =
    type State<'s, 'a> = State of ('s -> 'a * 's)
    let runState (State f) s = f s
    let bind (State f) g = State (fun s -> let a, s' = f s in runState (g a) s')
    let pure' a = State (fun s -> a, s)
    let get = State (fun s -> s, s)
    let put s = State (fun _ -> (), s)
    let modify f = bind get (fun s -> put (f s))
    let rec chain n = if n = 0 then pure' () else bind (modify (fun x -> x + 1)) (fun _ -> chain (n - 1))
    let runStateMonadFFI (n: obj) =
        let rec loop i acc =
            if i = 0 then acc
            else
                let _, s' = runState (chain 60) 0
                loop (i - 1) (acc + s')
        loop (unbox<int> n) 0 :> obj
    

let Test_StateMonadFFI_runStateMonadFFI = box (fun (arg0: obj) -> box (Test_StateMonadFFI_FFI.``runStateMonadFFI`` (unbox arg0)))


let Test_StateMonadFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "State Monad FFI (1.2k Binds, 60 Stack Depth):"))))

let Test_StateMonadFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 60))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_StateMonadFFI_runStateMonadFFI))) (box ((box dummy)))))))))))))))
