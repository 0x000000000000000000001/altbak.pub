[<AutoOpen>]
module PureScript_Control_Bind

open System
open System.Collections.Generic

module Control_Bind_FFI =
    let arrayBind =
        fun (xs: obj) -> fun (f: obj) ->
            let arr = unbox<obj[]> xs
            let result = System.Collections.Generic.List<obj>()
            for x in arr do
                let res = unbox<obj[]> (sharpurs_apply f x)
                result.AddRange(res)
            result.ToArray() :> obj
    

let Control_Bind_arrayBind = box (Control_Bind_FFI.``arrayBind``)


let Control_Bind_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Control_Bind_Bindusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Bind_Discardusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Bind_discard  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "discard" (unbox<Map<string, obj>> ((box v))))))))

let Control_Bind_bindProxy  = (sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Type_Proxy_Proxyusd_Ctor))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Control_Apply_applyProxy))))) Map.empty)))))))

let Control_Bind_bindFn  = (sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box (fun (m: obj) -> (box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box f))) (box ((sharpurs_apply (box ((box m))) (box ((box x))))))))) (box ((box x)))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Control_Apply_applyFn))))) Map.empty)))))))

let Control_Bind_bindArray  = (sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box Control_Bind_arrayBind))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Control_Apply_applyArray))))) Map.empty)))))))

let Control_Bind_bind  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "bind" (unbox<Map<string, obj>> ((box v))))))))

let Control_Bind_bindFlipped  = (box (fun (dictBind: obj) -> (sharpurs_apply (box ((box Data_Function_flip))) (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box dictBind)))))))))

let Control_Bind_composeKleisliFlipped  = (box (fun (dictBind: obj) -> (box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bindFlipped))) (box ((box dictBind)))))) (box ((box f)))))) (box ((sharpurs_apply (box ((box g))) (box ((box a)))))))))))))))

let Control_Bind_composeKleisli  = (box (fun (dictBind: obj) -> (box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box dictBind)))))) (box ((sharpurs_apply (box ((box f))) (box ((box a))))))))) (box ((box g))))))))))))

let Control_Bind_discardProxy  = (sharpurs_apply (box ((box Control_Bind_Discardusd_Dict))) (box ((box ((Map.add "discard" (box ((box (fun (dictBind: obj) -> (sharpurs_apply (box ((box Control_Bind_bind))) (box ((box dictBind)))))))) Map.empty))))))

let Control_Bind_discardUnit  = (sharpurs_apply (box ((box Control_Bind_Discardusd_Dict))) (box ((box ((Map.add "discard" (box ((box (fun (dictBind: obj) -> (sharpurs_apply (box ((box Control_Bind_bind))) (box ((box dictBind)))))))) Map.empty))))))

let Control_Bind_ifM  = (box (fun (dictBind: obj) -> (box (fun (cond: obj) -> (box (fun (t: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box dictBind)))))) (box ((box cond)))))) (box ((box (fun (cond_prime: obj) -> (match ((unbox ((box cond_prime)))) with | LitBool true () -> ((box t)) | _ -> ((box f))))))))))))))))

let Control_Bind_join  = (box (fun (dictBind: obj) -> (box (fun (m: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box dictBind)))))) (box ((box m)))))) (box ((box Control_Bind_identity))))))))
