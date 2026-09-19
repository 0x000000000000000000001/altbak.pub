[<AutoOpen>]
module PureScript_Data_BooleanAlgebra

open System
open System.Collections.Generic

let Data_BooleanAlgebra_heytingAlgebraRecord  = (sharpurs_apply (box ((box Data_HeytingAlgebra_heytingAlgebraRecord))) (box ((box Prim_undefined))))

let Data_BooleanAlgebra_BooleanAlgebraRecordusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_BooleanAlgebra_BooleanAlgebrausd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_BooleanAlgebra_booleanAlgebraUnit  = (sharpurs_apply (box ((box Data_BooleanAlgebra_BooleanAlgebrausd_Dict))) (box ((box ((Map.add "HeytingAlgebra0" (box ((box (fun (usd__unused: obj) -> (box Data_HeytingAlgebra_heytingAlgebraUnit))))) Map.empty))))))

let Data_BooleanAlgebra_booleanAlgebraRecordNil  = (sharpurs_apply (box ((box Data_BooleanAlgebra_BooleanAlgebraRecordusd_Dict))) (box ((box ((Map.add "HeytingAlgebraRecord0" (box ((box (fun (usd__unused: obj) -> (box Data_HeytingAlgebra_heytingAlgebraRecordNil))))) Map.empty))))))

let Data_BooleanAlgebra_booleanAlgebraRecordCons  = (box (fun (dictIsSymbol: obj) -> (let heytingAlgebraRecordCons = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_heytingAlgebraRecordCons))) (box ((box dictIsSymbol)))))) (box ((box Prim_undefined)))) in (box (fun (usd__unused: obj) -> (box (fun (dictBooleanAlgebraRecord: obj) -> (let heytingAlgebraRecordCons1 = (sharpurs_apply (box ((box heytingAlgebraRecordCons))) (box ((sharpurs_apply (box ((Map.find "HeytingAlgebraRecord0" (unbox<Map<string, obj>> ((box dictBooleanAlgebraRecord)))))) (box ((box Prim_undefined))))))) in (box (fun (dictBooleanAlgebra: obj) -> (let heytingAlgebraRecordCons2 = (sharpurs_apply (box ((box heytingAlgebraRecordCons1))) (box ((sharpurs_apply (box ((Map.find "HeytingAlgebra0" (unbox<Map<string, obj>> ((box dictBooleanAlgebra)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_BooleanAlgebra_BooleanAlgebraRecordusd_Dict))) (box ((box ((Map.add "HeytingAlgebraRecord0" (box ((box (fun (usd__unused: obj) -> (box heytingAlgebraRecordCons2))))) Map.empty)))))))))))))))))

let Data_BooleanAlgebra_booleanAlgebraRecord  = (box (fun (usd__unused: obj) -> (box (fun (dictBooleanAlgebraRecord: obj) -> (let heytingAlgebraRecord1 = (sharpurs_apply (box ((box Data_BooleanAlgebra_heytingAlgebraRecord))) (box ((sharpurs_apply (box ((Map.find "HeytingAlgebraRecord0" (unbox<Map<string, obj>> ((box dictBooleanAlgebraRecord)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_BooleanAlgebra_BooleanAlgebrausd_Dict))) (box ((box ((Map.add "HeytingAlgebra0" (box ((box (fun (usd__unused: obj) -> (box heytingAlgebraRecord1))))) Map.empty)))))))))))

let Data_BooleanAlgebra_booleanAlgebraProxy  = (sharpurs_apply (box ((box Data_BooleanAlgebra_BooleanAlgebrausd_Dict))) (box ((box ((Map.add "HeytingAlgebra0" (box ((box (fun (usd__unused: obj) -> (box Data_HeytingAlgebra_heytingAlgebraProxy))))) Map.empty))))))

let Data_BooleanAlgebra_booleanAlgebraFn  = (box (fun (dictBooleanAlgebra: obj) -> (let heytingAlgebraFunction = (sharpurs_apply (box ((box Data_HeytingAlgebra_heytingAlgebraFunction))) (box ((sharpurs_apply (box ((Map.find "HeytingAlgebra0" (unbox<Map<string, obj>> ((box dictBooleanAlgebra)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_BooleanAlgebra_BooleanAlgebrausd_Dict))) (box ((box ((Map.add "HeytingAlgebra0" (box ((box (fun (usd__unused: obj) -> (box heytingAlgebraFunction))))) Map.empty)))))))))

let Data_BooleanAlgebra_booleanAlgebraBoolean  = (sharpurs_apply (box ((box Data_BooleanAlgebra_BooleanAlgebrausd_Dict))) (box ((box ((Map.add "HeytingAlgebra0" (box ((box (fun (usd__unused: obj) -> (box Data_HeytingAlgebra_heytingAlgebraBoolean))))) Map.empty))))))
