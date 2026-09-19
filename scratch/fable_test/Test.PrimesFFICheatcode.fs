[<AutoOpen>]
module PureScript_Test_PrimesFFICheatcode

open System
open System.Collections.Generic

module Test_PrimesFFICheatcode_FFI =
    let isPrime n =
        if n < 2 then false
        else
            let mutable p = true
            let mutable i = 2
            while p && i * i <= n do
                if n % i = 0 then p <- false
                i <- i + 1
            p
    let runPrimesFFICheatcode (n: obj) =
        let n' = unbox<int> n
        let mutable sum = 0
        for i in 2 .. n' do
            if isPrime i then sum <- sum + i
        sum :> obj
    

let Test_PrimesFFICheatcode_runPrimesFFICheatcode = box (fun (arg0: obj) -> box (Test_PrimesFFICheatcode_FFI.``runPrimesFFICheatcode`` (unbox arg0)))


let Test_PrimesFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Prime Sieve FFICheatcode (sum primes up to 500):"))))

let Test_PrimesFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 500))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_PrimesFFICheatcode_runPrimesFFICheatcode))) (box ((box dummy)))))))))))))))
