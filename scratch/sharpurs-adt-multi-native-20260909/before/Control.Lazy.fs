[<AutoOpen>]
module PureScript_Control_Lazy

open System
open System.Collections.Generic

let Control_Lazy_Lazyusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Lazy_lazyUnit  = (sharpurs_apply (box ((box Control_Lazy_Lazyusd_Dict))) (box ((box ((Map.add "defer" (box ((box (fun (v: obj) -> (box Data_Unit_unit))))) Map.empty))))))

let Control_Lazy_lazyFn  = (sharpurs_apply (box ((box Control_Lazy_Lazyusd_Dict))) (box ((box ((Map.add "defer" (box ((box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit)))))) (box ((box x)))))))))) Map.empty))))))

let Control_Lazy_defer  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "defer" (unbox<Map<string, obj>> ((box v))))))))

let Control_Lazy_fix  = (box (fun (dictLazy: obj) -> (box (fun (f: obj) -> (
                                                                                                                                                                                                        
                                                                                                                                                                                                        let rec go : obj = ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Lazy_defer))) (box ((box dictLazy)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box f))) (box ((box go)))))))))) 
                                                                                                                                                                                                        in
                                                                                                                                                                                                        (box go)
                                                                                                                                                                                                        )))))
