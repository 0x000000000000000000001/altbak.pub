[<AutoOpen>]
module PureScript_Control_Apply

open System
open System.Collections.Generic

module Control_Apply_FFI =
    let arrayApply (fs: obj) (xs: obj) : obj =
        let fsArr = unbox<obj[]> fs
        let xsArr = unbox<obj[]> xs
        let lenF = fsArr.Length
        let lenX = xsArr.Length
        let result = Array.zeroCreate (lenF * lenX)
        let mutable n = 0
        for i = 0 to lenF - 1 do
            let f = fsArr.[i]
            for j = 0 to lenX - 1 do
                result.[n] <- sharpurs_apply f xsArr.[j]
                n <- n + 1
        result :> obj
    

let Control_Apply_arrayApply = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Control_Apply_FFI.``arrayApply`` (unbox arg0) (unbox arg1))))


let Control_Apply_Applyusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Apply_applyProxy  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Type_Proxy_Proxyusd_Ctor))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Functor_functorProxy))))) Map.empty)))))))

let Control_Apply_applyFn  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box f))) (box ((box x)))))) (box ((sharpurs_apply (box ((box g))) (box ((box x))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Functor_functorFn))))) Map.empty)))))))

let Control_Apply_applyArray  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box Control_Apply_arrayApply))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Functor_functorArray))))) Map.empty)))))))

let Control_Apply_apply  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "apply" (unbox<Map<string, obj>> ((box v))))))))

let Control_Apply_applyFirst  = (box (fun (dictApply: obj) -> (let Functor0 = (sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictApply)))))) (box ((box Prim_undefined)))) in (box (fun (a: obj) -> (box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Functor0)))))) (box ((box Data_Function_const)))))) (box ((box a))))))))) (box ((box b)))))))))))

let Control_Apply_applySecond  = (box (fun (dictApply: obj) -> (let Functor0 = (sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictApply)))))) (box ((box Prim_undefined)))) in (box (fun (a: obj) -> (box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Functor0)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn)))))))))))) (box ((box a))))))))) (box ((box b)))))))))))

let Control_Apply_lift2  = (box (fun (dictApply: obj) -> (let Functor0 = (sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictApply)))))) (box ((box Prim_undefined)))) in (box (fun (f: obj) -> (box (fun (a: obj) -> (box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Functor0)))))) (box ((box f)))))) (box ((box a))))))))) (box ((box b)))))))))))))

let Control_Apply_lift3  = (box (fun (dictApply: obj) -> (let Functor0 = (sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictApply)))))) (box ((box Prim_undefined)))) in (box (fun (f: obj) -> (box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Functor0)))))) (box ((box f)))))) (box ((box a))))))))) (box ((box b))))))))) (box ((box c)))))))))))))))

let Control_Apply_lift4  = (box (fun (dictApply: obj) -> (let Functor0 = (sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictApply)))))) (box ((box Prim_undefined)))) in (box (fun (f: obj) -> (box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Functor0)))))) (box ((box f)))))) (box ((box a))))))))) (box ((box b))))))))) (box ((box c))))))))) (box ((box d)))))))))))))))))

let Control_Apply_lift5  = (box (fun (dictApply: obj) -> (let Functor0 = (sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictApply)))))) (box ((box Prim_undefined)))) in (box (fun (f: obj) -> (box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (box (fun (e: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Functor0)))))) (box ((box f)))))) (box ((box a))))))))) (box ((box b))))))))) (box ((box c))))))))) (box ((box d))))))))) (box ((box e)))))))))))))))))))
