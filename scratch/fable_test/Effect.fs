[<AutoOpen>]
module PureScript_Effect

open System
open System.Collections.Generic

module Effect_FFI =
    let pureE = box (fun (a: obj) -> box (fun _ -> a))
    let bindE = box (fun (a: obj) -> box (fun (f: obj) -> box (fun _ -> 
        let a' = a :?> (obj -> obj)
        let f' = f :?> (obj -> obj)
        let a_res = a' null
        let f_res = f' a_res :?> (obj -> obj)
        f_res null
    )))
    
    let untilE = box (fun (f: obj) -> box (fun _ ->
        let f' = f :?> (obj -> obj)
        let mutable condition = false
        while not condition do
            condition <- unbox<bool> (f' null)
        null
    ))
    
    let whileE = box (fun (f: obj) -> box (fun (a: obj) -> box (fun _ ->
        let f' = f :?> (obj -> obj)
        let a' = a :?> (obj -> obj)
        let mutable condition = unbox<bool> (f' null)
        while condition do
            a' null |> ignore
            condition <- unbox<bool> (f' null)
        null
    )))
    
    let forE = box (fun (lo: obj) -> box (fun (hi: obj) -> box (fun (f: obj) -> box (fun _ ->
        let f' = f :?> (obj -> obj)
        let l = unbox<int> lo
        let h = unbox<int> hi
        for i = l to h - 1 do
            let step = f' (box i) :?> (obj -> obj)
            step null |> ignore
        null
    ))))
    
    let foreachE = box (fun (arr: obj) -> box (fun (f: obj) -> box (fun _ ->
        let f' = f :?> (obj -> obj)
        let arr' = unbox<obj[]> arr
        for v in arr' do
            let step = f' v :?> (obj -> obj)
            step null |> ignore
        null
    )))
    

let Effect_bindE = box (Effect_FFI.``bindE``)
let Effect_forE = box (Effect_FFI.``forE``)
let Effect_foreachE = box (Effect_FFI.``foreachE``)
let Effect_pureE = box (Effect_FFI.``pureE``)
let Effect_untilE = box (Effect_FFI.``untilE``)
let Effect_whileE = box (Effect_FFI.``whileE``)


let rec Effect_monadEffect : obj = ((sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Effect_applicativeEffect))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Effect_bindEffect))))) Map.empty))))))))
 and Effect_bindEffect : obj = ((sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box Effect_bindE))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Effect_applyEffect))))) Map.empty))))))))
 and Effect_applyEffect : obj = ((sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((sharpurs_apply (box ((box Control_Monad_ap))) (box ((box Effect_monadEffect)))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Effect_functorEffect))))) Map.empty))))))))
 and Effect_applicativeEffect : obj = ((sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box Effect_pureE))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Effect_applyEffect))))) Map.empty))))))))
 and Effect_functorEffect : obj = ((sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((sharpurs_apply (box ((box Control_Applicative_liftA1))) (box ((box Effect_applicativeEffect)))))) Map.empty)))))))


let Effect_semigroupEffect  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_lift2))) (box ((box Effect_applyEffect)))))) (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup))))))))) Map.empty))))))))

let Effect_monoidEffect  = (box (fun (dictMonoid: obj) -> (let semigroupEffect1 = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_pureE))) (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box dictMonoid))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffect1))))) Map.empty))))))))))
