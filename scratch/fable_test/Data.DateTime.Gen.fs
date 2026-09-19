[<AutoOpen>]
module PureScript_Data_DateTime_Gen

open System
open System.Collections.Generic

let Data_DateTime_Gen_genDateTime  = (box (fun (dictMonadGen: obj) -> (let Bind1 = (sharpurs_apply (box ((Map.find "Bind1" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Monad0" (unbox<Map<string, obj>> ((box dictMonadGen)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined)))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((sharpurs_apply (box ((Map.find "Apply0" (unbox<Map<string, obj>> ((box Bind1)))))) (box ((box Prim_undefined))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Apply0" (unbox<Map<string, obj>> ((box Bind1)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined))))))))) (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_DateTime_DateTimeusd_Ctor(usd__arg1, usd__arg2)))))))))))) (box ((sharpurs_apply (box ((box Data_Date_Gen_genDate))) (box ((box dictMonadGen)))))))))))) (box ((sharpurs_apply (box ((box Data_Time_Gen_genTime))) (box ((box dictMonadGen))))))))))
